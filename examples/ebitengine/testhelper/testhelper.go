package testhelper

import (
	"log"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	m *testing.M
}

func (g *game) Update() error {
	g.m.Run()
	return ebiten.Termination
}

func (g *game) Draw(screen *ebiten.Image) {}

func (g *game) Layout(int, int) (int, int) {
	return 1, 1
}

func RunGame(m *testing.M) {
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowSize(1, 1)
	ebiten.SetWindowPosition(-10, -10)

	g := &game{m: m}

	op := &ebiten.RunGameOptions{
		InitUnfocused:     true,
		ScreenTransparent: true,
		SkipTaskbar:       true,
	}
	if err := ebiten.RunGameWithOptions(g, op); err != nil {
		log.Fatal(err)
	}
}
