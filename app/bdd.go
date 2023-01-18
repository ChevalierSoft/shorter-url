package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func connectDB() *bun.DB {
	dsn := "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT") + "/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(pgdb, pgdialect.New())
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func createSchema(db *bun.DB) error {
	// ? list of models
	models := []interface{}{
		(*Link)(nil),
	}

	// ? enable uuid in postgres
	createExtention := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	_, err := db.Exec(createExtention)
	if err != nil {
		panic(err)
	}

	// ? create tables with models
	for _, model := range models {
		_, err = db.NewCreateTable().
			Model(model).
			IfNotExists().
			Varchar(500).
			Exec(context.Background())
		if err != nil {
			panic(err)
		}
	}

	return nil
}
