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

	// gin.SetMode(gin.ReleaseMode)

	db := connectDB()
	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := NewHttpController(db)
	router.Use(cors.Default())

	// add a get route to get all links

	router.GET("/", router.getLinks)
	router.POST("/", router.setNewLink)

	router.Run(":80")
}
