package server

import (
	"log"
	"math/rand"
	"server/messages"
	"sort"
	"time"

	"github.com/gogo/protobuf/proto"
)

type GameCommand struct {
	Player  Player
	Message *messages.Message
}

type GameInfo struct {
	Count       int64
	PlayerInfos []*PlayerInfo
}

type Game struct {
	Started     bool
	TimeStarted time.Time
	Players     map[int32]Player
	// mutexPlayers sync.RWMutex

	Command chan *GameCommand

	// 5000, 3000
	WorldWidth  int64
	WorldHeight int64

	PlayerNearValue float64

	quit chan bool
}

func NewGame() *Game {
	g := &Game{
		quit:        make(chan bool),
		Players:     make(map[int32]Player),
		WorldWidth:  5000,
		WorldHeight: 3000,
		// mutexPlayers: sync.RWMutex{},
		PlayerNearValue: 1500,
	}

	return g
}

func (g *Game) Start() {
	g.Started = true
	g.TimeStarted = time.Now()

	go g.gameLoop()
}

func (g *Game) Stop() {
	g.Started = false

	g.quit <- true
}

func (g *Game) gameLoop() {
	tick := time.Tick((1000 / 30) * time.Millisecond)

	last := time.Now()
	var dt float64

	for {
		select {
		case <-g.quit:
			// stop
			return
		case <-tick:
			dt = float64(time.Since(last).Microseconds()/1000) / 1000.0

			g.processInputs(dt)

			last = time.Now()
		}
	}
}

func (g *Game) processInputs(dt float64) {
	// g.mutexPlayers.Lock()
	// defer g.mutexPlayers.Unlock()

	for _, p := range g.Players {
		//g.mutexPlayers.Unlock()

		p.Tick(dt, g.WorldWidth, g.WorldHeight)

		// Collision check
		for _, otherPlayer := range g.Players {
			if p.GetPlayerId() != otherPlayer.GetPlayerId() && p.IsNear(otherPlayer, g.PlayerNearValue) {
				if p.GetTerminated() == false && otherPlayer.GetTerminated() == false && p.Overlaps(otherPlayer) {
					p.RevertDirection()
					otherPlayer.RevertDirection()
				}
			}
		}

		position := p.GetPosition(dt)
		if position == nil {
			//g.mutexPlayers.Lock()
			continue
		}

		yv, xy := p.GetCommand().GetYv(), p.GetCommand().GetXv()

		d := &messages.Data{
			Y:     position.y,
			X:     position.x,
			Yv:    yv,
			Xv:    xy,
			Color: p.GetColor(),
			Time:  p.GetCommand().GetTime(),
			Delta: dt,
		}

		// reset command identification
		if p.GetCommand() != nil {
			p.GetCommand().Time = 0
		}

		data, err := proto.Marshal(d)
		if err != nil {
			log.Println("marshaling error: ", err)
			//g.mutexPlayers.Lock()
			return
		}

		//g.mutexPlayers.Lock()
		g.WriteNear(p, &messages.Message{
			PlayerId: p.GetPlayerId(),
			Type:     messages.Message_DATA,
			Data:     data,
		})

		//g.mutexPlayers.Unlock()
	}
}

func (g *Game) sendGameState() {

}

func (g *Game) WriteAll(m *messages.Message) {
	// g.mutexPlayers.Lock()
	// defer g.mutexPlayers.Unlock()

	for i := range g.Players {
		g.Players[i].Write(m)
	}
}

func (g *Game) WriteNear(p Player, m *messages.Message) {
	for i := range g.Players {
		if p.IsNear(g.Players[i], g.PlayerNearValue) {
			g.Players[i].Write(m)
		}
	}
}

func (g *Game) AddPlayer(p Player) {
	p.SetPosition(&Position{
		x: float64(rand.Int63n(g.WorldWidth)),
		y: float64(rand.Int63n(g.WorldHeight)),
	})
	g.Players[p.GetPlayerId()] = p
}

func (g *Game) TerminatedPlayers() {
	for i := range g.Players {
		if g.Players[i].GetTerminated() {
			g.WriteAll(&messages.Message{
				PlayerId: g.Players[i].GetPlayerId(),
				Type:     messages.Message_QUIT,
			})

			delete(g.Players, i)
		}
	}
}

func (g *Game) GetInfo() *GameInfo {
	infoPlayers := []*PlayerInfo{}
	for i := range g.Players {
		if g.Players[i].GetScore() == 0 {
			continue
		}
		infoPlayers = append(infoPlayers, &PlayerInfo{
			Id:    g.Players[i].GetPlayerId(),
			Score: g.Players[i].GetScore(),
		})
	}

	sort.Slice(infoPlayers, func(i, j int) bool {
		return infoPlayers[i].Score > infoPlayers[j].Score
	})

	if len(infoPlayers) > 5 {
		infoPlayers = infoPlayers[0:5]
	}

	return &GameInfo{
		Count:       int64(len(g.Players)),
		PlayerInfos: infoPlayers,
	}
}
