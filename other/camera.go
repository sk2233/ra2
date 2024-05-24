/*
@author: sk
@date: 2023/9/2
*/
package other

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	tool.Camera = &camera{}
}

type camera struct {
	offset model.Vec2
}

func (c *camera) GetPos() model.Vec2 {
	return c.offset
}

func (c *camera) SetPos(pos model.Vec2) {
	c.offset = pos
}

func (c *camera) Screen2World(pos model.Vec2) model.Vec3 {
	zNum := tool.TileManager.GetZNum()
	for z := zNum; z > 0; z-- { // 高使用方块高，更符合视觉效果
		temp := c.Screen2WorldAtZ(pos, float64(z)*global.BlockSize)
		grid := utils.Pos2Grid(temp.Sub(model.NewVec3(0, 0, 1))) // 转 grid 必须恢复高
		if tool.TileManager.GetTile(grid) != nil {
			return temp
		}
	}
	return global.InvalidVec3
}

func (c *camera) Update() {
	c.offset.X += utils.GetAxis(ebiten.KeyA, ebiten.KeyD) * 5
	c.offset.Y += utils.GetAxis(ebiten.KeyW, ebiten.KeyS) * 5
	//if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	//	fmt.Println(c.xOffset, c.yOffset)
	//}
}

func (c *camera) World2Screen(pos model.Vec3) model.Vec2 {
	return model.NewVec2(pos.Y-pos.X-c.offset.X, (pos.X+pos.Y)/2-pos.Z-c.offset.Y)
}

func (c *camera) Screen2WorldAtZ(pos model.Vec2, z float64) model.Vec3 {
	y := (pos.X + pos.Y*2 + c.offset.X + c.offset.Y*2 + z*2) / 2
	x := y - pos.X - c.offset.X
	return model.NewVec3(x, y, z)
}
