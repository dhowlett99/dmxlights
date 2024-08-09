package sequence

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func startStatic(mySequenceNumber int, sequence *common.Sequence, channels common.Channels, fixtureStepChannels []chan common.FixtureCommand) {

	if debug {
		fmt.Printf("%d: Sequence Static mode StaticFadeUpOnce %t\n", mySequenceNumber, sequence.StaticFadeUpOnce)
	}

	sequence.Static = true
	sequence.PlayStaticOnce = false

	// Turn off any music trigger for this sequence.
	sequence.MusicTrigger = false
	// this.Functions[common.Function8_Music_Trigger].State = false
	channels.SoundTriggers[mySequenceNumber].State = false

	// Now send the Fade up command to the fixture.
	if sequence.StaticFadeUpOnce {
		if debug {
			fmt.Printf("%d: Sequence Fade up static \n", mySequenceNumber)
		}
		// Prepare a message to be sent to the fixtures in the sequence.
		command := common.FixtureCommand{
			Master:          sequence.Master,
			Blackout:        sequence.Blackout,
			Type:            sequence.Type,
			Label:           sequence.Label,
			SequenceNumber:  sequence.Number,
			RGBStaticFadeUp: true,
			RGBFade:         sequence.RGBFade,
			RGBStaticColors: sequence.StaticColors,
			Hidden:          false,
			StrobeSpeed:     sequence.StrobeSpeed,
			Strobe:          sequence.Strobe,
			ScannerChaser:   sequence.ScannerChaser,
		}

		// Now tell all the fixtures what they need to do.
		sendToAllFixtures(fixtureStepChannels, command)

		// Done fading for this static scene only reset when we set a static scene again.
		sequence.StaticFadeUpOnce = false
	} else {
		// else just play the static scene.
		if debug {
			fmt.Printf("%d: Sequence Turn on static \n", mySequenceNumber)
		}
		command := common.FixtureCommand{
			Master:          sequence.Master,
			Blackout:        sequence.Blackout,
			Type:            sequence.Type,
			Label:           sequence.Label,
			SequenceNumber:  sequence.Number,
			Hidden:          false,
			StrobeSpeed:     sequence.StrobeSpeed,
			Strobe:          sequence.Strobe,
			ScannerChaser:   sequence.ScannerChaser,
			RGBStaticOn:     true,
			RGBStaticColors: sequence.StaticColors,
		}

		// Now tell all the fixtures what they need to do.
		sendToAllFixtures(fixtureStepChannels, command)
	}
	sequence.PlayStaticOnce = false
}

func stopStatic(mySequenceNumber int, sequence *common.Sequence, channels common.Channels, fixtureStepChannels []chan common.FixtureCommand) {
	if debug {
		fmt.Printf("%d: Sequence RGB Static mode OFF Type %s Label %s \n", mySequenceNumber, sequence.Type, sequence.Label)
	}

	channels.SoundTriggers[mySequenceNumber].State = false

	// Prepare a message to be sent to the fixtures in the sequence.
	command := common.FixtureCommand{
		Master:          sequence.Master,
		Blackout:        sequence.Blackout,
		Type:            sequence.Type,
		Label:           sequence.Label,
		SequenceNumber:  sequence.Number,
		Hidden:          sequence.Hidden,
		StrobeSpeed:     sequence.StrobeSpeed,
		Strobe:          sequence.Strobe,
		ScannerChaser:   sequence.ScannerChaser,
		RGBStaticOff:    true,
		RGBStaticColors: sequence.StaticColors,
		RGBFade:         sequence.RGBFade,
	}

	// Now tell all the fixtures what they need to do.
	sendToAllFixtures(fixtureStepChannels, command)
	sequence.PlayStaticOnce = false
}
