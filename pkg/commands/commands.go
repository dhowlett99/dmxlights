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

	// Create an empty command.
	command := common.Command{}

	// Wait for a trigger: sound, command or timeout.
	select {
	case command = <-soundTriggerChannel:
		if sequence.MusicTrigger {
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
		sequence.Hide = false
		return sequence
	}
	if command.MusicTrigger {
		if debug {
			fmt.Printf("%d: Command Music Trigger On\n", mySequenceNumber)
		}
		sequence.MusicTrigger = true
		sequence.CurrentSpeed = time.Duration(12 * time.Hour)
		sequence.Run = true
		return sequence
	}
	if command.MusicTriggerOff {
		if debug {
			fmt.Printf("%d: Command Music Trigger Off\n", mySequenceNumber)
		}
		sequence.MusicTrigger = false
		sequence.Speed = command.Speed
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		sequence.Run = true
		return sequence
	}
	if command.UpdateSpeed {
		if debug {
			fmt.Printf("%d: Command Update Speed\n", mySequenceNumber)
		}
		sequence.Speed = command.Speed
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		return sequence
	}
	if command.UpdatePatten {
		if debug {
			fmt.Printf("%d: Command Update Patten\n", mySequenceNumber)
		}
		savePattenName := command.Patten.Name
		sequence.Patten.Name = savePattenName
		return sequence
	}
	if command.UpdateSize {
		if debug {
			fmt.Printf("%d: Command Update Size\n", mySequenceNumber)
		}
		sequence.Size = command.Size
		return sequence
	}
	if command.IncreaseFade {
		if debug {
			fmt.Printf("%d: Command Increase Fade\n", mySequenceNumber)
		}
		newFadeTime := SetSpeed(command.FadeSpeed)
		sequence.Steps = sequence.Steps + 1
		sequence.FadeTime = newFadeTime
		sequence.FadeSpeed = command.FadeSpeed
		return sequence
	}
	if command.DecreaseFade {
		if debug {
			fmt.Printf("%d: Command Decrease Fade\n", mySequenceNumber)
		}
		newFadeTime := SetSpeed(command.FadeSpeed)
		sequence.Steps = sequence.Steps - 1
		sequence.FadeTime = newFadeTime
		sequence.FadeSpeed = command.FadeSpeed
		return sequence
	}
	if command.UpdateColor {
		if debug {
			fmt.Printf("%d: Command Update Color\n", mySequenceNumber)
		}
		color := command.Color
		sequence.Color = color
		return sequence
	}
	if command.Start {
		if debug {
			fmt.Printf("%d: Command Start\n", mySequenceNumber)
		}
		sequence.MusicTrigger = command.MusicTrigger
		sequence.Static = false
		sequence.Run = true
		return sequence
	}
	if command.Stop {
		if debug {
			fmt.Printf("%d: Command Stop\n", mySequenceNumber)
		}
		sequence.Run = false
		sequence.Static = false
		return sequence
	}
	if command.Blackout {
		if debug {
			fmt.Printf("%d: Command Blackout\n", mySequenceNumber)
		}
		sequence.Blackout = true
		return sequence
	}
	if command.Normal {
		if debug {
			fmt.Printf("%d: Command Normal\n", mySequenceNumber)
		}
		sequence.Blackout = false
		return sequence
	}
	if command.UpdateFunctions {
		if debug {
			fmt.Printf("%d: Command Update Functions\n", mySequenceNumber)
		}
		sequence.Functions = command.Functions
		return sequence
	}

	if command.UpdateStatic {
		if debug {
			fmt.Printf("%d: Command Update Static\n", mySequenceNumber)
		}
		sequence.Static = command.Static
		return sequence
	}

	if command.UpdateStaticColor {
		if debug {
			fmt.Printf("%d: Command Update Static Color\n", mySequenceNumber)
			fmt.Printf("Lamp Color   R:%d  G:%d  B:%d\n", command.StaticColor.R, command.StaticColor.G, command.StaticColor.B)
		}
		sequence.Static = command.Static
		sequence.StaticColors[command.StaticLamp] = command.StaticColor
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
