package telegram

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

type TelegramBot struct {
	bot                 *telebot.Bot
	logger              *logrus.Logger
	mapChats            map[string]int64
	defaultChat         int64
	disableNotification bool
}

func (t *TelegramBot) SendMesage(
	text string,
	channel string,
) error {
	chat, ok := t.mapChats[channel]
	if !ok {
		chat = t.defaultChat
	}
	txt := normalizeMessage(text)
	message, err := t.bot.Send(telebot.ChatID(chat), txt, &telebot.SendOptions{
		DisableNotification:   t.disableNotification,
		DisableWebPagePreview: true,
	})

	if err != nil {
		return err
	}

	t.logger.Debugf("sent message id: %v, chat: %v", message.ID, message.Chat.ID)

	return nil
}

func NewTelegram(
	token string,
	mapChats map[string]int64,
	defaultChat int64,
	disableNotification bool,
	logger *logrus.Logger,
) (*TelegramBot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:     token,
		ParseMode: telebot.ModeHTML,
		Offline:   true,
	})
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		bot:                 bot,
		logger:              logger,
		mapChats:            mapChats,
		defaultChat:         defaultChat,
		disableNotification: disableNotification,
	}, nil
}

func normalizeMessage(str string) string {
	re := regexp.MustCompile(`(?m)(<.*?>)`)

	processed := re.ReplaceAllStringFunc(str, func(s string) string {
		s = strings.ReplaceAll(s, "<", "")
		s = strings.ReplaceAll(s, ">", "")

		new := strings.Split(s, "|")

		return `<a href="` + new[0] + `">` + new[1] + "</a>"
	})

	return processed
}
