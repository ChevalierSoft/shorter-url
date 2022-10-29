package main

// import (
// "fmt"
// "os"

// "github.com/go-pg/pg/v10"
// "github.com/go-pg/pg/v10/orm"
// "github.com/google/uuid"
// )

// var G_db = pg.Connect(&pg.Options{
// 	User:     os.Getenv("POSTGRES_USER"),
// 	Password: os.Getenv("POSTGRES_PASSWORD"),
// 	Database: os.Getenv("POSTGRES_DB"),
// 	Addr:     os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT"),
// })

// type Link struct {
// 	Uuid uuid.UUID `json:"uuid" pg:",pk,type:uuid,default:uuid_generate_v4()"`
// 	Link string
// }

// createSchema creates database schemas
// func createSchema(db *pg.DB) error {

// 	models := []interface{}{
// 		(*Link)(nil),
// 	}

// 	for _, model := range models {
// 		exists, _ := db.Model(model).Exists()
// 		if exists {
// 			fmt.Printf("table <%T> exists -> %v\n", model, exists)
// 			continue
// 		}

// 		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
