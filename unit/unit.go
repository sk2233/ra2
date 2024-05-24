/*
@author: sk
@date: 2023/9/9
*/
package unit

import (
	"ra3/global/collection"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type UnitExtra struct {
	Move       bool       // 是否移动中
	Deploy     bool       // 是否已经部署
	InHyper    bool       // 是否在超时空影响中，被超时空武器攻击的对象，超时空步兵移动后的后摇都是改状态
	RallyPoint model.Grid // 集结点，默认大门集合
	Passengers []*Unit    // 载具对象的乘客
	OtherUnit  *Unit
}

type Unit struct { // 士兵，载具，建筑都是Unit
	// 基础数据
	Pos                model.Vec3
	Hp                 int
	Hover              bool            // 鼠标悬停
	Select             bool            // 被选择
	AttackDir, MoveDir float64         // 攻击方向，移动方向 必须 影响绘制 0 ~ 2PI
	Data               *model.UnitData // 只读 单位元数据
	Extra              *UnitExtra      // 扩展数据 一般用于多Action间的数据交互
	// 碰撞相关
	Mask interfaces.IUnitMask
	// 绘制相关
	Anim     interfaces.IUnitAnim
	Shadow   interfaces.IUnitShadow
	BloodBar interfaces.IUnitBloodBar
	// 行动相关
	Idle           interfaces.IUnitAction                                 // 默认执行的行为，不会被弹出栈
	Actions        *collection.Stack[interfaces.IUnitAction]              // 当前行为栈
	ActionCreators map[model.UnitActionType]interfaces.IUnitActionCreator // 行为责任链
}

func (u *Unit) Draw(screen *ebiten.Image) {
	u.Shadow.DrawShadow(screen, u)
	u.Anim.DrawUnit(screen, u)
}

func (u *Unit) DrawAfter(screen *ebiten.Image) {
	if u.Actions.IsEmpty() { // 当前行为绘制
		utils.InvokeUnitDrawAfter(u.Idle, screen, u)
	} else {
		utils.InvokeUnitDrawAfter(u.Actions.Peek(), screen, u)
	}
	u.BloodBar.DrawBloodBar(screen, u)
}

func (u *Unit) GetOrder() float64 { // Unit 与其他对象重叠时先显示
	return u.Pos.Z*10000 + u.Pos.Y + u.Pos.X + 1 // X,Y 等权
}

func (u *Unit) DrawBefore(screen *ebiten.Image) {
}

func (u *Unit) Update() {
	u.Anim.UpdateAnim(u)
	if u.Actions.IsEmpty() {
		u.Idle.UpdateAction(u)
	} else {
		u.Actions.Peek().UpdateAction(u)
	}
}

func (u *Unit) PushAction(action interfaces.IUnitAction) {
	if u.Actions.IsEmpty() {
		utils.InvokeOnBelow(u.Idle, u)
	} else {
		utils.InvokeOnBelow(u.Actions.Peek(), u)
	}
	u.Actions.Push(action)
	utils.InvokeInit(action, u)
}

func (u *Unit) PopAction() {
	action := u.Actions.Pop()
	utils.InvokeDestroy(action, u)
	if u.Actions.IsEmpty() {
		utils.InvokeOnTop(u.Idle, u)
	} else {
		utils.InvokeOnTop(u.Actions.Peek(), u)
	}
}

// 立即动作可以立即执行，否则可以返回action入栈执行
func (u *Unit) Execute(cmd *model.UnitCmd) bool {
	if action := u.CreateAction(cmd); action != nil {
		u.PushAction(action)
		return true
	}
	return false
}

func (u *Unit) CreateAction(cmd *model.UnitCmd) interfaces.IUnitAction {
	for typ, creator := range u.ActionCreators {
		if res := creator.CreateAction(u, u.Data.ActionCreators[typ], cmd); res != nil {
			return res
		}
	}
	return nil
}

func (u *Unit) PopAllAction() {
	// 弹出所有，一般是为了执行新命令，并不是为了回到默认状态，只需要让他们进行销毁处理即可无需触发初始化
	for !u.Actions.IsEmpty() {
		action := u.Actions.Pop()
		utils.InvokeDestroy(action, u)
	}
}
