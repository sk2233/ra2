/*
@author: sk
@date: 2023/9/9
*/
package interfaces

import (
	"ra3/global/model"

	"github.com/hajimehoshi/ebiten/v2"
)

type ICursorHandler interface {
	HandleCursor(pos model.Vec2) bool
}

type IDrawManager interface {
	IBeforeDraw
	IDraw
	IAfterDraw
	AddDraw(draw IOrderDraw)
	RemoveDraw(draws ...IOrderDraw)
}

type ITileManager interface {
	GetXNum() int
	GetYNum() int
	GetZNum() int
	GetTile(grid model.Grid) any
	// 获取对应位置的高度，有些tile 不同位置高度不一 必须 传详细位置
	GetHeight(x float64, y float64) float64
	CheckTileType(grid model.Grid, tile model.TileType) bool
	// units []*unit.Unit    返回值 map[*unit.Unit]model.Grid
	GetNearGrids(grid model.Grid, units any) any
	GetLandPath(grid model.Grid, unit any) []model.Grid
}

type IDataManager interface {
	GetUnitData(name string) *model.UnitData
	GetMoney() int64
	GetPower() (int64, int64)
	GetIconData(name string) *model.IconData
	FilterIconDatas(filter func(icon *model.IconData) bool) []*model.IconData
	ChangeMoney(value int64) bool
	UpdatePower()
	GetEffectData(name string) *model.EffectData
}

type IUnitManager interface {
	IUpdate
	IAfterDraw

	CreateUnit(name string, grid model.Grid) any
	RegisterActionCreator(actionType model.UnitActionType, actionCreator IUnitActionCreator)
	GetUnits(grid model.Grid, land bool) any
	CheckGrids(grid model.Grid, data *model.UnitData) bool
	CheckGrid(grid model.Grid, data *model.UnitData) bool
	AddUnit(unit any)
	// 返回 unit 数组
	FilterUnits(filter func(unit any) bool) any
	// 获取某点的高度，若是有单位，获取单位的基点高度+它的高度，否则获取tile高度
	GetHeight(x float64, y float64) float64
	HandleCursor(pos model.Vec2) bool
	GetCapacity(grid model.Grid, land bool) int
	AddToGrid(grid model.Grid, temp any)
	BuildUnit(grid model.Grid, data *model.UnitData) any
	RemoveFromGrid(grid model.Grid, temp any)
	GetPos(grid model.Grid, data *model.UnitData) model.Vec3
	BroadcastCmd(filter func(unit any) bool, cmd *model.UnitCmd)
	RemoveSelect(unit any)
	SetSelectUnit(units any)
}

type IUiManager interface {
	IAfterDraw
	IUpdate
	ICursorHandler

	UpdateIconItems()
	UpdateMiniMap()
}

type ICursorManager interface {
	IUpdate
	IAfterDraw

	GetType() model.CursorType
	StartConstruct(iconItem any)
}

type IEffectManager interface {
	Update()
	CreateEffect(name string, pos model.Vec3) any
}

type IAudioManager interface {
	PlayBgm(path string, second int64)
	PlaySe(path string)
}

type IFogManager interface {
	Draw(screen *ebiten.Image)
	ClearFog(grid model.Grid, vision float64)
	ClearFogs(grids []model.Grid, vision float64)
	GetMiniMap() *ebiten.Image
}
