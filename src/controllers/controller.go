package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var Upgrader = websocket.Upgrader{}

func InitControllers(e *echo.Echo) {
	e.GET("/api/chat", ChatRoomController)
}
