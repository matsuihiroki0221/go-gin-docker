package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func DefineRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "SUCCESS!",
			})
			fmt.Println("hello")
			log.Println("hello")
		})
	}
}
