package testhelper

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidth  = 320
	screenHeight = 240
)

type TestGame interface {
	Run() error
}

type game[T any] struct {
	tests            []T
	currentTestIndex int
	draw             func(tt T) func(screen *ebiten.Image)
	hasDone          bool
}

func NewGame[T any](tests []T, draw func(tt T) func(screen *ebiten.Image)) TestGame {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	g := &game[T]{tests: tests}
	tc := len(tests)

	if tc == 0 {
		g.hasDone = true
	}

	drawFn := func(tt T) func(screen *ebiten.Image) {
		return func(screen *ebiten.Image) {
			draw(tt)(screen)
			g.currentTestIndex++
			if g.currentTestIndex >= tc {
				g.hasDone = true
			}
		}
	}

	g.draw = drawFn

	return g
}

func NewGameWithSingleCase(draw func(screen *ebiten.Image)) TestGame {
	return NewGame([]struct{}{{}}, func(struct{}) func(screen *ebiten.Image) { return draw })
}

func (g *game[T]) Update() error {
	if g.hasDone {
		return ebiten.Termination
	}

	return nil
}

func (g *game[T]) Draw(screen *ebiten.Image) {
	t := g.tests[g.currentTestIndex]
	g.draw(t)(screen)
}

func (g *game[T]) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func (g *game[T]) Run() error {
	return ebiten.RunGame(g)
}
