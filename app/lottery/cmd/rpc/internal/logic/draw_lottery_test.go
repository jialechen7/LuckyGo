package logic

import (
	"fmt"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func drawLottery(lotteryId int64, prizes []*model.Prize, participatorIds []int64) ([]Winner, error) {
	winners := make([]Winner, 0)

	// Fisher–Yates shuffle算法，实现纯随机性，参考链接：https://gaohaoyang.github.io/2016/10/16/shuffle-algorithm/
	for i := len(participatorIds) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		participatorIds[i], participatorIds[j] = participatorIds[j], participatorIds[i]
	}

	// 简易版本算法
	for _, prize := range prizes {
		// 从参与者中随机选取中奖者
		for i := 0; i < int(prize.Count); i++ {
			if len(participatorIds) == 0 {
				break
			}
			idx := rand.Intn(len(participatorIds))
			winner := Winner{
				LotteryId: lotteryId,
				UserId:    participatorIds[idx],
				PrizeId:   prize.Id,
			}
			winners = append(winners, winner)
		}
	}

	return winners, nil
}

func TestDrawLotteryFairness(t *testing.T) {
	// 初始化参与者和奖品
	totalParticipators := 1000
	participators := make([]int64, totalParticipators)
	for i := 0; i < totalParticipators; i++ {
		participators[i] = int64(i)
	}

	prizes := []*model.Prize{
		{
			Id:    1,
			Name:  "一等奖x1",
			Level: 1,
			Count: 1,
		},
		{
			Id:    2,
			Name:  "二等奖x2",
			Level: 2,
			Count: 2,
		},
		{
			Id:    3,
			Name:  "三等奖x5",
			Level: 3,
			Count: 5,
		},
	}

	// 定义测试参数
	iterations := 1000000           // 测试运行的次数
	winCount := make(map[int64]int) // 统计每个参与者获奖的次数

	for i := 0; i < iterations; i++ {
		winners, err := drawLottery(1, prizes, append([]int64{}, participators...))
		require.NoError(t, err)

		for _, winner := range winners {
			winCount[winner.UserId]++
		}
	}

	// 验证每位参与者的获奖分布是否接近均匀
	expectedProbability := float64(8) / float64(totalParticipators)
	tolerance := 0.01 // 公差范围，用于判定公平性
	for userId, count := range winCount {
		actualProbability := float64(count) / float64(iterations)
		if userId < int64(8) {
			fmt.Printf("User %d: Count=%d, Actual Probability=%.5f\n", userId, count, actualProbability)
		}
		require.InDelta(t, expectedProbability, actualProbability, tolerance,
			fmt.Sprintf("UserId %d has unexpected probability %.5f", userId, actualProbability))
	}

	fmt.Println("Lottery fairness test passed.")
}
