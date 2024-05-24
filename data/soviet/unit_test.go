/*
@author: sk
@date: 2023/10/3
*/
package soviet

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestParseData(t *testing.T) {
	str := "soundOnAttackOrder:  ROOT:\\音频\\语音+音效\\vwasata.wav:2, ROOT:\\音频\\语音+音效\\vwasatb.wav:2, ROOT:\\音频\\语音+音效\\vwasatc.wav:2\nsoundOnMoveOrder: ROOT:\\音频\\语音+音效\\vwasmoa.wav:2, ROOT:\\音频\\语音+音效\\vwasmob.wav:2, ROOT:\\音频\\语音+音效\\vwasmoc.wav:2, ROOT:\\音频\\语音+音效\\vwasmod.wav:2\nsoundOnNewSelection: ROOT:\\音频\\语音+音效\\vwassea.wav:2, ROOT:\\音频\\语音+音效\\vwasseb.wav:2, ROOT:\\音频\\语音+音效\\vwassec.wav:2\neffectOnDeath: CUSTOM:水花, CUSTOM:水花2, CUSTOM:水花3, CUSTOM:水花4, CUSTOM:死亡爆炸, CUSTOM:沉没\n"

	lines := strings.Split(str, "\n")
	buff := bytes.Buffer{}
	buff.WriteString("Audios: map[string][]*model.AudioRate{\n")
	reg, _ := regexp.Compile("[:,]")
	for _, line := range lines {
		temp := reg.Split(line, -1)
		switch temp[0] {
		case "soundOnAttackOrder":
			buff.WriteString("global.AudioAttackCmd: {\n")
		case "soundOnMoveOrder":
			buff.WriteString("global.AudioMoveCmd: {\n")
		case "soundOnNewSelection":
			buff.WriteString("global.AudioSelectCmd: {\n")
		case "soundOnDeath":
			buff.WriteString("global.AudioDeath: {\n")
		case "playSoundAtUnit":
			buff.WriteString("global.AudioMove: {\n")
		default:
			continue
		}
		for i := 1; i < len(temp); i++ {
			if strings.HasSuffix(temp[i], ".wav") {
				buff.WriteString(fmt.Sprintf("{Path: \"res/audio/effect/%v\", Rate: 1},\n", temp[i][22:]))
			}
		}
		buff.WriteString("},\n")
	}
	buff.WriteString("},\n")
	fmt.Println(buff.String())
}

//Audios: map[string][]*model.AudioRate{
//global.AudioAttackCmd: {
//{Path: "iflaata.wav", Rate: 1},
//{Path: "iflaatb.wav", Rate: 1},
//{Path: "iflaatc.wav", Rate: 1},
//{Path: "iflaatd.wav", Rate: 1},
//},
//global.AudioMoveCmd: {
//{Path: "iflamoa.wav", Rate: 1},
//{Path: "iflamob.wav", Rate: 1},
//{Path: "iflamoc.wav", Rate: 1},
//{Path: "iflamod.wav", Rate: 1},
//},
//global.AudioSelectCmd: {
//{Path: "iflasea.wav", Rate: 1},
//{Path: "iflaseb.wav", Rate: 1},
//{Path: "iflasec.wav", Rate: 1},
//{Path: "iflased.wav", Rate: 1},
//},
//global.AudioDeath: {
//{Path: "ifladia.wav", Rate: 1},
//{Path: "ifladib.wav", Rate: 1},
//{Path: "ifladic.wav", Rate: 1},
//{Path: "ifladid.wav", Rate: 1},
//},
//},
