package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"github.com/zjrb/OpeningTrainer/internal/core/services"
	"github.com/zjrb/OpeningTrainer/internal/logger"
)

type WebsocketHandler struct {
	svc      *services.ChessService
	upgrader *websocket.Upgrader
	l        *logger.Logger
	ctx      context.Context
	clients  map[*websocket.Conn]bool
	mu       sync.Mutex
}

func NewWebSocketHandler(svc *services.ChessService, l *logger.Logger) *WebsocketHandler {
	return &WebsocketHandler{
		svc: svc,
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		l:       l,
		ctx:     context.Background(),
		clients: make(map[*websocket.Conn]bool),
	}
}

func (ws *WebsocketHandler) HandleConnections() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := ws.upgrader.Upgrade(w, r, nil)
		if err != nil {
			ws.l.Error("Error", err)
		}
		ws.l.Info("HANDLING A CONNECTION")
		defer conn.Close()
		ws.mu.Lock()
		ws.clients[conn] = true
		ws.mu.Unlock()
		ws.handleConnection(conn)
	})
}

func (ws *WebsocketHandler) processMessage(message []byte) (*domain.GameSession, error) {
	var gameSesh domain.GameSession
	err := json.Unmarshal(message, &gameSesh)
	if err != nil {
		ws.l.Error("Error parsing message", err)
		return nil, err
	}
	return &gameSesh, nil
}

func (ws *WebsocketHandler) handleConnection(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			ws.l.Error("Error reading messsage", err)
			break
		}
		ws.l.Debug("Received %s\n", message)
		gameSession, err := ws.processMessage(message)
		if err != nil {
			ws.l.Error("Error parsing messgae", err)
		}
		gameSession = ws.svc.HandleMessage(gameSession)
		msg, _ := json.Marshal(gameSession)
		conn.WriteMessage(websocket.TextMessage, msg)
	}

}
