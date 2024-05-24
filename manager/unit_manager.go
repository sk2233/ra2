/*
@author: sk
@date: 2023/9/9
*/
package manager

import (
	"fmt"
	"math"
	"ra3/global"
	"ra3/global/collection"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/tile"
	"ra3/unit"
	"ra3/unit/anim"
	"ra3/unit/blood_bar"
	"ra3/unit/mask"
	"ra3/unit/shadow"
	"ra3/utils"
	"sort"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	unitManagerImpl = &unitManager{units: make([]*unit.Unit, 0),
		actionCreators: make(map[model.UnitActionType]interfaces.IUnitActionCreator)}
)

func init() {
	tool.UnitManager = unitManagerImpl
}

func initUnitManager() {
	//  单位在tile上也算在该tile内，主要方便处理斜坡的情况
	xNum, yNum, zNum := tool.TileManager.GetXNum(), tool.TileManager.GetYNum(), tool.TileManager.GetZNum()
	landUnits := make([][][][]*unit.Unit, xNum)
	airUnits := make([][][]*unit.Unit, xNum)
	for x := 0; x < xNum; x++ {
		landUnits[x] = make([][][]*unit.Unit, yNum)
		airUnits[x] = make([][]*unit.Unit, yNum)
		for y := 0; y < yNum; y++ {
			landUnits[x][y] = make([][]*unit.Unit, zNum)
			airUnits[x][y] = make([]*unit.Unit, 4)
			for z := 0; z < zNum; z++ {
				landUnits[x][y][z] = make([]*unit.Unit, 4)
			}
		}
	}
	unitManagerImpl.landUnits = landUnits
	unitManagerImpl.airUnits = airUnits
	unitManagerImpl.loadUnits()
}

type unitManager struct {
	units          []*unit.Unit       // 汇总unit 方便遍历
	landUnits      [][][][]*unit.Unit // 地面上的unit方便判断位置  X,Y,Z   前3个给小兵用，最后一个给载具或者建筑使用
	airUnits       [][][]*unit.Unit   // 空中的unit   X,Y,Z   这些位置存储都是使用等距坐标系外面一般传递笛卡尔坐标系
	actionCreators map[model.UnitActionType]interfaces.IUnitActionCreator
	touchUnit      *unit.Unit
	selectUnits    []*unit.Unit
	selectRange    *model.Range
}

func (u *unitManager) BroadcastCmd(filter func(unit any) bool, cmd *model.UnitCmd) {
	for _, temp := range u.units {
		if !filter(temp) {
			continue
		}
		temp.Execute(cmd)
		if cmd.Abort {
			break
		}
	}
}

func (u *unitManager) GetCapacity(grid model.Grid, land bool) int {
	temp, ok := u.GetUnits(grid, land).([]*unit.Unit)
	if !ok || temp == nil {
		return 0
	}
	if temp[3] != nil {
		return 0
	}
	res := 0
	for i := 0; i < 3; i++ {
		if temp[i] == nil {
			res++
		}
	}
	return res
}

func (u *unitManager) DrawAfter(screen *ebiten.Image) {
	if u.selectRange != nil {
		utils.StrokeRect(screen, u.selectRange, 1, colornames.White, global.WindowRegion)
	}
}

func (u *unitManager) HandleCursor(pos model.Vec2) bool {
	sort.Slice(u.units, func(i, j int) bool { // 点击处理前，排一下序，先处理最先看到的对象
		return u.units[i].GetOrder() > u.units[j].GetOrder()
	})
	if u.handleHover(pos) {
		return true
	}
	if u.handleMove(pos) {
		return true
	}
	if u.handleDeploy(pos) {
		return true
	}
	if u.handleUnit(pos) {
		return true
	}
	if u.handleSelect(pos) {
		return true
	}
	return false
}

