/*
@author: sk
@date: 2023/10/3
*/
package utils

import (
	"ra3/global"
	"ra3/global/tool"
)

func PlayDeputySe(name string) {
	tool.AudioManager.PlaySe(global.DeputyAudioPath + name)
}

func PlaySe(name string) {
	tool.AudioManager.PlaySe("res/audio/effect/" + name)
}
