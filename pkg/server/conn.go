package server

import (
	"fmt"
	"log"
	"server/messages"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

const CommandQueueSize = 10

type Command int

type GameCommandTransport struct {
	Command Command
	Data    []byte
}

type GameCommandInterface interface {
	Command() Command
	Source() int
	SetSource(int)
}

type Conn struct {
	conn         *websocket.Conn
	LastTransfer time.Time
	Terminated   bool

	Player     int
	In         chan GameCommandInterface
	out        chan GameCommandInterface
	forwardOut chan GameCommandInterface

	*sync.WaitGroup
}

func NewServerConn(conn *websocket.Conn, forwardOut chan GameCommandInterface) *Conn {
	c := Conn{conn: conn, WaitGroup: new(sync.WaitGroup)}

	c.In = make(chan GameCommandInterface, CommandQueueSize)
	c.out = make(chan GameCommandInterface, CommandQueueSize)
	c.forwardOut = forwardOut

	c.LastTransfer = time.Now()

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
		messageType, msg, err := s.conn.ReadMessage()
		if err != nil || messageType != websocket.BinaryMessage {
			// c.Logger().Error(err)
			// fmt.Println("error", messageType != websocket.BinaryMessage)
			break
		}

		message := &messages.Message{}
		err = proto.Unmarshal(msg, message)
		if err != nil {
			break
		}

		switch message.Type {
		case messages.Message_HELLO:
			hello := &messages.Hello{}
			err = proto.Unmarshal(message.Data, hello)
			if err != nil {
				break
			}

			fmt.Println("hello", hello)
		default:
			log.Println("unknown serverconn command", message.Type.String())
			continue
		}

		// Write
		// err = ws.WriteMessage(websocket.BinaryMessage, data)
		// if err != nil {
		// 	// c.Logger().Error(err)
		// 	return nil
		// }
	}

	s.Close()
}

func (s *Conn) handleWrite() {
	if s.conn == nil {
		for range s.out {
			s.Done()
		}
		return
	}

	// var (
	// 	msg GameCommandTransport
	// 	j   []byte
	// 	err error
	// )

	// for e := range s.out {
	// 	if s.Terminated {
	// 		s.Done()
	// 		continue
	// 	}

	// 	msg = GameCommandTransport{Command: e.Command()}

	// 	msg.Data, err = json.Marshal(e)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	j, err = json.Marshal(msg)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	j = append(j, '\n')

	// 	err = s.conn.SetWriteDeadline(time.Now().Add(ConnTimeout))
	// 	if err != nil {
	// 		s.Close()
	// 	}

	// 	_, err = s.conn.Write(j)
	// 	if err != nil {
	// 		s.Close()
	// 	}

	// 	s.LastTransfer = time.Now()
	// 	//s.conn.SetWriteDeadline(time.Time{})
	// 	s.Done()
	// }
}

func (s *Conn) Close() {
	if s.Terminated {
		return
	}

	s.Terminated = true

	s.conn.Close()

	go func() {
		s.Wait()
		close(s.In)
		close(s.out)
	}()
}