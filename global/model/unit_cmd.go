/*
@author: sk
@date: 2023/9/9
*/
package model

type UnitCmd struct { // 命令
	Type        UnitCmdType
	Abort       bool
	Anim        string
	Grid        Grid // 要移动到的位置 不同行为中语义不同
	NextGrid    Grid
	Angle       float64
	TurnMove    bool    // 移动转向，还是攻击转向，攻击转向有炮塔的不需要旋转车身
	Height      float64 // 爬升或下降的目标高度
	UnitName    string
	PowerEnough bool // 电量是否足够
	Unit        any  // 处理单位交互时对应的单位
}
