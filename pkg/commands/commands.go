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
			break
		}
	case command = <-commandChannel:
		break

	case <-time.After(speed):
		break
	}

	// Now process any command.
	if command.Hide {
		sequence.Hide = true
		return sequence
	}
	if command.UnHide {

		sequence.Hide = false
		return sequence
	}
	if command.MusicTrigger {
		sequence.MusicTrigger = true
		if sequence.MusicTrigger {
			sequence.CurrentSpeed = time.Duration(12 * time.Hour)
		}
		return sequence
	}
	if command.MusicTriggerOff {
		sequence.MusicTrigger = false
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		return sequence
	}
	if command.SoftFadeOn {
		sequence.SoftFade = true
		return sequence
	}
	if command.SoftFadeOff {
		sequence.SoftFade = false
		return sequence
	}
	if command.UpdateSpeed {
		sequence.Speed = command.Speed
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		sequence.Run = true
		return sequence
	}
	if command.UpdatePatten {
		savePattenName := command.Patten.Name
		sequence.Patten.Name = savePattenName
		return sequence
	}
	if command.UpdateSize {
		sequence.Size = command.Size
		return sequence
	}
	if command.IncreaseFade {
		newFadeTime := SetSpeed(command.FadeSpeed)
		sequence.Steps = sequence.Steps + 1
		sequence.FadeTime = newFadeTime
		sequence.FadeSpeed = command.FadeSpeed
		return sequence
	}
	if command.DecreaseFade {
		newFadeTime := SetSpeed(command.FadeSpeed)
		sequence.Steps = sequence.Steps - 1
		sequence.FadeTime = newFadeTime
		sequence.FadeSpeed = command.FadeSpeed
		return sequence
	}
	if command.UpdateColor {
		color := command.Color
		sequence.Color = color
		return sequence
	}
	if command.Start {
		sequence.MusicTrigger = command.MusicTrigger
		sequence.Run = true
		return sequence
	}
	if command.Stop {
		sequence.Run = false
		return sequence
	}
	if command.Blackout {
		sequence.Blackout = true
		return sequence
	}
	if command.Normal {
		sequence.Blackout = false
		return sequence
	}
	if command.UpdateFunctions {
		sequence.Functions = command.Functions
		return sequence
	}

	// If we are being asked for our config we must reply with our current sequence.
	if command.ReadConfig {
		replyChannel <- sequence
		return sequence
	}

	// If we are being asekd to load a config, use the new sequence.
	if command.LoadConfig {
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
