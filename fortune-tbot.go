package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "6859893331:AAGJKPLoM0YRiYyIJpdR9m6xFKWZtm2u3lY"

var bot *tgbotapi.BotAPI

var faqbotnames = [2]string{"джейсон", "стетхем"}
var chatID int64
var answers = []string{
	"Да",
	"Нет",
	"Ешь растишку и данон, пися будет как бетон.",
	"Делу время, потехе час. Пиво, водочка и квас.",
	"Водка, пиво,коньячок. Я иду на турничок.",
	"Если тебе где-то не рады в рваных носках, то и в целых туда идти не стоит.",
	"Кто носит кофту Адидас и душиться одекалоном Шик, тому любая баба даст, а " +
		"может даже и мужик.",
	"Ты не ты, когда за тобой бегут менты.",
	"Когда поднимаешь бычок, то смотришь на длину окурка, а не сорт табака.",
	"Спрятать чувства легко, а вот спрятать стояк сложнее.",
	"Лучше 5 см спереди, чем 25 см сзади.",
	"Шкуры как спагетти - мокрые уже не ломаются.",
	"Хороший человек плохой воздух в себе держать не будет.",
	"И лампы не горят и врут календари. Если баба не дает,то напои.",
	"Шнурки на кросах завязывют только лохи, засунул их и ништяк.",
	" Легче жопой съехать с терки, чем учиться на пятерки.",
	"Яжка, сижки,турничок - через месяц я качок.",
	"Когда меня рожали, собственно, тогда я и родился.",
	"Базар как вода: чем лучше фильтруешь - тем меньше проблем с почками.",
	"Усли мужчина просит руки женщины, значит своя устала.",
	"Роллы для лохов, пельмашки для пацанов.",
	"Бегать за овцам - удел баранов. Я бегаю только за пивом и сигами.",
	"Почесал яица - понюхал руку.",
	"Сегодня люди намного дешевле, чем их одежда.",
	"Сколько писей не тряси - одна капелька в трусы.",
	"На чужом мопеде по ямам, как на торпеде.",
}

func connectWithTelegram() {

	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}
}

func sendMessage(msg string) {
	msgconfig := tgbotapi.NewMessage(chatID, msg)
	bot.Send(msgconfig)
}
func isMessageForUs(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}
	msginLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range faqbotnames {
		if strings.Contains(msginLowerCase, name) {
			return true
		}
	}
	return false
}

func getFAQanswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatID, getFAQanswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()
	updateconfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateconfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatID = update.Message.Chat.ID

			sendMessage("Задай свой вопрос, назвав меня по имени. Ответом на вопрос должен быть \"да\"либо \"нет\".Например,\"Джейсон," +
				" я готов сменить работу?\"либо\"Стетхем,я действительно хочу отправиться на этот вечер?\"")
		}
		if isMessageForUs(&update) {
			sendAnswer(&update)
		}
	}
}
