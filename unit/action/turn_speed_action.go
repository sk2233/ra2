/*
@author: sk
@date: 2023/9/24
*/
package action

import (
	"math"
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"
)

func initTurnSpeedAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeTurnSpeed, &turnSpeedActionCreator{})
}

type turnSpeedActionCreator struct {
}

func (o *turnSpeedActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeTurn {
		return nil
	}
	unit0 := temp.(*unit.Unit)
	dir := utils.If(cmd.TurnMove, unit0.MoveDir, unit0.AttackDir)
	speed := param.TurnSpeed * utils.AngleDir(dir, cmd.Angle)
	return &turnSpeedAction{currentAngle: dir, targetAngle: cmd.Angle, bindAngle: param.BindAngle,
		turnMove: cmd.TurnMove, speed: speed}
}

type turnSpeedAction struct {
	currentAngle, targetAngle float64
	speed                     float64 // 有方向的速度
	bindAngle                 bool
	turnMove                  bool
}

func (o *turnSpeedAction) UpdateAction(temp any) {
	unit0 := temp.(*unit.Unit)
	if math.Abs(o.targetAngle-o.currentAngle) < math.Abs(o.speed) {
		o.setAngle(unit0, o.targetAngle)
		unit0.PopAction()
	} else {
		o.currentAngle = utils.TrimAngle(o.currentAngle + o.speed)
		o.setAngle(unit0, o.currentAngle)
	}
}

func (o *turnSpeedAction) setAngle(unit0 *unit.Unit, angle float64) {
	if o.bindAngle { // 绑定情况下，两个转向是一致的
		unit0.MoveDir = angle
		unit0.AttackDir = angle
	} else if o.turnMove { // 非绑定情况下移动转向会带动炮塔
		offset := utils.TrimAngle(angle - unit0.MoveDir)
		unit0.MoveDir = angle
		unit0.AttackDir = utils.TrimAngle(unit0.AttackDir + offset)
	} else { // 非绑定情况下旋转炮塔仅旋转炮塔
		unit0.AttackDir = angle
	}
}
