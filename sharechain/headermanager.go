package sharechain

import "github.com/NebulousLabs/Sia/types"

// HeaderForWork returns a header that is ready for nonce grinding.
func (sc *ShareChain) HeaderForWork(payoutaddress string) (blockheader types.BlockHeader, target types.Target, err error) {
	if err := sc.tg.Add(); err != nil {
		return types.BlockHeader{}, types.Target{}, err
	}
	defer sc.tg.Done()

	sc.mu.Lock()
	defer sc.mu.Unlock()

	return
}
