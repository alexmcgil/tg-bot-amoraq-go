package main

import (
	"amoraq-bot/answers"
	"amoraq-bot/auth"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const chatToOffer int64 = -1001793205533

func main() {

	bot, err := tgbotapi.NewBotAPI(auth.Token())

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		switch {
		case update.Message != nil:
			OnMessage(bot, update.Message)
		}

	}
}

func OnMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	var msg tgbotapi.MessageConfig

	msg.ReplyMarkup = answers.DefaultKeyboard()

	switch {

	case answers.SwitchReg(`(?i)интересное`, message.Text):
		msg = tgbotapi.NewMessage(message.Chat.ID, "Только никому не показывай, что ты здесь найдёшь.")
		msg.ReplyMarkup = answers.InterestingKeyboard()

	case answers.SwitchReg(`(?i)dolbogram`, message.Text):
		msg = tgbotapi.NewMessage(message.Chat.ID, "Вот что я для тебя нашёл: https://t.me/setlanguage/dlgram")
		msg.ReplyMarkup = answers.DefaultKeyboard()

	case answers.SwitchReg(`(?i)контакты`, message.Text):
		msg = tgbotapi.NewMessage(message.Chat.ID, "Админ -> @alexmcgil")

	case answers.SwitchReg(`(?i)предложить`, message.Text):
		msg = tgbotapi.NewMessage(message.Chat.ID, "Отправь мне мем или видосик")

	case len(message.Text) == 0:
		msgF := tgbotapi.NewForward(chatToOffer, message.Chat.ID, message.MessageID)
		msg = tgbotapi.NewMessage(message.Chat.ID, "Ваш мем принят.")
		bot.Send(msgF)

	default:
		msg = tgbotapi.NewMessage(message.Chat.ID, "Донт андерстенд")
		msg.ReplyMarkup = answers.DefaultKeyboard()
	}

	switch message.Command() {

	case "start":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
		msg.Text = "Меню"
		msg.ReplyMarkup = answers.DefaultKeyboard()

	case "info":
		msg.Text = message.Chat.UserName + ", я уже что-то умею :)"

	case "help":
		msg.Text = "Никто тебе не поможет :( Но ты можешь написать разработчику данного бота @alexmcgil"

	case "status":
		msg.Text = "Еррор"

	case "secretcommand":
		msg.Text = strconv.Itoa(int(message.Chat.ID))

	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}
