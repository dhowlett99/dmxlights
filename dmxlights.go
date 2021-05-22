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
var presetsStore map[string]bool
var savePreset bool
var flashButtons [][]bool

// main thread is used to get commands from the lauchpad.
func main() {

	selectedSequence := 0
	presetsStore = make(map[string]bool)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		presets.SavePresets(presetsStore)
		os.Exit(1)
	}()

	fmt.Println("Derek common.Lighting")

	fmt.Println("Loading Presets")
	presetsStore = presets.LoadPresets()
	fmt.Println("Loading Presets Done")

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
	go listenAndSendToLaunchPad(eventsForLauchpad, pad)

	// Start off by turning off all of the Lights
	//pad.Clear()

	// Create a channel to listen for buttons being pressed.
	buttonChannel := pad.Listen()

	// Build the default set of Pattens.
	Pattens := patten.MakePatterns()

	// Make a channel to communicate with each sequence.
	// Create a channel for par cans.
	sequence1 := make(chan common.Sequence)
	// sequence2 := make(chan Sequence)
	// sequence3 := make(chan Sequence)
	// sequence4 := make(chan Sequence)

	sequences := []chan common.Sequence{}
	sequences = append(sequences, sequence1)
	// sequences = append(sequences, sequence2)
	// sequences = append(sequences, sequence3)
	// sequences = append(sequences, sequence4)

	// Make channels for each sequence to talk back to us on.
	readSequence1 := make(chan common.Sequence)
	// readSequence2 := make(chan Sequence)
	// readSequence3 := make(chan Sequence)
	// readSequence4 := make(chan Sequence)

	readSequences := []chan common.Sequence{}
	readSequences = append(readSequences, readSequence1)
	// readSequences = append(readSequences, readSequence2)
	// readSequences = append(readSequences, readSequence3)
	// readSequences = append(readSequences, readSequence4)

	// Start threads for each sequence.
	go sequence.CreateSequence(1, pad, eventsForLauchpad, sequence1, readSequence1, Pattens)
	// go CreateSequence(2, pad, eventsForLauchpad, sequence2, readSequence2, Pattens)
	// go CreateSequence(3, pad, eventsForLauchpad, sequence3, readSequence3, Pattens)
	// go CreateSequence(4, pad, eventsForLauchpad, sequence4, readSequence4, Pattens)

	// common.Light up any existing presets.
	presets.InitPresets(eventsForLauchpad, presetsStore)

	fmt.Println("Setup Presets Done")

	// Initialize a ten length slice of empty slices
	flashButtons = make([][]bool, 9)

	// Initialize those 10 empty slices
	for i := 0; i < 9; i++ {
		flashButtons[i] = make([]bool, 9)
	}

	// common.Light the logo blue.
	pad.Light(8, -1, 79)
	// err = pad.Logo()
	// if err != nil {
	// 	fmt.Printf("Error setting logo %s\n", err.Error())
	// }

	for {
		select {

		case hit := <-buttonChannel:

			pad.Light(hit.X, hit.Y, 79)
			fmt.Printf("Pad X:%d Y:%d\n", hit.X, hit.Y)
			//time.Sleep(1 * time.Second)

			if hit.X == 0 && hit.Y == -1 {
				clearAll(pad, eventsForLauchpad, sequences)
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
				flashButton(pad, 8, 4, eventsForLauchpad, 1, 3, 0, 0)
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
				} else {
					// Load config, but only if it exists in the presets map.
					if presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {
						fmt.Printf("Read Config:")
						fmt.Printf(" OK \n")
						clearAll(pad, eventsForLauchpad, sequences)
						common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
						config.AskToLoadConfig(sequences, hit.X, hit.Y)
						// Reset flash buttons
						for x := 0; x < 9; x++ {
							for y := 0; y < 9; y++ {
								flashButtons[x][y] = false
							}
						}
						flashButtons[hit.X][hit.Y] = true
						flashButton(pad, hit.X, hit.Y, eventsForLauchpad, 1, 0, 3, 0)
					}
				}
			}

			// Select standard Patten.
			if hit.X == 2 && hit.Y == 7 {
				cmd := common.Sequence{
					UpdatePatten: true,
					Patten: common.Patten{
						Name: "rgbchase",
					},
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
				continue
			}
			// Select pairs Patten.
			if hit.X == 3 && hit.Y == 7 {
				cmd := common.Sequence{
					UpdatePatten: true,
					Patten: common.Patten{

						Name: "pairs",
					},
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
				continue
			}

			// Decrease speed of selected sequence.
			if hit.X == 0 && hit.Y == 7 {
				sequenceSpeed--
				if sequenceSpeed < 0 {
					sequenceSpeed = 1
				}
				fmt.Printf("Seq Speed %d\n", sequenceSpeed)
				cmd := common.Sequence{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

				continue
			}

			// Increase speed of selected sequence.
			if hit.X == 1 && hit.Y == 7 {
				sequenceSpeed++
				if sequenceSpeed > 12 {
					sequenceSpeed = 12
				}
				cmd := common.Sequence{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

				continue
			}

			// Select sequence 1.
			if hit.X == 8 && hit.Y == 0 {
				selectedSequence = 1
				common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 0, Brightness: full, Red: 3, Green: 3, Blue: 0})
				event := common.ALight{X: 8, Y: 1, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 2, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 3, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 2.
			if hit.X == 8 && hit.Y == 1 {
				selectedSequence = 2
				event := common.ALight{X: 8, Y: 1, Brightness: full, Red: 3, Green: 3, Blue: 0}
				eventsForLauchpad <- event

				event = common.ALight{X: 8, Y: 0, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 2, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 3, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 3.
			if hit.X == 8 && hit.Y == 2 {
				selectedSequence = 3
				event := common.ALight{X: 8, Y: 2, Brightness: full, Red: 3, Green: 3, Blue: 0}
				eventsForLauchpad <- event

				event = common.ALight{X: 8, Y: 1, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 3, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 0, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 4.
			if hit.X == 8 && hit.Y == 3 {
				selectedSequence = 4
				event := common.ALight{X: 8, Y: 3, Brightness: full, Red: 3, Green: 3, Blue: 0}
				eventsForLauchpad <- event

				event = common.ALight{X: 8, Y: 1, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 2, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.ALight{X: 8, Y: 4, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Start sequence.
			if hit.X == 8 && hit.Y == 5 {
				cmd := common.Sequence{
					Start: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
			}
			// Stop sequence.
			if hit.X == 8 && hit.Y == 6 {
				cmd := common.Sequence{
					Stop: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
			}

			// Fade time decrease.
			if hit.X == 4 && hit.Y == 7 {
				fadeSpeed--
				if fadeSpeed < 0 {
					fadeSpeed = 0
				}
				fmt.Printf("Fade down speed:%d", fadeSpeed)
				fadeTime := commands.SetFade(fadeSpeed)
				cmd := common.Sequence{
					UpdateFade: true,
					FadeTime:   fadeTime,
					Number:     selectedSequence,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

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
				cmd := common.Sequence{
					UpdateFade: true,
					FadeTime:   fadeTime,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

				continue
			}

			// // common.Light a button is pressed.
			// if !button[hit.X][hit.Y] {
			// 	event := common.ALight{hit.X, hit.Y, 0, 3}
			// 	eventsForLauchpad <- event
			// 	button[hit.X][hit.Y] = true
			// } else {
			// 	event := common.ALight{hit.X, hit.Y, 0, 0}
			// 	eventsForLauchpad <- event
			// 	button[hit.X][hit.Y] = false
			// }
		}

	}
}

func clearAll(pad *mk2.Launchpad, eventsForLauchpad chan common.ALight, sequences []chan common.Sequence) {
	fmt.Printf("C L E A R\n")
	pad.Reset()
	cmd := common.Sequence{
		Stop: true,
	}
	for _, seq := range sequences {
		seq <- cmd
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presetsStore[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				common.LightOn(eventsForLauchpad, common.ALight{X: x, Y: y, Brightness: full, Red: 3, Green: 0, Blue: 0})
			}
		}
	}
}

// listenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func listenAndSendToLaunchPad(eventsForLauchpad chan common.ALight, pad *mk2.Launchpad) {
	var green int
	var red int
	var blue int

	for {
		event := <-eventsForLauchpad
		switch event.Green {
		case 0:
			green = 0
		case 1:
			green = 19
		case 2:
			green = 22
		case 3:
			green = 21
		}

		switch event.Red {
		case 0:
			red = 0
		case 1:
			red = 7
		case 2:
			red = 6
		case 3:
			red = 5
		}

		switch event.Blue {
		case 0:
			blue = 0
		case 1:
			blue = 37
		case 2:
			blue = 38
		case 3:
			blue = 79
		}

		pad.Light(event.X, event.Y, green+red+blue)
	}
}

func flashButton(pad *mk2.Launchpad, x int, y int, eventsForLauchpad chan common.ALight, seqNumber int, green int, red int, blue int) {
	go func(pad *mk2.Launchpad, x int, y int) {

		for {
			// fmt.Printf("Flash X:%d Y:%d is %t\n", x, y, flashButtons[x][y])
			if !flashButtons[x][y] {
				break
			}
			event := common.ALight{X: x, Y: y, Brightness: full, Red: red, Green: green, Blue: blue}
			eventsForLauchpad <- event

			time.Sleep(1 * time.Second)
			event = common.ALight{X: x, Y: y, Brightness: full, Red: 0, Green: 0, Blue: 0}
			eventsForLauchpad <- event
			time.Sleep(1 * time.Second)
		}
	}(pad, x, y)
}
