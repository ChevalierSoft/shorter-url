package main

import (
	docs "github.com/ChevalierSoft/shorter-url/docs"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	bun "github.com/uptrace/bun"
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

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		g1 := v1.Group("/links")
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
	router.Run(":80")
}
