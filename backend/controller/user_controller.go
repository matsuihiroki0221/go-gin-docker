package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matsuihiroki0221/gin-docker/repository"
)

type Controller struct{}

func (ctrl Controller) Index(c *gin.Context) {
	var r repository.UserRepository
	users, err := r.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, users)
}
