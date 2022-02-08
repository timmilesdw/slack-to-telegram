package main

import (
	"fmt"
	"os"

	"github.com/timmilesdw/slack-to-telegram/config"
	"github.com/timmilesdw/slack-to-telegram/server"
	"github.com/timmilesdw/slack-to-telegram/telegram"
	"github.com/timmilesdw/slack-to-telegram/template"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"sigs.k8s.io/yaml"
)

func run(logger *logrus.Logger) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		configData, err := os.ReadFile(c.String("config"))
		if err != nil {
			return fmt.Errorf("can't parse config file %s", err)
		}

		var conf config.Config

		if err := yaml.Unmarshal(configData, &conf); err != nil {
			return err
		}

		template, err := template.NewTemplate(conf.Template)
		if err != nil {
			return err
		}

		telegram, err := telegram.NewTelegram(
			conf.Telegram.Token,
			conf.Telegram.MapChats,
			conf.Telegram.DefaultChat,
			conf.Telegram.DisableNotification,
			logger,
		)
		if err != nil {
			return err
		}

		http := server.NewHttpServer(conf.Server.Address, template, telegram, logger)

		http.SetupRoutes()

		return http.Listen()
	}
}
