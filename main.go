package main

import (
	"net/http"
	"log"
	"bytes"
	"fmt"
)

const URL = "https://api.telegram.org/bot197287389:AAGjR6JVSXLAv-qI-mTeVv6P3bYLHDUe6M8/"
const MyURL = "http://dmitrydorofeev.ru/vladjokes/"

func main() {
	fmt.Println("babab1")
	_, err := http.Get(URL + "setWebhook?url=" + MyURL)
	if err != nil {
		log.Fatal("Error")
	}

	fmt.Println("babab2")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("babab")
		_, err := http.Post(URL + "sendMessage", "application/x-www-form-urlencoded", bytes.NewBufferString("hello"))
		if err != nil {
			log.Fatal("ba")
		}
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("hi"))
	})

	http.ListenAndServe(":7356", nil)
}
