package server

import (
	"server/messages"
	"time"
)

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
}

type Player interface {
	Close()
	GetIn() chan *messages.Message
}

func NewPlayer(name string, conn *Conn) Player {
	if conn == nil {
		conn = &Conn{}
	}

	p := &player{Name: name, Conn: conn, Moved: time.Now()}

	return p
}

func (p *player) GetIn() chan *messages.Message {
	return p.Conn.In
}

func (p *player) Close() {

}
