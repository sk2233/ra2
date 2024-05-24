/*
@author: sk
@date: 2023/10/2
*/
package manager

import (
	"ra3/global/tool"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func init() {
	tool.AudioManager = &audioManager{ctx: audio.NewContext(44100), ses: make(map[string]*audio.Player)}
}

type audioManager struct {
	ctx *audio.Context
	bgm *audio.Player
	ses map[string]*audio.Player
}

func (a *audioManager) PlayBgm(path string, second int64) {
	if a.bgm != nil {
		a.bgm.Close()
	}
	file := utils.OpenFile(path) // bgm比较大流式加载，不要立即关闭文件，只播放一次，没有关闭文件还好
	reader, err := vorbis.DecodeWithSampleRate(44100, file)
	utils.HandleErr(err)
	a.bgm, err = a.ctx.NewPlayer(audio.NewInfiniteLoop(reader, second*4*44100))
	utils.HandleErr(err)
	a.bgm.SetVolume(0.6)
	a.bgm.Play()
}

func (a *audioManager) PlaySe(path string) {
	if _, ok := a.ses[path]; !ok {
		file := utils.OpenFile(path)
		reader, err := wav.DecodeWithSampleRate(44100, file)
		utils.HandleErr(err)
		player, err := a.ctx.NewPlayer(reader)
		utils.HandleErr(err)
		a.ses[path] = player
	}
	player := a.ses[path]
	err := player.Rewind()
	utils.HandleErr(err)
	player.Play()
}