// 处理选择单位与第二次选择单位的交互，例如小兵上船，小兵入驻
func (u *unitManager) handleUnit(pos model.Vec2) bool {
	if len(u.selectUnits) == 0 || !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false
	}
	// 点击的对象必须不在选择列表
	temp := u.clickUnit(pos)
	if temp == nil || temp.Select {
		return false
	}
	ok := false
	for _, unit0 := range u.selectUnits {
		unit0.PopAllAction()
		ok = unit0.Execute(&model.UnitCmd{
			Type: global.UnitCmdTypeHandleUnit,
			Unit: temp,
		}) || ok // 成功一次就算处理了，但是ok放后面防止短路
	}
	return ok
}

func (u *unitManager) handleDeploy(pos model.Vec2) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// 点击的对象必须已经被选择了
		if temp := u.clickUnit(pos); temp != nil && temp.Select {
			temp.PopAllAction()
			temp.Execute(&model.UnitCmd{
				Type: global.UnitCmdTypeDeploy,
			})
			return true
		}
	}
	return false
}

// 对于步兵，载具是移动，生产建筑是设置集结点，部分建筑也是移动(例如基地车，里面通过特殊的建筑handler处理)
func (u *unitManager) handleMove(pos2 model.Vec2) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if len(u.selectUnits) > 0 && u.clickUnit(pos2) == nil { // 必须选择了单位，且目的地没有单位
			pos3 := tool.Camera.Screen2World(pos2)
			if pos3 == global.InvalidVec3 {
				return false
			}
			temp := utils.Pos2Grid(pos3).Down() // 获取设置的目标点
			landUnits := make([]*unit.Unit, 0)  // 在地面的载具与步兵，移动需要预先算好位置
			units := make([]*unit.Unit, 0)
			for _, item := range u.selectUnits {
				if item.Data.Type == global.UnitTypeConstruct {
					item.Execute(&model.UnitCmd{ // 建筑设置集合点，(注意基地车比较特殊，会移动过来)
						Type: global.UnitCmdTypeRallyPoint,
						Grid: temp,
					})
				} else if item.Data.StandExtra.Land { // 地面单位
					landUnits = append(landUnits, item)
					units = append(units, item)
				} else {
					item.PopAllAction()          // 手动指定的指令，会取消前面所有的指令
					item.Execute(&model.UnitCmd{ // 飞行单位移动不计算体积
						Type: global.UnitCmdTypeMove,
						Grid: temp,
					})
					units = append(units, item)
				}
			}
			unit.PlayRandSe(units, global.AudioMoveCmd) // 播放音效
			if len(landUnits) > 0 {                     // 处理地面的步兵与载具移动设置，主要需要预选位置
				res := tool.TileManager.GetNearGrids(temp, landUnits).(map[*unit.Unit]model.Grid)
				for _, landUnit := range landUnits {
					if grid, ok := res[landUnit]; ok {
						landUnit.PopAllAction() // 手动指定的指令，会取消前面所有的指令
						landUnit.Execute(&model.UnitCmd{
							Type: global.UnitCmdTypeMove,
							Grid: grid,
						})
					}
				}
			}
			return true
		}
	}
	return false
}

func (u *unitManager) RemoveSelect(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Select = false
	u.selectUnits = utils.RemoveSlice(u.selectUnits, []*unit.Unit{unit0})
}

