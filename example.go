package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("", func(c *gin.Context) {
		var json struct {
			Message string `json:"message"`
		}

		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := gin.H{
			"status": http.StatusOK,
			"data":   json.Message,
		}

		c.JSON(http.StatusOK, response)
	})

	r.Run()
}
