package main

import (
	// _ "shorter-url/bdd"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	err := createSchema(G_db)
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.Use(cors.Default())

	router.GET("/")

	router.Run("80")
}
