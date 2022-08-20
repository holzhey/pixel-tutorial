package main

import (
	"fmt"
	"image/png"
	"math"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var gopherimg *pixel.Sprite

var uTime, uSpeed float32

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(100, 500), basicAtlas)
	basicTxt.Color = colornames.Red

	fmt.Fprintln(basicTxt, "Hello, text!")
	fmt.Fprintln(basicTxt, "I support multiple lines!")
	fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")

	f, err := os.Open("./gopher.png")
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

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	fs, err := os.ReadFile("./shader.glsl")
	if err != nil {
		panic(err)
	}

	win.Canvas().SetUniform("uTime", &uTime)
	win.Canvas().SetUniform("uSpeed", &uSpeed)
	uSpeed = 5.0
	win.Canvas().SetFragmentShader(string(fs))

	start := time.Now()
	angle := math.Pi

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 4))
		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center())
		mat = mat.Rotated(win.Bounds().Center(), angle)
		gopherimg.Draw(win, mat)
		uTime = float32(time.Since(start).Seconds())
		angle += 0.01
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
