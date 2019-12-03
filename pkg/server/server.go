package server

import (
	"log"
	"server/messages"

	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	// sync.RWMutex

	logLevel int

	In  chan *messages.Message
	Out chan *messages.Message

	Games map[int]*Game

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
	in := make(chan *messages.Message, CommandQueueSize)
	out := make(chan *messages.Message, CommandQueueSize)

	s := server{
		logLevel: logLevel,
		In:       in,
		Out:      out,

		Games: make(map[int]*Game),
	}

	s.NewPlayers = make(chan *IncomingPlayer, CommandQueueSize)

	go s.accept()
	go s.handle()

	return &s
}

func (s *server) Listen(address string) {
	s.e = echo.New()

	var i int32 = 0

	s.e.Use(middleware.Logger())
	// s.e.Use(middleware.Recover())
	s.e.Static("/", "./front/dist")
	s.e.GET("/ws", func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

		if err == nil {
			serverConn := NewServerConn(conn, nil)
			serverConn.SetPlayerId(i)
			s.NewPlayers <- &IncomingPlayer{Conn: serverConn}

			i++
		}

		return err
	})
	s.e.Start(address)
}

func (s *server) StopListening() {
	// s.Lock()
	// defer s.Unlock()
	s.e.Close()
}

func (s *server) accept() {
	for {
		np := <-s.NewPlayers

		// s.Lock()
		p := NewPlayer("name", np.Conn)
		// s.Unlock()

		go s.handleNewPlayer(p)
	}
}

func (s *server) handleNewPlayer(pl Player) {
	handled := false
	go func() {
		time.Sleep(10 * time.Second)
		if !handled {
			pl.Close()
		}
	}()

	for e := range pl.GetIn() {

		switch e.GetType() {
		case messages.Message_HELLO:
			log.Println(e)
		case messages.Message_JOIN:
			g := s.FindGame(pl)
			if g == nil {
				return
			}

			g.AddPlayer(pl)

			// send player id to front
			pl.Write(&messages.Message{
				PlayerId: pl.GetPlayerId(),
				Type:     messages.Message_JOINED,
			})

			go s.handleGameCommands(pl, g)
			return
		}
	}
}

func (s *server) FindGame(p Player) *Game {
	var game *Game
	if len(s.Games) == 0 {
		game = NewGame()
		game.Start()

		s.Games[0] = game
	} else {
		for _, g := range s.Games {
			if len(g.Players) < 50 {
				game = g
				break
			}
		}

		if game == nil {
			game = NewGame()
			game.Start()

			s.Games[len(s.Games)+1] = game
		}
	}

	return game
}

func (s *server) handle() {
	for {
		time.Sleep(1 * time.Second)

		// s.Lock()
		s.removeTerminatedPlayers()
		// s.Unlock()

		s.updateGameInfo()
	}
}

func (s *server) updateGameInfo() {
	for i := range s.Games {
		info := s.Games[i].GetInfo()

		playerInfos := []*messages.PlayerInfo{}
		for _, playerInfo := range info.PlayerInfos {
			playerInfos = append(playerInfos, &messages.PlayerInfo{
				Id:    playerInfo.Id,
				Score: playerInfo.Score,
			})
		}

		data, err := proto.Marshal(&messages.Info{
			Count:   info.Count,
			Players: playerInfos,
		})
		if err != nil {
			return
		}

		s.Games[i].WriteAll(&messages.Message{
			// PlayerId: pl.GetPlayerId(),
			Type: messages.Message_INFO,
			Data: data,
		})
	}
}

func (s *server) removeTerminatedPlayers() {
	for i := range s.Games {
		s.Games[i].TerminatedPlayers()

		if len(s.Games[i].Players) == 0 {
			s.Games[i].Stop()
			delete(s.Games, i)
		}
	}
}

func (s *server) handleGameCommands(pl Player, g *Game) {
	for e := range pl.GetIn() {
		switch e.GetType() {
		case messages.Message_QUIT:
			pl.Close()
		case messages.Message_COMMAND:
			cmd := &messages.Command{}
			err := proto.Unmarshal(e.GetData(), cmd)
			if err != nil {
				continue
			}

			pl.AddCommand(cmd)
		}
	}
}
