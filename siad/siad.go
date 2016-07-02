package siad

import (
	"github.com/NebulousLabs/Sia/api"
	"github.com/NebulousLabs/Sia/crypto"
	"github.com/NebulousLabs/Sia/modules"
	"github.com/NebulousLabs/Sia/modules/consensus"
	"github.com/NebulousLabs/Sia/modules/gateway"
	"github.com/NebulousLabs/Sia/modules/transactionpool"
	log "github.com/Sirupsen/logrus"
)

//Siad is the reference to the siad modules
type Siad struct {
	RPCAddr string
	APIAddr string
}

//Start starts the siad daemon with the consensus, gateway and transactionpool modules
func (s *Siad) Start() (err error) {

	log.Infoln("Loading gateway...")
	g, err := gateway.New(s.RPCAddr, modules.GatewayDir)
	if err != nil {
		return
	}

	log.Infoln("Loading consensus...")
	cs, err := consensus.New(g, modules.ConsensusDir)
	if err != nil {
		return
	}

	log.Infoln("Loading transaction pool...")
	tpool, err := transactionpool.New(cs, g, modules.TransactionPoolDir)
	if err != nil {
		return err
	}

	srv, err := api.NewServer(s.APIAddr, "SIA-Agent", "", cs, nil, g, nil, nil, nil, tpool, nil)
	if err != nil {
		return
	}

	// connect to 3 random bootstrap nodes
	perm, err := crypto.Perm(len(modules.BootstrapPeers))
	if err != nil {
		return err
	}
	for _, i := range perm[:3] {
		go g.Connect(modules.BootstrapPeers[i])
	}

	// Start serving api requests.
	err = srv.Serve()

	return
}
