// Modified based on https://github.com/hajimehoshi/ebiten/blob/main/examples/animation/main.go

package animation

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

type animation struct {
	runnerImage *ebiten.Image
	count       int
}

func New() ebiten.Game {
	a := &animation{}

	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	a.runnerImage = ebiten.NewImageFromImage(img)

	return a
}

func (a *animation) Update() error {
	a.count++
	return nil
}

func (a *animation) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (a.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(a.runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (a *animation) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
