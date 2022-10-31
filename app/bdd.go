package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// connectDB connects to the database using bun
func connectDB() *bun.DB {
	// pgconn := pgdriver.NewConnector(
	// 	pgdriver.WithAddr(os.Getenv("POSTGRES_HOST")+"=:"+os.Getenv("POSTGRES_PORT")),
	// 	pgdriver.WithUser(os.Getenv("POSTGRES_USER")),
	// 	pgdriver.WithPassword(os.Getenv("POSTGRES_PASSWORD")),
	// 	pgdriver.WithDatabase(os.Getenv("POSTGRES_DB")),
	// )
	// db := bun.NewDB(pgdriver.NewConnector(pgdriver.WithDSN(pgconn)), pgdialect.New())
	// db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose()))

	dsn := "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT") + "/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(pgdb, pgdialect.New())
	// if err := db.Ping(); err != nil {
	// 	panic(err)
	// }
	return db
}

func createSchema(db *bun.DB) error {
	// models := []interface{}{
	// 	(*Link)(nil),
	// }
	// for _, model := range models {
	// 	err := db.CreateTable(model, &orm.CreateTableOptions{
	// 		Temp:        false,
	// 		IfNotExists: true,
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// for _, model := range models {
	// 	db.RegisterModel(model)
	// }

	// create table if not exists using bun
	// res, err := db.NewCreateTable().Model((*Link)(nil)).IfNotExists().Exec(context.Background())

	createExtention := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	_, err := db.Exec(createExtention)
	if err != nil {
		panic(err)
	}

	_, err = db.NewCreateTable().
		Model((*Link)(nil)).
		IfNotExists().
		Varchar(500).
		Exec(context.Background())
	if err != nil {
		panic(err)
	}

	return nil
}
