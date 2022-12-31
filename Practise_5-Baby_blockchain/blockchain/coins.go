package blockchain

var startEmit uint64 = 7000000000

func GetCoinsForEmition() uint64 {
	return uint64(startEmit / ((uint64(len(Blockchain)+1)/2105280)/2 + 1))
}
