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

	var functionButtons [][]bool
	var functionMode [][]bool
	var selectButtons [][]bool
	var staticLamps [][]bool
	fadeSpeed = 11 // Default start at 50ms.

	// Make an empty presets store.
	presetsStore := make(map[string]bool)

	// Save the presets on exit.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		presets.SavePresets(presetsStore)
		os.Exit(1)
	}()

	fmt.Println("DMX Lighting")
	presetsStore = presets.LoadPresets()

	// Setup DMX card.
	dmxController := dmx.NewDmXController()

	// Setup a connection to the launchpad.
	pad, err := mk2.Open()
	if err != nil {
		log.Fatalf("error initializing launchpad: %v", err)
	}
	defer pad.Close()

	// We need to be in programmers mode to use the launchpad.
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
	sequenceChannels := common.Channels{}
	sequenceChannels.CommmandChannels = commandChannels
	sequenceChannels.ReplyChannels = replyChannels
	sequenceChannels.SoundTriggerChannels = soundTriggerChannels

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

	sequence1 := sequence.CreateSequence("standard", 0, pattens, sequenceChannels)
	sequence2 := sequence.CreateSequence("standard", 1, pattens, sequenceChannels)
	sequence3 := sequence.CreateSequence("scanner", 2, pattens, sequenceChannels)
	sequence4 := sequence.CreateSequence("standard", 3, pattens, sequenceChannels)

	// Add Sequence to an array.
	sequences := []*common.Sequence{}
	sequences = append(sequences, &sequence1)
	sequences = append(sequences, &sequence2)
	sequences = append(sequences, &sequence3)
	sequences = append(sequences, &sequence4)

	staticButtons := []common.StaticColorButtons{}
	staticButton1 := common.StaticColorButtons{}
	staticButton2 := common.StaticColorButtons{}
	staticButton3 := common.StaticColorButtons{}
	staticButton4 := common.StaticColorButtons{}

	staticButtons = append(staticButtons, staticButton1)
	staticButtons = append(staticButtons, staticButton2)
	staticButtons = append(staticButtons, staticButton3)
	staticButtons = append(staticButtons, staticButton4)

	// Create a sound trigger object and give it the sequences so it can access their configs.
	sound.NewSoundTrigger(sequences, sequenceChannels)

	// Start threads for each sequence.
	go sequence.PlayNewSequence(sequence1, 0, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels)
	go sequence.PlayNewSequence(sequence2, 1, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels)
	go sequence.PlayNewSequence(sequence3, 2, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels)
	go sequence.PlayNewSequence(sequence4, 3, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels)

	// common.Light up any existing presets.
	presets.InitPresets(eventsForLauchpad, presetsStore)

	// Light the function buttons at the top and bottom.
	//common.ShowFunctionButtons(sequence1, 0, eventsForLauchpad, functionButtons)
	common.ShowFunctionButtons(sequence1, 7, eventsForLauchpad, functionButtons)

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

	// Initialize a ten length slice of empty slices for static lamps.
	staticLamps = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		staticLamps[i] = make([]bool, 9)
	}

	// Light the logo blue.
	pad.Light(8, -1, 0, 0, 255)

	// Light the clear button purple.
	common.LightOn(eventsForLauchpad, common.ALight{X: 0, Y: -1, Brightness: 255, Red: 200, Green: 0, Blue: 255})

	// Light the static color buttons.
	common.LightOn(eventsForLauchpad, common.ALight{X: 1, Y: -1, Brightness: 255, Red: 255, Green: 0, Blue: 0})
	common.LightOn(eventsForLauchpad, common.ALight{X: 2, Y: -1, Brightness: 255, Red: 0, Green: 255, Blue: 0})
	common.LightOn(eventsForLauchpad, common.ALight{X: 3, Y: -1, Brightness: 255, Red: 0, Green: 0, Blue: 255})

	// Light top functions.
	common.LightOn(eventsForLauchpad, common.ALight{X: 4, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 5, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 6, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 7, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})

	// Light the save, start, stop and blackout buttons.
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 5, Brightness: full, Red: 255, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 6, Brightness: full, Red: 255, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 7, Brightness: full, Red: 255, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 8, Brightness: full, Red: 255, Green: 255, Blue: 255})

	// Light the first sequence as the default selected.
	selectedSequence := 0
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
				presets.ClearPresets(eventsForLauchpad, presetsStore)
				presets.InitPresets(eventsForLauchpad, presetsStore)
				// Make sure we stop all sequences.
				cmd := common.Command{
					Stop: true,
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
				continue
			}

			// Save mode.
			if hit.X == 8 && hit.Y == 4 {
				if savePreset {
					savePreset = false
					presets.InitPresets(eventsForLauchpad, presetsStore)
					//launchpad.FlashLight(8, 4, 0, 0, eventsForLauchpad)
					common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 4, Brightness: full, Red: 255, Green: 255, Blue: 255})
					continue
				}
				presets.InitPresets(eventsForLauchpad, presetsStore)
				launchpad.FlashLight(8, 4, 0x03, 0x5f, eventsForLauchpad)
				savePreset = true
			}

			// Ask all sequences for their current config and save in a file.
			if hit.X < 8 && (hit.Y > 3 && hit.Y < 7) {
				if savePreset {
					presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] = true
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
					config.AskToSaveConfig(commandChannels, replyChannels, hit.X, hit.Y)
					savePreset = false
					launchpad.FlashLight(8, 4, 0, 0, eventsForLauchpad) // turn off the save button from flashing.
					presets.SavePresets(presetsStore)
					presets.ClearPresets(eventsForLauchpad, presetsStore)
					presets.InitPresets(eventsForLauchpad, presetsStore)
					launchpad.FlashLight(hit.X, hit.Y, 0x0d, 0x78, eventsForLauchpad)
				} else {
					// Load config, but only if it exists in the presets map.
					if presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {
						launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, commandChannels)
						common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
						// Stop everything so that we start the recalled config in sync.
						cmd := common.Command{
							Stop: true,
						}
						common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

						// Clear all the fixtures down ready for the next scene.
						allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)

						// Wait for all the sequences to stop.
						time.Sleep(850 * time.Millisecond)

						// Load the config.
						config.AskToLoadConfig(commandChannels, hit.X, hit.Y)

						// Restore the blackout setting.
						cmd = common.Command{
							Blackout: blackout,
						}
						common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

						presets.InitPresets(eventsForLauchpad, presetsStore)
						launchpad.FlashLight(hit.X, hit.Y, 0x0d, 0x78, eventsForLauchpad)

						// Get a copy of the function button settings for all the sequences.
						for sequence := 0; sequence < len(sequences); sequence++ {
							// Get an upto date copy of the sequence.
							cmd = common.Command{
								ReadConfig: true,
							}
							common.SendCommandToSequence(sequence, cmd, commandChannels)

							// Listen for the reply and set the newSequence with the values.
							newSequence := common.Sequence{}
							replyChannel := sequenceChannels.ReplyChannels[sequence]
							newSequence = <-replyChannel

							// Make sure the music trigger is set.
							checkMusicTrigger(sequence, newSequence, sequences, commandChannels)

							// Keep track of the new run setting.
							sequences[selectedSequence].Run = newSequence.Run

							// Keep track of bounce setting.
							sequences[selectedSequence].Bounce = newSequence.Bounce

							// Make sure Static is set correctly
							if newSequence.Functions[common.Function6_Static].State {
								sequences[sequence].Static = newSequence.Functions[common.Function6_Static].State
								cmd = common.Command{
									UpdateStatic: true,
									Static:       newSequence.Functions[common.Function6_Static].State,
								}
								common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
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
				if !sequences[selectedSequence].MusicTrigger {
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
				if !sequences[selectedSequence].MusicTrigger {
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

			// S E L E C T    S E Q U E N C E.
			// Select sequence 1.
			if hit.X == 8 && hit.Y == 0 {
				selectedSequence = 0
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, sequenceChannels)
				continue
			}

			// Select sequence 2.
			if hit.X == 8 && hit.Y == 1 {
				selectedSequence = 1
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, sequenceChannels)
				continue
			}

			// Select sequence 3.
			if hit.X == 8 && hit.Y == 2 {
				selectedSequence = 2
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, sequenceChannels)
				continue
			}

			// Select sequence 4.
			if hit.X == 8 && hit.Y == 3 {
				selectedSequence = 3
				common.HandleSelect(sequences, selectedSequence, hit.X, hit.Y, eventsForLauchpad,
					selectButtons, functionButtons, functionMode, commandChannels, sequenceChannels)
				continue
			}

			// Start sequence.
			if hit.X == 8 && hit.Y == 5 {
				sequences[selectedSequence].MusicTrigger = false
				cmd := common.Command{
					Start:          true,
					Speed:          sequenceSpeed,
					MusicTriggerOn: false,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
				time.Sleep(100 * time.Millisecond)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
			}
			// Stop sequence.
			if hit.X == 8 && hit.Y == 6 {
				cmd := common.Command{
					Stop:  true,
					Speed: sequenceSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
				time.Sleep(100 * time.Millisecond)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
			}

			// Size decrease.
			if hit.X == 4 && hit.Y == 7 {
				size--
				if size < 0 {
					size = 0
				}
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
				// Send fade update command.
				cmd := common.Command{
					IncreaseFade: true,
					FadeSpeed:    fadeSpeed,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				continue
			}

			// Function buttons
			if hit.X >= 0 && hit.X < 8 && functionMode[8][selectedSequence] {
				// Get an upto date copy of the sequence.
				cmd := common.Command{
					ReadConfig: true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				// Create a temporary sequence.
				newSequence := common.Sequence{}
				replyChannel := sequenceChannels.ReplyChannels[selectedSequence]

				// Wait for sequence.
				newSequence = <-replyChannel

				// We've pushed a function key, this is where we set the value inside the temporary sequence.
				for _, functions := range newSequence.Functions {
					if hit.Y == functions.SequenceNumber {
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

				// Send update functions command. This sets the temporary representation of
				// the function keys in the real sequence.
				cmd = common.Command{
					UpdateFunctions: true,
					Functions:       newSequence.Functions,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				// Make sure the music trigger is set.
				checkMusicTrigger(selectedSequence, newSequence, sequences, commandChannels)

				// Keep track of the new run setting.
				sequences[selectedSequence].Run = newSequence.Run

				// Keep track of bounce setting.
				sequences[selectedSequence].Bounce = newSequence.Bounce

				// If we're unsetting static, then drop any presets as they would no longer apply.
				if !newSequence.Functions[common.Function6_Static].State {
					presets.ClearPresets(eventsForLauchpad, presetsStore)
					presets.InitPresets(eventsForLauchpad, presetsStore)
				}
				// Always make sure Static flag is set correctly
				cmd = common.Command{
					UpdateStatic: true,
					Static:       newSequence.Functions[common.Function6_Static].State,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				// And locally.
				sequences[selectedSequence].Functions[common.Function6_Static].State = newSequence.Functions[common.Function6_Static].State

				// If we are setting static then stop all the sequences and clear the launchpad.
				if newSequence.Functions[common.Function6_Static].State {
					cmd = common.Command{
						Stop: true,
					}
					common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
					allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)
				}

				// Light the correct function key.
				common.ShowFunctionButtons(newSequence, selectedSequence, eventsForLauchpad, functionButtons)
			}

			// FLASH BUTTONS - Briefly light (flash) the fixtures based on current patten.
			if hit.X >= 0 && hit.X < 8 && !functionMode[8][selectedSequence] && hit.Y >= 0 && hit.Y < 4 &&
				!sequences[selectedSequence].Functions[common.Function6_Static].State {

				flashSequence := common.Sequence{
					Patten: common.Patten{
						Name:  "colors",
						Steps: pattens["colors"].Steps,
					},
				}

				red := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].R
				green := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].G
				blue := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].B
				pan := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Pan
				tilt := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Tilt
				shutter := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Shutter
				gobo := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Gobo

				common.LightOn(eventsForLauchpad, common.ALight{
					X:          hit.X,
					Y:          hit.Y,
					Brightness: full,
					Red:        red,
					Green:      green,
					Blue:       blue,
				})
				fixture.MapFixtures(hit.Y, dmxController, hit.X, red, green, blue, pan, tilt, shutter, gobo, fixturesConfig, blackout, 255, 255)
				time.Sleep(200 * time.Millisecond)
				common.LightOff(eventsForLauchpad, hit.X, hit.Y)
				fixture.MapFixtures(hit.Y, dmxController, hit.X, 0, 0, 0, pan, tilt, shutter, gobo, fixturesConfig, blackout, 255, 255)
			}

			// C H O O S E   S T A T I C    C O L O R
			if hit.X == 1 && hit.Y == -1 {
				if staticButtons[selectedSequence].Color.R > 254 {
					staticButtons[selectedSequence].Color.R = 0
				} else {
					staticButtons[selectedSequence].Color.R = staticButtons[selectedSequence].Color.R + 10
				}
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: staticButtons[selectedSequence].Color.R, Green: 0, Blue: 0})
				updateStaticLamps(selectedSequence, staticButtons, staticLamps, commandChannels)
			}

			if hit.X == 2 && hit.Y == -1 {
				if staticButtons[selectedSequence].Color.G > 254 {
					staticButtons[selectedSequence].Color.G = 0
				} else {
					staticButtons[selectedSequence].Color.G = staticButtons[selectedSequence].Color.G + 10
				}
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: staticButtons[selectedSequence].Color.G, Blue: 0})
				updateStaticLamps(selectedSequence, staticButtons, staticLamps, commandChannels)
			}

			if hit.X == 3 && hit.Y == -1 {

				if staticButtons[selectedSequence].Color.B > 254 {
					staticButtons[selectedSequence].Color.B = 0
				} else {
					staticButtons[selectedSequence].Color.B = staticButtons[selectedSequence].Color.B + 10
				}
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: staticButtons[selectedSequence].Color.B})
				updateStaticLamps(selectedSequence, staticButtons, staticLamps, commandChannels)
			}

			// S E T    S T A T I C   C O L O R
			if hit.X >= 0 && hit.X < 8 && !functionMode[8][selectedSequence] && hit.Y != -1 &&
				sequences[selectedSequence].Functions[common.Function6_Static].State {

				// Remember which color we are setting in this sequence.
				red := staticButtons[selectedSequence].Color.R
				green := staticButtons[selectedSequence].Color.G
				blue := staticButtons[selectedSequence].Color.B

				// Static is set to true in the functions and this key is set to
				// the selected color.
				cmd := common.Command{
					UpdateStaticColor: true,
					Static:            true,
					StaticLamp:        hit.X,
					StaticColor:       common.Color{R: red, G: green, B: blue},
				}

				// Toggle the state of the lamp.
				if staticLamps[hit.X][hit.Y] {
					staticLamps[hit.X][hit.Y] = false
					// Turn the lamp off
					cmd = common.Command{
						UpdateStaticColor: true,
						Static:            true,
						StaticLamp:        hit.X,
						StaticColor:       common.Color{R: 0, G: 0, B: 0},
					}
				} else {
					// Remember which static lamp we just set.
					staticLamps[hit.X][hit.Y] = true
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}

			// B L A C K O U T   B U T T O N.
			if hit.X == 8 && hit.Y == 7 {
				if !blackout {
					blackout = true
					cmd := common.Command{
						Blackout: true,
					}
					common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: 255})
				} else {
					blackout = false
					cmd := common.Command{
						Normal: true,
					}
					common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
				}
			}
		}
	}
}

