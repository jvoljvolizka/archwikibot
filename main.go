package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
)

func main() {
	key := os.Getenv("BOTKEY")
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.archlinux.org/feeds/news/")

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "type /news or cry."
		//case "bok":
		//	msg.Text = "kel başa şimşir tarak"
		case "news":
			msg.Text = feed.Items[0].Title + "\n\n" + feed.Items[0].Link +  feed.Items[1].Title + "\n\n" + feed.Items[1].Link +  feed.Items[2].Title + "\n\n" + feed.Items[2].Link
		default:
			log.Printf("dalyarak")
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
