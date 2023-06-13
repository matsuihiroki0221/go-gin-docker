package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matsuihiroki0221/gin-docker/routes"
)

// var db = make(map[string]string)

func main() {
	r := gin.Default()
	routes.DefineRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
