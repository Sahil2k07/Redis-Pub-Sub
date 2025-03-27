package controllers

import (
	"github.com/Sahil2k07/Redis-Pub-Sub/src/services"
	"github.com/labstack/echo/v4"
)

func ChatRoomController(c echo.Context) error {
	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	channel := c.QueryParam("roomID")
	cs := services.NewChatService(c)

	return cs.HandleChatroomConnection(channel, ws)
}
