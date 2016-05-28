package main

import (
	"net/http"
	"log"
	"gopkg.in/telegram-bot-api.v4"
    "math/rand"
)

const URL = "https://api.telegram.org/bot197287389:AAGjR6JVSXLAv-qI-mTeVv6P3bYLHDUe6M8/"
const MyURL = "https://dmitrydorofeev.ru/vladjokes/"

var jokes = [...]string{
    "Кек",
    "Куплю себе дрипку, буду делать 7-8 затяжек",
    "Водка тёплая, потому что она с юга",
    "На улице такой минус, что Санёк под него зачитал",
    "Новый год. Пробили курсанты",
    "И не сложно, и не легко",
    "ААА! Сложна!!",
    "Для счастья нужны: дошик, насвай и кумыс",
    "Пойдут часы, пойдут года, Саня бросил Иру, но в душе она у него навсегда",
    "То чувство, когда не дрочил под партой",
    "Вась, тебе сейчас свиные амулеты принесут",
    "Легкоооооо",
    "Книга в уме, карта в рукаве",
    "Ты будешь джин как смузи разливать?",
    "А ты знал, что бумагу из хлеба делают",
    "Твоя жопа в невесомости",
    "У Санька все не по жизни, а по кайфу",
    "Гоша - геолог-проктолог",
    "Папин дворняга, мамин симпатяга",
    "Мемасы",
    "Гашло, от тебя зеленеют",
}

func main() {
	bot, err := tgbotapi.NewBotAPI("197287389:AAGjR6JVSXLAv-qI-mTeVv6P3bYLHDUe6M8")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(MyURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:7356", nil)

	for update := range updates {
        message := update.Message
        bot.Send(tgbotapi.NewMessage(message.Chat.ID, jokes[rand.Intn(len(jokes))]))
	}
}
