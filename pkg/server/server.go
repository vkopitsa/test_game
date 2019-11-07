package server

import (
	"log"
	"server/messages"

	//"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// const (
// 	DefaultPort = 8080
// )

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
	s.e.Static("/", "./public")
	s.e.GET("/ws", func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

		if err == nil {
			serverConn := NewServerConn(conn, nil)
			serverConn.Player = i
			s.NewPlayers <- &IncomingPlayer{Conn: serverConn}

			i++
		}

		return err
	})
	s.e.Start(address)
}

func (s *server) StopListening() {
	s.e.Close()
}

func (s *server) accept() {
	for {
		np := <-s.NewPlayers

		p := NewPlayer("name", np.Conn)

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

		// log.Println("handleNewPlayer", e)

		switch e.GetType() {
		case messages.Message_HELLO:
			log.Println(e)
		// 	if _, ok := e.(*GameCommandListGames); ok {
		// 		var gl []*ListedGame

		// 		for _, g := range s.Games {
		// 			if g.Terminated {
		// 				continue
		// 			}

		// 			gl = append(gl, &ListedGame{ID: g.ID, Name: g.Name, Players: len(g.Players), MaxPlayers: g.MaxPlayers, SpeedLimit: g.SpeedLimit})
		// 		}

		// 		sort.Slice(gl, func(i, j int) bool {
		// 			if gl[i].Players == gl[j].Players {
		// 				return gl[i].Name < gl[j].Name
		// 			}

		// 			return gl[i].Players > gl[j].Players
		// 		})

		// 		pl.Write(&GameCommandListGames{Games: gl})
		// 	}
		case messages.Message_JOIN:
			// 	if p, ok := e.(*GameCommandJoinGame); ok {
			// 		pl.Name = Nickname(p.Name)

			g := s.FindGame(pl)
			if g == nil {
				return
			}

			// 		if p.Listing.Name == "" {
			// 			g.Logf(LogStandard, "Player %s joined %s", pl.Name, g.Name)
			// 		} else {
			// 			g.Logf(LogStandard, "Player %s created new game %s", pl.Name, g.Name)
			// 		}

			g.AddPlayer(pl)

			go s.handleGameCommands(pl, g)
			return

			// 		handled = true
			// 		return
			//}
		}
	}
}

func (s *server) FindGame(p Player) *Game {
	if len(s.Games) == 0 {
		s.Games[0] = NewGame()
	}
	return s.Games[0]
}

func (s *server) handle() {
	for {
		time.Sleep(1 * time.Minute)

		// s.Lock()
		s.removeTerminatedGames()
		// s.Unlock()
	}
}

func (s *server) removeTerminatedGames() {
	// for gameID, g := range s.Games {
	// 	if g != nil && !g.Terminated {
	// 		continue
	// 	}

	// 	delete(s.Games, gameID)
	// 	g = nil
	// }
}

