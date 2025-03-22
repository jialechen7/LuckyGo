package constants

const (
	PublishTypeNormal         = 1
	PublishTypeTest           = 2
	LotteryNotAnnounced       = 0
	LotteryHasAnnounced       = 1
	PrizeNotWon               = 0
	PrizeHasWon               = 1
	AnnounceTypeTimeLottery   = 1
	AnnounceTypePeopleLottery = 2
	AnnounceTypeInstant       = 3
	InstantLotteryRedisKey    = "InstantLottery:"
	ProbabilityMax            = 10000
)

var ProbabilityMap = map[int][2]int{
	1: [2]int{0, 100},
	2: [2]int{100, 500},
	3: [2]int{500, 2000},
	4: [2]int{2000, 5000},
}
