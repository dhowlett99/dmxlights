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
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/rakyll/launchpad/mk2"
)

const (
	full = 3
)

var sequenceSpeed int
var fadeSpeed int
var savePreset bool

// main thread is used to get commands from the lauchpad.
func main() {

	var flashButtons [][]bool

	selectedSequence := 0
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

	// Create a channel to send events to the launchpad.
	eventsForLauchpad := make(chan common.ALight)

	// Now create a thread to handle those events.
	go launchpad.ListenAndSendToLaunchPad(eventsForLauchpad, pad)

	// Start off by turning off all of the Lights
	pad.Reset()

	// Create a channel to listen for buttons being pressed.
	buttonChannel := pad.Listen()

	// Build the default set of Pattens.
	Pattens := patten.MakePatterns()

	// Make a channel to send commands to each sequence.
	sequence1 := make(chan common.Command)
	sequence2 := make(chan common.Command)
	sequence3 := make(chan common.Command)
	sequence4 := make(chan common.Command)

	// Add command channels to an array.
	sequences := []chan common.Command{}
	sequences = append(sequences, sequence1)
	sequences = append(sequences, sequence2)
	sequences = append(sequences, sequence3)
	sequences = append(sequences, sequence4)

	// Make channels for each sequence to talk back to us on.
	readSequence1 := make(chan common.Command)
	readSequence2 := make(chan common.Command)
	readSequence3 := make(chan common.Command)
	readSequence4 := make(chan common.Command)

	readSequences := []chan common.Command{}
	readSequences = append(readSequences, readSequence1)
	readSequences = append(readSequences, readSequence2)
	readSequences = append(readSequences, readSequence3)
	readSequences = append(readSequences, readSequence4)

	// Start threads for each sequence.
	go sequence.CreateSequence(1, pad, eventsForLauchpad, sequence1, readSequence1, Pattens, dmxController)
	go sequence.CreateSequence(2, pad, eventsForLauchpad, sequence2, readSequence2, Pattens, dmxController)
	go sequence.CreateSequence(3, pad, eventsForLauchpad, sequence3, readSequence3, Pattens, dmxController)
	go sequence.CreateSequence(4, pad, eventsForLauchpad, sequence4, readSequence4, Pattens, dmxController)

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
	pad.Light(8, -1, 79)

	//var button [][]bool

	for {
		select {

		case hit := <-buttonChannel:

			// pad.Light(hit.X, hit.Y, 79)
			// fmt.Printf("Pad X:%d Y:%d\n", hit.X, hit.Y)

			if hit.X == 0 && hit.Y == -1 {
				launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, sequences)
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
				launchpad.FlashButton(presetsStore, pad, flashButtons, 8, 4, eventsForLauchpad, 1, 3, 0, 0)
				savePreset = true
			}

			// Ask all sequences for their current config and save in a file.
			if hit.X < 8 && (hit.Y > 3 && hit.Y < 7) {
				if savePreset {
					fmt.Printf("Write Config\n")
					presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] = true
					common.LightOn(eventsForLauchpad, common.ALight{
						X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
					fmt.Printf("Save Preset in X:%d Y:%d \n", hit.X, hit.Y)
					config.AskToSaveConfig(sequences, readSequences, hit.X, hit.Y)
					savePreset = false
					flashButtons[8][4] = false
					presets.SavePresets(presetsStore)
				} else {
					// Load config, but only if it exists in the presets map.
					if presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {
						fmt.Printf("Read Config:")
						fmt.Printf(" OK \n")
						launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, sequences)
						common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
						// Stop everything so that we start the recalled config in sync.
						cmd := common.Command{
							Stop: true,
						}
						sequence1 <- cmd
						sequence2 <- cmd
						sequence3 <- cmd
						sequence4 <- cmd
						time.Sleep(850 * time.Millisecond)
						// Load the config.
						config.AskToLoadConfig(sequences, hit.X, hit.Y)
						// Reset flash buttons
						for x := 0; x < 9; x++ {
							for y := 0; y < 9; y++ {
								flashButtons[x][y] = false
							}
						}
						flashButtons[hit.X][hit.Y] = true
						launchpad.FlashButton(presetsStore, pad, flashButtons, hit.X, hit.Y, eventsForLauchpad, 1, 0, 3, 0)
					}
				}
			}

			// Select standard Patten.
			if hit.X == 2 && hit.Y == 7 {
				cmd := common.Command{
					UpdatePatten: true,
					Patten: common.Patten{
						Name: "colors",
					},
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}
				continue
			}
			// Select pairs Patten.
			if hit.X == 3 && hit.Y == 7 {
				cmd := common.Command{
					UpdatePatten: true,
					Patten: common.Patten{

						Name: "pairs",
					},
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}
				continue
			}

			// Decrease speed of selected sequence.
			if hit.X == 0 && hit.Y == 7 {
				sequenceSpeed--
				if sequenceSpeed < 0 {
					sequenceSpeed = 1
				}
				fmt.Printf("Seq Speed %d\n", sequenceSpeed)
				cmd := common.Command{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}

				continue
			}

			// Increase speed of selected sequence.
			if hit.X == 1 && hit.Y == 7 {
				sequenceSpeed++
				if sequenceSpeed > 12 {
					sequenceSpeed = 12
				}
				cmd := common.Command{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}

				continue
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
					Speed: 11,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}
			}
			// Stop sequence.
			if hit.X == 8 && hit.Y == 6 {
				cmd := common.Command{
					Stop: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}
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
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}

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
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				if selectedSequence == 2 {
					sequence2 <- cmd
				}
				if selectedSequence == 3 {
					sequence3 <- cmd
				}
				if selectedSequence == 4 {
					sequence4 <- cmd
				}

				continue
			}

			// Light the flash buttons based on current patten.
			sequence := common.Sequence{
				Patten: common.Patten{
					Name:  "colors",
					Steps: Pattens["colors"].Steps,
				},
			}
			if hit.X >= 0 && hit.X < 8 {
				red := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].R
				green := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].G
				blue := sequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].B
				common.LightOn(eventsForLauchpad, common.ALight{
					X:          hit.X,
					Y:          hit.Y,
					Brightness: full,
					Red:        red,
					Green:      green,
					Blue:       blue,
				})
				dmx.Fixtures(dmxController, hit.X, red, green, blue)
				time.Sleep(200 * time.Millisecond)
				common.LightOff(eventsForLauchpad, hit.X, hit.Y)
				dmx.Fixtures(dmxController, hit.X, 0, 0, 0)
			}
		}
	}
}
