package server

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
	server *http.Server
}

func (app *application) StartServer() {
	app.logger.Info("Server Starting")
	err := app.server.ListenAndServe()
	if err != nil {

		app.logger.Error(err.Error())
		os.Exit(1)
	}

}

func CreateServer() *application {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello world"))
	})

	srv := http.Server{Addr: ":4000", Handler: mux}

	app := application{logger: slog.New(slog.NewTextHandler(os.Stdout, nil)), server: &srv}
	return &app
}
