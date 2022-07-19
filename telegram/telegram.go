package telegram

import (
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"

	"github.com/timmilesdw/slack-to-telegram/config"
)

type TelegramBot struct {
	bot                 *telebot.Bot
	mapChats            map[string]int64
	defaultChat         int64
	disableNotification bool
}

func (t *TelegramBot) SendMesage(text string, channel string,) error {
	chat, ok := t.mapChats[channel]
	if !ok {
		chat = t.defaultChat
	}

	txt := normalizeMessage(text)
	options := &telebot.SendOptions{
		DisableNotification:   t.disableNotification,
		DisableWebPagePreview: true,
	}

	message, err := t.bot.Send(
		telebot.ChatID(chat),
		txt,
		options,
	)
	if err != nil {
		return err
	}

	log.Debugf("sent message id: %d, chat: %d", message.ID, message.Chat.ID)

	return nil
}

func NewTelegram(cfg *config.Telegram) (*TelegramBot, error) {
	bot, err := telebot.NewBot(
		telebot.Settings{
			Token:     cfg.Token,
			ParseMode: telebot.ModeHTML,
			Offline:   true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		bot:                 bot,
		mapChats:            cfg.MapChats,
		defaultChat:         cfg.DefaultChat,
		disableNotification: cfg.DisableNotification,
	}, nil
}

func normalizeMessage(str string) string {
	re := regexp.MustCompile(`(?m)(<.*?>)`)
	processed := re.ReplaceAllStringFunc(str, repl)

	return processed
}

func repl(s string) string {
	s = strings.ReplaceAll(s, "<", "")
	s = strings.ReplaceAll(s, ">", "")

	new := strings.Split(s, "|")
	if len(new) < 2 {
		log.Errorf("replaceAllStringFunc error: invalid string '%s'", s)
		return ""
	}

	return `<a href="` + new[0] + `">` + new[1] + "</a>"
}