func checkMusicTrigger(selectedSequence int, newSequence common.Sequence, sequences []*common.Sequence, commandChannels []chan common.Command) {

	// Make sure the music trigger is set.
	if newSequence.Functions[common.Function8_Music_Trigger].State {
		sequences[selectedSequence].MusicTrigger = true
		cmd := common.Command{
			MusicTriggerOn: true,
		}
		common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
	} else {
		sequences[selectedSequence].MusicTrigger = false
		sequenceSpeed = 14 //Default to 25 Millisecond
		cmd := common.Command{
			MusicTriggerOff: true,
			Speed:           sequenceSpeed,
		}
		common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
	}
}

func updateStaticLamps(selectedSequence int, staticButtons []common.StaticColorButtons, staticLamps [][]bool, commandChannels []chan common.Command) {

	// Remember which color we are setting in this sequence.
	red := staticButtons[selectedSequence].Color.R
	green := staticButtons[selectedSequence].Color.G
	blue := staticButtons[selectedSequence].Color.B

	for X := 0; X < 8; X++ {
		if staticLamps[X][selectedSequence] {
			// Static is set to true in the functions and this key is set to
			// the selected color.
			cmd := common.Command{
				UpdateStaticColor: true,
				Static:            true,
				StaticLamp:        X,
				StaticColor:       common.Color{R: red, G: green, B: blue},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
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

// func refreshSequence(sequences []*common.Sequence, sequenceChannels common.Channels, commandChannels []chan common.Command) (newSequences []*common.Sequence) {

// 	for selectedSequence := range sequences {
// 		cmd := common.Command{
// 			ReadConfig: true,
// 		}
// 		common.SendCommandToSequence(selectedSequence+1, cmd, commandChannels)

// 		newSequence := common.Sequence{}
// 		replyChannel := sequenceChannels.ReplyChannels[selectedSequence]
// 		newSequence = <-replyChannel

// 		newSequences = append(newSequences, &newSequence)
// 	}

// 	return newSequences
// }
