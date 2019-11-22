package server

import (
	"server/messages"
	"time"
)

type Position struct {
	y float64
	x float64
}

type player struct {
	Name string

	*Conn

	// Score int
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
}

type Player interface {
	Close()
	GetIn() chan *messages.Message
	Write(m *messages.Message)
	GetPlayerId() int32
	AddCommand(message *messages.Command)
	Tick(dt float64, worldWidth int64, worldHeight int64)
	GetPosition(dt float64) *Position
}

func NewPlayer(name string, conn *Conn) Player {
	if conn == nil {
		conn = &Conn{}
	}

	p := &player{
		Name:  name,
		Conn:  conn,
		Moved: time.Now(),
		// commands: make([]*messages.Command, 0, 100),
		// command:  &messages.Command{},
		// position: Position{},
		speed:  200,
		radius: 50,
	}

	return p
}

func (p *player) GetIn() chan *messages.Message {
	return p.Conn.In
}

func (p *player) AddCommand(message *messages.Command) {
	p.command = message
}

func (p *player) Close() {

}

func (p *player) Tick(dt float64, worldWidth int64, worldHeight int64) {

	if p.command == nil {
		return
	}

	if p.position == nil {
		p.position = &Position{}
	}

	// if (this.direction == Direction.Left)
	// this.x -= this.speed * dt;
	// if (this.direction == Direction.Up)
	// this.y -= this.speed * dt;
	// if (this.direction == Direction.Right)
	// this.x += this.speed * dt;
	// if (this.direction == Direction.Down)
	// this.y += this.speed * dt;
	if p.command.GetYv() != 0 {
		p.position.y = p.position.y + (dt * p.speed * p.command.GetYv())
	}
	if p.command.GetXv() != 0 {
		p.position.x = p.position.x + (dt * p.speed * p.command.GetXv())
	}

	// if (this.y + this.radius > worldHeight || this.y - this.radius < 0) {
	// 	this.direction = (this.y + this.radius) > worldHeight ? Direction.Up : Direction.Down;
	//   }

	//   if (this.x + this.radius > worldWidth || this.x - this.radius < 0) {
	// 	this.direction = (this.x + this.radius) > worldWidth ? Direction.Left : Direction.Right;
	//   }

	if p.position.y+p.radius > float64(worldHeight) || p.position.y-p.radius < -p.radius {
		// p.direction = (p.position.y + p.radius) > worldHeight ? Direction.Up : Direction.Down;

		// Up
		if (p.position.y + p.radius) > float64(worldHeight) {
			p.command.Yv = -1
		} else {
			// Down
			p.command.Yv = 1
		}
	}

	if p.position.x+p.radius > float64(worldWidth) || p.position.x-p.radius < -p.radius {
		// p.direction = (p.position.x + p.radius) > worldWidth ? Direction.Left : Direction.Right;

		// Left
		if (p.position.x + p.radius) > float64(worldWidth) {
			p.command.Xv = -1
		} else {
			// Right
			p.command.Xv = 1
		}
	}

	// fmt.Println(p.command, "command")
	// fmt.Println(p.position, "position")
	// fmt.Println((dt * p.speed), dt, "(dt * p.speed)")
}

func (p *player) GetPosition(dt float64) *Position {
	return p.position
}
