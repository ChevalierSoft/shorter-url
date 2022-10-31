package main

import (
	"log"
	"net/http"
	"net/url"

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
	if link.Url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Url is required"})
		return
	}
	_, err := url.ParseRequestURI(link.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Url is not valid"})
		return
	}

	_, err = controller.Database.NewInsert().Model(&link).Exec(c)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": link})
}
