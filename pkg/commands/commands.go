package commands

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
)

const debug = false

// listenCommandChannelAndWait listens on channel for instructions or timeout and go to next step of sequence.
func ListenCommandChannelAndWait(mySequenceNumber int, speed time.Duration, sequence common.Sequence, channels common.Channels) common.Sequence {

	// Setup channels.
	commandChannel := channels.CommmandChannels[mySequenceNumber]
	replyChannel := channels.ReplyChannels[mySequenceNumber]
	soundTriggerChannel := channels.SoundTriggerChannels[mySequenceNumber]
	updateChannel := channels.UpdateChannels[mySequenceNumber]

	// Create an empty command.
	command := common.Command{}

	// Wait for a trigger: sound, command or timeout.
	select {
	case command = <-soundTriggerChannel:
		if sequence.MusicTrigger {
			// if debug {
			// 	fmt.Printf("%d: BEAT\n", mySequenceNumber)
			// }
			break
		}
	case command = <-commandChannel:
		break

	case <-time.After(speed):
		break
	}

	switch command.Action {
	// Now process any command.
	case common.Hide:
		if debug {
			fmt.Printf("%d: Command Hide\n", mySequenceNumber)
		}
		sequence.Hide = true
		return sequence

	case common.UnHide:
		if debug {
			fmt.Printf("%d: Command UnHide\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Hide = false
		return sequence

	case common.UpdateSpeed:
		const SPEED = 0
		if debug {
			fmt.Printf("%d: Command Update %s to %d\n", mySequenceNumber, command.Args[SPEED].Name, command.Args[SPEED].Value)
		}
		sequence.Speed = command.Args[SPEED].Value.(int)
		sequence.CurrentSpeed = SetSpeed(command.Args[SPEED].Value.(int))
		return sequence

	case common.UpdatePatten:
		const PATTEN_NAME = 0
		if debug {
			fmt.Printf("%d: Command Update Patten to %s\n", mySequenceNumber, command.Args[PATTEN_NAME].Value)
		}
		sequence.Patten.Name = command.Args[PATTEN_NAME].Value.(string)

		return sequence

	case common.SelectPatten:
		const SELECTED_PATTEN = 0
		if debug {
			fmt.Printf("%d: Command Select Patten to %d\n", mySequenceNumber, command.Args[SELECTED_PATTEN].Value)
		}
		sequence.UpdatePatten = true
		sequence.SelectedRGBPatten = command.Args[SELECTED_PATTEN].Value.(int)
		sequence.ScannerPatten = command.Args[SELECTED_PATTEN].Value.(int)
		return sequence

	case common.UpdateShift:
		const SHIFT = 0
		if debug {
			fmt.Printf("%d: Command Update Shift to %d\n", mySequenceNumber, command.Args[SHIFT].Value)
		}
		sequence.UpdateShift = true
		sequence.ScannerShift = command.Args[SHIFT].Value.(int)
		return sequence

	case common.UpdateSize:
		const START = 0
		if debug {
			fmt.Printf("%d: Command Update Size to %d\n", mySequenceNumber, command.Args[START].Value)
		}
		sequence.Size = command.Args[START].Value.(int)
		return sequence

	case common.UpdateScannerSize:
		const SCANNER_SIZE = 0
		if debug {
			fmt.Printf("%d: Command Update Scanner Size to %d\n", mySequenceNumber, command.Args[SCANNER_SIZE].Value)
		}
		sequence.ScannerSize = command.Args[SCANNER_SIZE].Value.(int)
		return sequence

	case common.SetFadeSpeed:
		const FADE_SPEED = 0
		if debug {
			fmt.Printf("%d: Command Set Fade to %d\n", mySequenceNumber, command.Args[FADE_SPEED].Value)
		}
		sequence.FadeSpeed = command.Args[FADE_SPEED].Value.(int)
		sequence.FadeTime = SetSpeed(command.Args[FADE_SPEED].Value.(int))
		return sequence

	case common.UpdateColor:
		const COLOR = 0
		if debug {
			fmt.Printf("%d: Command Update Color to %d\n", mySequenceNumber, command.Args[COLOR].Value)
		}
		sequence.Color = command.Args[COLOR].Value.(int)
		return sequence

	case common.Start:
		if debug {
			fmt.Printf("%d: Command Start\n", mySequenceNumber)
		}
		sequence.Mode = "Sequence"
		sequence.Static = false
		sequence.Run = true
		return sequence

	case common.Stop:
		if debug {
			fmt.Printf("%d: Command Stop\n", mySequenceNumber)
		}
		sequence.Functions[common.Function8_Music_Trigger].State = false
		sequence.Functions[common.Function6_Static_Gobo].State = false
		sequence.MusicTrigger = false
		sequence.Run = false
		sequence.Static = false
		return sequence

	case common.PlayStaticOnce:
		if debug {
			fmt.Printf("%d: Command PlayStaticOnce\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		return sequence

	case common.Blackout:
		if debug {
			fmt.Printf("%d: Command Blackout\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Blackout = true
		return sequence

	case common.Flood:
		if debug {
			fmt.Printf("%d: Command to Start Flood\n", mySequenceNumber)
		}
		sequence.StartFlood = true
		sequence.StopFlood = false
		sequence.FloodPlayOnce = true
		return sequence

	case common.StopFlood:
		if debug {
			fmt.Printf("%d: Command to Stop Flood\n", mySequenceNumber)
		}
		sequence.StartFlood = false
		sequence.StopFlood = true
		sequence.FloodPlayOnce = true
		return sequence

	case common.Normal:
		if debug {
			fmt.Printf("%d: Command Normal\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Blackout = false
		return sequence

	case common.UpdateFunctions:
		const FUNCTIONS = 0
		if debug {
			fmt.Printf("%d: Command Update Functions\n", mySequenceNumber)
			for _, function := range command.Args[FUNCTIONS].Value.([]common.Function) {
				fmt.Printf(" Function:%d: Name:%s State:%t\n", function.Number, function.Name, function.State)
			}
		}
		// Setup the actions based on the state of the function keys.
		sequence = common.SetFunctionKeyActions(command.Args[FUNCTIONS].Value.([]common.Function), sequence)

		// Always bounce the pattern if we're a scanner. Except if we're a circle.
		if sequence.Type == "scanner" && sequence.Patten.Name != "circle" {
			sequence.Bounce = true
		}
		if sequence.Type == "scanner" && sequence.Patten.Name == "circle" {
			sequence.Bounce = false
		}
		return sequence

	case common.UpdateStatic:
		const STATIC = 0
		if debug {
			fmt.Printf("%d: Command Update Static to %t\n", mySequenceNumber, command.Args[STATIC].Value)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Static = command.Args[STATIC].Value.(bool)

	case common.UpdateMode:
		const MODE = 0
		if debug {
			fmt.Printf("%d: Command Update Mode to %s\n", mySequenceNumber, command.Args[MODE].Value)
		}
		sequence.Mode = command.Args[MODE].Value.(string)
		if sequence.Mode == "Static" {
			sequence.Static = true
		}
		if sequence.Mode == "Sequence" {
			sequence.Static = false
		}
		return sequence

	case common.UpdateStaticColor:
		const STATIC = 0            // Boolean
		const STATIC_LAMP = 1       // Integer
		const STATIC_LAMP_FLASH = 2 // Boolean
		const SELECTED_COLOR = 3    // Integer
		const STATIC_COLOR = 4      // Color
		if debug {
			fmt.Printf("%d: Command Update Static Color\n", mySequenceNumber)
			fmt.Printf("Lamp Color   %+v\n", command.Args[STATIC_COLOR].Value.(common.Color))
			fmt.Printf("Selected Color:%d Flash:%t\n", command.Args[SELECTED_COLOR].Value, command.Args[STATIC_LAMP_FLASH].Value)
		}
		sequence.PlayStaticOnce = true
		sequence.Static = command.Args[STATIC].Value.(bool)
		sequence.Hide = true
		sequence.StaticColors[command.Args[STATIC_LAMP].Value.(int)].SelectedColor = command.Args[SELECTED_COLOR].Value.(int)
		sequence.StaticColors[command.Args[STATIC_LAMP].Value.(int)].Color = command.Args[STATIC_COLOR].Value.(common.Color)
		sequence.StaticColors[command.Args[STATIC_LAMP].Value.(int)].Flash = command.Args[STATIC_LAMP_FLASH].Value.(bool)

	case common.UpdateSequenceColor:
		const SELECTED_COLOR = 0
		if debug {
			fmt.Printf("%d: Command Update Sequence Color to %d\n", mySequenceNumber, command.Args[SELECTED_COLOR].Value)
		}
		sequence.UpdateSequenceColor = true
		sequence.SaveColors = true
		sequence.SequenceColors = append(sequence.SequenceColors, common.GetColorButtonsArray(command.Args[SELECTED_COLOR].Value.(int)))
		return sequence

	case common.UpdateScannerColor:
		const SELECTED_COLOR = 0
		const FIXTURE_NUMBER = 1
		if debug {
			fmt.Printf("%d: Command Update Scanner Color for fixture %d to %d\n", mySequenceNumber, command.Args[FIXTURE_NUMBER].Value, command.Args[SELECTED_COLOR].Value)
		}
		sequence.SaveColors = true
		sequence.ScannerColor[command.Args[FIXTURE_NUMBER].Value.(int)] = command.Args[SELECTED_COLOR].Value.(int)
		return sequence

	case common.ClearSequenceColor:
		if debug {
			fmt.Printf("%d: Command Clear Sequence Color \n", mySequenceNumber)
		}
		sequence.UpdateSequenceColor = false
		sequence.SequenceColors = []common.Color{}
		sequence.CurrentColors = []common.Color{}
		return sequence

	// If we are being asked for our config we must reply with our current sequence.
	case common.ReadConfig:
		if debug {
			fmt.Printf("%d: Command Read Config\n", mySequenceNumber)
		}
		replyChannel <- sequence
		return sequence

	// We are setting the Master brightness in this sequence.
	case common.MasterBrightness:
		const MASTER = 0
		if debug {
			fmt.Printf("%d: Command Master Brightness set to %d\n", mySequenceNumber, command.Args[MASTER].Value)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Master = command.Args[MASTER].Value.(int)
		return sequence

	// If we are being asked for a updated config we must reply with our current sequence.
	case common.GetUpdatedSequence:
		if debug {
			fmt.Printf("%d: Command Get Updated Sequence\n", mySequenceNumber)
		}
		updateChannel <- sequence
		return sequence

	// Update function mode for the current sequence.
	case common.UpdateFunctionMode:
		const FUNCTION_MODE = 0
		if debug {
			fmt.Printf("%d: Command Update Function Mode %t\n", mySequenceNumber, command.Args[FUNCTION_MODE].Value)
		}
		sequence.FunctionMode = command.Args[FUNCTION_MODE].Value.(bool)
		return sequence

	// Clear switch positions for this sequence
	case common.ClearAllSwitchPositions:
		// Loop through all the switchies.
		for X := 0; X < len(sequence.Switches); X++ {
			sequence.Switches[X].CurrentState = 0
		}
		sequence.PlaySwitchOnce = true
		return sequence

	// Update the named switch position for the current sequence.
	case common.UpdateSwitch:
		const SWITCH_NUMBER = 0   // Integer
		const SWITCH_POSITION = 1 // Integer
		if debug {
			fmt.Printf("%d: Command Update Switch %d to Position %d\n", mySequenceNumber, command.Args[SWITCH_NUMBER].Value, command.Args[SWITCH_POSITION].Value)
		}
		sequence.Switches[command.Args[SWITCH_NUMBER].Value.(int)].CurrentState = command.Args[SWITCH_POSITION].Value.(int)
		sequence.PlaySwitchOnce = true
		sequence.Run = false
		sequence.Type = "switch"
		return sequence

	// Update switch positions so they get displayed.
	case common.UpdateSwitchPositions:
		if debug {
			fmt.Printf("%d: Command Update Switch Positions \n", mySequenceNumber)
		}
		sequence.PlaySwitchOnce = true
		sequence.Run = false
		sequence.Type = "switch"
		return sequence

	case common.EnableAllScanners:
		const SEQUENCE_NUMBER = 0 // Integer
		if command.Args[SEQUENCE_NUMBER].Value == mySequenceNumber {
			sequence.FixtureDisabledMutex.Lock()
			for scanner := 0; scanner < sequence.ScannersTotal; scanner++ {
				sequence.FixtureDisabled[command.Args[scanner].Value.(int)] = false
			}
			sequence.FixtureDisabledMutex.Unlock()
		}

	// Here we want to disable/enable the selected scanner.
	case common.ToggleFixtureState:
		const SEQUENCE_NUMBER = 0 // Integer
		const FIXTURE_NUMBER = 1  // Integer
		const FIXTURE_STATE = 2   // Boolean
		if debug {
			fmt.Printf("%d: Command ToggleFixtureState for fixture number %d on sequence %d\n", mySequenceNumber, command.Args[FIXTURE_NUMBER].Value, command.Args[SEQUENCE_NUMBER].Value)
		}
		if command.Args[SEQUENCE_NUMBER].Value == mySequenceNumber {
			if command.Args[FIXTURE_NUMBER].Value.(int) < sequence.ScannersTotal {
				sequence.FixtureDisabledMutex.Lock()
				sequence.FixtureDisabled[command.Args[FIXTURE_NUMBER].Value.(int)] = command.Args[FIXTURE_STATE].Value.(bool)
				sequence.FixtureDisabledMutex.Unlock()
			}
		}

		// When we disable a fixture we send a off command to the shutter to make it go off.
		// We only want to do this once to avoid flooding the universe with DMX commands.
		sequence.DisableOnceMutex.Lock()
		sequence.DisableOnce[command.Args[FIXTURE_NUMBER].Value.(int)] = true
		sequence.DisableOnceMutex.Unlock()
		// it will be the fixtures resposiblity to unset this when it's played the stop command.

		return sequence

	case common.UpdateGobo:
		const SELECTED_GOBO = 0
		if debug {
			fmt.Printf("%d: Command Update Gobo to Number %d\n", mySequenceNumber, command.Args[SELECTED_GOBO].Value)
		}
		sequence.ScannerGobo = command.Args[SELECTED_GOBO].Value.(int)
		sequence.Static = false
		return sequence

	case common.UpdateAutoColor:
		const AUTO_COLOR = 0
		if debug {
			fmt.Printf("%d: Command Update Auto Color to  %t\n", mySequenceNumber, command.Args[AUTO_COLOR].Value)
		}
		sequence.AutoColor = command.Args[AUTO_COLOR].Value.(bool)
		sequence.ScannerGobo = 1
		if !command.Args[AUTO_COLOR].Value.(bool) {
			sequence.RecoverSequenceColors = true
		} else {
			sequence.RecoverSequenceColors = false
		}
		return sequence

	case common.UpdateAutoPatten:
		const AUTO_PATTEN = 0
		if debug {
			fmt.Printf("%d: Command Update Auto Patten to  %t\n", mySequenceNumber, command.Args[AUTO_PATTEN].Value)
		}
		sequence.AutoPatten = command.Args[AUTO_PATTEN].Value.(bool)
		if !command.Args[AUTO_PATTEN].Value.(bool) {
			if sequence.Type == "rgb" {
				sequence.Patten.Name = "standard"
			}
			if sequence.Type == "scanner" {
				sequence.Patten.Name = "circle"
				sequence.ScannerPatten = 1
			}
		}
		return sequence

	case common.UpdateNumberCoordinates:
		const NUMBER_COORDINATES = 0
		if debug {
			fmt.Printf("%d: Command Update Number Coordinates to  %d\n", mySequenceNumber, command.Args[NUMBER_COORDINATES].Value)
		}
		sequence.ScannerSelectedCoordinates = command.Args[NUMBER_COORDINATES].Value.(int)
		return sequence

	case common.UpdateOffsetPan:
		const OFFSET_PAN = 0
		if debug {
			fmt.Printf("%d: Command Update Offset Pan to  %d\n", mySequenceNumber, command.Args[OFFSET_PAN].Value)
		}
		sequence.ScannerOffsetPan = command.Args[OFFSET_PAN].Value.(int)
		return sequence

	case common.UpdateOffsetTilt:
		const OFFSET_TILT = 0
		if debug {
			fmt.Printf("%d: Command Update Offset Tilt to  %d\n", mySequenceNumber, command.Args[OFFSET_TILT].Value)
		}
		sequence.ScannerOffsetTilt = command.Args[OFFSET_TILT].Value.(int)
		return sequence

	// If we are being asekd to load a config, use the new sequence.
	case common.LoadConfig:
		const X = 0
		const Y = 1
		if debug {
			fmt.Printf("%d: Command Load Config\n", mySequenceNumber)
		}
		x := command.Args[X].Value.(int)
		y := command.Args[Y].Value.(int)
		config := config.LoadConfig(fmt.Sprintf("config%d.%d.json", x, y))
		seq := common.Sequence{}
		for _, seq = range config {
			if seq.Number == sequence.Number {
				sequence = seq
				// Assume we're blacked out.
				sequence.Blackout = true
				sequence.PlayStaticOnce = true
				sequence.PlaySwitchOnce = true
				return sequence
			}
		}

	}
	return sequence
}

// Used to convert a speed to a millisecond time.
func SetSpeed(commandSpeed int) (Speed time.Duration) {
	if commandSpeed == 0 {
		Speed = 3500
	}
	if commandSpeed == 1 {
		Speed = 3000
	}
	if commandSpeed == 2 {
		Speed = 2500
	}
	if commandSpeed == 3 {
		Speed = 1800
	}
	if commandSpeed == 4 {
		Speed = 1500
	}
	if commandSpeed == 5 {
		Speed = 1000
	}
	if commandSpeed == 6 {
		Speed = 750
	}
	if commandSpeed == 7 {
		Speed = 500
	}
	if commandSpeed == 8 {
		Speed = 250
	}
	if commandSpeed == 9 {
		Speed = 150
	}
	if commandSpeed == 10 {
		Speed = 125
	}
	if commandSpeed == 11 {
		Speed = 100
	}
	if commandSpeed == 12 {
		Speed = 75
	}
	if commandSpeed == 13 {
		Speed = 50
	}
	if commandSpeed == 14 {
		Speed = 25
	}
	if commandSpeed == 15 {
		Speed = 20
	}
	if commandSpeed == 16 {
		Speed = 15
	}
	if commandSpeed == 17 {
		Speed = 10
	}
	if commandSpeed == 18 {
		Speed = 7
	}
	return Speed * time.Millisecond
}
