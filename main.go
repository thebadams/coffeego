package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thebadams/coffeego/database"
)

type Coffee struct {
	Name    string `json:"name"`
	Roaster string `json:"roaster"`
}

func main() {
	log.Println("Hello World")
	seed := flag.Bool("seed", false, "A boolean flag used to selectively run the seedDB flag")

	flag.Parse()

	if *seed {
		err := seedDB()
		if err != nil {

			log.Panicf("Seed Error: %v", err)
		}

	}

}

//go:embed schema.sql
var ddl string

func seedDB() error {
	log.Println("SeedDB Running")
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./coffeego.db")
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := database.New(db)
	//read the json source file
	coffeeFile, err := os.Open("./coffees.json")
	if err != nil {

		return err
	}
	defer coffeeFile.Close()

	coffeeBytes, err := io.ReadAll(coffeeFile)
	if err != nil {

		return err
	}

	var coffees []Coffee

	err = json.Unmarshal(coffeeBytes, &coffees)
	if err != nil {
		return err
	}
	for i, v := range coffees {

		// first, find roaster by name
		roaster, err := queries.FindRoasterByName(ctx, v.Roaster)
		log.Printf("Roaster: %v", roaster.ID)
		if roaster.ID < 1 {
			roaster, err = queries.CreateRoaster(ctx, v.Roaster)

		}
		if err != nil {
			log.Printf("Error Finding Roaster By Name: %s", v.Roaster)
			return err

		}
		// then create coffee
		coffee, err := queries.CreateCoffee(ctx, database.CreateCoffeeParams{Name: v.Name, RoasterID: roaster.ID})
		log.Printf("Coffee Successfully Created: %v", coffee)
		if err != nil {
			log.Printf("Error creating a coffee at index %d", i)
			return err

		}
	}
	log.Println("Successfully Seeded Database")
	return nil
}
