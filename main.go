package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("hello"))
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Println("start arver")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
