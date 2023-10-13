package animation_test

import (
	"fmt"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"

	snap "github.com/midorimici/goimagesnapshot"
	animation "github.com/midorimici/goimagesnapshot/examples/ebitengine"
	"github.com/midorimici/goimagesnapshot/examples/ebitengine/testhelper"
)

func Test_animation_Draw(t *testing.T) {
	// Draw should be called in Ebitengine game loop,
	// but image constructor functions such as ebiten.NewImage panics if they are called after ebiten.RunGame finishes.
	// In order to avoid that, you may need to add t.Parallel() to execute Draw tests
	// after all other tests in the package.
	// t.Parallel()

	type test struct {
		updateCount int
	}
	tests := []test{}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{5 * i})
	}

	testhelper.RunGame(t, tests, func(tt test) func(screen *ebiten.Image) {
		a := animation.New()

		for i := 0; i < tt.updateCount; i++ {
			if err := a.Update(); err != nil {
				t.Errorf("animation.Update() error = %v", err)
			}
		}

		return func(screen *ebiten.Image) {
			name := fmt.Sprintf("renders unchanged when updateCount = %d", tt.updateCount)
			t.Run(name, func(t *testing.T) {
				a.Draw(screen)

				snap.Match(t, screen, snap.WithName(fmt.Sprint(tt.updateCount)))
			})
		}
	})
}
