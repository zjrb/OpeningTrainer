package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zjrb/OpeningTrainer/internal/core/services"
	"github.com/zjrb/OpeningTrainer/internal/logger"
)

type WebsocketHandler struct {
	svc      *services.ChessService
	upgrader *websocket.Upgrader
	l        *logger.Logger
}

func NewWebSocketHandler(svc *services.ChessService) *WebsocketHandler {
	return &WebsocketHandler{
		svc: svc,
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (ws *WebsocketHandler) HandleConnections() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := ws.upgrader.Upgrade(w, r, nil)
		if err != nil {
			ws.l.Error("Error", err)
		}
		defer conn.Close()
		go ws.handleConnection(conn)
	})
}

func (ws *WebsocketHandler) handleConnection(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			ws.l.Error("Error reading messsage", err)
			break
		}
		ws.l.Debug("Received %s\n", message)
		ws.svc.
	}

}
