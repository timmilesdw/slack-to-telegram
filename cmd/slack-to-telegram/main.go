package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetReportCaller(true)
	// Stderr might cause some problems with docker
	// logger.SetOutput(os.Stderr)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)

	app := cli.App{
		Name:  "slack-to-telegram",
		Usage: "Send slack incoming webhook notifications to telegram",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				EnvVars:  []string{"STT_CONFIG"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "log_level",
				Aliases:  []string{"log"},
				EnvVars:  []string{"LOG_LEVEL"},
				Required: false,
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
