package animation_test

import (
	"fmt"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	snap "github.com/midorimici/goimagesnapshot"
	animation "github.com/midorimici/goimagesnapshot/examples/ebitengine"
	"github.com/midorimici/goimagesnapshot/examples/ebitengine/testhelper"
)

func TestMain(m *testing.M) {
	// Draw should be called in Ebitengine game loop.
	testhelper.RunGame(m)
}

func Test_animation_Draw(t *testing.T) {
	const (
		sw = 320
		sh = 240
	)

	type test struct {
		updateCount int
	}
	tests := []test{}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{5 * i})
	}

	for _, tt := range tests {
		name := fmt.Sprintf("renders unchanged when updateCount = %d", tt.updateCount)
		t.Run(name, func(t *testing.T) {
			a := animation.New()

			for i := 0; i < tt.updateCount; i++ {
				if err := a.Update(); err != nil {
					t.Fatalf("animation.Update() error = %v", err)
				}
			}

			i := ebiten.NewImage(sw, sh)

			a.Draw(i)

			snap.Match(t, i, snap.WithName(fmt.Sprint(tt.updateCount)))
		})
	}
}
