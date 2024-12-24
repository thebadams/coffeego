package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thebadams/coffeego/database"
)

func main() {
	seed := flag.Bool("seed", false, "Run seed command or not?")

}

//go:embed schema.sql
var ddl string

func seed() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./coffeego.db")
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	return nil
}
