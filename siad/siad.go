package siad

import (
	"github.com/NebulousLabs/Sia/api"
	"github.com/NebulousLabs/Sia/build"
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
	srv     *Server
}

//Start starts the siad daemon with the consensus, gateway and transactionpool modules
func (s *Siad) Start() (err error) {

	// Create the server and start serving daemon routes immediately.
	log.Infoln("Loading siad...")
	s.srv, err = NewServer(s.APIAddr)
	if err != nil {
		return err
	}

	servErrs := make(chan error)
	go func() {
		servErrs <- s.srv.Serve()
	}()

	log.Infoln("Loading siad/gateway...")
	g, err := gateway.New(s.RPCAddr, true, modules.GatewayDir)
	if err != nil {
		return
	}

	log.Infoln("Loading siad/consensus...")
	cs, err := consensus.New(g, true, modules.ConsensusDir)
	if err != nil {
		return
	}

	log.Infoln("Loading siad/transaction pool...")
	tpool, err := transactionpool.New(cs, g, modules.TransactionPoolDir)
	if err != nil {
		return err
	}

	a := api.New("Sia-Agent", "", cs, nil, g, nil, nil, nil, tpool, nil)

	// connect the API to the server
	s.srv.Handle("/", a)

	select {
	case err = <-servErrs:
		build.Critical(err)
	default:
	}
	return
}

//Close stops the siad daemon
func (s *Siad) Close() (err error) {
	err = s.srv.Close()
	return
}
