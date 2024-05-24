/*
@author: sk
@date: 2023/9/16
*/
package manager

import (
	"ra3/global"
	"ra3/global/collection"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/ui"
	"ra3/unit"
	"ra3/utils"
	"sort"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	uiManagerImpl = &uiManager{money: ui.NewMoneyUI(),
		power:        ui.NewPowerUI(),
		currentGroup: global.IconGroupConstruct,
		icons: map[model.IconGroup][]*ui.IconItem{
			global.IconGroupConstruct: make([]*ui.IconItem, 0),
			global.IconGroupDefense:   make([]*ui.IconItem, 0),
			global.IconGroupSoldier:   make([]*ui.IconItem, 0),
			global.IconGroupVehicle:   make([]*ui.IconItem, 0),
		},
		iconQueue: map[model.IconGroup]*collection.Queue[*ui.IconItem]{
			global.IconGroupSoldier: collection.NewQueue[*ui.IconItem](),
			global.IconGroupVehicle: collection.NewQueue[*ui.IconItem](),
		},
		iconSingle: make(map[model.IconGroup]any),
	}
)

func init() {
	tool.UiManager = uiManagerImpl
}

func initUiManager() {
	pos := model.NewVec2(global.WindowW-144, 16+88)
	size := model.NewVec2(72, 32)
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagRepair, "Repair", uiManagerImpl))
	pos.X += 72
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagSell, "Sell", uiManagerImpl))
	pos = model.NewVec2(global.WindowW-144, 16+88+32)
	size = model.NewVec2(36, 32)
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagConstruct, "Const", uiManagerImpl))
	pos.X += 36
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagDefense, "Defen", uiManagerImpl))
	pos.X += 36
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagSoldier, "Soldi", uiManagerImpl))
	pos.X += 36
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagVehicle, "Vehic", uiManagerImpl))
	pos = model.NewVec2(global.WindowW-144, global.WindowH-32)
	size = model.NewVec2(72, 32)
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagDown, "Down", uiManagerImpl))
	pos.X += 72
	uiManagerImpl.stateBtns = append(uiManagerImpl.stateBtns, ui.NewStateButton(
		pos, size, global.ButtonTagUp, "Up", uiManagerImpl))
	uiManagerImpl.UpdateIconItems()
	uiManagerImpl.map0 = ui.NewMap()
}

//func initFont() {
//	bs, err := ioutil.ReadFile("res/font/VonwaonBitmap-12px.ttf")
//	utils.HandleErr(err)
//	raw, err := opentype.Parse(bs)
//	utils.HandleErr(err)
//	global.Font16 = createFont(raw, 32)
//}

//func createFont(raw *opentype.Font, size float64) font.Face {
//	size = (size - 0.5) * 72 / 64
//	face, err := opentype.NewFace(raw, &opentype.FaceOptions{DPI: 12, Size: size})
//	utils.HandleErr(err)
//	//ebitenutil.DebugPrintAt()
//	return face
//}

type uiManager struct {
	money     *ui.MoneyUI //  144 * 16
	map0      *ui.Map     //  144 * 88
	stateBtns []*ui.StateButton
	//  UI     64 * 52  内边距 2
	icons    map[model.IconGroup][]*ui.IconItem // 一页显示 10个
	iconHash int                                // 存储现在显示的 item hash  | 集
	// 最大翻页数目  	max := (len(temps)+1)/2-10  注意翻页时限制
	iconOffset   int // icons 多的时候可以使用上下翻页 每次按2个翻页 从0开始，切换tab 复位到0
	iconQueue    map[model.IconGroup]*collection.Queue[*ui.IconItem]
	iconSingle   map[model.IconGroup]any
	power        *ui.PowerUI //  16 * 520
	currentGroup model.IconGroup
}

// 每次建筑变动，要主动更新小地图
func (u *uiManager) UpdateMiniMap() {
	u.map0.UpdateMiniMap()
}

func (u *uiManager) AddQueueData(temp any) {
	icon := temp.(*ui.IconItem)
	queue := u.iconQueue[icon.Data.Group]
	queue.Offer(icon)
	if queue.Count() == 1 {
		icon.Init() // 第一个只能在这里初始化
	}
}

func (u *uiManager) PollQueueData(group model.IconGroup) any {
	queue := u.iconQueue[group]
	temp := queue.Poll()
	temp.Clear() // 清理脏数据
	if !queue.IsEmpty() {
		queue.Peek().Init()
	}
	return temp
}

func (u *uiManager) SetSingleData(group model.IconGroup, temp any) {
	last := u.iconSingle[group]
	if last != nil {
		last.(*ui.IconItem).Clear()
	}
	u.iconSingle[group] = temp
	if temp != nil {
		temp.(*ui.IconItem).Init()
	}
}

