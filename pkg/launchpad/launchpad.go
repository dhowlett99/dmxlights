package launchpad

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/rakyll/launchpad/mk2"
)

func ClearAll(pad *mk2.Launchpad, presetsStore map[string]bool, eventsForLauchpad chan common.ALight, sequences []chan common.Command) {
	fmt.Printf("C L E A R\n")
	pad.Reset()
	cmd := common.Command{
		Stop: true,
	}
	for _, sequence := range sequences {
		sequence <- cmd
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presetsStore[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				common.LightOn(eventsForLauchpad, common.ALight{X: x, Y: y, Brightness: 255, Red: 255, Green: 0, Blue: 0})
			}
		}
	}
}

// ListenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func ListenAndSendToLaunchPad(eventsForLauchpad chan common.ALight, pad *mk2.Launchpad) {

	for {
		event := <-eventsForLauchpad
		// For the math to work we need to convert our ints to floats and then back again.
		Red := ((float64(event.Red) / 2) / 100) * (float64(event.Brightness) / 2.55)
		Green := ((float64(event.Green) / 2) / 100) * (float64(event.Brightness) / 2.55)
		Blue := ((float64(event.Blue) / 2) / 100) * (float64(event.Brightness) / 2.55)
		pad.Light(event.X, event.Y, int(Red), int(Green), int(Blue))
	}
}

// FlashButton creates a thread which loops forever flash a position X Y
// until the flashButton flag for that location is cleared.
func FlashButton(presetsStore map[string]bool, pad *mk2.Launchpad, flashButtons [][]bool, x int, y int, eventsForLauchpad chan common.ALight, seqNumber int, green int, red int, blue int) {
	go func(pad *mk2.Launchpad, x int, y int) {
		for {
			presets.InitPresets(eventsForLauchpad, presetsStore)
			if !flashButtons[x][y] {
				break
			}
			event := common.ALight{X: x, Y: y, Brightness: 255, Red: red, Green: green, Blue: blue}
			eventsForLauchpad <- event

			time.Sleep(500 * time.Millisecond)
			event = common.ALight{X: x, Y: y, Brightness: 255, Red: 0, Green: 0, Blue: 0}
			eventsForLauchpad <- event
			time.Sleep(500 * time.Millisecond)
		}
	}(pad, x, y)
}

func LightLamp(X, Y, R, G, B int, eventsForLauchpad chan common.ALight) {

	// Now trigger the fixture lamp on the launch pad by sending an event.
	e := common.ALight{
		X:          Y,
		Y:          X - 1,
		Brightness: 255,
		Red:        R,
		Green:      G,
		Blue:       B,
	}
	eventsForLauchpad <- e
}
