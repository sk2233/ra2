/*
@author: sk
@date: 2023/9/2
*/
package manager

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"ra3/global"
	"ra3/global/collection"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/tile"
	"ra3/unit"
	"ra3/utils"
	"strings"

	"github.com/beevik/etree"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	tileManagerImpl = &tileManager{}
)

func init() {
	tool.TileManager = tileManagerImpl
}

func initTileManager() {
	tileManagerImpl.loadTile()
	tileManagerImpl.tidyTile()
}

type tileManager struct {
	tiles [][][]*tile.Tile // X,Y,Z
}

type gridNode struct {
	Grid model.Grid
	Pre  *gridNode
	Cost float64 // 已经消耗的
	Len  float64 // 将来估算消耗的(曼哈顿距离)
}

func (m *tileManager) GetLandPath(target model.Grid, temp any) []model.Grid {
	unit0 := temp.(*unit.Unit)
	grids := m.getNearGrids(target, []*unit.Unit{unit0}, true)
	if len(grids) == 0 {
		return nil
	}
	target = grids[unit0]
	queue := collection.NewPriorityQueue[*gridNode](func(n1, n2 *gridNode) bool {
		return n1.Cost+n1.Len < n2.Cost+n2.Len
	})
	set := collection.NewSet[string]()
	node := &gridNode{Grid: unit.GetGrid(unit0), Cost: 0}
	node.Len = getLen(node.Grid, target)
	queue.Offer(node)
	set.Add(node.Grid.String())
	for !queue.IsEmpty() {
		node = queue.Poll()
		nexts, costs := m.getNextGrids(node.Grid, true)
		for i, next := range nexts {
			if next.Equal(target) {
				res := make([]model.Grid, 0)
				res = append(res, next)
				for node.Pre != nil {
					res = append(res, node.Grid)
					node = node.Pre
				}
				return utils.ReverseSlice(res)
			}
			if !set.Has(next.String()) {
				set.Add(next.String())
				item := m.GetTile(next).(*tile.Tile)
				capacity := tool.UnitManager.GetCapacity(next, true)
				if unit.CanStand(unit0, item.Type, capacity) {
					queue.Offer(&gridNode{Grid: next, Pre: node, Cost: node.Cost + costs[i], Len: getLen(next, target)})
				}
			}
		}
	}
	return nil
}

func getLen(src model.Grid, tar model.Grid) float64 {
	return float64(utils.Abs(src.X-tar.X) + utils.Abs(src.Y-tar.Y))
}

// temp 只会是步兵与载具的数组
func (m *tileManager) GetNearGrids(grid model.Grid, temp any) any {
	units := temp.([]*unit.Unit)
	landUnits := make([]*unit.Unit, 0)
	airUnits := make([]*unit.Unit, 0)
	for _, item := range units {
		if item.Data.StandExtra.Land { // 载具与小兵看其land区分空军与陆军
			landUnits = append(landUnits, item)
		} else {
			airUnits = append(airUnits, item)
		}
	}
	landRes := m.getNearGrids(grid, landUnits, true)
	airRes := m.getNearGrids(grid, airUnits, false)
	return utils.MergeMap(landRes, airRes)
}

// 注意这里的入参是笛卡尔坐标系下的
func (m *tileManager) getNearGrids(grid model.Grid, units []*unit.Unit, land bool) map[*unit.Unit]model.Grid {
	res := make(map[*unit.Unit]model.Grid) // 这里的 units 都是步兵与载具
	if len(units) == 0 {
		return res
	}
	queue := collection.NewPriorityQueue[*gridNode](func(n1, n2 *gridNode) bool {
		return n1.Cost < n2.Cost
	})
	set := collection.NewSet[string]()
	node := &gridNode{Grid: grid, Cost: 0}
	queue.Offer(node)
	set.Add(grid.String())
	for !queue.IsEmpty() && len(units) > 0 {
		node = queue.Poll()
		capacity := tool.UnitManager.GetCapacity(node.Grid, land)
		if capacity > 0 { // 还有容量尝试分配一下空间
			temp, _ := m.GetTile(node.Grid).(*tile.Tile) // 只有空军这里为nil，但是空军不会使用该变量
			removeUnits := make([]*unit.Unit, 0)
			for i := 0; i < len(units) && capacity > 0; i++ {
				if land && !unit.CanStand(units[i], temp.Type, capacity) {
					continue
				}
				res[units[i]] = node.Grid
				capacity -= unit.GetSize(units[i])
				removeUnits = append(removeUnits, units[i])
			}
			units = utils.RemoveSlice(units, removeUnits)
		}
		grids, _ := m.getNextGrids(node.Grid, land)
		for _, temp := range grids {
			if !set.Has(temp.String()) {
				// 优先找距离grid 最近的
				queue.Offer(&gridNode{Grid: temp, Cost: float64(temp.Sub(grid).Len2())})
				set.Add(temp.String())
			}
		}
	}
	return res
}

