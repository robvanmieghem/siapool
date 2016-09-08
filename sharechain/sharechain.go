package sharechain

import (
	"github.com/NebulousLabs/Sia/persist"
	siasync "github.com/NebulousLabs/Sia/sync"
	"github.com/NebulousLabs/Sia/types"
	"github.com/NebulousLabs/demotemutex"
	"github.com/robvanmieghem/siapool/siad"
)

const (
	//ShareChainLength is the number of shares the chain can hold, given an average share every 10 seconds, it holds 4 days worth of shares
	ShareChainLength = 8640 * 4
	//ShareTime is the target time between two shares (like block time in a normal blockchain)
	ShareTime = 30
	//ShareDifficulty is the required difficiculty for a share to be accepted, it is currently fixed for a 1Gh/s miner to find 2 shares per day
	ShareDifficulty = 1 * 1000 * 1000 * 1000 * 3600 * 12
)

//ShareChain holds the previous shares of the pool
type ShareChain struct {

	//Siad is the handler towards the sia daemon
	Siad *siad.Siad

	// Utilities
	db         *persist.BoltDatabase
	log        *persist.Logger
	mu         demotemutex.DemoteMutex
	persistDir string

	// tg signals the Miner's goroutines to shut down and blocks until all
	// goroutines have exited before returning from Close().
	tg siasync.ThreadGroup
}

// New returns a new ShareChain.
// If there is an existing sharechain database present in the persist directory, it is loaded.
func New(siadaemon *siad.Siad, persistDir string) (sc *ShareChain, err error) {

	sc = &ShareChain{
		Siad: siadaemon,

		persistDir: persistDir,
	}

	// Initialize the persistence structures.
	err = sc.initPersist()

	return
}

//Share is a block with a lower difficulty target
type Share struct {
	BlockID   types.BlockID
	ParentID  types.BlockID
	Timestamp types.Timestamp
	Miner     string
}

//GetPPLNSSummary returns a mapping between miner addresses and the number of shares they found (within the ShareChainLength last number of shares)
func (sc *ShareChain) GetPPLNSSummary() (sharesummary map[string]int, err error) {
	//TODO
	return
}
