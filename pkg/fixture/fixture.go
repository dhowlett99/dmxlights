package fixture

import (
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

	var green int

	saveColor := make(map[int]common.Color)

	// Start the step counter so we know where we are in the sequence.
	stepCount := 0

	// Start the color counter.
	currentColor := 0

	//fmt.Printf("Now Listening on channel %d\n", fixture)
	for {

		<-channel

		// for issue := 0; issue < fixture; issue++ {
		// 	<-channel
		// }

		// if fixture == 0 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 1 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 2 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 3 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 4 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 5 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 6 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 7 {
		// 	for issue := 0; issue < fixture; issue++ {
		// 		<-channel
		// 	}
		// }
		// }
		// }
		// if fixture == 5 {
		// 	for issue := 0; issue < fixture*2; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 6 {
		// 	for issue := 0; issue < fixture*2; issue++ {
		// 		<-channel
		// 	}
		// }
		// if fixture == 7 {
		// 	for issue := 0; issue < fixture*2; issue++ {
		// 		<-channel
		// 	}
		// }

		//fmt.Printf("My fixture %d\n", fixture)

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
			for green = 0; green <= step[stepCount].Fixtures[fixture].Colors[0].G; green++ {
				time.Sleep(command.CurrentSpeed / 4)
				e := common.ALight{X: fixture, Y: mySequenceNumber - 1, Brightness: 3, Red: R, Green: green, Blue: B}
				eventsForLauchpad <- e
			}
			// Save your previos colors.
			saveColor[fixture] = common.Color{R: R, G: green, B: B}
			time.Sleep(command.CurrentSpeed / 4)
		}
		// Fade down.
		if saveColor[fixture].R > 0 || saveColor[fixture].G > 0 || saveColor[fixture].B > 0 {
			time.Sleep(command.CurrentSpeed / 4)
			for green = saveColor[fixture].G; green >= 0; green-- {
				time.Sleep(command.CurrentSpeed / 4)
				e := common.ALight{X: fixture, Y: mySequenceNumber - 1, Brightness: 3, Red: R, Green: green, Blue: B}
				eventsForLauchpad <- e
			}
			saveColor[fixture] = common.Color{R: R, G: green, B: B}
			time.Sleep(command.CurrentSpeed / 4)
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
