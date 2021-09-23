package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"os/signal"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx/ft232"

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
var size int

//var color int
var savePreset bool
var selectedPatten = 0
var blackout bool = false

// main thread is used to get commands from the lauchpad.
func main() {

	var flashButtons [][]bool
	var functionButtons [][]bool
	var functionMode [][]bool
	var selectButtons [][]bool
	fadeSpeed = 11 // Default start at 50ms.
	var musicSequence bool

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

	sequence1 := sequence.CreateSequence("standard", 1, pattens, channels)
	sequence2 := sequence.CreateSequence("standard", 2, pattens, channels)
	sequence3 := sequence.CreateSequence("scanner", 3, pattens, channels)
	sequence4 := sequence.CreateSequence("standard", 4, pattens, channels)

	// Add Sequence to an array.
	sequences := []*common.Sequence{}
	sequences = append(sequences, &sequence1)
	sequences = append(sequences, &sequence2)
	sequences = append(sequences, &sequence3)
	sequences = append(sequences, &sequence4)

	// Create a sound trigger object and give it the sequences so it can access their configs.
	sound.NewSoundTrigger(sequences, channels)

	// Start threads for each sequence.
	go sequence.PlayNewSequence(sequence1, 1, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)
	go sequence.PlayNewSequence(sequence2, 2, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)
	go sequence.PlayNewSequence(sequence3, 3, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)
	go sequence.PlayNewSequence(sequence4, 4, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, channels)

	// common.Light up any existing presets.
	presets.InitPresets(eventsForLauchpad, presetsStore)

	// Light the function buttons at the top and bottom.
	common.ShowFunctionButtons(sequence1, 0, eventsForLauchpad, functionButtons)
	common.ShowFunctionButtons(sequence1, 8, eventsForLauchpad, functionButtons)

	fmt.Println("Setup Presets Done")

	// Initialize a ten length slice of empty slices for flash buttons.
	flashButtons = make([][]bool, 9)
	// Initialize those 10 empty flash button slices.
	for i := 0; i < 9; i++ {
		flashButtons[i] = make([]bool, 9)
	}

	// Initialize a ten length slice of empty slices for function buttons.
	functionButtons = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		functionButtons[i] = make([]bool, 9)
	}

	// Initialize a ten length slice of empty slices for select buttons.
	selectButtons = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		selectButtons[i] = make([]bool, 9)
	}

	// Initialize a ten length slice of empty slices for function mode state.
	functionMode = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		functionMode[i] = make([]bool, 9)
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
	availablePatten = append(availablePatten, "fade")

	// Clear the pad.
	allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)

	// Main loop reading commands from the Novation Launchpad.
	for {
		select {

		case hit := <-buttonChannel:

			// Clear all the lights on the launchpad.
			if hit.X == 0 && hit.Y == -1 {
				launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, commandChannels)
				allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)
				presets.ClearPresets(eventsForLauchpad, presetsStore, flashButtons)
				presets.InitPresets(eventsForLauchpad, presetsStore)
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
					// If its already set, then clear it.
					if presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {
						presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] = false
					} else {
						// Not already set then set it.
						fmt.Printf("Write Config\n")
						presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] = true
						common.LightOn(eventsForLauchpad, common.ALight{
							X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
						fmt.Printf("Save Preset in X:%d Y:%d \n", hit.X, hit.Y)
						config.AskToSaveConfig(commandChannels, replyChannels, hit.X, hit.Y)
						savePreset = false
						flashButtons[8][4] = false
					}
					presets.SavePresets(presetsStore)
					presets.ClearPresets(eventsForLauchpad, presetsStore, flashButtons)
					flashButtons[hit.X][hit.Y] = true
					launchpad.FlashButton(presetsStore, pad, flashButtons, hit.X, hit.Y, eventsForLauchpad, 1, 0, 255, 0)
				} else {
					// Load config, but only if it exists in the presets map.
					if presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {
						fmt.Printf("Read Config:")
						fmt.Printf(" OK \n")
						launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, commandChannels)
						common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
						// Stop everything so that we start the recalled config in sync.
						cmd := common.Command{
							Stop: true,
						}
						common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

						// Clear all the fixtures down ready for the next scene.
						allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)

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

						// Get a copy of the function button settings for all the sequences.
						for s := 1; s < len(sequences)+1; s++ {
							// Get an upto date copy of the sequence.
							cmd = common.Command{
								ReadConfig: true,
							}
							common.SendCommandToSequence(s, cmd, commandChannels)

							// Listen for the reply and set the newSequence with the values.
							newSequence := common.Sequence{}
							replyChannel := channels.ReplyChannels[s-1]
							newSequence = <-replyChannel
							fmt.Printf("found newSeq %d\n", newSequence.Number)

							// Make sure the music trigger is set.
							if newSequence.Functions[common.Function8_Music_Trigger].State {
								sequences[s-1].MusicTrigger = true
							} else {
								sequences[s-1].MusicTrigger = false
								sequenceSpeed = sequences[s-1].Speed // Put the speed back.
								cmd := common.Command{
									Speed:       sequenceSpeed,
									UpdateSpeed: true,
								}
								common.SendCommandToSequence(s, cmd, commandChannels)
							}
						}
					}
				}
			}

			// Increment Patten.
			if hit.X == 2 && hit.Y == 7 {
				if selectedPatten < 5 {
					selectedPatten = selectedPatten + 1
				}
				if selectedPatten > 5 {
					selectedPatten = 5
				}

				fmt.Printf("Selecting Patten %d selectedPatten %s\n", selectedPatten, availablePatten[selectedPatten])
				cmd := common.Command{
					Stop: true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				cmd = common.Command{
					UpdatePatten: true,
					Patten: common.Patten{
						Name: availablePatten[selectedPatten],
					},
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				cmd = common.Command{
					Stop:  true,
					Speed: sequenceSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				cmd = common.Command{
					Start: true,
					Speed: sequenceSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

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
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				cmd = common.Command{
					Stop:  true,
					Speed: sequenceSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				cmd = common.Command{
					Start: true,
					Speed: sequenceSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				continue
			}

			// Decrease speed of selected sequence.
			if hit.X == 0 && hit.Y == 7 {
				if !musicSequence {
					sequenceSpeed--
					if sequenceSpeed < 0 {
						sequenceSpeed = 1
					}
					cmd := common.Command{
						Speed:       sequenceSpeed,
						UpdateSpeed: true,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
					continue
				}
			}

			// Increase speed of selected sequence.
			if hit.X == 1 && hit.Y == 7 {
				if !musicSequence {
					sequenceSpeed++
					if sequenceSpeed > 20 {
						sequenceSpeed = 20
					}
					cmd := common.Command{
						Speed:       sequenceSpeed,
						UpdateSpeed: true,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
					continue
				}
			}

			// Set or unset music trigger on this sequence.
			if hit.X == 7 && hit.Y == -1 {
				if sequences[selectedSequence-1].MusicTrigger {
					musicSequence = false
					sequences[selectedSequence-1].MusicTrigger = false
					cmd := common.Command{
						MusicTriggerOff: true,
						Speed:           sequenceSpeed,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				} else {
					musicSequence = true
					sequences[selectedSequence-1].MusicTrigger = true
					cmd := common.Command{
						MusicTrigger: true,
						Speed:        sequenceSpeed,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				}
			}

			// Set or unset soft fade on this sequence.
			if hit.X == 6 && hit.Y == -1 {
				if sequences[selectedSequence-1].SoftFade {
					sequences[selectedSequence-1].SoftFade = false
					cmd := common.Command{
						SoftFadeOff: true,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				} else {
					sequences[selectedSequence-1].SoftFade = true
					cmd := common.Command{
						SoftFadeOn: true,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				}
			}

			// S E L E C T    S E Q U E N C E.
			// Select sequence 1.
			if hit.X == 8 && hit.Y == 0 {
				selectedSequence = 1
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, channels)
				continue
			}

			// Select sequence 2.
			if hit.X == 8 && hit.Y == 1 {
				selectedSequence = 2
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, channels)
				continue
			}

			// Select sequence 3.
			if hit.X == 8 && hit.Y == 2 {
				selectedSequence = 3
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, channels)
				continue
			}

			// Select sequence 4.
			if hit.X == 8 && hit.Y == 3 {
				selectedSequence = 4
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, channels)
				continue
			}

			// Start sequence.
			if hit.X == 8 && hit.Y == 5 {
				musicSequence = false
				cmd := common.Command{
					Start:        true,
					Speed:        sequenceSpeed,
					MusicTrigger: false,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}
			// Stop sequence.
			if hit.X == 8 && hit.Y == 6 {
				cmd := common.Command{
					Stop:  true,
					Speed: sequenceSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}

			// Size decrease.
			if hit.X == 4 && hit.Y == 7 {
				size--
				if size < 0 {
					size = 0
				}
				fmt.Printf("size:%d", size)

				// Send Update Fade Speed.
				cmd := common.Command{
					UpdateSize: true,
					Size:       size,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				continue
			}

			// Size increase.
			if hit.X == 5 && hit.Y == 7 {
				size++
				if size > 20 {
					size = 20
				}
				fmt.Printf("Size up :%d", size)

				// Send size update.
				cmd := common.Command{
					UpdateSize: true,
					Size:       size,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				continue
			}

			// Fade time decrease.
			if hit.X == 6 && hit.Y == 7 {
				fadeSpeed--
				if fadeSpeed < 0 {
					fadeSpeed = 0
				}
				fmt.Printf("Fade down speed:%d", fadeSpeed)
				// Send fade update command.
				cmd := common.Command{
					DecreaseFade: true,
					FadeSpeed:    fadeSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				continue
			}

			// Fade time increase.
			if hit.X == 7 && hit.Y == 7 {
				fadeSpeed++
				if fadeSpeed > 20 {
					fadeSpeed = 20
				}
				fmt.Printf("Fade up speed:%d", fadeSpeed)
				// Send fade update command.
				cmd := common.Command{
					IncreaseFade: true,
					FadeSpeed:    fadeSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				continue
			}

			// // Color decrease.
			// if hit.X == 6 && hit.Y == 7 {
			// 	color--
			// 	if color < 0 {
			// 		color = 0
			// 	}
			// 	fmt.Printf("color down:%d", color)
			// 	cmd := common.Command{
			// 		UpdateColor: true,
			// 		Color:       color,
			// 	}
			// 	sendCommandToSequence(selectedSequence, cmd, commandChannels)

			// 	continue
			// }

			// // Color increase.
			// if hit.X == 7 && hit.Y == 7 {
			// 	color++
			// 	if color > 5 {
			// 		color = 5
			// 	}
			// 	fmt.Printf("Color up :%d", color)
			// 	cmd := common.Command{
			// 		UpdateColor: true,
			// 		Color:       color,
			// 	}
			// 	sendCommandToSequence(selectedSequence, cmd, commandChannels)

			// 	continue
			// }

			// Function buttons
			if hit.X >= 0 && hit.X < 8 && functionMode[8][selectedSequence-1] {
				// get an upto date copy of the sequence.
				cmd := common.Command{
					ReadConfig: true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				newSequence := common.Sequence{}
				replyChannel := channels.ReplyChannels[selectedSequence-1]
				newSequence = <-replyChannel

				for _, f := range newSequence.Functions {
					fmt.Printf("f:%d state:%t\n", f.Number, f.State)
				}

				for _, functions := range newSequence.Functions {
					fmt.Printf("Y is %d  Func seq no is %d\n", hit.Y, functions.SequenceNumber)
					if hit.Y == functions.SequenceNumber {
						fmt.Printf("state is %t\n", newSequence.Functions[hit.X].State)
						if !newSequence.Functions[hit.X].State {
							newSequence.Functions[hit.X].State = true
							break
						}
						if newSequence.Functions[hit.X].State {
							newSequence.Functions[hit.X].State = false
							break
						}
					}
				}
				common.ShowFunctionButtons(newSequence, selectedSequence, eventsForLauchpad, functionButtons)
				// Send update functions command.
				cmd = common.Command{
					UpdateFunctions: true,
					Functions:       newSequence.Functions,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				// Make sure the music trigger is set.
				if newSequence.Functions[common.Function8_Music_Trigger].State {
					sequences[selectedSequence-1].MusicTrigger = true
				} else {
					sequences[selectedSequence-1].MusicTrigger = false
					sequenceSpeed = 14 //Default to 25 Millisecond
					cmd := common.Command{
						Speed:       sequenceSpeed,
						UpdateSpeed: true,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				}

				continue
			}

			// FLASH BUTTONS - Light the flash buttons based on current patten.
			var sequence common.Sequence
			sequence = common.Sequence{
				Patten: common.Patten{
					Name:  "colors",
					Steps: pattens["colors"].Steps,
				},
			}
			if hit.X >= 0 && hit.X < 8 && !functionMode[8][selectedSequence-1] {
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
				fixture.MapFixtures(hit.Y+1, dmxController, hit.X, red, green, blue, pan, tilt, shutter, gobo, fixturesConfig, blackout, 255, 255)
				time.Sleep(200 * time.Millisecond)
				common.LightOff(eventsForLauchpad, hit.X, hit.Y)
				fixture.MapFixtures(hit.Y+1, dmxController, hit.X, 0, 0, 0, pan, tilt, shutter, gobo, fixturesConfig, blackout, 255, 255)
			}

			// Blackout button.
			if hit.X == 8 && hit.Y == 7 {
				if !blackout {
					fmt.Printf("B L A C K O U T \n")
					blackout = true
					cmd := common.Command{
						Blackout: true,
					}
					common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: 255})
				} else {
					fmt.Printf("NORMAL\n")
					blackout = false

					cmd := common.Command{
						Normal: true,
					}
					common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: 0})
				}
			}
		}
	}
}

func allFixturesOff(eventsForLauchpad chan common.ALight, dmxController ft232.DMXController, fixturesConfig *fixture.Fixtures) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			common.LightOff(eventsForLauchpad, x, y)
			fixture.MapFixtures(y, dmxController, x, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0)
		}
	}
}
