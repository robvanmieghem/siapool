package sharechain

import "github.com/NebulousLabs/Sia/types"

// HeaderForWork returns a header that is ready for nonce grinding.
func (sc *ShareChain) HeaderForWork(payoutaddress string) (blockheader types.BlockHeader, target types.Target, err error) {
	//TODO
	return
}
