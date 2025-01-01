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

	app := &application{logger: slog.New(slog.NewTextHandler(os.Stdout, nil))}
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", app.home)
	mux.HandleFunc("GET /coffee", app.getCoffee)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	srv := http.Server{Addr: ":4000", Handler: mux}
	app.server = &srv

	return app
}
