/*
@author: sk
@date: 2023/8/20
*/
package manager

import (
	"math"
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	tool.FogManager = fogManagerImpl
}

var (
	fogManagerImpl = &fogManager{
		option: &ebiten.DrawImageOptions{
			Blend: ebiten.BlendSourceIn, // 因为对应的图片有黑边，只能进行叠加 只看α_out 取两者最小的保留
		},
	}
)

func initFogManager() {
	xNum, yNum := tool.TileManager.GetXNum(), tool.TileManager.GetYNum()
	visions := make([][]float64, xNum)
	for x := 0; x < xNum; x++ { // 默认全部初始化为0
		visions[x] = make([]float64, yNum)
	}
	fogManagerImpl.visions = visions
	img := tool.ImageLoader.LoadImage("res/image/fog_tile.png")
	fogManagerImpl.fogImgs = utils.SplitImage(img, 4, 4)
	fogManagerImpl.miniMap = ebiten.NewImage(xNum*global.MiniMapBlockSize-global.MiniMapBlockSize/2,
		(yNum-1)*global.MiniMapBlockSize/2/2)
	fogManagerImpl.miniMap.Fill(colornames.Black)
	img = tool.ImageLoader.LoadImage("res/image/minimap_grid.png")
	fogManagerImpl.visionImg = utils.SplitImage(img, 2, 3)[5]
}

type fogManager struct {
	fogImgs   []*ebiten.Image // 8张迷雾图片 按位运算来
	visions   [][]float64     // 记录在改点的最大视野值，若是后面有更大的进行更新，默认全是0，代表没有视野,大于0才有视野
	miniMap   *ebiten.Image   // 迷雾小地图遮罩
	visionImg *ebiten.Image   // 可以看到的部分绘制，绘制到miniMap上
	option    *ebiten.DrawImageOptions
}

func (m *fogManager) GetMiniMap() *ebiten.Image {
	return m.miniMap
}

var (
	//                 1        2        4       8
	fogDir1 = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}   // 只有横向竖向 代价为1的
	fogDir2 = [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} // 只有斜向，代价为 根号2的
)

// 传入的是世界坐标
func (m *fogManager) getImageIndex(x int, y int) int {
	res := 0
	for i, dir := range fogDir1 {
		tx, ty := x+dir[0], y+dir[1]
		tx, ty = utils.ToIsometric(tx, ty)                              // 转换为等距坐标
		if tx < 0 || tx >= m.GetXNum() || ty < 0 || ty >= m.GetYNum() { // 越界也算没有视野
			res += 1 << i
		} else if m.visions[tx][ty] == 0 {
			res += 1 << i
		}
	}
	return res
}

func (m *fogManager) GetXNum() int {
	return len(m.visions)
}

func (m *fogManager) GetYNum() int {
	return len(m.visions[0])
}

var (
	fogAnchor = model.NewVec2(32, 0)
	fogRect   = &model.Rect{Anchor: fogAnchor, Size: model.NewVec2(64, 32)}
)

func (m *fogManager) Draw(screen *ebiten.Image) {
	for x := 0; x < m.GetXNum(); x++ {
		for y := 0; y < m.GetYNum(); y++ {
			if m.visions[x][y] > 0 { // 已经有视野了
				continue
			}
			grid := model.NewGrid(x, y, 0)
			grid = utils.ToCommonGrid(grid) // 转换到世界坐标系
			pos3 := utils.Grid2Pos(grid)
			pos2 := tool.Camera.World2Screen(pos3)
			fogRect.Pos = pos2 // 提前优化一波
			if !utils.RectCollision(fogRect, global.WindowRegion) {
				continue
			}
			index := m.getImageIndex(grid.X, grid.Y) // 使用世界坐标系
			utils.DrawImage(screen, m.fogImgs[index], pos2, fogAnchor, global.WindowRegion)
		}
	}
}

// 传入的是世界坐标系下的要注意转换
func (m *fogManager) ClearFog(grid model.Grid, vision float64) {
	x, y := utils.ToIsometric(grid.X, grid.Y)
	if x < 0 || x >= m.GetXNum() || y < 0 || y >= m.GetYNum() {
		return // 出界裁剪
	}
	if m.visions[x][y] >= vision {
		return // 之前更新过，无需更新裁剪
	}
	if m.visions[x][y] == 0 { // 还没有动过，补充minimap
		m.option.GeoM.Reset()
		m.option.GeoM.Translate(utils.MiniMapGrid2Pos(x, y))
		m.miniMap.DrawImage(m.visionImg, m.option)
	}
	m.visions[x][y] = vision      // 更新最大值
	for _, dir := range fogDir1 { // 代价1  递归肯定会终止，因为vision最终会小于0
		m.ClearFog(grid.Add(model.NewGrid(dir[0], dir[1], 0)), vision-1)
	}
	for _, dir := range fogDir2 { // 代价 根号2的
		m.ClearFog(grid.Add(model.NewGrid(dir[0], dir[1], 0)), vision-math.Sqrt2)
	}
}

func (m *fogManager) ClearFogs(grids []model.Grid, vision float64) {
	for _, grid := range grids {
		m.ClearFog(grid, vision)
	}
}
