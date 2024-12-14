package server

import (
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	lib "server/internal/lib"
	"server/internal/model"
	"server/internal/repository"
	"strconv"
)

func (s *FiberServer) RegisterSendMessage() {
	s.SocketServer.OnEvent("", "SEND_MESSAGE", func(s socketio.Conn, msg string) {
		type Message struct {
			Token   string `json:"token"`
			Message string `json:"message"`
			ChatId  int    `json:"chatId"`
		}

		var message Message

		err := json.Unmarshal([]byte(msg), &message)
		if err != nil {
			fmt.Println(err)
			return
		}

		claims, err := lib.GetClaimsFromJWT(message.Token)

		messageRepository := repository.NewMessageRepository()

		chatId := uint(message.ChatId)
		userId := claims["ID"].(float64)

		messageEntity := model.Message{
			UserID:        uint(userId),
			Content:       message.Message,
			PrivateChatID: &chatId,
		}

		createdMessage, err := messageRepository.Create(messageEntity)
		if err != nil {
			log.Fatal(err)
			return
		}

		jsonData, err := json.Marshal(createdMessage)
		if err != nil {
			log.Fatal(err)
			return
		}

		s.Emit("SEND_MESSAGE:"+strconv.Itoa(message.ChatId), string(jsonData))
	})
}

func (s *FiberServer) RegisterFiberWsRoutes() {
	s.RegisterSendMessage()
}
