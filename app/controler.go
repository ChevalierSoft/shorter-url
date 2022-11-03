package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @BasePath /api/v1

// shorter-url godoc
// @Description get every link in database when the api in not in production mode
// @Produce json
// @Success 200 {json} []Link
// @Router /l [get]
func (controller *HttpController) getLinks(c *gin.Context) {
	var links []Link
	if err := controller.Database.NewSelect().Model(&links).Scan(c); err != nil {
		logrus.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": links})
}

// shorter-url godoc
// @Description get a link by id
// @Produce json
// @Param id path string true "Link ID"
// @Success 301 {string} string "Redirect to the link"
// @Failure 404 {string} string "Link not found"
// @Router /l/{id} [get]
func (controller *HttpController) getLink(c *gin.Context) {
	var link Link

	err := controller.Database.NewSelect().
		Model(&link).
		Column("url", "visits").
		Where("id = ?", c.Param("id")).
		Scan(c)
	if err != nil {
		logrus.Error(err)
	}

	_, err = controller.Database.NewUpdate().
		Model(&link).
		Set("visits = ?", link.Visits+1).
		Set("last_visit = ?", time.Now()).
		Where("id = ?", c.Param("id")).
		Exec(c)
	if err != nil {
		logrus.Error(err)
	}

	if link.Url != "" {
		c.Redirect(http.StatusMovedPermanently, link.Url)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
	}
}

// shorter-url godoc
// @Description register a new link
// @Produce json
// @Param link body LinkDto true "Link to register"
// @Success 200 {json} Link
// @Failure 400 {string} string "Url is required"
// @Failure 400 {string} string "Url is not valid"
// @Router /l [post]
func (controller *HttpController) setNewLink(c *gin.Context) {
	var input LinkDto

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
		logrus.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": link})
}
