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
	for _, seq := range sequences {
		seq <- cmd
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presetsStore[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				common.LightOn(eventsForLauchpad, common.ALight{X: x, Y: y, Brightness: 3, Red: 3, Green: 0, Blue: 0})
			}
		}
	}
}

// ListenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func ListenAndSendToLaunchPad(eventsForLauchpad chan common.ALight, pad *mk2.Launchpad) {
	var green int
	var red int
	var blue int

	for {
		event := <-eventsForLauchpad
		switch event.Green {
		case 0:
			green = 0
		case 1:
			green = 19
		case 2:
			green = 22
		case 3:
			green = 21
		}

		switch event.Red {
		case 0:
			red = 0
		case 1:
			red = 7
		case 2:
			red = 6
		case 3:
			red = 5
		}

		switch event.Blue {
		case 0:
			blue = 0
		case 1:
			blue = 37
		case 2:
			blue = 38
		case 3:
			blue = 79
		}

		pad.Light(event.X, event.Y, red+green+blue)
	}
}

func FlashButton(presetsStore map[string]bool, pad *mk2.Launchpad, flashButtons [][]bool, x int, y int, eventsForLauchpad chan common.ALight, seqNumber int, green int, red int, blue int) {
	go func(pad *mk2.Launchpad, x int, y int) {
		for {
			presets.InitPresets(eventsForLauchpad, presetsStore)
			// fmt.Printf("Flash X:%d Y:%d is %t\n", x, y, flashButtons[x][y])
			if !flashButtons[x][y] {
				break
			}
			event := common.ALight{X: x, Y: y, Brightness: 3, Red: red, Green: green, Blue: blue}
			eventsForLauchpad <- event

			time.Sleep(500 * time.Millisecond)
			event = common.ALight{X: x, Y: y, Brightness: 3, Red: 0, Green: 0, Blue: 0}
			eventsForLauchpad <- event
			time.Sleep(500 * time.Millisecond)
		}
	}(pad, x, y)
}
