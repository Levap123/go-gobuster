package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Levap123/go-gobuster/internal/domain/models"
	"github.com/Levap123/go-gobuster/internal/services"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handlers struct {
	services *services.Services
}

func New(services *services.Services) *Handlers {
	return &Handlers{
		services: services,
	}
}

func (h *Handlers) handleBust(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
	}

	urlChan := make(chan string)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
		}

		var msg models.GobusterRequest
		err = json.Unmarshal(message, &msg)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
		}
		go h.services.Bust(msg.Url, urlChan)

		for url := range urlChan {
			conn.WriteMessage(websocket.TextMessage, []byte(url))
		}

		break
	}
}
