# Mini Telegram Bot API

This is a very minimal bot API for telegram. It can be used to send and receive messages via a bot in Telegram. The user can implement a custom HTTP client to allow them to modify the request object for operations like modifying headers. The API only has support for long polling, not webhooks.
## Example Usage
This example will simply reply "Pong!" to any message sent to the bot.
```go
package main

import (
	"net/http"

	tgbot "github.com/matthew-hajec/mini-telegram-bot"
)

var bot tgbot.TelegramBot

func updateHandler(update *tgbot.TelegramUpdateResult) {
	for i := 0; i < len(update.Results); i++ {
		chatId := update.Results[i].Message.Chat.Id

		bot.SendMessage(&tgbot.TelegramOutgoingMessage{
			ChatID: chatId,
			Text:   "Pong!",
		})
	}
}

func main() {
	bot = *tgbot.CreateBot("0000000000:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", updateHandler, http.DefaultClient) // The first parameter would be your bot's API key
	bot.Start()
}

```

### MessageHandler 
```go
type TelegramMessageHandler func(update *TelegramUpdateResult)
```
Whenever a message update is recieved from Telegram, the ```handler``` passed to ```tgbot.CreateBot``` is called with an array containing all of the messages. When processing these updates, you should loop over each one.

### SendMessage 
```go
func (b *TelegramBot) SendMessage(message *TelegramOutgoingMessage) (bool, error)
```
To send a message, call ```TelegramBot.SendMessage``` with a message struct instance containing the chat ID and text content of the message. If the request is successful, the first boolean return value will be true, it will be false in all other cases. If an error occurred that was not recieved from the Telegram server, it will be contained in the error return value. If an error is returned from telegram's server, the boolean will be false and the error will be nil.
