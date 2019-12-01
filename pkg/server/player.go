package server

import (
	"fmt"
	"image/color"
	"image/color/palette"
	"math"
	"math/rand"
	"server/messages"
	"sync"
	"time"
)

type Position struct {
	y float64
	x float64
}

type PlayerInfo struct {
	Id    int32
	Score int64
}

type player struct {
	// sync.RWMutex

	Name string

	*Conn

	score int64
	// Preview *mino.Matrix
	// Matrix  *mino.Matrix
	Moved time.Time     // Time of last piece move
	Idle  time.Duration // Time spent idling

	// pendingGarbage       int
	// totalGarbageSent     int
	// totalGarbageReceived int

	command *messages.Command

	position *Position

	speed  float64
	radius float64
	color  color.Color
}

type Player interface {
	Close()
	GetIn() chan *messages.Message
	Write(m *messages.Message)
	GetPlayerId() int32
	AddCommand(message *messages.Command)
	Tick(dt float64, worldWidth int64, worldHeight int64)
	GetPosition(dt float64) *Position
	SetPosition(pos *Position)
	GetCommand() (float64, float64)
	GetColor() string
	IsNear(other Player, playerNearValue float64) bool
	GetTerminated() bool
	GetScore() int64
	GetRadius() float64
	Overlaps(other Player) bool
	RevertDirection()
}

var mutex sync.Mutex

func NewPlayer(name string, conn *Conn) Player {
	mutex.Lock()
	defer mutex.Unlock()

	if conn == nil {
		conn = &Conn{}
	}

	p := &player{
		Name:     name,
		Conn:     conn,
		Moved:    time.Now(),
		speed:    200,
		radius:   50,
		color:    palette.WebSafe[rand.Intn(len(palette.WebSafe))],
		position: &Position{},
	}

	return p
}

func (p *player) GetIn() chan *messages.Message {
	return p.Conn.In
}

func (p *player) GetColor() string {
	//return p.color.RGBA()
	R, G, B, _ := p.color.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", uint8(R*255.0), uint8(G*255.0), uint8(B*255.0))
}

func (p *player) AddCommand(message *messages.Command) {
	// p.Lock()
	// defer p.Unlock()

	p.command = message
}

func (p *player) Close() {

}

func (p *player) Tick(dt float64, worldWidth int64, worldHeight int64) {
	// p.Lock()
	// defer p.Unlock()

	if p.command == nil {
		return
	}

	if p.command.GetYv() != 0 {
		p.position.y = p.position.y + (dt * p.speed * p.command.GetYv())
	}
	if p.command.GetXv() != 0 {
		p.position.x = p.position.x + (dt * p.speed * p.command.GetXv())
	}

	if p.position.y+p.radius > float64(worldHeight) || p.position.y-p.radius < -p.radius {
		// Up
		if (p.position.y + p.radius) > float64(worldHeight) {
			p.command.Yv = -1
		} else {
			// Down
			p.command.Yv = 1
		}
	}

	if p.position.x+p.radius > float64(worldWidth) || p.position.x-p.radius < -p.radius {
		// Left
		if (p.position.x + p.radius) > float64(worldWidth) {
			p.command.Xv = -1
		} else {
			// Right
			p.command.Xv = 1
		}
	}

	// score
	// p.score++
}

func (p *player) RevertDirection() {
	if p.command == nil {
		return
	}
	p.command.Yv = p.command.Yv * -1
	p.command.Xv = p.command.Xv * -1

	// score
	p.score++
}

func (p *player) GetPosition(dt float64) *Position {
	// p.Lock()
	// defer p.Unlock()

	return p.position
}

func (p *player) SetPosition(pos *Position) {
	p.position = pos
}

func (p *player) GetCommand() (float64, float64) {
	// p.Lock()
	// defer p.Unlock()

	return p.command.GetYv(), p.command.GetXv()
}

func (p *player) GetScore() int64 {
	return p.score
}

func (p *player) GetRadius() float64 {
	return p.radius
}

// Determines is a player is near other player.
func (p *player) IsNear(other Player, playerNearValue float64) bool {
	xdiff := math.Abs(p.GetPosition(0).x - other.GetPosition(0).x)
	ydiff := math.Abs(p.GetPosition(0).y - other.GetPosition(0).y)
	mdiff := math.Max(xdiff, ydiff)
	return mdiff < playerNearValue
}

func (p *player) Overlaps(other Player) bool {
	dx := p.GetPosition(0).x - other.GetPosition(0).x
	dy := p.GetPosition(0).y - other.GetPosition(0).y
	distance := math.Sqrt(dx*dx + dy*dy)

	return distance < p.GetRadius()+other.GetRadius()
}
