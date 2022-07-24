package answers

import (
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DefaultKeyboard() tgbotapi.ReplyKeyboardMarkup {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🤙 Предложить мем/новость"),
			tgbotapi.NewKeyboardButton("💊 Интересное"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⛪ Контакты"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🌐 О боте"),
		),
	)
	return numericKeyboard
}

func InterestingKeyboard() tgbotapi.ReplyKeyboardMarkup {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🤪 Язык Dolbogram"),
			tgbotapi.NewKeyboardButton("🐸 Жабий язык"),
		),
	)
	return numericKeyboard
}

func SwitchReg(reg string, option string) bool {
	matched, _ := regexp.MatchString(reg, option)
	return matched
}
