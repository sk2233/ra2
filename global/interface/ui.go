/*
@author: sk
@date: 2023/9/16
*/
package interfaces

import "ra3/global/model"

type IClickListener interface {
	OnClick(tag model.ButtonTag)
}

type IStateButtonSupport interface {
	IClickListener
	GetState(tag model.ButtonTag) bool
}

type IIconItemSupport interface {
	AddQueueData(icon any)
	PollQueueData(group model.IconGroup) any
	SetSingleData(group model.IconGroup, icon any)
	GetSingleData(group model.IconGroup) any
}
