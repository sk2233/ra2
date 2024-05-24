/*
@author: sk
@date: 2023/9/2
*/
package global

import (
	"ra3/global/model"
)

//===========sys===========

var (
	WindowRegion = &model.Rect{Size: model.NewVec2(WindowW-144, WindowH)} // 只看展示的区域
	MenuRegion   = &model.Rect{Pos: model.NewVec2(WindowW-144, 0), Size: model.NewVec2(144, WindowH)}
)

const (
	WindowW = 1280
	WindowH = 720
)

const (
	MapPath         = "res/map/main.tmx"
	DeputyAudioPath = "res/audio/soviet/" // 语音提示阵营，相关音效使用这里的，方便统一修改
)

const (
	BlockSize        = 32
	MiniMapBlockSize = 8 // 小地图一个菱形格子的横向大小
)

const (
	FlyHeight = BlockSize * 4 // 飞行单位距离地面的高度
)

var (
	InvalidVec3 = model.Vec3{X: -1, Y: -1, Z: -1}
)

var (
	ZeroVec2 = model.Vec2{}
)

//============tile==============

const (
	TileTypeLand model.TileType = 1 << iota
	TileTypeSlope1
	TileTypeSlope2
	TileTypeBlock
	TileTypeWater
	TileTypeGround = TileTypeLand | TileTypeSlope1 | TileTypeSlope2
)

//=============unit================

const (
	UnitTypeSoldier model.UnitType = iota + 1
	UnitTypeVehicle
	UnitTypeConstruct
)

const (
	UnitAnimTypeDir   model.UnitAnimType = iota + 1 // 多朝向的动画 也可以是 1朝向  支持前置，后置绘制
	UnitAnimTypeImage                               // 只有单张图片的绘制
	// 有炮塔的动画，切换动画只是切换炮塔的动画(也可以没有炮塔，没有炮塔就不绘制炮塔)，身子不变 一般载具使用
	UnitAnimTypeTurret
	// 测试使用的
	UnitAnimTypeTurnTurret
)

const (
	UnitMaskTypeConstruct model.UnitMaskType = iota + 1
	UnitMaskTypeSoldier
	UnitMaskTypeVehicle
)

const (
	UnitShadowTypeNone              model.UnitShadowType = iota + 1
	UnitShadowTypeConstructBloodBar                      // 建筑选择后线框的补充绘制 严格来说不属于阴影
	UnitShadowTypeImage                                  // 使用图片绘制阴影
)

const (
	UnitBloodBarTypeConstruct model.UnitBloodBarType = iota + 1
	UnitBloodBarTypeSoldier
	UnitBloodBarTypeVehicle
)

const (
	UnitActionTypeNone model.UnitActionType = iota + 1
	UnitActionTypeConstruct
	UnitActionTypeIdleActive      // 部分对象激活与不激活时的动画切换处理
	UnitActionTypeWall            // 围墙有一次调整自己与四周方块的机会，初始化后进行或被其他方块调用后进行
	UnitActionTypeAnimOnce        // 接受指定指令，运行指定动画，运行完毕再切换回idle状态
	UnitActionTypeLandMove        // 载具,步兵等在地面(陆地水面)移动
	UnitActionTypeAirMove         // 载具,步兵等在天空移动
	UnitActionTypeHyperMove       // 超时空移动，超时空兵专用
	UnitActionTypeAirLift         // 飞行单位在空中的升降操作，来保持飞行高度
	UnitActionTypeTurnNoWait      // 立即的转向，无需时间，一般用于小兵
	UnitActionTypeTurnSpeed       // 有一定的转向速度，常用于载具，防御设施
	UnitActionTypeSwitchDeploy    // 切换部署状态，对于部署指令的部署反部署处理，一般用于常规部署的处理(例如美国大兵)
	UnitActionTypeNoneIdle        // idle行为，什么都不干的idle行为
	UnitActionTypeRallyPointIdle  // 生产建筑的idle主要负责选择后集结点的绘制
	UnitActionTypeFlyDisperseIdle // 飞行单位在执行完任意操作进入idle状态后需要进行分散操作
	UnitActionTypeMoveAnim        // 包含水面动画移动过程中的动画切换，可以使用参数指定是否包含水面移动动画
	// Deprecated  移动发声太吵了，尽量不要用
	UnitActionTypeMoveStart          // 开始移动的回调，尽量不要做太大的动作，可以播放音乐等
	UnitActionTypeWaterWave          // 水上载具，移动时产生水面波纹
	UnitActionTypeFlyHomeIdle        // 一进入idle就试图回家，若是没有家了回直接坠毁
	UnitActionTypeConstructComplete  // 监听建造完成的事件
	UnitActionTypeUnitReady          // 单位建造完成，广播给生产建筑进行生产
	UnitActionTypePowerChange        // 电量发生变化切换单位状态的行为
	UnitActionTypeRallyPoint         // 设置集合点行为
	UnitActionTypeEnterUnit          // 单位进入另一个单位的行为，挂在前一个单位上
	UnitActionTypeUnitEnterOut       // 单位自主控制，被进入，或放出对象
	UnitActionTypePassengerCountIdle // 绘制乘客情况的idle
	UnitActionTypePlayCreateAudio    // 播放出场音效
	UnitActionTypeDeployConstruct    // 载具变建筑
	UnitActionTypeTurnTurren         // 测试使用的
)

