// https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/bot.go
// Many of the ideas in this package come from this project.

package tgbot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// HTTPClient which makes requests, can be the default go http client or contain a custom "Do" function which can do things such as set headers.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type TelegramBot struct {
	apiKey  string
	offset  int
	handler TelegramMessageHandler
	client  HTTPClient
}

func CreateBot(apiKey string, handler TelegramMessageHandler, client HTTPClient) *TelegramBot {
	bot := &TelegramBot{
		apiKey:  apiKey,
		offset:  0,
		handler: handler,
		client:  client,
	}

	return bot
}

func (b *TelegramBot) Start() {
	for {
		strBody := `{
			"offset": ` + fmt.Sprintf("%d", b.offset) + `,
			"timeout": 300,
			"allowed_updates": ["messages"]
		}`

		rBody := strings.NewReader(strBody)

		resp, err := b.runTelegramMethod("getUpdates", rBody)

		defer resp.Body.Close()

		if err != nil {
			log.Println(err)
			log.Println("An error occured while attempting to fetch updates. Retrying in 3 seconds.")
			time.Sleep(3 * time.Second)

			continue
		}

		body, err := io.ReadAll(resp.Body)

		var updates TelegramUpdateResult

		err = json.Unmarshal(body, &updates)

		if err != nil {
			log.Fatalln(err)
			log.Fatalln("A fatal error occured. Stopping bot.")
		}

		if len(updates.Results) > 0 {
			b.handler(&updates)
			b.offset = updates.Results[len(updates.Results)-1].ID + 1
		}
	}
}

func (b *TelegramBot) SendMessage(message *TelegramOutgoingMessage) (bool, error) {
	rbBytes, err := json.Marshal(message)

	if err != nil {
		return false, err
	}

	rBody := strings.NewReader(string(rbBytes))

	resp, err := b.runTelegramMethod("sendMessage", rBody)

	if err != nil {
		return false, err
	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		return false, err
	}

	var ok TelegramOKResponse

	err = json.Unmarshal(body, &ok)

	if err != nil {
		return false, err
	}

	return ok.Ok, nil
}

func (b *TelegramBot) runTelegramMethod(method string, rBody io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+b.apiKey+"/"+method, rBody)

	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")

	return b.client.Do(req)
}
