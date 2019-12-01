package server

import (
	//"fmt"
	"log"
	"server/messages"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

const CommandQueueSize = 10

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	wait = 6000 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Conn struct {
	// sync.RWMutex
	conn         *websocket.Conn
	lastTransfer time.Time
	terminated   bool

	playerId   int32
	In         chan *messages.Message
	out        chan *messages.Message
	forwardOut chan *messages.Message
}

func NewServerConn(conn *websocket.Conn, forwardOut chan *messages.Message) *Conn {
	c := Conn{conn: conn}

	c.In = make(chan *messages.Message, CommandQueueSize)
	c.out = make(chan *messages.Message, CommandQueueSize)
	c.forwardOut = forwardOut

	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(wait))

	c.SetLastTransfer(time.Now())

	if conn == nil {
		// Local instance

		// go c.handleLocalWrite()
	} else {
		go c.handleRead()
		go c.handleWrite()
		// go c.handleSendKeepAlive()
	}

	return &c
}

func (s *Conn) handleRead() {
	if s.conn == nil {
		return
	}

	for {
		//s.RLock()

		messageType, msg, err := s.conn.ReadMessage()
		if err != nil || messageType != websocket.BinaryMessage {
			// c.Logger().Error(err)
			break
		}

		message := &messages.Message{}
		err = proto.Unmarshal(msg, message)
		if err != nil {
			break
		}

		s.In <- message
	}

	s.Close()
}

func (s *Conn) handleWrite() {
	//s.Lock()
	if s.conn == nil {
		for range s.out {
			//s.Done()
		}
		return
	}

	for e := range s.out {
		if s.GetTerminated() {
			// 	s.Done()
			continue
		}

		data, err := proto.Marshal(e)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		err = s.conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			// c.Logger().Error(err)
			s.Close()
		}

		//s.Done()
		s.SetLastTransfer(time.Now())
	}

	//s.Unlock()
}

func (s *Conn) Write(m *messages.Message) {
	if s == nil || s.GetTerminated() {
		return
	}

	//s.Add(1)
	s.out <- m
}

func (s *Conn) Close() {
	if s.GetTerminated() {
		return
	}

	// s.Lock()
	// defer s.Unlock()

	s.SetTerminated(true)

	s.conn.Close()

	go func() {
		//s.Wait()
		close(s.In)
		close(s.out)
	}()
}

func (s *Conn) SetLastTransfer(now time.Time) {
	// s.Lock()
	// defer s.Unlock()

	s.lastTransfer = now
}

func (s *Conn) SetPlayerId(i int32) {
	// s.Lock()
	// defer s.Unlock()

	s.playerId = i
}

func (s Conn) GetPlayerId() int32 {
	// s.Lock()
	// defer s.Unlock()

	return s.playerId
}

func (s Conn) GetTerminated() bool {
	// s.Lock()
	// defer s.Unlock()

	return s.terminated
}

func (s *Conn) SetTerminated(terminated bool) {
	// s.Lock()
	// defer s.Unlock()

	s.terminated = terminated
}
