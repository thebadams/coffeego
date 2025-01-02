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
	StatusCode int                       `json:"statusCode"`
	Success    bool                      `json:"success"`
	Data       []database.ListCoffeesRow `json:"data"`
	Message    string                    `json:"message"`
}

func (app *application) getCoffee(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Route Hit: %s", r.URL)

	ctx := context.Background()
	db, err := sql.Open("sqlite3", "./coffeego.db")
	if err != nil {
		app.logger.Error("Error opening database", slog.String("ERROR", err.Error()))
		app.serverError(w, r, err)
	}

	queries := database.New(db)
	coffees, err := queries.ListCoffees(ctx)
	if err != nil {
		app.logger.Error("Error Getting Coffees", slog.String("ERROR", err.Error()))
		app.serverError(w, r, err)

	}

	app.logger.Info("Coffees Found", slog.Any("Coffees", coffees))
	payload := GetCoffeePayload{StatusCode: http.StatusOK, Success: true, Data: coffees, Message: msg}
	j, err := json.Marshal(payload)
	if err != nil {
		app.serverError(w, r, err)

	}
	w.Write(j)

}

type PostCoffeePayload struct {
	StatusCode int             `json:"statusCode"`
	Success    bool            `json:"success"`
	Message    string          `json:"message"`
	Data       database.Coffee `json:"data"`
}

func (app *application) postCoffee(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {

		app.logger.Error("ERROR PARSING FORM", slog.String("ERROR STRING", err.Error()))
	}
	coffeeName := r.PostForm.Get("Coffee")
	roasterName := r.PostForm.Get("Roaster")

	app.logger.Info("ROUTE HIT", slog.String("Method", r.Method), slog.String("URL", r.URL.Path), slog.String("Coffee", coffeeName), slog.String("Roaster", roasterName))
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "coffeego.db")
	if err != nil {
		app.logger.Error("Error opening database", slog.String("ERROR", err.Error()))

	}

	queries := database.New(db)
	//first find roaster

	roaster, err := queries.FindRoasterByName(ctx, roasterName)
	if roaster.ID < 1 {
		roaster, err = queries.CreateRoaster(ctx, roasterName)
		if err != nil {
			app.logger.Error("Error Creating Roaster", slog.String("ERROR", err.Error()))
		}

	}
	if err != nil {
		app.logger.Error("Error Finding Roaster By Name", slog.String("ERROR", err.Error()))
	}

	coffee, err := queries.CreateCoffee(ctx, database.CreateCoffeeParams{Name: coffeeName, RoasterID: roaster.ID})
	if err != nil {
		app.logger.Error("Error Creating Coffee", slog.String("ERROR", err.Error()))

	}
	payload := PostCoffeePayload{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Coffee Successfully Created",
		Data:       coffee,
	}
	j, err := json.Marshal(payload)
	if err != nil {

		app.logger.Error("Error Marshalling JSON", slog.String("Error", err.Error()))
	}
	w.Write(j)

}
