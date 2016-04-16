package main

import (
	"net/http"
	"log"
	"bytes"
)

const URL = "https://api.telegram.org/bot197287389:AAGjR6JVSXLAv-qI-mTeVv6P3bYLHDUe6M8/"
const MyURL = "https://dmitrydorofeev.ru/vladjokes/"

func main() {
	_, err := http.Get(URL + "setWebhook?url=" + MyURL)
	if err != nil {
		log.Fatal("Error")
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := http.Post(URL + "sendMessage", "application/x-www-form-urlencoded", bytes.NewBufferString("hello"))
		if err != nil {
			log.Fatal("ba")
		}
		res.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":7356", nil)
}
