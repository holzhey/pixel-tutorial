package shader

import (
	"os"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

var uTime, uSpeed float32
var start time.Time

func DefineShader(shaderPath string, window *pixelgl.Window) {
	start = time.Now()
	uSpeed = 5.0
	window.Canvas().SetUniform("uTime", &uTime)
	window.Canvas().SetUniform("uSpeed", &uSpeed)
	window.Canvas().SetFragmentShader(loadShader(shaderPath))
}

func Increment() {
	uTime = float32(time.Since(start).Seconds())
}

func loadShader(shaderPath string) string {
	fs, err := os.ReadFile(shaderPath)
	if err != nil {
		panic(err)
	}
	return string(fs)
}
