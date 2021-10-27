package session

import (
	"encoding/json"
	"fmt"
	"time"
	"webstudy/store/types"

	"github.com/gorilla/websocket"
	"github.com/tinode/chat/server/logs"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// // Time allowed to read the next pong message from the peer.
	pongWait = time.Second * 55

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type Session struct {
	ws         *websocket.Conn
	remoteAddr string
	// Session ID
	sid string

	//发送消息
	send chan interface{}

	ver int
}

func (sess *Session) ReadLoop() {
	defer func() {
		// sess.closeWS()
		// sess.cleanUp(false)
	}()

	sess.ws.SetReadLimit(512)
	sess.ws.SetReadDeadline(time.Now().Add(time.Second * 55))
	sess.ws.SetPongHandler(func(string) error {
		sess.ws.SetReadDeadline(time.Now().Add(time.Second * 55))
		return nil
	})

	for {
		// Read a ClientComMessage
		_, raw, err := sess.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
				websocket.CloseNormalClosure) {
				logs.Err.Println("ws: readLoop", sess.sid, err)
			}
			return
		}

		sess.dispatchRaw(raw)
	}
}

func (sess *Session) dispatchRaw(message []byte) {
	now := types.TimeNow()
	var clientMessage ClientComMessage
	if err := json.Unmarshal(message, &clientMessage); err != nil {
		sess.queueOut(ErrMalformed("", "", now))
	}
	sess.dispatch(&clientMessage)
}

func (sess *Session) dispatch(clientMessage *ClientComMessage) {
	var handler func(*ClientComMessage)
	switch {
	case clientMessage.Login != nil:
		handler = sess.login
	}
	handler(clientMessage)
}

//发送信息
func (sess *Session) queueOut(msg *ServerComMessage) {
	select {
	case sess.send <- msg:
	}
}

func (sess *Session) WriteLoop() {
	ticker := time.NewTicker(pingPeriod) //定时器
	for {
		select {
		case msg, ok := <-sess.send:
			if !ok {
				return
			}
			switch v := msg.(type) {
			case []*ServerComMessage:
				for _, msg := range v {
					sess.sendMessage(msg)
				}
			case *ServerComMessage:
				sess.sendMessage(msg)
			}
		case <-ticker.C: //发送心跳
			if err := wsWrite(sess.ws, websocket.PingMessage, nil); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
					websocket.CloseNormalClosure) {
					logs.Err.Println("ws: writeLoop ping", sess.sid, err)
				}
				return
			}
		}
	}
}

//发送文字消息
func (sess *Session) sendMessage(msg interface{}) bool {
	if len(sess.send) > sendQueueLimit {
		logs.Err.Println("ws: outbound queue limit exceeded", sess.sid)
		return false
	}
	fmt.Println("sendMessage:", len(sess.send))

	if err := wsWrite(sess.ws, websocket.TextMessage, msg); err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
			websocket.CloseNormalClosure) {
			logs.Err.Println("ws: writeLoop", sess.sid, err)
		}
		return false
	}
	return true
}

// Writes a message with the given message type (mt) and payload.
func wsWrite(ws *websocket.Conn, mt int, msg interface{}) error {
	var bits []byte
	if msg != nil {
		bits = msg.([]byte)
	} else {
		bits = []byte{}
	}
	ws.SetWriteDeadline(time.Now().Add(writeWait))
	return ws.WriteMessage(mt, bits)
}

func (s *Session) login(clientMessage *ClientComMessage) {
	var serverMessage ServerComMessage
	serverMessage = ServerComMessage{
		Info: &MsgServerInfo{
			Src: clientMessage.Login.Scheme,
		},
	}

	s.queueOut(&serverMessage)
}
