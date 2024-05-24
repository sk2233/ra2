/*
@author: sk
@date: 2023/9/2
*/
package utils

import (
	"image"
	"image/color"
	interfaces "ra3/global/interface"
	"ra3/global/model"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	emptyImage     = ebiten.NewImage(1, 1)
	imageOption    = &colorm.DrawImageOptions{}
	triangleOption = &colorm.DrawTrianglesOptions{}
	colorModel     = colorm.ColorM{}
	rectVs         = make([]ebiten.Vertex, 4)
	imgRect        = &model.Rect{}
)

func init() {
	emptyImage.Set(0, 0, colornames.White)
	for i := 0; i < 4; i++ {
		rectVs[i] = ebiten.Vertex{ColorA: 1, ColorB: 1, ColorG: 1, ColorR: 1}
	}
}

/*
DrawImage
@param region 传入绘制的范围，减少绘制
*/
func DrawImage(screen, image *ebiten.Image, pos, anchor model.Vec2, region interfaces.IRect) {
	imgRect.Pos = pos
	imgRect.Anchor = anchor
	size := image.Bounds()
	imgRect.Size = model.NewVec2(float64(size.Dx()), float64(size.Dy()))
	if !RectCollision(imgRect, region) { // 优化绘图
		return
	}
	imageOption.GeoM.Reset()
	imageOption.GeoM.Translate(-anchor.X, -anchor.Y)
	imageOption.GeoM.Translate(pos.X, pos.Y)
	colorModel.Reset()
	colorModel.Scale(1, 1, 1, 1)
	colorm.DrawImage(screen, image, colorModel, imageOption)
}

/*
DrawLine
@param region 传入绘制的范围，减少绘制
*/
func DrawLine(screen *ebiten.Image, pos1, pos2 model.Vec2, w float64, clr color.Color, region interfaces.IRect) {
	min := region.GetMin()
	max := region.GetMax()
	if pos1.X < min.X && pos2.X < min.X {
		return
	}
	if pos1.Y < min.Y && pos2.Y < min.Y {
		return
	}
	if pos1.X > max.X && pos2.X > max.X {
		return
	}
	if pos1.Y > max.Y && pos2.Y > max.Y {
		return
	}
	vector.StrokeLine(screen, float32(pos1.X), float32(pos1.Y), float32(pos2.X), float32(pos2.Y), float32(w), clr, false)
}

/*
StrokeRect
@param region 传入绘制的范围，减少绘制
*/
func StrokeRect(screen *ebiten.Image, rect interfaces.IRect, w float64, clr color.Color, region interfaces.IRect) {
	if !RectCollision(rect, region) { // 优化绘图
		return
	}
	pos := rect.GetMin()
	size := rect.GetMax().Sub(pos)
	vector.StrokeRect(screen, float32(pos.X+w/2), float32(pos.Y+w/2), float32(size.X-w), float32(size.Y-w), float32(w), clr, false)
}

/*
FillRect
@param region 传入绘制的范围，减少绘制
*/
func FillRect(screen *ebiten.Image, rect interfaces.IRect, clr color.Color, region interfaces.IRect) {
	if !RectCollision(rect, region) { // 优化绘图
		return
	}
	pos := rect.GetMin()
	size := rect.GetMax().Sub(pos)
	x, y := float32(pos.X), float32(pos.Y)
	w, h := float32(size.X), float32(size.Y)
	rectVs[0].DstX, rectVs[0].DstY = x, y
	rectVs[1].DstX, rectVs[1].DstY = x+w, y
	rectVs[2].DstX, rectVs[2].DstY = x+w, y+h
	rectVs[3].DstX, rectVs[3].DstY = x, y+h
	FillFan(screen, rectVs, clr)
}

func FillFan(screen *ebiten.Image, vs []ebiten.Vertex, clr color.Color) { // 这里不再优化
	indices := make([]uint16, 0)
	for i := 0; i < len(vs)-2; i++ {
		indices = append(indices, 0, uint16(i+1), uint16(i+2))
	}
	colorModel.Reset()
	colorModel.ScaleWithColor(clr)
	colorm.DrawTriangles(screen, vs, indices, emptyImage, colorModel, triangleOption)
}

//func DrawText(screen *ebiten.Image, str string, pos, anchor model.Vec2, face font.Face, clr color.Color) {
//	bound := text.BoundString(face, str)                                               // 测量不会忽略换行
//	pos = pos.Sub(model.NewVec2(float64(bound.Min.X), float64(bound.Min.Y)))           // 找到 左上点
//	pos = pos.Sub(anchor.Mul(model.NewVec2(float64(bound.Dx()), float64(bound.Dy())))) // 移动到锚点
//	text.Draw(screen, str, face, int(pos.X), int(pos.Y), clr)
//}

func SplitImage(source *ebiten.Image, xNum, yNum int) []*ebiten.Image {
	res := make([]*ebiten.Image, 0, xNum*yNum)
	size := source.Bounds().Size()
	w := size.X / xNum
	h := size.Y / yNum
	for y := 0; y < yNum; y++ {
		for x := 0; x < xNum; x++ {
			temp := source.SubImage(image.Rect(x*w, y*h, (x+1)*w, (y+1)*h))
			res = append(res, temp.(*ebiten.Image))
		}
	}
	return res
}
