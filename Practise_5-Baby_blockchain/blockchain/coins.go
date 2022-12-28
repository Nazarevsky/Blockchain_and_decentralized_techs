package blockchain

var startEmit uint64 = 7000000000

func GetCoinsForEmition(chainDepth int) uint64 {
	return uint64(startEmit / ((uint64(chainDepth) / 2105280) + 1) / 2)
}
