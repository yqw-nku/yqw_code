package common

type GiftType int

const (
	Gift1 GiftType = iota // 自选6元礼包
	Gift2                 // 自选12元礼包
	Gift3                 // 弹窗30礼包
	Gift4                 // 弹窗68礼包
	Gift5                 // 弹窗128礼包
)

func (g GiftType) Name() string {
	switch g {
	case Gift1:
		return "自选6元礼包"
	case Gift2:
		return "自选12元礼包"
	case Gift3:
		return "弹窗30礼包"
	case Gift4:
		return "弹窗68礼包"
	case Gift5:
		return "弹窗128礼包"
	default:
		return "未知礼包"
	}
}

func (g GiftType) Cost() int {
	switch g {
	case Gift1:
		return 6
	case Gift2:
		return 12
	case Gift3:
		return 30
	case Gift4:
		return 68
	case Gift5:
		return 128
	default:
		return 0
	}
}

func (g GiftType) Value() int {
	switch g {
	case Gift1:
		return 35
	case Gift2:
		return 35
	case Gift3:
		return 90
	case Gift4:
		return 200
	case Gift5:
		return 400
	default:
		return 0
	}
}

const (
	BottleEqualTime = 5 // 1个瓶子等价5min
)

func GetEveryDayFreeBottle() int {
	var sum = 0
	// 首选是每天的自然流速
	sum += (24 * 60) / BottleEqualTime
	// 每日广告
	sum += (5 * 30) / BottleEqualTime
	// 妖盟
	sum += 25
	// 坊市
	sum += 100
	return sum
}
