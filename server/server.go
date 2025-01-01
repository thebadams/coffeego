package server

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	log.Println("Server Started")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)

	}

}
