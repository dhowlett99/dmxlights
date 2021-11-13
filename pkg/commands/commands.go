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

	// Now process any command.
	if command.Hide {
		if debug {
			fmt.Printf("%d: Command Hide\n", mySequenceNumber)
		}
		sequence.Hide = true
		return sequence
	}
	if command.UnHide {
		if debug {
			fmt.Printf("%d: Command UnHide\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Hide = false
		return sequence
	}
	if command.UpdateSpeed {
		if debug {
			fmt.Printf("%d: Command Update Speed to %d\n", mySequenceNumber, command.Speed)
		}
		sequence.Speed = command.Speed
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		return sequence
	}
	if command.UpdatePatten {
		if debug {
			fmt.Printf("%d: Command Update Patten to %s\n", mySequenceNumber, command.Patten.Name)
		}
		savePattenName := command.Patten.Name
		sequence.Patten.Name = savePattenName
		return sequence
	}
	if command.UpdateSize {
		if debug {
			fmt.Printf("%d: Command Update Size to %d\n", mySequenceNumber, command.Size)
		}
		sequence.Size = command.Size
		return sequence
	}
	if command.UpdateSequenceSize {
		if debug {
			fmt.Printf("%d: Command Update Sequence Size to %d\n", mySequenceNumber, command.SequenceSize)
		}
		sequence.SequenceSize = command.SequenceSize
		return sequence
	}
	if command.IncreaseFade {
		if debug {
			fmt.Printf("%d: Command Increase Fade to %d\n", mySequenceNumber, command.FadeSpeed)
		}
		sequence.FadeSpeed = command.FadeSpeed
		sequence.FadeTime = SetSpeed(command.FadeSpeed)
		return sequence
	}
	if command.DecreaseFade {
		if debug {
			fmt.Printf("%d: Command Decrease Fade to %d\n", mySequenceNumber, command.FadeSpeed)
		}
		sequence.FadeSpeed = command.FadeSpeed
		sequence.FadeTime = SetSpeed(command.FadeSpeed)
		return sequence
	}
	if command.UpdateColor {
		if debug {
			fmt.Printf("%d: Command Update Color to %d\n", mySequenceNumber, command.Color)
		}
		color := command.Color
		sequence.Color = color
		return sequence
	}
	if command.Start {
		if debug {
			fmt.Printf("%d: Command Start\n", mySequenceNumber)
		}
		sequence.Mode = "Sequence"
		sequence.Static = false
		sequence.Run = true
		return sequence
	}
	if command.Stop {
		if debug {
			fmt.Printf("%d: Command Stop\n", mySequenceNumber)
		}
		sequence.Functions[common.Function8_Music_Trigger].State = false
		sequence.Functions[common.Function6_Static].State = false
		sequence.MusicTrigger = false
		sequence.Run = false
		sequence.Static = false
		return sequence
	}
	if command.PlayStaticOnce {
		if debug {
			fmt.Printf("%d: Command PlayStaticOnce\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		return sequence
	}
	if command.Blackout {
		if debug {
			fmt.Printf("%d: Command Blackout\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Blackout = true
		return sequence
	}
	if command.UpdateFlood {
		if debug {
			fmt.Printf("%d: Command Flood to %t\n", mySequenceNumber, command.Flood)
		}
		sequence.Flood = command.Flood
		sequence.PlayFloodOnce = true
		return sequence
	}
	if command.Normal {
		if debug {
			fmt.Printf("%d: Command Normal\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Blackout = false
		return sequence
	}
	if command.UpdateFunctions {
		if debug {
			fmt.Printf("%d: Command Update Functions\n", mySequenceNumber)
			for _, function := range command.Functions {
				fmt.Printf(" Function:%d: Name:%s State:%t\n", function.Number, function.Name, function.State)
			}
		}
		// Setup the actions based on the state of the function keys.
		sequence := common.SetFunctionKeyActions(command.Functions, sequence)

		// Always bounce the pattern if we're a scanner. Except if we're a circle.
		if sequence.Type == "scanner" && sequence.Patten.Name != "circle" {
			sequence.Bounce = true
		}
		if sequence.Type == "scanner" && sequence.Patten.Name == "circle" {
			sequence.Bounce = false
		}
		return sequence
	}
	if command.SetEditColors {
		if debug {
			fmt.Printf("%d: Command EditColors Static to %t\n", mySequenceNumber, command.Static)
		}
		sequence.PlayStaticOnce = true
		sequence.EditColors = command.EditColors
		return sequence
	}
	if command.UpdateStatic {
		if debug {
			fmt.Printf("%d: Command Update Static to %t\n", mySequenceNumber, command.Static)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Static = command.Static
		if sequence.Mode == "Static" {
			sequence.Static = true
		}
		return sequence
	}
	if command.UpdateMode {
		if debug {
			fmt.Printf("%d: Command Update Mode to %s\n", mySequenceNumber, command.Mode)
		}
		sequence.Mode = command.Mode
		if sequence.Mode == "Static" {
			fmt.Printf("Setting Run to %t\n", sequence.Run)
			sequence.Run = false
		}
		return sequence
	}
	if command.UpdateStaticColor {
		if debug {
			fmt.Printf("%d: Command Update Static Color\n", mySequenceNumber)
			fmt.Printf("Lamp Color   R:%d  G:%d  B:%d\n", command.StaticColor.R, command.StaticColor.G, command.StaticColor.B)
			fmt.Printf("Selected Color:%d Flash:%t\n", command.SelectedColor, command.StaticLampFlash)
		}
		sequence.PlayStaticOnce = true
		sequence.Static = command.Static
		sequence.StaticColors[command.StaticLamp].SelectedColor = command.SelectedColor
		sequence.StaticColors[command.StaticLamp].Color = command.StaticColor
		sequence.StaticColors[command.StaticLamp].Flash = command.StaticLampFlash
		return sequence
	}
	if command.UpdateSequenceColor {
		if debug {
			fmt.Printf("%d: Command Update Sequence Color to %d\n", mySequenceNumber, command.SelectedColor)
		}
		sequence.UpdateSequenceColor = true
		sequence.SaveColors = true
		sequence.SequenceColors = append(sequence.SequenceColors, common.GetColorButtonsArray(command.SelectedColor))
		return sequence
	}
	if command.ClearSequenceColor {
		if debug {
			fmt.Printf("%d: Command Update Sequence Color to %d\n", mySequenceNumber, command.SelectedColor)
		}
		sequence.UpdateSequenceColor = false
		sequence.SequenceColors = []common.Color{}
		sequence.CurrentSequenceColors = []common.Color{}
		return sequence
	}
	// If we are being asked for our config we must reply with our current sequence.
	if command.ReadConfig {
		if debug {
			fmt.Printf("%d: Command Read Config\n", mySequenceNumber)
		}
		replyChannel <- sequence
		return sequence
	}
	// We are setting the Master brightness in this sequence.
	if command.MasterBrightness {
		if debug {
			fmt.Printf("%d: Command Master Brightness set to %d\n", mySequenceNumber, sequence.Master)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Master = command.Master
		return sequence
	}
	// If we are being asked for a updated config we must reply with our current sequence.
	if command.UpdateSequence {
		if debug {
			fmt.Printf("%d: Command Update Sequence\n", mySequenceNumber)
		}
		updateChannel <- sequence
		return sequence
	}
	// Update function mode for the current sequence.
	if command.UpdateFunctionMode {
		if debug {
			fmt.Printf("%d: Command Update Function Mode %t\n", mySequenceNumber, command.FunctionMode)
		}
		sequence.FunctionMode = command.FunctionMode
		return sequence
	}
	// Update the named switch position for the current sequence.
	if command.UpdateSwitch {
		if debug {
			fmt.Printf("%d: Command Update Switch %d to Position %d\n", mySequenceNumber, command.SwitchNumber, command.SwitchPosition)
		}
		sequence.Switches[command.SwitchNumber].CurrentPosition = command.SwitchPosition
		sequence.PlaySwitchOnce = true
		sequence.Run = false
		sequence.Type = "switch"
		return sequence
	}
	// Update switch positions so they get displayed.
	if command.UpdateSwitchPositions {
		if debug {
			fmt.Printf("%d: Command Update Switch Positions \n", mySequenceNumber)
		}
		sequence.PlaySwitchOnce = true
		sequence.Run = false
		sequence.Type = "switch"
		return sequence
	}
	if command.UpdateGobo {
		if debug {
			fmt.Printf("%d: Command Update Gobo to Number %d\n", mySequenceNumber, command.SelectedGobo)
		}
		sequence.SelectedGobo = command.SelectedGobo + 1
		return sequence
	}
	if command.UpdateAutoColor {
		if debug {
			fmt.Printf("%d: Command Update Auto Color to  %t\n", mySequenceNumber, command.AutoColor)
		}
		sequence.AutoColor = command.AutoColor
		sequence.SelectedGobo = 1
		if !command.AutoColor {
			sequence.RecoverSequenceColors = true
		} else {
			sequence.RecoverSequenceColors = false
		}
		return sequence
	}
	// If we are being asekd to load a config, use the new sequence.
	if command.LoadConfig {
		if debug {
			fmt.Printf("%d: Command Load Config\n", mySequenceNumber)
		}
		X := command.X
		Y := command.Y
		config := config.LoadConfig(fmt.Sprintf("config%d.%d.json", X, Y))
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
	if commandSpeed == 19 {
		Speed = 5
	}
	if commandSpeed == 20 {
		Speed = 3
	}
	if commandSpeed == 21 {
		Speed = 1
	}
	return Speed * time.Millisecond
}

// Fade time must be relative to the current speed.
func SetFade(commandSpeed int) (Fade time.Duration) {
	if commandSpeed == 0 {
		Fade = 1000
	}
	if commandSpeed == 1 {
		Fade = 900
	}
	if commandSpeed == 2 {
		Fade = 800
	}
	if commandSpeed == 3 {
		Fade = 700
	}
	if commandSpeed == 4 {
		Fade = 600
	}
	if commandSpeed == 5 {
		Fade = 500
	}
	if commandSpeed == 6 {
		Fade = 400
	}
	if commandSpeed == 7 {
		Fade = 300
	}
	if commandSpeed == 8 {
		Fade = 200
	}
	if commandSpeed == 9 {
		Fade = 150
	}
	if commandSpeed == 10 {
		Fade = 100
	}
	if commandSpeed == 11 {
		Fade = 50
	}
	if commandSpeed == 12 {
		Fade = 25
	}
	return Fade * time.Millisecond
}
