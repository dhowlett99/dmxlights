package fixture

import (
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func FixtureReceiver(channel chan common.Event,
	fixture int,
	command common.Sequence,
	commandChannel chan common.Sequence,
	mySequenceNumber int,
	eventsForLauchpad chan common.ALight) {

	for {

		event := <-channel

		if event.Fadeup {
			// Fade up.
			// for green := 0; green <= event.Color.G; green++ {
			// 	time.Sleep(event.FadeTime)
			e := common.ALight{
				X:          fixture,
				Y:          mySequenceNumber - 1,
				Brightness: 3,
				Red:        event.Color.B,
				Green:      event.Color.G,
				Blue:       event.Color.B,
			}
			eventsForLauchpad <- e
			// }
		}

		if event.Fadedown {
			for green := event.LastColor.G; green >= 0; green-- {
				time.Sleep(event.FadeTime)
				e := common.ALight{
					X:          fixture,
					Y:          mySequenceNumber - 1,
					Brightness: 3,
					Red:        event.Color.R,
					Green:      green,
					Blue:       event.Color.B}
				eventsForLauchpad <- e
			}
		}
	}
}
