package tgbot

import (
	"net/http"
	"strings"
)

var (
	bot    *TelegramBot
	apiKey = ""
)

func handleUpdate(update *TelegramUpdateResult) {
	for i := 0; i < len(update.Results); i++ {
		var text string

		command := strings.ToLower(extractCommand(update.Results[i].Message))

		switch command {
		case "ping":
			text = "Pong!"
		case "":
			text = "You need to specify a command with /"
		default:
			text = "Unrecognized command: /" + command
		}

		bot.SendMessage(&TelegramOutgoingMessage{
			ChatID: update.Results[i].Message.Chat.Id,
			Text:   text,
		})
	}
}

func extractCommand(message *TelegramMessage) string {
	command := ""

	var entity *TelegramMessageEntity

	for i := 0; i < len(message.Entities); i++ {
		entity = message.Entities[i]
		if message.Entities[i].EntityType == "bot_command" {
			command = message.Text[entity.Offset+1 : entity.Offset+entity.Length]
			break
		}
	}

	return command
}

type customClient struct{}

func (c *customClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("user-agnet", "mini-telegram-bot")

	return http.DefaultClient.Do(req)
}

func main() {
	bot = CreateBot(apiKey, handleUpdate, &customClient{})
	bot.Start()
}
