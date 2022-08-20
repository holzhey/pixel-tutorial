package canvas

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func GetWindowConfig(title string) pixelgl.WindowConfig {
	return pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
}

func OpenWindow(config pixelgl.WindowConfig) *pixelgl.Window {
	window, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}
	window.SetSmooth(true)
	return window
}
