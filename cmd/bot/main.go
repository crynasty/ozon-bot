package main

import (
	"log"
	"os"

	"github.com/crynasty/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	productService := product.NewService()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
		helpCommand(bot, update.Message)
		case "list":
		listCommand(bot, update.Message, productService)
		
		default:
		defaultBehaivor(bot, update.Message)
		}
	}
}


func defaultBehaivor(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	// msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		 "/help - help\n"+
		"/list - list products",
	)
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsgText := "Here all the products: \n\n"
	products := productService.List()
	for _, p := range products {
		outputMsgText += p.Title+"\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	bot.Send(msg)
}