package main

import (
	"os"

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
	DB       *bun.DB
	Producer *sarama.SyncProducer
}

func NewHttpController() *HttpController {
	return &HttpController{DB: connectDB(), Engine: gin.New(), Producer: setProducer()}
}

func checkEnv() {
	if os.Getenv("POSTGRES_USER") == "" {
		panic("env 'POSTGRES_USER' is not set")
	}
	if os.Getenv("POSTGRES_PASSWORD") == "" {
		panic("env 'POSTGRES_PASSWORD' is not set")
	}
	if os.Getenv("POSTGRES_HOST") == "" {
		panic("env 'POSTGRES_HOST' is not set")
	}
	if os.Getenv("POSTGRES_PORT") == "" {
		panic("env 'POSTGRES_PORT' is not set")
	}
	if os.Getenv("POSTGRES_DB") == "" {
		panic("env 'POSTGRES_DB' is not set")
	}
	// ? release mode for gin condition
	if os.Getenv("PRODUCTION") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
}

// @title shorter-url API
// @description This is a simple url shortener api
// @version 0.1.0
// @BasePath /api/v1
func main() {

	checkEnv()

	router := NewHttpController()
	router.Use(cors.Default())

	// ? init db
	err := createSchema(router.DB)
	if err != nil {
		panic(err)
	}
	defer router.DB.Close()

	// ? init red panda

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
