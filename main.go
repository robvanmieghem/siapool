package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Siapool node"
	app.Version = "0.1-Dev"

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	var debugLogging bool
	var bindAddress, siadAddress string

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
	}

	app.Before = func(c *cli.Context) error {
		if debugLogging {
			log.SetLevel(log.DebugLevel)
			log.Debug("Debug logging enabled")
			log.Debug(app.Name, "-", app.Version)
		}
		return nil
	}

	app.Action = func(c *cli.Context) {

	}

	app.Run(os.Args)
}
