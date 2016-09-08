package sharechain

import (
	"github.com/NebulousLabs/Sia/crypto"
	"github.com/NebulousLabs/Sia/types"
)

//GenerateMinerPayouts creates a list of pplns payouts
// An extra payment of 0 to a random address is added to have a unique merkleroot for every generated set of payouts
func (sc *ShareChain) GenerateMinerPayouts(minerAddress types.UnlockHash, subsidy types.Currency) (payouts []types.SiacoinOutput, err error) {

	//Create a random address to have a unique merkleroot for every generated set of payouts
	randomBytes, err := crypto.RandBytes(crypto.HashSize)
	if err != nil {
		return
	}
	randomAddress := types.UnlockHash{}
	copy(randomAddress[:], randomBytes)

	payouts = []types.SiacoinOutput{
		types.SiacoinOutput{
			Value:      subsidy,
			UnlockHash: minerAddress,
		},
		types.SiacoinOutput{
			Value:      types.ZeroCurrency,
			UnlockHash: randomAddress,
		},
	}
	return
}
