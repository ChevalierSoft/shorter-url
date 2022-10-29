package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func getLinkByID(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	user := &Link{Id: id}
// 	err := G_db.Model(user).WherePK().Select()
// 	if err != nil {
// 		log.Printf("getLinkByID: %v : %v", user.Id, err)
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Link not found"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, user)
// }

func getLinks(c *gin.Context) {
	var links []Link
	err := G_db.Model(&links).Select()
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Link not found"})
	}
	c.IndentedJSON(http.StatusOK, links)
}

func main() {

	// gin.SetMode(gin.ReleaseMode)

	// err :=
	createSchema(G_db)
	// if err != nil {
	// 	panic(err)
	// }
	defer G_db.Close()

	router := gin.New()
	router.Use(cors.Default())

	router.GET("/", getLinks)

	router.Run("8080")

}