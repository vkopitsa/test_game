package game

type game struct {
}

type Game interface {
}

func New() Game {
	g := game{}
	return &g
}
