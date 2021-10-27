package thread

import (
	"net/http"

	"webstudy/global"
	"webstudy/store"

	"github.com/gorilla/websocket"

	"github.com/tinode/chat/server/logs"
)

func WebSocketConnect(w http.ResponseWriter, r *http.Request) {
	// Handles websocket requests from peers
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow connections from any Origin
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		logs.Err.Println("ws: Not a websocket handshake")
		return
	} else if err != nil {
		logs.Err.Println("ws: failed to Upgrade ", err)
		return
	}

	id := store.Store.GetUidString()

	ss := global.Global.Sessionstore.NewSession(ws, id)
    
    go ss.ReadLoop()
	go ss.WriteLoop()
	
}
