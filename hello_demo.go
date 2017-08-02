package main 

import (
	tb "github.com/TelegramApi/bot"
	"encoding/json"
)

func main() {
	bot := tb.Create("316622512:AAE-o4aqdFBSI9bdclU90PEYQOFGuudDXrw")
	bot.Listen()

	var outputMessage string

	for update := range bot.Updates{
	switch update.Message.Text {
		case "/start":
			outputMessage = "I am your new Bot.\n\n"
		case "Hi, Bot!":
			outputMessage = "Hello, " + update.Message.From.FirstName
		default:
			outputMessage = ""
		}

		var chat tb.User
		json.Unmarshal(update.Message.Chat, &chat)

		var keyboard = tb.ReplyKeyboardMarkup{Keyboard: [][]string{[]string{"Hi, Bot!"}}}
		bot.SendMessage(chat.Id, outputMessage, false, 0, keyboard)
	}
}