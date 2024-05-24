/*
@author: sk
@date: 2023/9/16
*/
package other

import (
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	tool.TextDraw = textDrawImpl
}

var (
	textDrawImpl = &textDraw{
		option: &ebiten.DrawImageOptions{},
	}
	destRect = &model.Rect{}
)

func initTextDraw() {
	img := tool.ImageLoader.LoadImage("res/image/font.png")
	imgs := utils.SplitImage(img, 32, 8)
	textDrawImpl.imgs = imgs
}

type textDraw struct { // 每个文本 6 * 16 大小
	imgs   []*ebiten.Image
	option *ebiten.DrawImageOptions
}

func (t *textDraw) DrawText(screen *ebiten.Image, pos model.Vec2, str string, config *model.TextConfig) {
	size := t.GetTextSize(str).Scale(config.Size)
	pos = pos.Sub(size.Mul(config.Anchor))
	x, y := pos.X, pos.Y
	t.option.ColorScale.Reset()
	t.option.ColorScale.ScaleWithColor(config.Color)
	destRect.Size = size
	//ebitenutil.DrawRect(screen, x, y, size.X, size.Y, colornames.Red)
	for _, c := range str {
		if c == '\n' {
			x = pos.X
			y += 16 * config.Size
			continue
		}
		destRect.Pos = model.NewVec2(x, y)
		if utils.RectCollision(destRect, config.Region) {
			t.option.GeoM.Reset()
			t.option.GeoM.Scale(config.Size, config.Size)
			t.option.GeoM.Translate(x, y)
			if int(c) > len(t.imgs) {
				c = 0
			}
			screen.DrawImage(t.imgs[c], t.option)
		}
		x += 6 * config.Size
	}
}

func (t *textDraw) GetTextSize(str string) model.Vec2 {
	items := strings.Split(str, "\n")
	wNum := 0
	for _, item := range items {
		wNum = utils.Max(wNum, len(item))
	}
	return model.NewVec2(float64(wNum*6), float64(len(items)*16))
}
