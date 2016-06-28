package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"github.com/robvanmieghem/siapool/api"
	"github.com/robvanmieghem/siapool/siadclient"
)

func main() {

	app := cli.NewApp()
	app.Name = "Siapool node"
	app.Version = "0.1-Dev"

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	var debugLogging bool
	var bindAddress, siadAddress string
	var poolFee int

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug logging",
			Destination: &debugLogging,
		},
		cli.StringFlag{
			Name:        "bind, b",
			Usage:       "Bind address",
			Value:       ":9985",
			Destination: &bindAddress,
		},
		cli.StringFlag{
			Name:        "siad, s",
			Usage:       "SIA daemon address",
			Value:       "localhost:9980",
			Destination: &siadAddress,
		},
		cli.IntFlag{
			Name:        "fee, f",
			Usage:       "Pool fee, in 0.01%",
			Value:       200,
			Destination: &poolFee,
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
		dc := siadclient.NewSiadClient(siadAddress)
		poolapi := api.PoolAPI{Fee: poolFee, SiadClient: dc}
		r := mux.NewRouter()
		r.Path("/fee").Methods("GET").Handler(http.HandlerFunc(poolapi.FeeHandler))
		http.ListenAndServe(bindAddress, r)
	}

	app.Run(os.Args)
}
