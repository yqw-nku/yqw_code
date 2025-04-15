package plugins

import (
	"active/common"
	"active/config"
	"fmt"
	"math"
	"time"
)

func minCostGifts(n int, expectValue int) (int, []int) {
	gifts := []common.GiftType{common.Gift1, common.Gift2, common.Gift3, common.Gift4, common.Gift5}
	maxValue := 0
	for _, gift := range gifts {
		maxValue += gift.Value() * n
	}

	// dp[i]表示达到价值i所需的最小消耗
	dp := make([]int, maxValue+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	// 记录每种礼物的数量
	counts := make([][]int, maxValue+1)
	for i := range counts {
		counts[i] = make([]int, len(gifts))
	}

	for _, gift := range gifts {
		value := gift.Value()
		cost := gift.Cost()
		for j := value; j <= maxValue; j++ {
			for k := 1; k <= n; k++ {
				if j >= k*value {
					if dp[j] > dp[j-k*value]+k*cost {
						dp[j] = dp[j-k*value] + k*cost
						copy(counts[j], counts[j-k*value])
						counts[j][gift] += k
					}
				}
			}
		}
	}

	// 找到最小消耗值
	minCost := math.MaxInt32
	var resultCounts []int
	for i := expectValue; i <= maxValue; i++ {
		if dp[i] < minCost {
			minCost = dp[i]
			resultCounts = counts[i]
		}
	}

	return minCost, resultCounts
}

type Gift struct {
	value int
	cost  int
}

type State struct {
	cost   int
	counts [5]int
}

func splitNumber(n int) []int {
	res := make([]int, 0)
	k := 1
	for n > 0 {
		if k > n {
			res = append(res, n)
			break
		} else {
			res = append(res, k)
			n -= k
			k *= 2
		}
	}
	return res
}

func minCost(n, expectValue int) {
	gifts := []Gift{
		{common.Gift1.Value(), common.Gift1.Cost()},
		{common.Gift2.Value(), common.Gift2.Cost()},
		{common.Gift3.Value(), common.Gift3.Cost()},
		{common.Gift1.Value(), common.Gift4.Cost()},
		{common.Gift5.Value(), common.Gift5.Cost()},
	}

	// Initialize dp array
	dp := make([]State, expectValue+1)
	for i := range dp {
		dp[i].cost = math.MaxInt32
	}
	dp[0].cost = 0

	for i, gift := range gifts {
		sList := splitNumber(n)
		for _, s := range sList {
			valueAdd := s * gift.value
			costAdd := s * gift.cost
			// Reverse traversal to handle 0-1 knapsack
			for v := expectValue; v >= 0; v-- {
				if dp[v].cost == math.MaxInt32 {
					continue
				}
				newV := v + valueAdd
				if newV > expectValue {
					newV = expectValue
				}
				if newCost := dp[v].cost + costAdd; newCost < dp[newV].cost {
					// Update state
					dp[newV].cost = newCost
					newCounts := dp[v].counts
					newCounts[i] += s
					dp[newV].counts = newCounts
				}
			}
		}
	}

	if dp[expectValue].cost == math.MaxInt32 {
		fmt.Println("无法达到期望价值")
	} else {
		fmt.Printf("最小消耗值cost: %d\n", dp[expectValue].cost)
		fmt.Println("礼物搭配：")
		for i, count := range dp[expectValue].counts {
			fmt.Printf("%s 出现 %d 次\n", common.GiftType(i).Name(), count)
		}
	}
}

// GetGift 用户指定期望达成的天数，推算出需要的礼包搭配
func GetGift(conf config.Config) {
	// 把升级所需的天数转为瓶子数
	fmt.Printf("%+v\n", conf)
	var needBottleSum = (conf.NeedDays * 24 * 60) / common.BottleEqualTime

	// 计算在期望天数内升级，免费可以获得瓶子
	var freeBottleSum = common.GetEveryDayFreeBottle() * conf.ExpectDays

	// 剩下的问题演变成：每种gift最多出现conf.ExpectDays次，找到一个消耗最小的搭配，使得gift总值等于needBottleSum-freeBottleSum
	var giftMaxCount = conf.ExpectDays                                   // 每种礼包出现的最多次数
	var leftBottleSum = needBottleSum - freeBottleSum - conf.ExtraBottle // 需要用礼包弥补的瓶子数
	fmt.Printf("每种礼包最多出现次数：%d，需要弥补的瓶子数：%d\n", giftMaxCount, leftBottleSum)
	time.Sleep(10 * time.Second)
	minCost(giftMaxCount, leftBottleSum)
}
