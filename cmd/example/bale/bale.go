package main

import (
	"fmt"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var numri = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
	tgbotapi.NewInlineKeyboardButtonData("ورود", "hey"),
	tgbotapi.NewInlineKeyboardButtonURL("login", "http://localhost:3000/pending"),
))

func main() {
	_ = godotenv.Load()
	const baseApiUrl = "https://tapi.bale.ai/bot%s/%s"

	token := os.Getenv("BALE_TOKEN")

	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, baseApiUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	fmt.Println("runned 1")
	for update := range updates {

		if update.Message == nil {
			continue
		}
		//fmt.Println(update.CallbackQuery.Data)

		isPrivate := update.Message.Chat.IsPrivate()

		if !isPrivate {
			continue
		}

		//message := update.Message.Text
		//
		//userName := update.Message.From.UserName
		//
		//userId := update.Message.From.ID

		var msg tgbotapi.MessageConfig

		//if strings.Contains(message, "@") && strings.Contains(message, ".") {
		//	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "ایمیل شما تایید شد")
		//} else {
		//
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "لطفا ایمیل خود را وارد کنید")
		//}

		msg.ReplyMarkup = numri

		nMsg, errB := bot.Send(msg)

		if errB != nil {
			fmt.Println(err)
		}

		fmt.Println(nMsg)
	}

}
