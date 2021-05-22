package fixture

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func FixtureReceiver(channel chan common.Event,
	fixture int,
	command common.Sequence,
	commandChannel chan common.Sequence,
	replyChannel chan common.Sequence,
	mySequenceNumber int,
	Pattens map[string]common.Patten,
	eventsForLauchpad chan common.ALight) {

	// Start the step counter so we know where we are in the sequence.
	stepCount := 0

	// Start the color counter.
	currentColor := 0

	fmt.Printf("Now Listening on channel %d\n", fixture)
	for {

		// Listen on this fixtures channel for the step events.
		<-channel

		step := Pattens[command.Patten.Name].Steps
		totalSteps := len(command.Patten.Steps)
		tolalColors := len(step[stepCount].Fixtures[fixture].Colors)

		R := step[stepCount].Fixtures[fixture].Colors[currentColor].R
		G := step[stepCount].Fixtures[fixture].Colors[currentColor].G
		B := step[stepCount].Fixtures[fixture].Colors[currentColor].B

		if currentColor <= tolalColors {
			currentColor++
		}
		// Fade up.
		if R > 0 || G > 0 || B > 0 {
			for green := 0; green <= step[stepCount].Fixtures[fixture].Colors[0].G; green++ {
				time.Sleep(command.CurrentSpeed / 4)
				e := common.ALight{X: fixture, Y: mySequenceNumber - 1, Brightness: 3, Red: R, Green: green, Blue: B}
				eventsForLauchpad <- e
			}
		}
		// Fade down.
		if R == 0 || G == 0 || B == 0 {
			time.Sleep(command.CurrentSpeed / 4)
			for green := step[stepCount].Fixtures[fixture].Colors[0].G; green >= 0; green-- {
				time.Sleep(command.CurrentSpeed / 4)
				e := common.ALight{X: fixture, Y: mySequenceNumber - 1, Brightness: 3, Red: R, Green: green, Blue: B}
				eventsForLauchpad <- e
			}
		}

		if currentColor == tolalColors {
			stepCount++
			currentColor = 0
		}

		if stepCount >= totalSteps {
			stepCount = 0
			currentColor = 0
		}
	}
}
