/*
@author: sk
@date: 2023/9/16
*/
package ui

import (
	"fmt"
	"image/color"
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"
	"strconv"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type IconItem struct {
	Data    *model.IconData
	support interfaces.IIconItemSupport
	img     *ebiten.Image // 60 * 48
	count   int           // 当前待生成数量，只有 队列类型有
	active  bool          // 激活，倒计时不受该字段影响
	timer   int           // 当前执行进度
	index   int           // 倒计时类专用，防止绘制的倒计时重叠
}

func NewIconItem(data *model.IconData, support interfaces.IIconItemSupport, index int) *IconItem {
	timer := 0
	if data.Type == global.IconTypeCountdown {
		timer = data.Time
	}
	return &IconItem{Data: data, support: support, img: tool.ImageLoader.LoadImage(data.Path), timer: timer, index: index}
}

func (i *IconItem) Update() {
	switch i.Data.Type {
	case global.IconTypeCountdown:
		if i.timer > 0 {
			i.timer--
			if i.timer == 600 { // 进行警报
				tool.UnitManager.BroadcastCmd(func(temp any) bool {
					unit0 := temp.(*unit.Unit)
					return unit0.Data.Type == global.UnitTypeConstruct &&
						unit0.Data.ConstructExtra.Effect == i.Data.Name
				}, &model.UnitCmd{
					Type: global.UnitCmdTypeIdle2Active,
				})
			}
		}
	case global.IconTypeSingle:
		if i.active && i.timer > 0 && tool.DataManager.ChangeMoney(-1) {
			i.timer--
			if i.timer == 0 {
				utils.PlayDeputySe("deputy048.wav")
			}
		}
	case global.IconTypeQueue:
		if !i.active {
			return
		}
		if i.timer > 0 {
			if tool.DataManager.ChangeMoney(-1) {
				i.timer--
			}
		} else {
			i.support.PollQueueData(i.Data.Group)
			i.count--
			utils.PlayDeputySe("deputy062.wav")
			tool.UnitManager.BroadcastCmd(func(temp any) bool {
				unit0 := temp.(*unit.Unit)
				return unit0.Data.Type == global.UnitTypeConstruct && unit0.Data.ConstructExtra.Tag == global.ConstructTagProduce
			}, &model.UnitCmd{
				Type:     global.UnitCmdTypeUnitReady,
				UnitName: i.Data.Name,
			})
		}
	}
}

var (
	iconAnchor  = model.Vec2{X: -2, Y: -2}
	iconRect    = &model.Rect{Anchor: iconAnchor, Size: model.NewVec2(60, 48)}
	readyConfig = &model.TextConfig{
		Anchor: model.NewVec2(0.5, 0),
		Color:  colornames.Yellow,
		Size:   1,
		Region: global.MenuRegion,
	}
	iconMaskClr   = color.RGBA{R: 0xff, G: 0xff, A: 0xff / 2}
	textMaskRect  = &model.Rect{Size: model.Vec2{X: 32, Y: 16}, Anchor: model.Vec2{X: -16}}
	iconMaskRect  = &model.Rect{Size: model.Vec2{X: 64, Y: 52}}
	iconRangeRect = &model.Rect{Size: model.Vec2{X: 64, Y: 52}}
	maskClr       = color.RGBA{A: 0xff / 2}
	countConfig   = &model.TextConfig{
		Anchor: model.NewVec2(1, 0),
		Color:  colornames.Yellow,
		Size:   1,
		Region: global.MenuRegion,
	}
	countdownPos    = model.NewVec2(global.WindowW-144, 0)
	countdownConfig = &model.TextConfig{
		Anchor: model.NewVec2(1, 1),
		Color:  colornames.Green,
		Size:   1,
		Region: global.WindowRegion,
	}
)

func (i *IconItem) DrawPos(screen *ebiten.Image, pos model.Vec2) { // 2的绘制边距
	utils.DrawImage(screen, i.img, pos, iconAnchor, global.MenuRegion)
	switch i.Data.Type {
	case global.IconTypeSingle:
		if i.support.GetSingleData(i.Data.Group) == nil { // 改组没有建造行为
			return
		}
		if !i.active { // 有建造行为 但不是自己
			iconMaskRect.Pos = pos
			utils.FillRect(screen, iconMaskRect, maskClr, global.MenuRegion)
			return
		} // 有建造行为自己在建造
	case global.IconTypeQueue:
		if i.count > 1 {
			tool.TextDraw.DrawText(screen, pos.Add(model.NewVec2(60, 4)), strconv.Itoa(i.count), countConfig)
		}
		if !i.active { // 自己没有激活
			return
		}
	case global.IconTypeCountdown: // 倒计时类特殊的绘制
		countdownPos.Y = global.WindowH - float64(i.index*16*2)
		if i.timer > 0 {
			tool.TextDraw.DrawText(screen, countdownPos, fmt.Sprintf("%s:%v", i.Data.Name, i.timer/60), countdownConfig)
		} else {
			tool.TextDraw.DrawText(screen, countdownPos, i.Data.Name+":Ready", countdownConfig)
		}
	} // 倒计时类不管
	if i.timer > 0 {
		height := 52 * float64(i.timer) / float64(i.Data.Time)
		iconRangeRect.Pos = pos
		iconRangeRect.Size.Y = height
		iconRangeRect.Anchor.Y = height - 52
		utils.FillRect(screen, iconRangeRect, iconMaskClr, global.MenuRegion)
	} else {
		textMaskRect.Pos = pos
		utils.FillRect(screen, textMaskRect, maskClr, global.MenuRegion)
		tool.TextDraw.DrawText(screen, pos.Add(model.NewVec2(32, 0)), "Ready", readyConfig)
	}
}

func (i *IconItem) HandleCursor(origin model.Vec2, pos model.Vec2) bool { // 暂时不处理右键暂停行为
	iconRect.Pos = origin
	if !utils.PointCollision(pos, iconRect) {
		return false
	}
	switch i.Data.Type {
	case global.IconTypeCountdown:
		if i.timer <= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			fmt.Println(i.Data)
		}
	case global.IconTypeQueue:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { // 添加任务到队列
			i.count++
			i.support.AddQueueData(i)
			utils.PlaySe("umenucl1.wav")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyC) { // 移除任务 不管队列中的数据，自然消除
			if i.count > 0 {
				i.count--
				utils.PlayDeputySe("deputy051.wav")
			}
		}
	case global.IconTypeSingle:
		current := i.support.GetSingleData(i.Data.Group)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { // 开始建造
			if current == nil {
				i.support.SetSingleData(i.Data.Group, i)
				utils.PlayDeputySe("deputy052.wav")
			} else if current == i && i.timer <= 0 {
				tool.CursorManager.StartConstruct(i)
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyC) { // 取消建造
			if current == i {
				i.Reset()
				utils.PlayDeputySe("deputy051.wav")
			}
		}
	}
	return true
}

func (i *IconItem) Init() {
	if i.Data.Type == global.IconTypeQueue && i.count <= 0 {
		i.support.PollQueueData(i.Data.Group)
		return
	}
	i.active = true
	i.timer = i.Data.Time
}

func (i *IconItem) Clear() {
	i.active = false
}

func (i *IconItem) Reset() {
	i.support.SetSingleData(i.Data.Group, nil)
	i.active = false
}
