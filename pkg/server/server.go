package server

import (
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

const (
	DefaultPort = 8080
)

const (
	LogStandard = iota
	LogDebug
	LogVerbose
)

var (
	upgrader = websocket.Upgrader{}
)

type server struct {
	logLevel int
	// listeners []net.Listener
	e          *echo.Echo
	NewPlayers chan *IncomingPlayer
}

type IncomingPlayer struct {
	Conn *Conn
}

type Server interface {
	Listen(address string)
	StopListening()
}

func New(logLevel int) Server {
	s := server{logLevel: logLevel}
	return &s
}

func (s *server) Listen(address string) {
	s.e = echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	s.e.Static("/", "../../public")
	s.e.GET("/ws", func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

		if err == nil {
			s.NewPlayers <- &IncomingPlayer{Conn: NewServerConn(conn, nil)}
		}

		return err
	})
	s.e.Start(address)
}

func (s *server) StopListening() {
	s.e.Close()
}

// func (s *server) hello(c echo.Context) error {
// 	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

// 	if err == nil {
// 		s.NewPlayers <- &IncomingPlayer{Conn: NewServerConn(conn, nil)}
// 	}

// 	return err
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// defer ws.Close()

// 	// msg := &messages.Message{
// 	// 	Id: proto.Int32(17),
// 	// 	Author: &messages.Message_Person{
// 	// 		Id:   proto.Int32(1),
// 	// 		Name: proto.String("othree"),
// 	// 	},
// 	// 	Text: proto.String("Hi, this is message."),
// 	// }

// 	// data, _ := proto.Marshal(msg)
// 	// fmt.Println(data)

// 	// for {

// 	// 	// Read
// 	// 	_, msg, err := ws.ReadMessage()
// 	// 	if err != nil {
// 	// 		// c.Logger().Error(err)
// 	// 		return nil
// 	// 	}

// 	// 	newTest := &messages.Message{}
// 	// 	_ = proto.Unmarshal(msg, newTest)

// 	// 	//fmt.Println("newTest", msg, newTest)

// 	// 	// Write
// 	// 	err = ws.WriteMessage(websocket.BinaryMessage, data)
// 	// 	if err != nil {
// 	// 		// c.Logger().Error(err)
// 	// 		return nil
// 	// 	}
// 	// }
// }

func NetworkAndAddress(address string) (string, string) {
	if !strings.Contains(address, `:`) {
		address = fmt.Sprintf("%s:%d", address, DefaultPort)
	}

	return "tcp", address
}
