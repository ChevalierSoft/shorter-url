package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *HttpController) getLinks(c *gin.Context) {
	var links []Link
	if err := controller.Database.NewSelect().Model(&links).Scan(c); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": links})
}

func (controller *HttpController) setNewLink(c *gin.Context) {
	var input Link
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	link := Link{Url: input.Url}
	// bite, err :=
	_, err := controller.Database.NewInsert().Model(&link).Exec(c)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": link})
}
