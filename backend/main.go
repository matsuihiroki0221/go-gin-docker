package main

import (
	_ "embed"

	"github.com/gin-gonic/gin"
	"github.com/matsuihiroki0221/gin-docker/db"
	"github.com/matsuihiroki0221/gin-docker/router"
	"github.com/matsuihiroki0221/gin-docker/settings"
)

// var db = make(map[string]string)

//go:embed config.json
var configByte []byte

func main() {

	settings.Init(configByte)

	db.Init()

	r := gin.Default()
	router.DefineRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
