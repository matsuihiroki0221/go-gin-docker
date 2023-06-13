package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func DefineRouter(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "SUCCESS!",
			})
			fmt.Println("hello")
			log.Println("hello")
		})
	}
}
