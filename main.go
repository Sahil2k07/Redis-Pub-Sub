package main

import (
	"github.com/Sahil2k07/Redis-Pub-Sub/src/configs"
	"github.com/Sahil2k07/Redis-Pub-Sub/src/controllers"
	"github.com/labstack/echo/v4"
)

func init() {
	configs.RedisInit()
}

func main() {
	e := echo.New()

	controllers.InitControllers(e)

	e.Logger.Fatal(e.Start(":1323"))
}
