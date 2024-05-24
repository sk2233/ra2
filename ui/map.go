/*
@author: sk
@date: 2023/9/16
*/
package ui

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

type Map struct {
	rect         *model.Rect
	miniMap      *ebiten.Image
	unitImg      *ebiten.Image
	unitOption   *ebiten.DrawImageOptions
	visionOption *ebiten.DrawImageOptions
	visionRect   *model.Rect
	visionScale  model.Vec2
	visionOffset model.Vec2
}

func (m *Map) HandleCursor(pos model.Vec2) bool {
	if !utils.PointCollision(pos, m.rect) { // 都是先检验区域
		return false
	}
	return true
}

func NewMap() *Map { // 144 * 88
	miniMap := tool.ImageLoader.LoadImage("res/image/minimap.png")
	option := &ebiten.DrawImageOptions{} // 设置好位置缩放
	bound := miniMap.Bounds()
	scale := utils.Min(144/float64(bound.Dx()/2), 88/float64(bound.Dy())) // x 先缩放到 0.5 地图纵横距离必须一致
	option.GeoM.Scale(scale/2, scale)
	realW := float64(bound.Dx()) * scale / 2
	realH := float64(bound.Dy()) * scale
	xOffset := (144 - realW) / 2
	yOffset := (88 - realH) / 2
	option.GeoM.Translate(global.WindowW-144+xOffset, 16+yOffset)
	img := tool.ImageLoader.LoadImage("res/image/minimap_grid.png")
	visionScale := model.NewVec2(realW/float64(tool.TileManager.GetXNum()*64-64/2),
		realH/float64((tool.TileManager.GetYNum()-1)*64/2/2))
	return &Map{
		rect: &model.Rect{
			Pos:  model.NewVec2(global.WindowW-144, 16),
			Size: model.NewVec2(144, 88),
		},
		miniMap:      miniMap,
		unitImg:      utils.SplitImage(img, 2, 3)[4],
		visionOption: option,
		visionScale:  visionScale,
		visionRect: &model.Rect{
			Size: global.WindowRegion.Size.Mul(visionScale),
		},
		visionOffset: model.NewVec2(global.WindowW-144+xOffset, 16+yOffset),
		unitOption:   &ebiten.DrawImageOptions{},
	}
}

func (m *Map) Draw(screen *ebiten.Image) { // 144 * 88
	utils.FillRect(screen, m.rect, colornames.Black, global.MenuRegion)
	screen.DrawImage(m.miniMap, m.visionOption)
	fogImg := tool.FogManager.GetMiniMap()
	screen.DrawImage(fogImg, m.visionOption)
	m.visionRect.Pos = tool.Camera.GetPos().Mul(m.visionScale).Add(m.visionOffset)
	utils.StrokeRect(screen, m.visionRect, 1, colornames.Yellow, global.MenuRegion)
	utils.StrokeRect(screen, m.rect, 1, colornames.Yellow, global.MenuRegion)
}

func (m *Map) UpdateMiniMap() { // 更新小地图
	m.miniMap = tool.ImageLoader.LoadImage("res/image/minimap.png")
	units := tool.UnitManager.FilterUnits(func(temp any) bool {
		unit0 := temp.(*unit.Unit)
		return unit0.Data.Type == global.UnitTypeConstruct // 只管建筑
	}).([]*unit.Unit)
	for _, item := range units {
		grid := unit.GetGrid(item)
		for x := 0; x < item.Data.ConstructExtra.WNum; x++ {
			for y := 0; y < item.Data.ConstructExtra.HNum; y++ {
				temp := grid.Add(model.NewGrid(x, y, 0))
				temp = utils.ToIsometricGrid(temp)
				m.unitOption.GeoM.Reset()
				m.unitOption.GeoM.Translate(utils.MiniMapGrid2Pos(temp.X, temp.Y))
				m.miniMap.DrawImage(m.unitImg, m.unitOption)
			}
		}
	}
}
