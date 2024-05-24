/*
@author: sk
@date: 2023/9/9
*/
package model

type UnitData struct {
	Name   string
	Hp     int
	Vision float64 // 视野范围
	// 单位类型与其特殊参数
	Type           UnitType
	ConstructExtra *UnitConstructExtra
	VehicleExtra   *UnitVehicleExtra
	// 可以站立的位置 等参数
	StandTile  TileType // 可以站立的 tile
	StandExtra *UnitStandExtra
	// 碰撞选择类型
	MaskType UnitMaskType
	// 主动画相关
	AnimType  UnitAnimType
	AnimExtra *UnitAnimExtra
	// 阴影相关   一般只有飞行单位需要，动画中一般包含阴影
	ShadowType  UnitShadowType
	ShadowExtra *UnitShadowExtra
	// 血条相关  一般需要配置悬浮，选择血条的显示
	BloodBarType  UnitBloodBarType
	BloodBarExtra *UnitBloodBarExtra
	// action相关与其对应参数
	ActionCreators map[UnitActionType]*UnitActionCreatorExtra
	// 音效相关内容
	Audios map[string][]*AudioRate
}

type AudioRate struct {
	Path string // 路径
	Rate int    // 播放概率
}

type UnitConstructExtra struct {
	WNum, HNum              int // 占地面积，一般只有建筑需要
	Height                  int // 高度(格子为单位) 只有建筑需要
	CostPower, ProductPower int64
	Tag                     ConstructTag // 特殊单位也可以打标记 使用 UnitTag
	Door, Dome              Grid         // 作为生产建筑的出生点，是相对于建筑的基点,Door是门陆地部队使用，Dome是穹顶空军使用
	Capacity                int          // 当为进人建筑时，进人容量
	Effect                  string       // 建筑对应的效果(暂时只有超级武器使用)
}

type UnitVehicleExtra struct {
	Tag      VehicleTag
	Capacity int // 当为进人载具时，进人容量
}

type UnitStandExtra struct {
	Land bool // 是否在地面上 占据地面空间，否则占用天空空间，一般只有步兵，载具有
}

type UnitAnimExtra struct {
	// 方向动画图片的属性
	Dir1       map[string]*ImageExtra // image_path -> 图片属性
	AfterDir1  map[string]*ImageExtra // 在Dir1绘制后可选的绘制装饰
	BeforeDir1 map[string]*ImageExtra // 在Dir1绘制前可选的绘制装饰
	Dir8       map[string]*ImageExtra // image_path -> 图片属性
	Dir32      map[string]*ImageExtra // image_path -> 图片属性
	// 简单图片的锚点
	Images map[string]string // 图片名称 -> 对应图片
	Anchor Vec2              // 锚点
	// 炮塔相关动画属性
	TurretDir32 map[string]*ImageExtra // image_path -> 图片属性
	// 测试使用的内容
	TurnSpeed float64
}

type ImageExtra struct {
	Count     int  // 横向数目
	Anchor    Vec2 // 锚点
	AnimSpeed float64
	Frames    map[string]*FrameRange // 图片区域
}

type FrameRange struct {
	Start int
	End   int
}

type UnitShadowExtra struct {
}

type UnitBloodBarExtra struct {
	BloodBarSize UnitBloodBarSize // 暂时只有建筑的血条区分大小
}

type UnitActionCreatorExtra struct {
	// Move相关
	MoveSpeed float64
	// 转向相关
	TurnSpeed float64
	BindAngle bool // 是否绑定攻击方向与移动方向，没有炮塔的一般需要绑定
	// 是否具备水面上相关动画
	WaterAnim bool
	// flyDisperseIdleAction 分散到对应位置后是否需要降落
	NeedLand     bool
	ProduceUnits []string
	// 绘制入驻人员，被入住单位是建筑吗？因为没有对应的图片，只能绘制为一样的
	IsConstruct bool
	// 部分单位部署后产生的基地
	ConstructName string
	// 部分单位变成另一个单位，哪个单位的名称
	UnitName string
}
