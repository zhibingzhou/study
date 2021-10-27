package session

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tinode/chat/server/logs"
)

var sendQueueLimit = 128

type SessionStore struct {
	lock      sync.Mutex
	sessCache map[string]*Session
	lifeTime  time.Duration
}


// NewSessionStore initializes a session store.
func NewSessionStore(lifetime time.Duration) *SessionStore {
	ss := &SessionStore{
		lifeTime:  lifetime,
		sessCache: make(map[string]*Session),
	}

	return ss
}

func (ss *SessionStore) NewSession(conn interface{}, sid string) *Session {

	ss.lock.Lock()

	if _, found := ss.sessCache[sid]; found {
		logs.Err.Fatalln("ERROR! duplicate session ID", sid)
	}
	ss.lock.Unlock()

	var s Session
	switch c := conn.(type) {
	case *websocket.Conn:
		s.ws = c
	}
	s.send = make(chan interface{}, sendQueueLimit)
	s.sid = sid
	ss.lock.Lock()
	ss.sessCache[sid] = &s
	ss.lock.Unlock()

	return &s
}
