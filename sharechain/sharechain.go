package sharechain

import "github.com/robvanmieghem/siapool/siad"

const (
	//ShareChainLength is the number of shares the chain can hold, given an average share every 30 seconds, it holds 3 * 4 days worth of shares
	ShareChainLength = 8640 * 4
	//ShareTime is the target time between two shares (like block time in a normal blockchain)
	ShareTime = 30
	//ShareDifficulty is the required difficiculty for a share to be accepted, it is currently fixed for a 1Gh/s miner to find 2 shares per day
	// This means that payout is done over
	ShareDifficulty = 1 * 1000 * 1000 * 1000 * 3600 * 12
)

//ShareChain holds the previous shares of the pool
type ShareChain struct {

	//Siad is the handler towards the sia daemon
	Siad *siad.Siad
}
