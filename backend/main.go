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

	// configファイルを読み込み、設定の初期化
	settings.Init()

	// dbのインスタンス作成
	db.Init()

	r := gin.Default()
	router.DefineRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
