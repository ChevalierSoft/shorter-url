package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type HttpController struct {
	*gin.Engine
	Database *bun.DB
}

func NewHttpController(db *bun.DB) *HttpController {
	return &HttpController{Database: connectDB(), Engine: gin.New()}
}

func main() {
	// if os.Getenv("PRODUCTION") == "true" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	db := connectDB()
	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := NewHttpController(db)
	router.Use(cors.Default())

	// ? if not in production
	if gin.Mode() != gin.ReleaseMode {
		// todo: add pagination
		router.GET("/", router.getLinks) // ? debug : get all links
	}
	router.POST("/", router.setNewLink)
	router.GET("/:id", router.getLink)

	router.Run(":80")
}
