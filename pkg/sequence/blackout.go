package sequence

import "github.com/dhowlett99/dmxlights/pkg/common"

func blackout(fixtureStepChannels []chan common.FixtureCommand) {

	command := common.FixtureCommand{
		Type:      "lastColor",
		LastColor: common.Black,
	}
	sendToAllFixtures(fixtureStepChannels, command)
}
