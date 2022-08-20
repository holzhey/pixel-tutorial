package sound

import (
	"math/rand"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

type Noise struct{}

func (n Noise) Stream(samples [][2]float64) (num int, ok bool) {
	for i := range samples {
		samples[i][0] = rand.Float64()*2 - 1
		samples[i][1] = rand.Float64()*2 - 1
	}
	return len(samples), true
}

func (n Noise) Err() error {
	return nil
}

func GetNoiseGenerator() beep.Streamer {
	return Noise{}
}

func GetSineGenerator(sampleRate int) beep.Streamer {
	sr := beep.SampleRate(sampleRate)
	speaker.Init(sr, 4100)
	sine, err := generators.SinTone(sr, 1800)
	if err != nil {
		panic(err)
	}
	two := sr.N(2 * time.Second)

	sounds := []beep.Streamer{
		beep.Take(two, sine),
	}

	return beep.Seq(sounds...)
}

func Play(streamer beep.Streamer) {
	speaker.Play(streamer)
}