const (
	AnimIdle         = "AnimIdle"
	AnimDeActive     = "AnimDeActive" // 断电等停止运行的动画
	AnimWaterIdle    = "AnimWaterIdle"
	AnimDeployIdle   = "AnimDeployIdle"
	AnimMove         = "AnimMove"
	AnimWaterMove    = "AnimWaterMove"
	AnimAttack       = "AnimAttack"
	AnimWaterAttack  = "AnimAttack"
	AnimDeployAttack = "AnimDeployAttack"
	AnimConstruct    = "AnimConstruct"
	AnimActive       = "AnimActive" // 超级武器，建筑是否进人都使用这个
	AnimIdle2Active  = "AnimIdle2Active"
	AnimActive2Idle  = "AnimActive2Idle"
	AnimHorizontal   = "AnimHorizontal"
	AnimVertical     = "AnimVertical"
	AnimBuild        = "AnimBuild"        // 建造完成动画 一般只有船厂，基地有
	AnimTurretIdle   = "AnimTurretIdle"   // turretAnim 的默认炮塔动画
	AnimTurretAttack = "AnimTurretAttack" // 主炮塔攻击
	AnimTurretGun    = "AnimTurretGun"    // 多功能专用，枪类炮塔
	AnimTurretRepair = "AnimTurretRepair" // 多功能专用，修复炮塔
	AnimTurretOther  = "AnimTurretOther"  // 多功能专用，其他类型炮塔
	AnimOpenDoor     = "AnimOpenDoor"     // 打开大门，一般生产建筑生产陆地单位使用
	AnimOpenDome     = "AnimOpenDome"     // 打开穹顶，一般生产建筑生产空军使用
)

const (
	AudioDeath         = "AudioDeath"
	AudioRepair        = "AudioRepair"
	AudioPrepareAttack = "AudioPrepareAttack" // 例如 磁暴线圈 先响再攻击
	AudioAttack        = "AudioAttack"
	AudioPowerAttack   = "AudioPowerAttack" // 增幅攻击，例如磁暴线圈带充电宝
	AudioCreate        = "AudioCreate"      // 部分单位出现全图的提醒
	AudioActive        = "AudioActive"      // 超级武器就绪提醒
	AudioAttackCmd     = "AudioAttackCmd"
	AudioMoveCmd       = "AudioMoveCmd"
	AudioSelectCmd     = "AudioSelectCmd"
	// Deprecated  移动发声太吵了，尽量不要用
	AudioMove = "AudioMove" // 移动音效
)

const (
	UnitBloodBarSizeSmall model.UnitBloodBarSize = iota + 1
	UnitBloodBarSizeMiddle
	UnitBloodBarSizeBig
)

const (
	UnitCmdTypeInit model.UnitCmdType = iota + 1
	UnitCmdTypeIdle2Active
	UnitCmdTypeActive2Idle
	UnitCmdTypeTidy
	UnitCmdTypeAnimOnce
	UnitCmdTypeMove // 只有步兵与载具使用，移动到指定点
	// UnitCmdTypeMove派生事件
	UnitCmdTypeMoveStart  // 开始移动
	UnitCmdTypeMoveToNext // 移动到下一个格子事件
	UnitCmdTypeMovePause  // 移动暂停事件
	UnitCmdTypeMoveEnd    // 移动结束事件
	UnitCmdTypeTurn       // 转向
	UnitCmdTypeLift       // 飞行单位，爬升或下降
	UnitCmdTypeRallyPoint // 设置集结点，注意基地类型的设置集结点是移动基地
	UnitCmdTypeDeploy
	UnitCmdTypeCreateIdle        // 创建默认的idle行为，由系统在构建对象时完成，不要在其他地方调用
	UnitCmdTypeConstructComplete // 建筑建造完成
	UnitCmdTypeUnitReady         // 单位训练完成
	UnitCmdTypePowerChange       // 电量发生变化
	UnitCmdTypeHandleUnit        // 处理选择单位后，再点击其他单位时，单位之间的交互事件
	UnitCmdTypeEnterUnit         // 一个单位进入另一个单位
)

//==============UI=================

const (
	ButtonTagConstruct model.ButtonTag = iota + 1 // 前4个必须与IconGroup对齐
	ButtonTagDefense
	ButtonTagSoldier
	ButtonTagVehicle
	ButtonTagRepair
	ButtonTagSell
	ButtonTagDown
	ButtonTagUp
)

const (
	CursorTypeNormal model.CursorType = iota + 1
	CursorTypeRepair
	CursorTypeSell
	CursorTypeConstruct
)

const (
	IconTypeSingle model.IconType = iota + 1
	IconTypeQueue
	IconTypeCountdown
)

const (
	IconGroupConstruct model.IconGroup = iota + 1 // 必须与 ButtonTag 对齐
	IconGroupDefense
	IconGroupSoldier
	IconGroupVehicle
)

const (
	// 部分功能使用可以中断的广播实现
	ConstructTagBase    model.ConstructTag = iota + 1 // 基地标签
	ConstructTagExtend                                // 可扩展建筑标签，主要用于围墙
	ConstructTagProduce                               // 生产类建筑
	ConstructTagEnter                                 // 可以进兵的建筑
)

const (
	VehicleTagEnter model.VehicleTag = iota + 1 // 可以进兵的载具
)

//===============Effect==============

const (
	EffectTypeTimer model.EffectType = iota + 1 // 到指定时间消失
	EffectTypeAnim                              // 动画结束消失
)
