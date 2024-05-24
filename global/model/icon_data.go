/*
@author: sk
@date: 2023/9/16
*/
package model

type IconData struct {
	Name                string
	Conditions          []string
	Hash, ConditionHash int       // 自身的 hash值，与显示条件的hash值 | 集 使用位运算判断是否可以显示，无需手动赋值自动赋值
	Order               int       // 展示顺序，自动赋值，Hash值是经过压缩的，只有有当做条件的才会占用bit，其他的都在前面 1024个顺序排放，不能按Hash排序
	Type                IconType  // 如何运作的，最多一个?（例如建筑），多个使用栈排序？(例如步兵)，倒计时？(例如超级武器)
	Group               IconGroup // 如果类型限定最多一个，或多个使用栈排序，如何分组 也使用该字段显示在对应的 tab下 因此超级武器的计时也需要标明
	Path                string    // 图片位置
	Time                int       // 冷却总时间
}
