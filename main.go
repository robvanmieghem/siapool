package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"github.com/robvanmieghem/siapool/api"
	"github.com/robvanmieghem/siapool/sharechain"
	"github.com/robvanmieghem/siapool/siad"
)

func main() {

	app := cli.NewApp()
	app.Name = "Siapool node"
	app.Version = "0.1-Dev"

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	var debugLogging bool
	var bindAddress, apiAddr, rpcAddr string
	var poolFee int

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug logging",
			Destination: &debugLogging,
		},
		cli.StringFlag{
			Name:        "bind, b",
			Usage:       "Pool bind address",
			Value:       ":9985",
			Destination: &bindAddress,
		},
		cli.IntFlag{
			Name:        "fee, f",
			Usage:       "Pool fee, in 0.01%",
			Value:       200,
			Destination: &poolFee,
		},
		cli.StringFlag{
			Name:  "api-addr",
			Value: "localhost:9980", Usage: "which host:port the API server listens on",
			Destination: &apiAddr,
		},
		cli.StringFlag{
			Name:        "rpc-addr",
			Value:       ":9981",
			Usage:       "which port the gateway listens on",
			Destination: &rpcAddr,
		},
	}

	app.Before = func(c *cli.Context) error {
		log.Infoln(app.Name, "-", app.Version)
		if debugLogging {
			log.SetLevel(log.DebugLevel)
			log.Debugln("Debug logging enabled")
		}
		return nil
	}

	app.Action = func(c *cli.Context) {
		// Print a startup message.
		log.Infoln("Loading...")

		dc := &siad.Siad{RPCAddr: rpcAddr, APIAddr: apiAddr}
		go dc.Start()

		sc, err := sharechain.New(dc, "p2pool")
		if err != nil {
			log.Fatal("Error initializing sharechain:", err)
		}
		poolapi := api.PoolAPI{Fee: poolFee, ShareChain: sc}
		r := mux.NewRouter()
		r.Path("/fee").Methods("GET").Handler(http.HandlerFunc(poolapi.FeeHandler))
		r.Path("/{payoutaddress}/miner/header").Methods("GET").Handler(http.HandlerFunc(poolapi.GetWorkHandler))
		r.Path("/{payoutaddress}/miner/header").Methods("POST").Handler(http.HandlerFunc(poolapi.SubmitHeaderHandler))

		log.Infoln("Finished loading")

		http.ListenAndServe(bindAddress, r)
	}

	app.Run(os.Args)
}
