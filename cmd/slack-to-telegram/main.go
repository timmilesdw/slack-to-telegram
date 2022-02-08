package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)

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
		},
		Action: run(logger),
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
