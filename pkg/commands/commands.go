package commands

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
)

// listenCommandChannelAndWait listens on channel for instructions or timeout and go to next step of sequence.
func ListenCommandChannelAndWait(mySequenceNumber int, speed time.Duration, sequence common.Sequence, channels common.Channels) common.Sequence {

	// Setup channels.
	commandChannel := channels.CommmandChannels[mySequenceNumber-1]
	replyChannel := channels.ReplyChannels[mySequenceNumber-1]
	soundTriggerChannel := channels.SoundTriggerChannels[mySequenceNumber-1]

	// Create an empty command.
	command := common.Command{}

	// Wait for a trigger: sound, command or timeout.
	select {
	case command = <-soundTriggerChannel:
		if sequence.MusicTrigger {
			fmt.Printf("Music Trigger on seq %d\n", mySequenceNumber)
			break
		}
	case command = <-commandChannel:
		break

	case <-time.After(speed):
		break
	}

	// Now process any command.
	if command.Hide {
		fmt.Printf("Command Hide\n")
		sequence.Hide = true
		return sequence
	}
	if command.UnHide {
		fmt.Printf("Command UNHide\n")
		sequence.Hide = false
		return sequence
	}
	if command.MusicTrigger {
		fmt.Printf("Command Music Trigger\n")
		sequence.MusicTrigger = true
		sequence.CurrentSpeed = time.Duration(12 * time.Hour)
		sequence.Run = true
		return sequence
	}
	if command.MusicTriggerOff {
		fmt.Printf("Command Music Trigger Off\n")
		sequence.MusicTrigger = false
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		sequence.Run = true
		return sequence
	}
	if command.UpdateSpeed {
		fmt.Printf("Command Update Speed\n")
		sequence.Speed = command.Speed
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		return sequence
	}
	if command.UpdatePatten {
		fmt.Printf("Command Update Patten\n")
		savePattenName := command.Patten.Name
		sequence.Patten.Name = savePattenName
		return sequence
	}
	if command.UpdateSize {
		fmt.Printf("Command Update Size\n")
		sequence.Size = command.Size
		return sequence
	}
	if command.IncreaseFade {
		fmt.Printf("Command Increase Fade\n")
		newFadeTime := SetSpeed(command.FadeSpeed)
		sequence.Steps = sequence.Steps + 1
		sequence.FadeTime = newFadeTime
		sequence.FadeSpeed = command.FadeSpeed
		return sequence
	}
	if command.DecreaseFade {
		fmt.Printf("Command Decrease Fade\n")
		newFadeTime := SetSpeed(command.FadeSpeed)
		sequence.Steps = sequence.Steps - 1
		sequence.FadeTime = newFadeTime
		sequence.FadeSpeed = command.FadeSpeed
		return sequence
	}
	if command.UpdateColor {
		fmt.Printf("Command Update Color\n")
		color := command.Color
		sequence.Color = color
		return sequence
	}
	if command.Start {
		fmt.Printf("Command Start\n")
		sequence.MusicTrigger = command.MusicTrigger
		sequence.Static = false
		sequence.Run = true
		return sequence
	}
	if command.Stop {
		fmt.Printf("Command Stop\n")
		sequence.Run = false
		sequence.Static = false
		return sequence
	}
	if command.Blackout {
		fmt.Printf("Command Blackout\n")
		sequence.Blackout = true
		return sequence
	}
	if command.Normal {
		fmt.Printf("Command Normal\n")
		sequence.Blackout = false
		return sequence
	}
	if command.UpdateFunctions {
		fmt.Printf("Command Update Functions\n")
		sequence.Functions = command.Functions
		return sequence
	}

	if command.UpdateStatic {
		fmt.Printf("Command Update Static\n")
		sequence.Static = command.Static
		return sequence
	}

	if command.UpdateStaticColor {
		fmt.Printf("Command Update Static Color\n")
		sequence.Static = command.Static
		fmt.Printf("Lamp Color   R:%d  G:%d  B:%d\n", command.StaticColor.R, command.StaticColor.G, command.StaticColor.B)
		sequence.StaticColors[command.StaticLamp] = command.StaticColor
		return sequence
	}

	// If we are being asked for our config we must reply with our current sequence.
	if command.ReadConfig {
		fmt.Printf("Command Read Config\n")
		replyChannel <- sequence
		return sequence
	}

	// If we are being asekd to load a config, use the new sequence.
	if command.LoadConfig {
		fmt.Printf("Command Load Config\n")
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