func (s *server) handleGameCommands(pl Player, g *Game) {
	// var (
	// 	msgJSON []byte
	// 	err     error
	// )
	for e := range pl.GetIn() {
		// c := e.Command()
		// if (c != CommandPing && c != CommandPong && c != CommandUpdateMatrix) || g.LogLevel >= LogVerbose {
		// 	msgJSON, err = json.Marshal(e)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	g.Logf(LogStandard, "%d -> %s %s", e.Source(), e.Command(), msgJSON)
		// }

		// g.Lock()

		// log.Println("handleGameCommands", e)

		switch e.GetType() {
		case messages.Message_QUIT:
			//g.RemovePlayerL(p.SourcePlayer)
			log.Println("handleGameCommands quit", e)
		case messages.Message_DATA:
			data := &messages.Data{}
			err := proto.Unmarshal(e.GetData(), data)
			if err != nil {
				log.Println(err)
				continue
			}

			if _, ok := g.Players[e.GetPlayerId()]; ok {
				// newNick := Nickname(p.Nickname)
				// if newNick != "" && newNick != player.Name {
				// 	oldNick := player.Name
				// 	player.Name = newNick

				// 	g.Logf(LogStandard, "* %s is now known as %s", oldNick, newNick)
				g.WriteAll(&messages.Message{
					PlayerId: e.GetPlayerId(),
					Type:     messages.Message_DATA,
					Data:     e.GetData(),
				})
				// }
			}
		}
		// case *GameCommandDisconnect:
		// 	g.RemovePlayerL(p.SourcePlayer)
		// case *GameCommandMessage:
		// 	if player, ok := g.Players[p.SourcePlayer]; ok {
		// 		s.Logf("<%s> %s", player.Name, p.Message)

		// 		msg := strings.ReplaceAll(strings.TrimSpace(p.Message), "\n", "")
		// 		if msg != "" {
		// 			g.WriteAllL(&GameCommandMessage{Player: p.SourcePlayer, Message: msg})
		// 		}
		// 	}
		// case *GameCommandNickname:
		// 	if player, ok := g.Players[p.SourcePlayer]; ok {
		// 		newNick := Nickname(p.Nickname)
		// 		if newNick != "" && newNick != player.Name {
		// 			oldNick := player.Name
		// 			player.Name = newNick

		// 			g.Logf(LogStandard, "* %s is now known as %s", oldNick, newNick)
		// 			g.WriteAllL(&GameCommandNickname{Player: p.SourcePlayer, Nickname: newNick})
		// 		}
		// 	}
		// case *GameCommandUpdateMatrix:
		// 	if pl, ok := g.Players[p.SourcePlayer]; ok {
		// 		for _, m := range p.Matrixes {
		// 			pl.Matrix.Replace(m)

		// 			if g.SpeedLimit > 0 && m.Speed > g.SpeedLimit+5 && time.Since(g.TimeStarted) > 7*time.Second {
		// 				pl.Matrix.SetGameOver()

		// 				g.WriteMessage(fmt.Sprintf("%s went too fast and crashed", pl.Name))
		// 				g.WriteAllL(&GameCommandGameOver{Player: p.SourcePlayer})
		// 			}
		// 		}

		// 		m := pl.Matrix
		// 		spawn := m.SpawnLocation(m.P)
		// 		if m.P != nil && spawn.X >= 0 && spawn.Y >= 0 && m.P.X != spawn.X {
		// 			pl.Moved = time.Now()
		// 			pl.Idle = 0
		// 		}
		// 	}
		// case *GameCommandGameOver:
		// 	g.Players[p.SourcePlayer].Matrix.SetGameOver()

		// 	g.WriteMessage(fmt.Sprintf("%s was knocked out", g.Players[p.SourcePlayer].Name))
		// 	g.WriteAllL(&GameCommandGameOver{Player: p.SourcePlayer})
		// case *GameCommandSendGarbage:
		// 	leastGarbagePlayer := -1
		// 	leastGarbage := -1
		// 	for playerID, player := range g.Players {
		// 		if playerID == p.SourcePlayer || player.Matrix.GameOver {
		// 			continue
		// 		}

		// 		if leastGarbage == -1 || player.totalGarbageReceived < leastGarbage {
		// 			leastGarbagePlayer = playerID
		// 			leastGarbage = player.totalGarbageReceived
		// 		}
		// 	}

		// 	if leastGarbagePlayer != -1 {
		// 		g.Players[leastGarbagePlayer].totalGarbageReceived += p.Lines
		// 		g.Players[leastGarbagePlayer].pendingGarbage += p.Lines

		// 		g.Players[p.SourcePlayer].totalGarbageSent += p.Lines
		// 	}
		// case *GameCommandStats:
		// 	players := 0
		// 	games := 0

		// 	for _, g := range s.Games {
		// 		players += len(g.Players)
		// 		games++
		// 	}

		// 	g.Players[p.SourcePlayer].Write(&GameCommandStats{Created: s.created, Players: players, Games: games})
		// }

		// g.Unlock()
	}
}

// func NetworkAndAddress(address string) (string, string) {
// 	if !strings.Contains(address, `:`) {
// 		address = fmt.Sprintf("%s:%d", address, DefaultPort)
// 	}

// 	return "tcp", address
// }
