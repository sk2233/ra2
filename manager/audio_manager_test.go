/*
@author: sk
@date: 2023/10/2
*/
package manager

import (
	"ra3/global/tool"
	"testing"
	"time"
)

func TestPlayerBgm(t *testing.T) {
	for i := 0; i < 100; i++ {
		tool.AudioManager.PlayBgm("/Users/bytedance/Documents/go/ra3/res/audio/bg.ogg", 3*60+44)
		time.Sleep(time.Second * 5)
	}
}

func TestPlayerEffect(t *testing.T) {
	for i := 0; i < 100; i++ {
		tool.AudioManager.PlaySe("/Users/bytedance/Documents/go/ra3/res/audio/effect/abirj01b.wav")
		time.Sleep(time.Second * 3)
	}
}
