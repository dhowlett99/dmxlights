package fixture

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

// newMiniSequencer is a simple sequencer which can be attached to a switch and a fixture to allow simple effects.
func newMiniSequencer(fixtureName string, switchNumber int, switchPosition int, action Action, dmxController *ft232.DMXController, fixturesConfig *Fixtures,
	switchChannels map[int]common.SwitchChannel, soundConfig *sound.SoundConfig,
	blackout bool, master int, dmxInterfacePresent bool) {

	switchName := fmt.Sprintf("switch%d", switchNumber)
	fixture := findFixtureByName(fixtureName, fixturesConfig)

	mySequenceNumber := fixture.Group - 1
	myFixtureNumber := fixture.Number - 1

	cfg := getConfig(action)

	if debug {
		fmt.Printf("Action %+v\n", action)
	}

	if action.Mode == "Off" {
		if debug {
			fmt.Printf("Stop mini sequence for switch number %d\n", switchNumber)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(switchChannels, switchNumber, switchPosition, false, blackout, master)

		// Disable this mini sequencer with the sound service.
		// Use the switch number as the unique sequence name.
		soundConfig.DisableSoundTrigger(switchName)

		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		return
	}

	if action.Mode == "Static" {
		if debug {
			fmt.Printf("Static mini sequence for switch number %d\n", switchNumber)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(switchChannels, switchNumber, switchPosition, false, blackout, master)

		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		color, err := common.GetRGBColorByName(action.Colors[0])
		if err != nil {
			fmt.Printf("error %d\n", err)
		}
		MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, color.R, color.G, color.B, 0, 0, 0, 0, 0, 0, cfg.RotateSpeed, cfg.Music, cfg.Program, 0, nil, fixturesConfig, blackout, master, master, cfg.Strobe, dmxInterfacePresent)
		return
	}

	if action.Mode == "Chase" {

		if debug {
			fmt.Printf("Chase mini sequence for switch number %d\n", switchNumber)
		}

		// Don't stop this mini sequencer if there's one running already.
		// Unless we are changing switch positions.
		if getSwitchState(switchChannels, switchNumber) &&
			switchChannels[switchNumber].SwitchPosition == switchPosition {
			setSwitchState(switchChannels, switchNumber, switchPosition, true, blackout, master)
			return
		}

		// Remember that we have started this mini sequencer.
		setSwitchState(switchChannels, switchNumber, switchPosition, true, blackout, master)

		// DeRegister this mini sequencer with the sound service.
		// Use the switch number as the unique sequence name.
		soundConfig.DisableSoundTrigger(switchName)

		// Turn off the fixture.
		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Register this mini sequencer with the sound service.
		if cfg.MusicTrigger {
			soundConfig.EnableSoundTrigger(switchName)
		}

		// Stop any left over sequence left over for this switch.
		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		sequence := common.Sequence{
			ScannerInvert: false,
			RGBInvert:     false,
			Bounce:        false,
			ScannerChase:  false,
			RGBShift:      1,
		}
		sequence.Pattern = pattern.MakeSingleFixtureChase(cfg.Colors)
		sequence.Steps = sequence.Pattern.Steps
		sequence.NumberFixtures = 1
		// Calculate fade curve values.
		slopeOn, slopeOff := common.CalculateFadeValues(cfg.Fade, cfg.Size)
		// Calulate positions for each RGB fixture.
		optimisation := false
		scannerState := map[int]common.ScannerState{
			0: {
				Enabled: true,
			},
			1: {
				Enabled: true,
			},
			2: {
				Enabled: true,
			},
			3: {
				Enabled: true,
			},
			4: {
				Enabled: true,
			},
		}

		sequence.RGBPositions, sequence.NumberSteps = position.CalculatePositions(sequence, slopeOn, slopeOff, optimisation, scannerState)

		var rotateCounter int

		go func() {

			if cfg.Rotatable {
				// Thread to run the rotator
				go func(switchNumber int, switchChannels map[int]common.SwitchChannel) {

					for {
						rotateChannel, err := FindChannel("Rotate", myFixtureNumber, mySequenceNumber, fixturesConfig)
						if err != nil {
							fmt.Printf("rotator: %s,", err)
							return
						}
						masterChannel, err := FindChannel("Master", myFixtureNumber, mySequenceNumber, fixturesConfig)
						if err != nil {
							fmt.Printf("rotator: %s,", err)
							return
						}

						select {
						case <-switchChannels[switchNumber].StopRotate:
							time.Sleep(1 * time.Millisecond)
							setChannel(fixture.Address+int16(rotateChannel), byte(0), dmxController, dmxInterfacePresent)
							return
						case <-switchChannels[switchNumber].KeepRotateAlive:
							time.Sleep(1 * time.Millisecond)
							continue
						case <-time.After(1500 * time.Millisecond):
							setChannel(fixture.Address+int16(rotateChannel), byte(0), dmxController, dmxInterfacePresent)
							time.Sleep(250 * time.Millisecond)
							setChannel(fixture.Address+int16(masterChannel), byte(0), dmxController, dmxInterfacePresent)
						}
					}
				}(switchNumber, switchChannels)
			}

			// Wait for rotator thread to start.
			time.Sleep(100 * time.Millisecond)

			for {

				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < sequence.NumberSteps; step++ {

					blackout = switchChannels[switchNumber].Blackout
					master = switchChannels[switchNumber].Master

					if cfg.Rotatable {
						select {
						case switchChannels[switchNumber].KeepRotateAlive <- true:
						case <-time.After(10 * time.Millisecond):
						}
					}

					if rotateCounter > 500 {
						rotateCounter = 1
					}

					if rotateCounter < 128 {
						cfg.RotateSpeed = 127
					} else {
						cfg.RotateSpeed = 128
					}

					if debug {
						fmt.Printf("switch:%d waiting for beat on %d with speed %d\n", switchNumber, switchNumber+10, cfg.Speed)
						fmt.Printf("switch:%d speed %d\n", switchNumber, cfg.Speed)
					}

					// This is were we wait for a beat or a time out equivalent to the speed.
					select {
					case <-soundConfig.SoundTriggers[switchNumber+3].Channel:
					case <-switchChannels[switchNumber].Stop:
						if cfg.Rotatable {
							switchChannels[switchNumber].StopRotate <- true
						}
						MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixturesConfig, blackout, master, master, cfg.Strobe, dmxInterfacePresent)
						return
					case <-time.After(cfg.Speed):
					}

					// Play out fixture to DMX channels.
					position := sequence.RGBPositions[step]

					fixtures := position.Fixtures

					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
						fixture := fixtures[fixtureNumber]
						for _, color := range fixture.Colors {
							MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, color.R, color.G, color.B, color.W, 0, 0, 0, 0, 0, cfg.RotateSpeed, 0, 0, 0, nil, fixturesConfig, blackout, master, master, cfg.Strobe, dmxInterfacePresent)
						}
					}

					rotateCounter++
				}
			}
		}()
	}
}
