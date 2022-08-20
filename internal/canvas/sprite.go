package canvas

import (
	"image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var gopherimg *pixel.Sprite

func LoadSprite(spritePath string) {
	f, err := os.Open(spritePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	pd := pixel.PictureDataFromImage(img)
	gopherimg = pixel.NewSprite(pd, pd.Bounds())
}

func DrawSprite(window *pixelgl.Window, angle float64) {
	mat := pixel.IM
	mat = mat.Moved(window.Bounds().Center())
	mat = mat.Rotated(window.Bounds().Center(), angle)
	gopherimg.Draw(window, mat)
}
