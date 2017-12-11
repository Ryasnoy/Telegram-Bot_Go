package main
import "fmt"

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const token = "YOUR TOKEN"

func main() {

	// Connect to Bot via Token
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "NOT FOUND"))
			continue
		}

		if update.Message.NewChatMembers != nil {
			member := (*update.Message.NewChatMembers)[0]
			
			var name string
			if member.UserName != "" {
				name = member.UserName
			} else {
				name =  member.FirstName + " " + member.LastName
			}
			var greetings = fmt.Sprintf("Hello %s!", name)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, greetings)
			bot.Send(msg)
		}
		
	}
}