func (u *unitManager) handleSelect(pos model.Vec2) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { // 处理点选事件
		if temp := u.clickUnit(pos); temp != nil {
			u.setSelectUnit(temp)
			return true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) { // 清除选择
		u.clearSelectUnit()
		return true
	} // 框选处理
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { // 开始选择
		if u.selectRange == nil {
			u.selectRange = &model.Range{Pos1: pos, Pos2: pos}
			return true
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) { // 选择结束
		if u.selectRange != nil {
			units := make([]*unit.Unit, 0)
			for _, temp := range u.units {
				if temp.Mask.InRect(temp, u.selectRange) { // 建筑不能框选，是在实现决定的
					units = append(units, temp)
				}
			}
			u.setSelectUnit(units...)
			u.selectRange = nil
			return true
		}
	} else { // 选择中
		if u.selectRange != nil {
			u.selectRange.Pos2 = pos
			return true
		}
	}
	return false
}

func (u *unitManager) clickUnit(pos model.Vec2) *unit.Unit {
	for _, temp := range u.units {
		if temp.Mask.PointIn(temp, pos) {
			return temp
		}
	}
	return nil
}

func (u *unitManager) clearSelectUnit() {
	for _, temp := range u.selectUnits {
		temp.Select = false
	}
	u.selectUnits = make([]*unit.Unit, 0)
}

func (u *unitManager) SetSelectUnit(units any) {
	u.setSelectUnit(units.([]*unit.Unit)...)
}

func (u *unitManager) setSelectUnit(units ...*unit.Unit) {
	u.clearSelectUnit()
	for _, temp := range units {
		temp.Select = true
		u.selectUnits = append(u.selectUnits, temp)
	}
	unit.PlayRandSe(units, global.AudioSelectCmd)
}

func (u *unitManager) handleHover(pos model.Vec2) bool {
	if u.touchUnit != nil {
		u.touchUnit.Hover = false
	}
	for _, temp := range u.units {
		if temp.Mask.PointIn(temp, pos) {
			temp.Hover = true
			if u.touchUnit == temp { // 这时相当于无效处理，不阻塞后续操作
				return false
			}
			u.touchUnit = temp
			return true
		}
	}
	return false
}

func (u *unitManager) GetHeight(x float64, y float64) float64 {
	xIndex := utils.GetAdjustIndex(x)
	yIndex := utils.GetAdjustIndex(y)
	xIndex, yIndex = utils.ToIsometric(xIndex, yIndex) // 转换到 格子坐标系
	units := u.landUnits[xIndex][yIndex]
	height := float64(0)
	for z := u.GetZNum() - 1; z >= 0; z-- {
		temp := units[z][3] // 只有建筑单位增加地形高度
		if temp != nil && temp.Data.Type == global.UnitTypeConstruct {
			height = temp.Pos.Z + float64(temp.Data.ConstructExtra.Height*global.BlockSize)
			break
		}
	} // 取建筑高度与地形高度中最高的 防止有些建筑在地形下面
	return utils.Max(height, tool.TileManager.GetHeight(x, y))
}

func (u *unitManager) FilterUnits(filter func(unit any) bool) any {
	res := make([]*unit.Unit, 0)
	for _, temp := range u.units {
		if filter(temp) {
			res = append(res, temp)
		}
	}
	return res
}

func (u *unitManager) RegisterActionCreator(actionType model.UnitActionType, actionCreator interfaces.IUnitActionCreator) {
	u.actionCreators[actionType] = actionCreator
}

// 注意，围墙建立多个是调用多次 CreateUnit 生成的，在预览界面显示的grid会依次进行创建 调用该方法
func (u *unitManager) CreateUnit(name string, grid model.Grid) any {
	data := tool.DataManager.GetUnitData(name)
	if data == nil {
		return nil
	}
	if !u.CheckGrids(grid, data) {
		return nil
	}
	temp := u.BuildUnit(grid, data).(*unit.Unit)
	u.AddToGrid(grid, temp)
	u.AddUnit(temp)
	return temp
}

func (u *unitManager) Update() {
	for _, temp := range u.units {
		temp.Update()
	}
}

// CheckGrids 根据对象类型检查所有单元格
func (u *unitManager) CheckGrids(grid model.Grid, data *model.UnitData) bool {
	switch data.Type {
	case global.UnitTypeSoldier, global.UnitTypeVehicle:
		return u.CheckGrid(grid, data)
	case global.UnitTypeConstruct:
		wNum, hNum := data.ConstructExtra.WNum, data.ConstructExtra.HNum
		for x := 0; x < wNum; x++ {
			for y := 0; y < hNum; y++ {
				temp := grid.Add(model.Grid{X: x, Y: y})
				if !u.CheckGrid(temp, data) {
					return false
				}
			}
		}
		return true
	default:
		panic(fmt.Sprintf("can't CheckGrids of unitType of : %v", data.Type))
	}
}

func (u *unitManager) CheckGrid(grid model.Grid, data *model.UnitData) bool {
	land := data.StandExtra.Land
	if data.Type == global.UnitTypeConstruct {
		land = true
	}
	if land && !tool.TileManager.CheckTileType(grid, data.StandTile) { // 空军不需要检查格子
		return false
	}
	units, ok := u.GetUnits(grid, land).([]*unit.Unit)
	if !ok || units[3] != nil {
		return false
	}
	if data.Type == global.UnitTypeSoldier {
		for i := 0; i < 3; i++ { // 小兵任意一个为nil 就行
			if units[i] == nil {
				return true
			}
		}
		return false
	}
	for i := 0; i < 3; i++ { // 建筑，载具都必须全为nil
		if units[i] != nil {
			return false
		}
	}
	return true
}

var (
	soldierOffsets = []model.Vec3{{X: 24, Y: 24}, {X: 8, Y: 24}, {X: 24, Y: 8}}
)

func (u *unitManager) GetPos(grid model.Grid, data *model.UnitData) model.Vec3 {
	var pos model.Vec3
	switch data.Type {
	case global.UnitTypeConstruct:
		pos = utils.Grid2Pos(grid)
	case global.UnitTypeVehicle:
		pos = utils.Grid2Pos(grid).Add(model.Vec3{X: 16, Y: 16})
	case global.UnitTypeSoldier:
		// 如果没有空位置，就使用中间为默认位置
		pos = utils.Grid2Pos(grid).Add(model.Vec3{X: 16, Y: 16})
		units := u.GetUnits(grid, data.StandExtra.Land).([]*unit.Unit)
		for i := 0; i < 3; i++ {
			if units[i] == nil {
				pos = utils.Grid2Pos(grid).Add(soldierOffsets[i])
				break
			}
		}
	default:
		panic(fmt.Sprintf("GetPos can't handle unitType of : %v", data.Type))
	}
	land := data.StandExtra.Land
	if data.Type == global.UnitTypeConstruct {
		land = true
	}
	if land { // 地面单位可能需要跨越隧道等地形，不一定是获取最高点，获取当前格子位置的高度即可
		temp := tool.TileManager.GetTile(grid).(*tile.Tile)
		pos.Z = temp.GetHeight(pos.X, pos.Y)
	} else { // 只有天空中飞行的对象才考虑重算高度，空中单位一直处于其所在 X,Y的最高点
		pos.Z = u.GetHeight(pos.X, pos.Y)
	}
	return pos
}

func (u *unitManager) GetUnits(grid model.Grid, land bool) any {
	grid = utils.ToIsometricGrid(grid)
	if grid.X < 0 || grid.X >= u.GetXNum() || grid.Y < 0 || grid.Y >= u.GetYNum() {
		return nil
	}
	if !land {
		return u.airUnits[grid.X][grid.Y]
	}
	if grid.Z < 0 || grid.Z >= u.GetZNum() {
		return nil
	}
	return u.landUnits[grid.X][grid.Y][grid.Z]
}

func (u *unitManager) CreateMask(data *model.UnitData) interfaces.IUnitMask {
	switch data.MaskType {
	case global.UnitMaskTypeConstruct:
		return mask.NewConstructMask(data.ConstructExtra)
	case global.UnitMaskTypeSoldier:
		return mask.NewSoldierMask()
	case global.UnitMaskTypeVehicle:
		return mask.NewVehicleMask()
	default:
		panic(fmt.Sprintf("CreateMask can't handle MaskType of : %v", data.MaskType))
	}
}

func (u *unitManager) CreateAnim(data *model.UnitData) interfaces.IUnitAnim {
	switch data.AnimType {
	case global.UnitAnimTypeDir:
		return anim.NewDirAnim(data.AnimExtra)
	case global.UnitAnimTypeImage:
		return anim.NewImageAnim(data.AnimExtra)
	case global.UnitAnimTypeTurret:
		return anim.NewTurretAnim(data.AnimExtra)
	case global.UnitAnimTypeTurnTurret:
		return anim.NewTurnTurretAnim(data.AnimExtra)
	default:
		panic(fmt.Sprintf("CreateAnim can't handle AnimType of : %v", data.AnimType))
	}
}

func (u *unitManager) CreateShadow(data *model.UnitData) interfaces.IUnitShadow {
	switch data.ShadowType {
	case global.UnitShadowTypeNone:
		return shadow.NewNoneShadow()
	case global.UnitShadowTypeConstructBloodBar:
		return shadow.NewConstructBloodBarShadow()
	case global.UnitShadowTypeImage:
		return shadow.NewImageShadow(data)
	default:
		panic(fmt.Sprintf("CreateShadow can't handle ShadowType of : %v", data.ShadowType))
	}
}

func (u *unitManager) CreateBloodBar(data *model.UnitData) interfaces.IUnitBloodBar {
	switch data.BloodBarType {
	case global.UnitBloodBarTypeConstruct:
		return blood_bar.NewConstructBloodBar(data)
	case global.UnitBloodBarTypeSoldier:
		return blood_bar.NewSoldierBloodBar()
	case global.UnitBloodBarTypeVehicle:
		return blood_bar.NewVehicleBloodBar()
	default:
		panic(fmt.Sprintf("CreateBloodBar can't handle BloodBarType of : %v", data.BloodBarType))
	}
}

func (u *unitManager) AddUnit(temp any) {
	unit0 := temp.(*unit.Unit)
	u.units = append(u.units, unit0)
	tool.DrawManager.AddDraw(unit0)
	unit0.Execute(&model.UnitCmd{Type: global.UnitCmdTypeInit})
}

func (u *unitManager) RemoveFromGrid(grid model.Grid, temp any) {
	unit0 := temp.(*unit.Unit)
	data := unit0.Data
	switch data.Type {
	case global.UnitTypeConstruct:
		wNum, hNum := data.ConstructExtra.WNum, data.ConstructExtra.HNum
		for x := 0; x < wNum; x++ {
			for y := 0; y < hNum; y++ {
				tempGrid := grid.Add(model.Grid{X: x, Y: y})
				tempGrid = utils.ToIsometricGrid(tempGrid)
				if u.landUnits[tempGrid.X][tempGrid.Y][tempGrid.Z][3] == unit0 {
					u.landUnits[tempGrid.X][tempGrid.Y][tempGrid.Z][3] = nil
				}
			}
		}
	case global.UnitTypeVehicle:
		units := u.GetUnits(grid, data.StandExtra.Land).([]*unit.Unit)
		if units[3] == unit0 {
			units[3] = nil
		}
	case global.UnitTypeSoldier:
		units := u.GetUnits(grid, data.StandExtra.Land).([]*unit.Unit)
		for i := 0; i < 3; i++ {
			if units[i] == unit0 {
				units[i] = nil
				break
			}
		}
	default:
		panic(fmt.Sprintf("AddUnit can't add unit of unitType : %v", data.Type))
	}
}

func (u *unitManager) AddToGrid(grid model.Grid, temp any) {
	unit0 := temp.(*unit.Unit)
	data := unit0.Data
	grids := make([]model.Grid, 0)
	grids = append(grids, grid)
	switch data.Type {
	case global.UnitTypeConstruct:
		wNum, hNum := data.ConstructExtra.WNum, data.ConstructExtra.HNum
		for x := 0; x < wNum; x++ {
			for y := 0; y < hNum; y++ {
				tempGrid := grid.Add(model.Grid{X: x, Y: y})
				grids = append(grids, tempGrid) // 建筑需要额外添加格子
				tempGrid = utils.ToIsometricGrid(tempGrid)
				u.landUnits[tempGrid.X][tempGrid.Y][tempGrid.Z][3] = unit0
			}
		}
	case global.UnitTypeVehicle:
		units := u.GetUnits(grid, data.StandExtra.Land).([]*unit.Unit)
		units[3] = unit0
	case global.UnitTypeSoldier:
		units := u.GetUnits(grid, data.StandExtra.Land).([]*unit.Unit)
		for i := 0; i < 3; i++ {
			if units[i] == nil {
				units[i] = unit0
				break
			}
		}
	default:
		panic(fmt.Sprintf("AddUnit can't add unit of unitType : %v", data.Type))
	}
	tool.FogManager.ClearFogs(grids, data.Vision) // 清迷雾
}

func (u *unitManager) GetXNum() int {
	return len(u.landUnits)
}

func (u *unitManager) GetYNum() int {
	return len(u.landUnits[0])
}

func (u *unitManager) GetZNum() int {
	return len(u.landUnits[0][0])
}

func (u *unitManager) BuildUnit(grid model.Grid, data *model.UnitData) any {
	actionCreators := make(map[model.UnitActionType]interfaces.IUnitActionCreator, len(data.ActionCreators))
	for creatorType := range data.ActionCreators {
		if actionCreator, ok := u.actionCreators[creatorType]; ok {
			actionCreators[creatorType] = actionCreator
		} else {
			panic(fmt.Sprintf("CreateUnit don't has creatorType of : %v", creatorType))
		}
	}
	res := &unit.Unit{
		Pos:            u.GetPos(grid, data), // 飞行单位若是通过放置 默认是在地面上
		Hp:             data.Hp,
		Data:           data,
		Extra:          u.CreateExtra(grid, data),
		Mask:           u.CreateMask(data),
		Anim:           u.CreateAnim(data),
		Shadow:         u.CreateShadow(data),
		BloodBar:       u.CreateBloodBar(data),
		Actions:        collection.NewStack[interfaces.IUnitAction](),
		ActionCreators: actionCreators,
		//Select:         true,    // TODO TEST
		AttackDir: math.Pi, // 主要是各种防御建筑默认朝向 这个位置
		MoveDir:   math.Pi,
	}
	res.Idle = res.CreateAction(&model.UnitCmd{ // 这里构建的对象必须保证完整性
		Type: global.UnitCmdTypeCreateIdle,
	})
	utils.InvokeInit(res.Idle, res)
	return res
}

const (
	objTypeUnit = "unit"
)

func (u *unitManager) loadUnits() {
	root := utils.OpenXml(global.MapPath)
	layers := root.FindElements("objectgroup")
	for _, layer := range layers {
		objs := layer.FindElements("object")
		for _, obj := range objs {
			if utils.XmlStr(obj, "type", "") != objTypeUnit {
				continue
			}
			x := utils.XmlFloat(obj, "x", 0)
			y := utils.XmlFloat(obj, "y", 0)
			name := utils.XmlStr(obj, "name", "")
			x, y = utils.ToWorldPos(x, y) // 默认是等距坐标系下的位置
			z := u.GetHeight(x, y)
			// -1 防止单位放在非斜坡方块上
			u.CreateUnit(name, utils.Pos2Grid(model.NewVec3(x, y, z-1)))
		}
	}
}

func (u *unitManager) CreateExtra(grid model.Grid, data *model.UnitData) *unit.UnitExtra {
	res := &unit.UnitExtra{}
	if data.Type == global.UnitTypeConstruct && data.ConstructExtra.Tag == global.ConstructTagProduce {
		res.RallyPoint = grid.Add(data.ConstructExtra.Door)
	}
	return res
}
