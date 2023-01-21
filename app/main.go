package main

import (
	docs "github.com/ChevalierSoft/shorter-url/docs"
	"github.com/Shopify/sarama"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	bun "github.com/uptrace/bun"
)

type HttpController struct {
	*gin.Engine
	Database *bun.DB
	Producer *sarama.SyncProducer
}

func NewHttpController() *HttpController {
	return &HttpController{Database: connectDB(), Engine: gin.New(), Producer: setProducer()}
}

// @title shorter-url API
// @description This is a simple url shortener api
// @version 0.1.0
// @BasePath /api/v1
func main() {
	// ? init kafka
	defer kafkaWriter.Close()

	router := NewHttpController()
	router.Use(cors.Default())

	// ? init db
	err := createSchema(router.Database)
	if err != nil {
		panic(err)
	}
	defer router.Database.Close()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		g1 := v1.Group("/l")
		{
			// ? if not in production
			if gin.Mode() != gin.ReleaseMode {
				// todo: add pagination
				g1.GET("/", router.getLinks) // ? debug : get all links
			}
			g1.POST("/", router.setNewLink)
			g1.GET("/:id", router.getLink)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":12345")
}
