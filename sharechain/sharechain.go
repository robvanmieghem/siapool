package sharechain

const (
	//ShareChainLength is the number of shares the chain can hold, given an average share every 30 seconds, it holds 3 days worth of shares
	ShareChainLength = 8640
	//ShareTime is the target time between two shares (like block time in a normal blockchain)
	ShareTime = 30
)

//ShareChain holds the previous shares of the pool
type ShareChain struct {
}
