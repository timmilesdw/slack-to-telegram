package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/timmilesdw/slack-to-telegram/config"
	"github.com/timmilesdw/slack-to-telegram/server"
	"github.com/timmilesdw/slack-to-telegram/telegram"
	"github.com/timmilesdw/slack-to-telegram/template"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"sigs.k8s.io/yaml"
)

func run(c *cli.Context) error {
	configData, err := os.ReadFile(c.String("config"))
	if err != nil {
		return fmt.Errorf("can't parse config file %s", err)
	}

	var conf config.Config

	if err := yaml.Unmarshal(configData, &conf); err != nil {
		return errors.Wrap(err, "yaml.Unmarshal")
	}

	log.Debug("Loaded configuration")

	l, err := log.ParseLevel(conf.LogLevel)
	if err != nil {
		log.Errorf("logrus.ParseLevel: %v")
	} else {
		log.SetLevel(l)
	}
	

	template, err := template.NewTemplate(conf.Template)
	if err != nil {
		return errors.Wrap(err, "template.NewTemplate")
	}

	telegram, err := telegram.NewTelegram(conf.Telegram)
	if err != nil {
		return errors.Wrap(err, "telegram.NewTelegram")
	}
	log.Debug("Created telegram client")

	http := server.NewHttpServer(conf.Server.Address, template, telegram)
	log.Debug("Created http server")

	http.SetupRoutes()

	return http.Listen()
}
