package server

import (
	"log"
	"server/messages"
	"time"

	"github.com/gogo/protobuf/proto"
)

type GameCommand struct {
	Player  Player
	Message *messages.Message
}

type Game struct {
	Started     bool
	TimeStarted time.Time
	Players     map[int32]Player

	Command chan *GameCommand

	// 5000, 3000
	WorldWidth  int64
	WorldHeight int64
}

func NewGame() *Game {
	// minos, err := mino.Generate(rank)
	// if err != nil {
	// 	return nil, err
	// }

	g := &Game{
		// Name:       "netris",
		// Rank:       rank,
		// Minos:      minos,
		// nextPlayer: 1,
		Players: make(map[int32]Player),
		// Event:      make(chan interface{}, CommandQueueSize),
		// draw:       draw,
		// logger:     logger,
		// Mutex:      new(sync.Mutex)
		WorldWidth:  5000,
		WorldHeight: 3000,
	}

	// if out != nil {
	// 	g.out = out
	// } else {
	// 	g.LocalPlayer = PlayerHost
	// 	g.out = func(commandInterface GameCommandInterface) {
	// 		// Do nothing
	// 	}
	// }

	// g.FallTime = 850 * time.Millisecond

	// go g.handleDropTerminatedPlayers()

	return g
}

func (g *Game) Start() {
	g.Started = true
	g.TimeStarted = time.Now()

	go g.gameLoop()
}

func (g *Game) gameLoop() {
	tick := time.Tick((1000 / 30) * time.Millisecond)

	last := time.Now()
	var dt float64

	for {
		select {
		case <-tick:
			dt = float64(time.Since(last).Microseconds()/1000) / 1000.0
			// var now = Date.now();
			// var dt = (now - this.lastTime) / 1000.0;
			// log.Println("FPS", dt)
			// _ = dt

			g.processInputs(dt)
			// g.sendGameState(dt)

			last = time.Now()
		}
	}
}

func (g *Game) processInputs(dt float64) {
	for _, p := range g.Players {
		p.Tick(dt, g.WorldWidth, g.WorldHeight)

		position := p.GetPosition(dt)
		if position == nil {
			continue
		}

		d := &messages.Data{
			Y: position.y,
			X: position.x,
		}
		data, err := proto.Marshal(d)
		if err != nil {
			log.Println("marshaling error: ", err)
			return
		}

		g.WriteAll(&messages.Message{
			PlayerId: p.GetPlayerId(),
			Type:     messages.Message_DATA,
			Data:     data,
		})
	}
}

func (g *Game) sendGameState() {

}

func (g *Game) WriteAll(m *messages.Message) {
	for i := range g.Players {
		g.Players[i].Write(m)
	}
}

func (g *Game) AddPlayer(p Player) {
	// if p.Player == PlayerUnknown {
	// 	if g.LocalPlayer != PlayerHost {
	// 		return
	// 	}

	// 	p.Player = g.nextPlayer
	// 	g.nextPlayer++
	// }

	g.Players[p.GetPlayerId()] = p

	// TODO Verify rank-2 is valid for all playable rank previews
	// p.Preview = mino.NewMatrix(g.Rank, g.Rank-2, 0, 1, g.Event, g.draw, mino.MatrixPreview)
	// p.Preview.PlayerName = p.Name

	// p.Matrix = mino.NewMatrix(10, 20, 4, 1, g.Event, g.draw, mino.MatrixStandard)
	// p.Matrix.PlayerName = p.Name

	// if g.Started {
	// 	p.Matrix.GameOver = true
	// }

	// if g.LocalPlayer == PlayerHost {
	// 	p.Write(&GameCommandJoinGame{PlayerID: p.Player})

	// 	var players = make(map[int]string)
	// 	for _, player := range g.Players {
	// 		players[player.Player] = player.Name
	// 	}

	// 	g.WriteAllL(&GameCommandUpdateGame{Players: players})

	// 	if g.Started {
	// 		p.Write(&GameCommandStartGame{Seed: g.Seed, Started: g.Started})
	// 	}

	// 	if len(g.Players) > 1 {
	// 		g.WriteMessage(fmt.Sprintf("%s has joined the game", p.Name))
	// 	}
	// }
}
