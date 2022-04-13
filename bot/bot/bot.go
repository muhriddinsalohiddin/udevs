package bot

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muhriddinsalohiddin/udevs/bot/config"
)

func BotService(text string) (err error) {
	conf := config.Load()
	bot, err := tgbotapi.NewBotAPI(conf.BotTonken)
	if err != nil {
		log.Fatalf("Problem with connecting to bot: %v", err)
		return
	}

	id, err := strconv.ParseInt(conf.ChatId, 10, 64)
	if err != nil {
		log.Printf("Problem with converting chatId to int64: %v", err)
		return
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(id, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Problem with sending message: %v", err)
		return err
	}

	return
}
