package fixture

import (
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func FixtureReceiver(
	channel chan common.Event,
	fixture int,
	command common.Sequence,
	commandChannel chan common.Command,
	mySequenceNumber int,
	eventsForLauchpad chan common.ALight) {

	for {

		event := <-channel

		e := common.ALight{
			X:          fixture,
			Y:          mySequenceNumber - 1,
			Brightness: 3,
			Red:        event.Color.R,
			Green:      event.Color.G,
			Blue:       event.Color.B,
		}
		eventsForLauchpad <- e
	}
}
