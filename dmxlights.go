package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"os/signal"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"

	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/rakyll/launchpad/mk2"
)

const (
	full = 255
)

var sequenceSpeed int = 12
var fadeSpeed int
var savePreset bool
var selectedPatten = 0
var blackout bool = false

// main thread is used to get commands from the lauchpad.
func main() {

	var flashButtons [][]bool

	presetsStore := make(map[string]bool)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		presets.SavePresets(presetsStore)
		os.Exit(1)
	}()

	fmt.Println("Derek Lighting")

	fmt.Println("Loading Presets")
	presetsStore = presets.LoadPresets()
	fmt.Println("Loading Presets Done")

	// Setup DMX card.
	dmxController := dmx.NewDmXController()

	// Setup a connection to the launchpad.
	pad, err := mk2.Open()
	if err != nil {
		log.Fatalf("Error initializing launchpad: %v", err)
	}
	defer pad.Close()

	fmt.Println("Set Programmers Mode")
	pad.Program()

	// Create all the channels I need.
	commandChannels := []chan common.Command{}
	replyChannels := []chan common.Sequence{}
	soundTriggerChannels := []chan common.Command{}

	// Make channels for commands.
	for sequence := 0; sequence < 4; sequence++ {
		commandChannel := make(chan common.Command)
		commandChannels = append(commandChannels, commandChannel)
		replyChannel := make(chan common.Sequence)
		replyChannels = append(replyChannels, replyChannel)
		soundTriggerChannel := make(chan common.Command)
		soundTriggerChannels = append(soundTriggerChannels, soundTriggerChannel)
	}

	// Now add them all to a handy channels struct.
	channels := common.Channels{}
	channels.CommmandChannels = commandChannels
	channels.ReplyChannels = replyChannels
	channels.SoundTriggerChannels = soundTriggerChannels

	// Get a list of all the fixtures in the groups.
	fixturesConfig := fixture.LoadFixtures()

	// Create a channel to send events to the launchpad.
	eventsForLauchpad := make(chan common.ALight)

	// Now create a thread to handle those events.
	go launchpad.ListenAndSendToLaunchPad(eventsForLauchpad, pad)

	// Start off by turning off all of the Lights
	pad.Reset()

	// Create a channel to listen for buttons being pressed.
	buttonChannel := pad.Listen()

	// Build the default set of Pattens.
	pattens := patten.MakePatterns()

	sequence1 := sequence.CreateSequence("colors", 1, pattens, channels)
	sequence2 := sequence.CreateSequence("standard", 2, pattens, channels)
	sequence3 := sequence.CreateSequence("scanner", 3, pattens, channels)
	sequence4 := sequence.CreateSequence("colors", 4, pattens, channels)

	// Add Sequence to an array.
	sequences := []*common.Sequence{}
	sequences = append(sequences, &sequence1)
	sequences = append(sequences, &sequence2)
	sequences = append(sequences, &sequence3)
	sequences = append(sequences, &sequence4)

	// Create a sound trigger object and give it the sequences so it can access their configs.
	sound.NewSoundTrigger(sequences, channels)

	// Start threads for each sequence.
	go sequence.PlaySequence(sequence1, 1, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)
	go sequence.PlaySequence(sequence2, 2, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)
	go sequence.PlaySequence(sequence3, 3, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)
	go sequence.PlaySequence(sequence4, 4, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)

	// common.Light up any existing presets.
	presets.InitPresets(eventsForLauchpad, presetsStore)

	fmt.Println("Setup Presets Done")

	// Initialize a ten length slice of empty slices
	flashButtons = make([][]bool, 9)

	// Initialize those 10 empty slices
	for i := 0; i < 9; i++ {
		flashButtons[i] = make([]bool, 9)
	}

	// Light the logo blue.
	pad.Light(8, -1, 0, 0, 255)

	// Light the first sequence as the default selected.
	selectedSequence := 1
	common.SequenceSelect(eventsForLauchpad, selectedSequence)

	// Initialise the pattens.
	availablePatten := []string{}
	availablePatten = append(availablePatten, "standard")
	availablePatten = append(availablePatten, "inverted")
	availablePatten = append(availablePatten, "rgbchase")
	availablePatten = append(availablePatten, "pairs")
	availablePatten = append(availablePatten, "colors")

	for {
		select {

		case hit := <-buttonChannel:

			// Clear all the lights on the launchpad.
			if hit.X == 0 && hit.Y == -1 {
				launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, commandChannels)
				continue
			}

			// Save mode.
			if hit.X == 8 && hit.Y == 4 {
				if savePreset {
					savePreset = false
					flashButtons[8][4] = false
					continue
				}
				flashButtons[8][4] = true
				launchpad.FlashButton(presetsStore, pad, flashButtons, 8, 4, eventsForLauchpad, 1, 255, 0, 0)
				savePreset = true
			}

			// Ask all sequences for their current config and save in a file.
			if hit.X < 8 && (hit.Y > 3 && hit.Y < 7) {
				if savePreset {
					fmt.Printf("Write Config\n")
					presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] = true
					common.LightOn(eventsForLauchpad, common.ALight{
						X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
					fmt.Printf("Save Preset in X:%d Y:%d \n", hit.X, hit.Y)
					config.AskToSaveConfig(commandChannels, replyChannels, hit.X, hit.Y)
					savePreset = false
					flashButtons[8][4] = false
					presets.SavePresets(presetsStore)
					presets.InitPresets(eventsForLauchpad, presetsStore)
				} else {
					// Load config, but only if it exists in the presets map.
					if presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {
						fmt.Printf("Read Config:")
						fmt.Printf(" OK \n")
						launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, commandChannels)
						common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
						// Stop everything so that we start the recalled config in sync.
						// cmd := common.Command{
						// 	Stop: true,
						// }
						// sendCommandToAllSequence(selectedSequence, cmd, commandChannels)

						time.Sleep(850 * time.Millisecond)
						// Load the config.
						config.AskToLoadConfig(commandChannels, hit.X, hit.Y)

						// Reset flash buttons
						for x := 0; x < 9; x++ {
							for y := 0; y < 9; y++ {
								flashButtons[x][y] = false
							}
						}
						flashButtons[hit.X][hit.Y] = true
						launchpad.FlashButton(presetsStore, pad, flashButtons, hit.X, hit.Y, eventsForLauchpad, 1, 0, 255, 0)

						// cmd := common.Command{
						// 	Wakeup: true,
						// }
						// sendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					}
				}
			}

			// Increment Patten.
			if hit.X == 2 && hit.Y == 7 {
				if selectedPatten < 4 {
					selectedPatten = selectedPatten + 1
				}

				fmt.Printf("Selecting Patten %d selectedPatten %s\n", selectedPatten, availablePatten[selectedPatten])
				cmd := common.Command{
					Stop: true,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)

				cmd = common.Command{
					UpdatePatten: true,
					Patten: common.Patten{
						Name: availablePatten[selectedPatten],
					},
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)

				cmd = common.Command{
					Start: true,
					Speed: sequenceSpeed,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)

				continue
			}

			// Decrement Patten.
			if hit.X == 3 && hit.Y == 7 {
				if selectedPatten > 0 {
					selectedPatten = selectedPatten - 1
				}

				fmt.Printf("Selecting Patten %d selectedPatten %s\n", selectedPatten, availablePatten[selectedPatten])

				cmd := common.Command{
					UpdatePatten: true,
					Patten: common.Patten{
						Name: availablePatten[selectedPatten],
					},
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)
				continue
			}

			// Decrease speed of selected sequence.
			if hit.X == 0 && hit.Y == 7 {
				sequenceSpeed--
				if sequenceSpeed < 0 {
					sequenceSpeed = 1
				}
				cmd := common.Command{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)

				continue
			}

			// Increase speed of selected sequence.
			if hit.X == 1 && hit.Y == 7 {
				sequenceSpeed++
				if sequenceSpeed > 20 {
					sequenceSpeed = 20
				}
				cmd := common.Command{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)

				continue
			}

			// Set or unset music trigger on this sequence.
			if hit.X == 7 && hit.Y == -1 {
				if sequences[selectedSequence-1].MusicTrigger {
					sequences[selectedSequence-1].MusicTrigger = false
					cmd := common.Command{
						MusicTriggerOff: true,
						Speed:           sequenceSpeed,
					}
					sendCommandToSequence(selectedSequence, cmd, commandChannels)
				} else {
					sequences[selectedSequence-1].MusicTrigger = true
					cmd := common.Command{
						MusicTrigger: true,
						Speed:        sequenceSpeed,
					}
					sendCommandToSequence(selectedSequence, cmd, commandChannels)
				}
			}

			// Select sequence 1.
			if hit.X == 8 && hit.Y == 0 {
				selectedSequence = 1
				common.SequenceSelect(eventsForLauchpad, selectedSequence)
				continue
			}

			// Select sequence 2.
			if hit.X == 8 && hit.Y == 1 {
				selectedSequence = 2
				common.SequenceSelect(eventsForLauchpad, selectedSequence)
				continue
			}

			// Select sequence 3.
			if hit.X == 8 && hit.Y == 2 {
				selectedSequence = 3
				common.SequenceSelect(eventsForLauchpad, selectedSequence)
				continue
			}

			// Select sequence 4.
			if hit.X == 8 && hit.Y == 3 {
				selectedSequence = 4
				common.SequenceSelect(eventsForLauchpad, selectedSequence)
				continue
			}

			// Start sequence.
			if hit.X == 8 && hit.Y == 5 {
				cmd := common.Command{
					Start: true,
					Speed: sequenceSpeed,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)
			}
			// Stop sequence.
			if hit.X == 8 && hit.Y == 6 {
				cmd := common.Command{
					Stop:  true,
					Speed: sequenceSpeed,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)
			}

			// Fade time decrease.
			if hit.X == 4 && hit.Y == 7 {
				fadeSpeed--
				if fadeSpeed < 0 {
					fadeSpeed = 0
				}
				fmt.Printf("Fade down speed:%d", fadeSpeed)
				fadeTime := commands.SetFade(fadeSpeed)
				cmd := common.Command{
					UpdateFade: true,
					FadeTime:   fadeTime,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)

				continue
			}

			// Fade time increase.
			if hit.X == 5 && hit.Y == 7 {
				fadeSpeed++
				if fadeSpeed > 12 {
					fadeSpeed = 12
				}
				fmt.Printf("Fade up speed:%d", fadeSpeed)
				fadeTime := commands.SetSpeed(fadeSpeed)
				cmd := common.Command{
					UpdateFade: true,
					FadeTime:   fadeTime,
				}
				sendCommandToSequence(selectedSequence, cmd, commandChannels)

				continue
			}

			// Light the flash buttons based on current patten.
			var sequence common.Sequence
			sequence = common.Sequence{
				Patten: common.Patten{
					Name:  "colors",
					Steps: pattens["colors"].Steps,
				},
			}
			if hit.X >= 0 && hit.X < 8 {
				fmt.Printf("X=%d   Y=%d\n", hit.X, hit.Y)
				red := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].R
				green := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].G
				blue := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].B
				pan := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Pan
				tilt := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Tilt
				shutter := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Shutter
				gobo := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Gobo
				common.LightOn(eventsForLauchpad, common.ALight{
					X:          hit.X,
					Y:          hit.Y,
					Brightness: full,
					Red:        red,
					Green:      green,
					Blue:       blue,
				})
				dmx.Fixtures(hit.Y+1, dmxController, hit.X, red, green, blue, pan, tilt, shutter, gobo, fixturesConfig, blackout)
				time.Sleep(200 * time.Millisecond)
				common.LightOff(eventsForLauchpad, hit.X, hit.Y)
				dmx.Fixtures(hit.Y+1, dmxController, hit.X, 0, 0, 0, pan, tilt, shutter, gobo, fixturesConfig, blackout)

			}

			// Blackout button.
			if hit.X == 8 && hit.Y == 7 {
				if !blackout {
					fmt.Printf("B L A C K O U T \n")
					blackout = true
					cmd := common.Command{
						Blackout: true,
					}
					sendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: 255})
				} else {
					fmt.Printf("NORMAL\n")
					blackout = false

					cmd := common.Command{
						Normal: true,
					}
					sendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: 0})
				}
			}
		}
	}
}

func sendCommandToSequence(selectedSequence int, command common.Command, commandChannels []chan common.Command) {
	fmt.Printf("Send Command %+v\n", command)
	commandChannels[selectedSequence-1] <- command
}

func sendCommandToAllSequence(selectedSequence int, command common.Command, commandChannels []chan common.Command) {
	commandChannels[0] <- command
	commandChannels[1] <- command
	commandChannels[2] <- command
	commandChannels[3] <- command
}
