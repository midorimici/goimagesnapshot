package testhelper

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type game[T any] struct {
	tests            []T
	currentTestIndex int
	draw             func(tt T) func(screen *ebiten.Image)
	hasDone          bool
}

func RunGame[T any](t *testing.T, tests []T, draw func(tt T) func(screen *ebiten.Image)) {
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

	if err := ebiten.RunGame(g); err != nil {
		t.Errorf("ebiten.RunGame() error = %v", err)
	}
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