func (u *uiManager) GetSingleData(group model.IconGroup) any {
	return u.iconSingle[group]
}

func (u *uiManager) HandleCursor(pos model.Vec2) bool { // 处理点击事件
	if u.map0.HandleCursor(pos) {
		return true
	}
	for _, btn := range u.stateBtns {
		if btn.HandleCursor(pos) {
			return true
		}
	}
	origin := model.NewVec2(global.WindowW-128, 16+88+64)
	icons := u.getActiveIcons()
	for i, icon := range icons {
		if icon.HandleCursor(origin, pos) {
			return true
		}
		if i%2 == 0 {
			origin.X += 64
		} else {
			origin.X = global.WindowW - 128
			origin.Y += 52
		}
	}
	return false
}

func (u *uiManager) GetState(tag model.ButtonTag) bool {
	switch tag {
	case global.ButtonTagRepair:
		return tool.CursorManager.GetType() == global.CursorTypeRepair
	case global.ButtonTagSell:
		return tool.CursorManager.GetType() == global.CursorTypeSell
	case global.ButtonTagConstruct, global.ButtonTagDefense, global.ButtonTagSoldier, global.ButtonTagVehicle:
		return u.currentGroup == model.IconGroup(tag)
	default:
		return false
	}
}

func (u *uiManager) OnClick(tag model.ButtonTag) {
	switch tag {
	case global.ButtonTagRepair:
		// TODO Repair
	case global.ButtonTagSell:
		// TODO Sell
	case global.ButtonTagConstruct, global.ButtonTagDefense, global.ButtonTagSoldier, global.ButtonTagVehicle:
		u.currentGroup = model.IconGroup(tag)
	case global.ButtonTagUp:
		u.iconOffset = utils.Max(0, u.iconOffset-1)
	case global.ButtonTagDown:
		max := (len(u.icons[u.currentGroup])+1)/2 - 10
		if max <= 0 { // 一点点都不能翻
			return
		}
		u.iconOffset = utils.Min(u.iconOffset+1, max)
	}
}

func (u *uiManager) Update() {
	for _, icons := range u.icons { // 每个都要
		for _, icon := range icons {
			icon.Update()
		}
	}
}

var (
	iconBgRect = &model.Rect{
		Pos:  model.NewVec2(global.WindowW-128, 16+88+64),
		Size: model.NewVec2(128, 520),
	}
)

func (u *uiManager) DrawAfter(screen *ebiten.Image) {
	u.money.Draw(screen)
	u.map0.Draw(screen)
	for _, btn := range u.stateBtns {
		btn.Draw(screen)
	}
	u.power.Draw(screen)
	utils.FillRect(screen, iconBgRect, colornames.Black, global.MenuRegion) // 绘制item的背景
	icons := u.getActiveIcons()
	pos := iconBgRect.Pos
	for i, icon := range icons {
		icon.DrawPos(screen, pos)
		if i%2 == 0 {
			pos.X += 64
		} else {
			pos.X = global.WindowW - 128
			pos.Y += 52
		}
	}
}

// UpdateIconItems 当有建筑变化时进行调用刷新，item列表
func (u *uiManager) UpdateIconItems() {
	temps := tool.UnitManager.FilterUnits(func(temp any) bool {
		return temp.(*unit.Unit).Data.Type == global.UnitTypeConstruct
	}).([]*unit.Unit)
	currentHash := 0
	for _, temp := range temps {
		data := tool.DataManager.GetIconData(temp.Data.Name)
		currentHash |= data.Hash
	}
	if u.iconHash == currentHash { // 没有变化
		return
	}
	datas := tool.DataManager.FilterIconDatas(func(data *model.IconData) bool {
		return data.ConditionHash&u.iconHash != data.ConditionHash &&
			data.ConditionHash&currentHash == data.ConditionHash // 原来没有，但是当前可以有
	}) // 暂时只处理新增
	index := 0 // 倒计时类的索引防止绘制文本重叠
	for _, data := range datas {
		u.icons[data.Group] = append(u.icons[data.Group], ui.NewIconItem(data, u, index))
		if data.Type == global.IconTypeCountdown {
			index++
		}
	}
	for _, icons := range u.icons {
		sort.Slice(icons, func(i, j int) bool {
			return icons[i].Data.Order < icons[j].Data.Order
		})
	}
	u.iconHash = currentHash
	u.iconOffset = 0 // 取消翻页
}

func (u *uiManager) getActiveIcons() []*ui.IconItem {
	temps := u.icons[u.currentGroup]
	temps = temps[u.iconOffset*2:] // 信任翻页索引
	if len(temps) > 20 {           // 太多的截断
		temps = temps[:20]
	}
	return temps
}
