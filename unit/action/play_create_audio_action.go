/*
@author: sk
@date: 2023/10/4
*/
package action

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
)

func initPlayCreateAudioAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypePlayCreateAudio, &playCreateAudioActionCreator{})
}

type playCreateAudioActionCreator struct {
}

func (p *playCreateAudioActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypeInit {
		unit0 := temp.(*unit.Unit)
		unit.PlaySe(unit0, global.AudioCreate)
	}
	return nil
}
