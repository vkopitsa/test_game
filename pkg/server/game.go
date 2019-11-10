package server

import (
	"server/messages"
	"time"
)

type Game struct {
	Started     bool
	TimeStarted time.Time
	Players     map[int32]Player
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
