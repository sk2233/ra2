/*
@author: sk
@date: 2023/9/2
*/
package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"ra3/global"
	"testing"
)

func TestZlib(t *testing.T) {
	str := "eJzt1jEKgEAMRNGgsOv9T2yl4BaWZoT3TjDwi6QKAAAAAAAAAAAAAAAAcozuAdxG6ZFklCYpRumRZO2hSS898miSZW1x9M6hnj325i1okUiP773danfje/6nHLP0SLOVJkn0yHM1IcPsHgAAAAAAAAAAAAAAAAAAAAAAAADAb52uuQE8"
	bs, _ := base64.StdEncoding.DecodeString(str)
	r := bytes.NewReader(bs)
	reader, err := zlib.NewReader(r)
	fmt.Println(err)
	bs = make([]byte, 16)
	reader.Read(bs)
	for i := 0; i < 4; i++ {
		fmt.Println(binary.LittleEndian.Uint32(bs[i*4 : (i+1)*4]))
	}
}

func TestDiv(t *testing.T) {
	num := (-1 - global.BlockSize) / global.BlockSize
	fmt.Println(int(num))
}

func TestTint(t *testing.T) {
	// shader可以使用梯度实现不使用if
	// 每次阵营加载颜色时都需要为其颜色维护一套图片，或使用shader，建议引入shader，方便后续其他效果实现
	// 只有绿色通道有颜色时把他占255的比例缩放到其他纯色并应用
	//bound := a.src.Bounds()
	//temp := image.NewRGBA(bound)
	//size := bound.Max
	//tar := a.clrs[a.index]
	//for x := 0; x < size.X; x++ {
	//	for y := 0; y < size.Y; y++ {
	//		clr := color.RGBAModel.Convert(a.src.At(x, y)).(color.RGBA)
	//		if clr.R == 0 && clr.G > 0 && clr.B == 0 {
	//			g := float64(clr.G)
	//			clr.R = uint8(float64(tar.R) * g / 0xFF)
	//			clr.G = uint8(float64(tar.R) * g / 0xFF)
	//			clr.B = uint8(float64(tar.B) * g / 0xFF)
	//		}
	//		temp.Set(x, y, clr)
	//	}
	//}
	//a.img = ebiten.NewImageFromImage(temp)
}
