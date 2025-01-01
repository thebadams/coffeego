package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/thebadams/coffeego/database"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Route Hit %s", r.URL)
	app.logger.Info(msg)
	w.Write([]byte("Hello World"))
}

type GetCoffeePayload struct {
	StatusCode int      `json:"statusCode"`
	Success    bool     `json:"success"`
	Data       []string `json:"data"`
	Message    string   `json:"message"`
}

func (app *application) getCoffee(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Route Hit: %s", r.URL)

	ctx := context.Background()
	db, err := sql.Open("sqlite3", "./coffeego.db")
	if err != nil {
		app.logger.Error("Error opening database", slog.String("ERROR", err.Error()))

	}

	queries := database.New(db)
	coffees, err := queries.ListCoffees(ctx)
	if err != nil {
		app.logger.Error("Error Getting Coffees", slog.String("ERROR", err.Error()))

	}

	app.logger.Info("Coffees Found", slog.Any("Coffees", coffees))
	payload := GetCoffeePayload{StatusCode: http.StatusOK, Success: true, Data: coffees, Message: msg}
	j, err := json.Marshal(payload)
	w.Write(j)

}
