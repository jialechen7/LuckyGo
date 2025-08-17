package constants

const (
	PublishTypeNormal               = 1
	PublishTypeTest                 = 2
	LotteryNotAnnounced             = 0
	LotteryHasAnnounced             = 1
	PrizeNotWon                     = 0
	PrizeHasWon                     = 1
	AnnounceTypeTimeLottery         = 1
	AnnounceTypePeopleLottery       = 2
	AnnounceTypeInstant             = 3
	InstantLotteryRedisKey          = "InstantLottery:"
	InstantLotteryPrizePoolRedisKey = "InstantLotteryPrizePool:"
	ProbabilityMax                  = 10000
)

var ProbabilityMap = map[int][2]int{
	1: [2]int{0, 100},
	2: [2]int{100, 500},
	3: [2]int{500, 2000},
	4: [2]int{2000, 5000},
}

const LotteryLuaScript = `
local prizePool = redis.call("HGETALL", KEYS[1])
local rand = tonumber(ARGV[1])

for i = 1, #prizePool, 2 do
    local prizeId = prizePool[i]
    local prizeData = cjson.decode(prizePool[i + 1])

    local left = tonumber(prizeData.left_probability)
    local right = tonumber(prizeData.right_probability)
    local stock = tonumber(prizeData.stock)

    if left and right and stock then
        if rand >= left and rand < right and stock > 0 then
            prizeData.stock = stock - 1
            redis.call("HSET", KEYS[1], prizeId, cjson.encode(prizeData))
            return prizeId
        end
    end
end

return 0
`
