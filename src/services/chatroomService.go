package services

import (
	"context"

	"github.com/Sahil2k07/Redis-Pub-Sub/src/configs"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type chatService struct {
	redisClient *redis.Client
	ctx         echo.Context
}

func NewChatService(c echo.Context) *chatService {
	return &chatService{
		redisClient: configs.RedisClient,
		ctx:         c,
	}
}

func (cs *chatService) HandleChatroomConnection(redisChannel string, wsConn *websocket.Conn) error {
	ctx := context.Background()

	sub := configs.RedisClient.Subscribe(ctx, redisChannel)
	defer sub.Close()

	subMessages := sub.Channel()

	go sendMessages(subMessages, wsConn)

	for {
		// Receive Message
		_, msg, err := wsConn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				return nil
			}

			cs.ctx.Logger().Error("the json parsing failed")
			return err
		}

		err = configs.RedisClient.Publish(ctx, redisChannel, msg).Err()
		if err != nil {
			cs.ctx.Logger().Error("Error publishing to Redis:", err)
			return err
		}

		// Save in the DB as well
	}
}

func sendMessages(msg <-chan *redis.Message, ws *websocket.Conn) {
	for m := range msg {
		if err := ws.WriteMessage(websocket.TextMessage, []byte(m.Payload)); err != nil {
			return
		}
	}
}
