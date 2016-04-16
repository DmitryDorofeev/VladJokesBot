package main

import (
	"net/http"
	"log"
	"bytes"
	"fmt"
	"net/url"
	"gopkg.in/telegram-bot-api.v4"
)

const URL = "https://api.telegram.org/bot197287389:AAGjR6JVSXLAv-qI-mTeVv6P3bYLHDUe6M8/"
const MyURL = "http://dmitrydorofeev.ru/vladjokes/"

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(MyURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:7356", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}