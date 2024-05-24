/*
@author: sk
@date: 2023/9/16
*/
package manager

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/ui"
	"ra3/unit"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	tool.CursorManager = &cursorManager{
		type0: global.CursorTypeNormal,
	}
}

var (
	okGridImg, notOkGridImg *ebiten.Image
	gridAnchor              = model.Vec2{X: 32}
)

func initCursorManager() {
	image := tool.ImageLoader.LoadImage("res/image/grid.png")
	images := utils.SplitImage(image, 2, 2)
	okGridImg = images[0]
	notOkGridImg = images[1]
}

type cursorManager struct {
	type0    model.CursorType
	iconItem *ui.IconItem
}

func (c *cursorManager) DrawAfter(screen *ebiten.Image) {
	pos := utils.GetCursorPos()
	switch c.type0 {
	case global.CursorTypeConstruct:
		c.constructDraw(screen, pos)
	}
}

func (c *cursorManager) StartConstruct(iconItem any) {
	c.iconItem = iconItem.(*ui.IconItem)
	c.type0 = global.CursorTypeConstruct
}

func (c *cursorManager) Update() {
	pos := utils.GetCursorPos()
	switch c.type0 {
	case global.CursorTypeNormal:
		c.normalUpdate(pos)
	case global.CursorTypeConstruct:
		c.constructUpdate(pos)
	}
}

func (c *cursorManager) GetType() model.CursorType {
	return c.type0
}

func (c *cursorManager) normalUpdate(pos model.Vec2) {
	if tool.UiManager.HandleCursor(pos) {
		return
	}
	if tool.UnitManager.HandleCursor(pos) {
		return
	}
}

func (c *cursorManager) constructUpdate(pos2 model.Vec2) {
	if !utils.PointCollision(pos2, global.WindowRegion) {
		return
	}
	pos3 := tool.Camera.Screen2World(pos2)
	if pos3 == global.InvalidVec3 {
		return
	}
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}
	// 收集要创建的位置，主要是扩展类建筑，会自动扩展，可能不止一个
	grid := utils.Pos2Grid(pos3).Down()
	grids := make([]model.Grid, 0)
	grids = append(grids, grid)
	name := c.iconItem.Data.Name
	data := tool.DataManager.GetUnitData(name)
	if data.ConstructExtra.Tag == global.ConstructTagExtend {
		grids = append(grids, extendGrid(grid, data)...)
	}
	// 尝试创建对象
	create := false
	for _, temp := range grids {
		if tool.UnitManager.CreateUnit(name, temp) != nil {
			create = true // 创建成功一个计算创建成功
		}
	}
	if create { // 创建成功，清楚数据
		utils.PlaySe("uplace.wav")
		c.iconItem.Reset()
		tool.UiManager.UpdateIconItems()
		tool.UiManager.UpdateMiniMap()
		tool.DataManager.UpdatePower()
		c.iconItem = nil
		c.type0 = global.CursorTypeNormal
		tool.UnitManager.BroadcastCmd(func(temp any) bool {
			unit0 := temp.(*unit.Unit)
			return unit0.Data.Type == global.UnitTypeConstruct && unit0.Data.ConstructExtra.Tag == global.ConstructTagBase
		}, &model.UnitCmd{
			Type:     global.UnitCmdTypeConstructComplete,
			UnitName: name,
		})
	}
}

func (c *cursorManager) constructDraw(screen *ebiten.Image, pos2 model.Vec2) {
	if !utils.PointCollision(pos2, global.WindowRegion) {
		return
	}
	pos3 := tool.Camera.Screen2World(pos2)
	if pos3 == global.InvalidVec3 {
		return
	}
	grid := utils.Pos2Grid(pos3).Down()
	data := tool.DataManager.GetUnitData(c.iconItem.Data.Name)
	grids := make([]model.Grid, 0)
	if data.ConstructExtra.Tag == global.ConstructTagExtend {
		grids = append(grids, grid)
		grids = append(grids, extendGrid(grid, data)...)
	} else {
		wNum, hNum := data.ConstructExtra.WNum, data.ConstructExtra.HNum
		for x := 0; x < wNum; x++ {
			for y := 0; y < hNum; y++ {
				grids = append(grids, grid.Add(model.Grid{X: x, Y: y}))
			}
		}
	}
	for _, temp := range grids {
		pos3 = utils.Grid2Pos(temp)
		pos3.Z += global.BlockSize
		pos2 = tool.Camera.World2Screen(pos3)
		if tool.UnitManager.CheckGrid(temp, data) {
			utils.DrawImage(screen, okGridImg, pos2, gridAnchor, global.WindowRegion)
		} else {
			utils.DrawImage(screen, notOkGridImg, pos2, gridAnchor, global.WindowRegion)
		}
	}
}

var (
	extendDirs = []model.Grid{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
)

func extendGrid(grid model.Grid, data *model.UnitData) []model.Grid {
	res := make([]model.Grid, 0)
	for _, extendDir := range extendDirs {
		res = append(res, extendGridDir(grid, extendDir, data)...)
	}
	return res
}

func extendGridDir(grid model.Grid, dir model.Grid, data *model.UnitData) []model.Grid {
	res := make([]model.Grid, 0)
	for i := 0; i < 5; i++ {
		grid = grid.Add(dir)
		if tool.UnitManager.CheckGrid(grid, data) {
			res = append(res, grid)
		} else {
			temp := tool.UnitManager.GetUnits(grid, true)
			if units, ok := temp.([]*unit.Unit); ok && units[3] != nil && units[3].Data == data {
				return res
			}
			break
		}
	}
	return make([]model.Grid, 0)
}
