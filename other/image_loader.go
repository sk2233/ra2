/*
@author: sk
@date: 2023/9/2
*/
package other

import (
	"ra3/global/tool"
	"ra3/utils"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	tool.ImageLoader = &imageLoader{images: make(map[string]*ebiten.Image)}
}

type imageLoader struct {
	images map[string]*ebiten.Image
}

func (i *imageLoader) LoadImage(path string) *ebiten.Image {
	index := strings.Index(path, "image")
	path = "res/" + path[index:]
	if !utils.HasKey(i.images, path) {
		img, _, err := ebitenutil.NewImageFromFile(path)
		utils.HandleErr(err)
		i.images[path] = img
	}
	return i.images[path]
}