var ( // 只是x,y方向上的偏移  优先距离最近的直边再是斜边
	tileOffsets = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, -1}, {1, 1}, {-1, -1}, {-1, 1}}
	tileCosts   = []float64{1, 1, 1, 1, math.Sqrt2, math.Sqrt2, math.Sqrt2, math.Sqrt2}
)

// grid 是笛卡尔坐标系下的
func (m *tileManager) getNextGrids(grid model.Grid, land bool) ([]model.Grid, []float64) {
	res := make([]model.Grid, 0)
	cost := make([]float64, 0)
	for i, offset := range tileOffsets {
		x, y := utils.ToIsometric(grid.X+offset[0], grid.Y+offset[1]) // 转换到 等距坐标系下
		if x < 0 || x >= m.GetXNum() || y < 0 || y >= m.GetYNum() {
			continue
		}
		if !land { // 天空中有就行
			res = append(res, grid.Add(model.NewGrid(offset[0], offset[1], 0))) // 注意返回的是笛卡尔坐标系
			cost = append(cost, tileCosts[i])
			continue
		} // 地面上的情况
		// 从上向下 找到一个合适的方块就可以结束了
		for z := grid.Z + 1; z >= grid.Z-1; z-- { // 陆地只需要看 z方向 上中下 3格即可(没有太陡峭的地形)
			if z < 0 || z >= m.GetZNum() {
				continue
			} // 只要不为空是什么都要结束
			if temp := m.tiles[x][y][z]; temp != nil {
				if temp.Type != global.TileTypeBlock {
					res = append(res, model.Grid{X: grid.X + offset[0], Y: grid.Y + offset[1], Z: z}) // 注意返回的是笛卡尔坐标系
					cost = append(cost, tileCosts[i])
				}
				break
			}
		}
	} // 乱序返回防止寻路，或找就近点过于统一
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
		cost[i], cost[j] = cost[j], cost[i]
	})
	return res, cost
}

func (m *tileManager) CheckTileType(grid model.Grid, tileType model.TileType) bool {
	grid = utils.ToIsometricGrid(grid)
	if grid.X < 0 || grid.X >= m.GetXNum() || grid.Y < 0 || grid.Y >= m.GetYNum() || grid.Z < 0 || grid.Z >= m.GetZNum() {
		return false
	}
	temp := m.tiles[grid.X][grid.Y][grid.Z]
	return temp != nil && int(temp.Type)&int(tileType) > 0
}

func (m *tileManager) GetHeight(x float64, y float64) float64 {
	xIndex := utils.GetAdjustIndex(x)
	yIndex := utils.GetAdjustIndex(y)
	xIndex, yIndex = utils.ToIsometric(xIndex, yIndex) // 转换到 格子坐标系
	if xIndex < 0 || yIndex < 0 {
		return 0
	}
	tiles := m.tiles[xIndex][yIndex]
	for z := m.GetZNum() - 1; z >= 0; z-- {
		if tiles[z] != nil {
			return tiles[z].GetHeight(x, y)
		}
	}
	return 0
}

// 传入的是笛卡尔坐标系下的
func (m *tileManager) GetTile(grid model.Grid) any {
	grid = utils.ToIsometricGrid(grid)
	if grid.X < 0 || grid.X >= m.GetXNum() || grid.Y < 0 || grid.Y >= m.GetYNum() || grid.Z < 0 || grid.Z >= m.GetZNum() {
		return nil
	}
	if m.tiles[grid.X][grid.Y][grid.Z] == nil { // 防止返回的只有接口信息 被认为非空
		return nil
	}
	return m.tiles[grid.X][grid.Y][grid.Z]
}

