package launchpad

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/rakyll/launchpad/mk3"
)

func ClearAll(pad *mk3.Launchpad, presetsStore map[string]bool, eventsForLauchpad chan common.ALight, sequences []chan common.Command) {
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
func ListenAndSendToLaunchPad(eventsForLauchpad chan common.ALight, pad *mk3.Launchpad) {

	for {
		event := <-eventsForLauchpad

		if event.Flash {
			pad.FlashLight(event.X, event.Y, event.OnColor, event.OffColor)
		} else {
			// For the math to work we need to convert our ints to floats and then back again.
			Red := ((float64(event.Red) / 2) / 100) * (float64(event.Brightness) / 2.55)
			Green := ((float64(event.Green) / 2) / 100) * (float64(event.Brightness) / 2.55)
			Blue := ((float64(event.Blue) / 2) / 100) * (float64(event.Brightness) / 2.55)
			pad.Light(event.X, event.Y, int(Red), int(Green), int(Blue))
		}
	}
}

func FlashLight(X int, Y int, onColor int, offColor int, eventsForLauchpad chan common.ALight) {

	// Now ask the fixture lamp to flash on the launch pad by sending an event.
	e := common.ALight{
		Flash:    true,
		X:        X,
		Y:        Y,
		OnColor:  onColor,
		OffColor: offColor,
	}
	eventsForLauchpad <- e
}

func LightLamp(Y, X, R, G, B, Master int, eventsForLauchpad chan common.ALight) {

	// Now trigger the fixture lamp on the launch pad by sending an event.
	e := common.ALight{
		X:          X,
		Y:          Y,
		Brightness: Master,
		Red:        R,
		Green:      G,
		Blue:       B,
	}
	eventsForLauchpad <- e
}
