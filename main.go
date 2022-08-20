package main

import (
	"math"
	"pixel/hello/internal/canvas"
	"pixel/hello/internal/shader"
	"pixel/hello/internal/sound"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const shaderOcean = "./assets/shader/ocean.glsl"
const spriteGopher = "./assets/image/gopher.png"

func run() {
	window := canvas.OpenWindow(canvas.GetWindowConfig("Rotating Gohper"))
	shader.DefineShader(shaderOcean, window)
	canvas.LoadSprite(spriteGopher)
	text := canvas.CreateText("Just some text")
	angle := math.Pi

	sound.Play(sound.GetSineGenerator(44100))

	for !window.Closed() {
		window.Clear(colornames.Skyblue)
		text.Draw(window, pixel.IM.Scaled(text.Orig, 4))
		canvas.DrawSprite(window, angle)
		angle += 0.01
		window.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