func (m *tileManager) loadTile() {
	root := utils.OpenXml(global.MapPath)
	images, types := m.getTileImage(root.FindElements("tileset"))
	layers := root.FindElements("layer")
	temp := make([][][]*tile.Tile, 0)
	for z, layer := range layers {
		temp = append(temp, m.createTile(layer, images, types, z))
	}
	// 纠正为 x,y,z
	zNum, yNum, xNum := len(temp), len(temp[0]), len(temp[0][0])
	m.tiles = make([][][]*tile.Tile, xNum)
	for x := 0; x < xNum; x++ {
		m.tiles[x] = make([][]*tile.Tile, yNum)
		for y := 0; y < yNum; y++ {
			m.tiles[x][y] = make([]*tile.Tile, zNum)
			for z := 0; z < zNum; z++ {
				m.tiles[x][y][z] = temp[z][y][x]
			}
		}
	}
}

var (
	tileTypes = map[string]model.TileType{
		"land":   global.TileTypeLand,
		"slope1": global.TileTypeSlope1,
		"slope2": global.TileTypeSlope2,
		"water":  global.TileTypeWater,
		"block":  global.TileTypeBlock,
	}
)

func (m *tileManager) getTileImage(tilesets []*etree.Element) (map[uint32]*ebiten.Image, map[uint32]model.TileType) {
	images := make(map[uint32]*ebiten.Image)
	types := make(map[uint32]model.TileType)
	for _, tileset := range tilesets {
		path := utils.XmlStr(tileset, "source", "")
		root := utils.OpenXml("res" + path[2:])
		elems := root.FindElements("tile")
		for _, elem := range elems {
			id := uint32(utils.XmlUint(elem, "id", 0) + 1)
			type0 := utils.XmlStr(elem, "type", "")
			if res, ok := tileTypes[type0]; ok {
				types[id] = res
			} else {
				panic(fmt.Sprintf("invalid tileType of %v , id : %v", type0, id))
			}
			source := utils.XmlStr(elem.FindElement("image"), "source", "")
			images[id] = tool.ImageLoader.LoadImage(source)
		}
	}
	return images, types
}

func (m *tileManager) createTile(layer *etree.Element, images map[uint32]*ebiten.Image, types map[uint32]model.TileType, z int) [][]*tile.Tile {
	wNum := int(utils.XmlInt(layer, "width", 0))
	hNum := int(utils.XmlInt(layer, "height", 0))
	res := make([][]*tile.Tile, hNum)
	data := layer.FindElement("data").Text()
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(data))
	utils.HandleErr(err)
	reader, err := zlib.NewReader(bytes.NewReader(bs))
	utils.HandleErr(err)
	for y := 0; y < hNum; y++ {
		res[y] = make([]*tile.Tile, wNum)
		for x := 0; x < wNum; x++ {
			id := utils.ReadUin32(reader)
			if id > 0 {
				grid := utils.ToCommonGrid(model.NewGrid(x, y, z))
				pos := utils.Grid2Pos(grid)
				res[y][x] = tile.NewTile(pos, images[id], types[id])
				tool.DrawManager.AddDraw(res[y][x])
			}
		}
	}
	return res
}

func (m *tileManager) tidyTile() {
	xNum, yNum, zNum := m.GetXNum(), m.GetYNum(), m.GetZNum()
	for z := 0; z < zNum-1; z++ {
		for y := 0; y < yNum-1; y++ {
			if y%2 == 0 { // 奇数层与偶数层规律不同
				for x := 1; x < xNum; x++ {
					if m.tiles[x][y][z] != nil && m.tiles[x-1][y+1][z] != nil && m.tiles[x][y+1][z] != nil && m.tiles[x][y][z+1] != nil {
						m.tiles[x][y][z].Enable = false
					}
				}
			} else {
				for x := 0; x < xNum-1; x++ {
					if m.tiles[x][y][z] != nil && m.tiles[x][y+1][z] != nil && m.tiles[x+1][y+1][z] != nil && m.tiles[x][y][z+1] != nil {
						m.tiles[x][y][z].Enable = false
					}
				}
			}
		}
	}
}

func (m *tileManager) GetZNum() int {
	return len(m.tiles[0][0])
}

func (m *tileManager) GetXNum() int {
	return len(m.tiles)
}

func (m *tileManager) GetYNum() int {
	return len(m.tiles[0])
}
