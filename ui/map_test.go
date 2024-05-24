/*
@author: sk
@date: 2023/10/5
*/
package ui

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"image"
	"image/png"
	"os"
	"ra3/global"
	"ra3/utils"
	"strings"
	"testing"

	"golang.org/x/image/draw"

	"github.com/beevik/etree"
)

func TestGenMiniMap(t *testing.T) {
	tiles := loadTiles("/Users/bytedance/Documents/go/ra3/res/map/main.tmx")
	xNum := len(tiles)
	yNum := len(tiles[0])
	miniMap := image.NewRGBA(image.Rect(0, 0,
		xNum*global.MiniMapBlockSize-global.MiniMapBlockSize/2,
		(yNum-1)*global.MiniMapBlockSize/2/2))
	imgMap := loadImageMap("/Users/bytedance/Documents/go/ra3/res/image/minimap_grid.png")
	for x := 0; x < xNum; x++ {
		for y := 0; y < yNum; y++ {
			tile := getTile(tiles, x, y)
			if tile > 0 {
				xPos := x * global.MiniMapBlockSize
				if y%2 == 0 {
					xPos -= global.MiniMapBlockSize / 2
				}
				yPos := (y - 1) * global.MiniMapBlockSize / 2 / 2
				draw.Draw(miniMap, image.Rect(xPos, yPos, xPos+global.MiniMapBlockSize, yPos+global.MiniMapBlockSize/2),
					imgMap[tile], imgMap[tile].Bounds().Min, draw.Over)
			}
		}
	}
	file, err := os.Create("/Users/bytedance/Documents/go/ra3/res/image/minimap.png")
	utils.HandleErr(err)
	err = png.Encode(file, miniMap)
	utils.HandleErr(err)
}

func loadImageMap(path string) map[uint32]image.Image {
	file := utils.OpenFile(path)
	defer file.Close()
	temp, err := png.Decode(file)
	utils.HandleErr(err)
	img := temp.(*image.NRGBA)
	img0 := img.SubImage(image.Rect(0, 0, global.MiniMapBlockSize, global.MiniMapBlockSize/2))
	img1 := img.SubImage(image.Rect(global.MiniMapBlockSize, 0, global.MiniMapBlockSize*2, global.MiniMapBlockSize/2))
	img2 := img.SubImage(image.Rect(0, global.MiniMapBlockSize/2, global.MiniMapBlockSize, global.MiniMapBlockSize/2*2))
	img3 := img.SubImage(image.Rect(global.MiniMapBlockSize, global.MiniMapBlockSize/2, global.MiniMapBlockSize*2, global.MiniMapBlockSize/2*2))
	return map[uint32]image.Image{
		1: img0,
		2: img0,
		3: img0,
		4: img1,
		5: img2,
		6: img3,
		7: img3,
		8: img3,
	}
}

func getTile(tiles [][][]uint32, x int, y int) uint32 {
	for z := len(tiles[x][y]) - 1; z >= 0; z-- {
		if tiles[x][y][z] > 0 {
			return tiles[x][y][z]
		}
	}
	return 0
}

func loadTiles(path string) [][][]uint32 {
	root := utils.OpenXml(path)
	layers := root.FindElements("layer")
	temp := make([][][]uint32, 0)
	for _, layer := range layers {
		temp = append(temp, loadLayer(layer))
	}
	// 纠正为 x,y,z
	zNum, yNum, xNum := len(temp), len(temp[0]), len(temp[0][0])
	tiles := make([][][]uint32, xNum)
	for x := 0; x < xNum; x++ {
		tiles[x] = make([][]uint32, yNum)
		for y := 0; y < yNum; y++ {
			tiles[x][y] = make([]uint32, zNum)
			for z := 0; z < zNum; z++ {
				tiles[x][y][z] = temp[z][y][x]
			}
		}
	}
	return tiles
}

func loadLayer(layer *etree.Element) [][]uint32 {
	wNum := int(utils.XmlInt(layer, "width", 0))
	hNum := int(utils.XmlInt(layer, "height", 0))
	res := make([][]uint32, hNum)
	data := layer.FindElement("data").Text()
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(data))
	utils.HandleErr(err)
	reader, err := zlib.NewReader(bytes.NewReader(bs))
	utils.HandleErr(err)
	for y := 0; y < hNum; y++ {
		res[y] = make([]uint32, wNum)
		for x := 0; x < wNum; x++ {
			res[y][x] = utils.ReadUin32(reader)
		}
	}
	return res
}
