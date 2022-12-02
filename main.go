package main

import (
	"github.com/gin-gonic/gin"

	"tribal/entities"
	"tribal/handlers"
	"tribal/repositories"
)

func main() {
	rep := repositories.NewChuckNorris(entities.ChuckNorrisURL)
	han := handlers.NewChuckNorris(rep)
	r := gin.Default()

	r.GET("/chuck-norris", func(c *gin.Context) {
		r, err := han.GetManyDistinct(entities.JokesNum)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(200, gin.H{
			"jokes": r,
		})
	})

	r.Run()
}
