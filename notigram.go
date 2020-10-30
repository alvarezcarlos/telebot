package notigram

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func NewTelegramBot(url, token string) *tb.Bot {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL:    url,
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func Notify(bot *tb.Bot, info *MessageInfo) error {
	chat, err := bot.ChatByID(info.ChatId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	message := fmt.Sprintf("%s says: \n %s", info.AppName, info.Message)
	_, err = bot.Send(chat, message)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}