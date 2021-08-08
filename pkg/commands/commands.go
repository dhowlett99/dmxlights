package commands

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/sound"
)

// listenCommandChannelAndWait listens on channel for instructions or timeout and go to next step of sequence.
func ListenCommandChannelAndWait(
	sequence common.Sequence,
	commandChannel chan common.Command,
	replyChannel chan common.Command,
	soundTriggerChannel chan common.Command,
	soundTriggerControls *sound.Sound) common.Sequence {

	// Create an empty command.
	command := common.Command{}

	currentSequence := sequence

	if sequence.Number == 1 && soundTriggerControls.SendSoundToSequence1 {
		sequence.CurrentSpeed = time.Duration(12 * time.Hour)
		sequence.MusicTrigger = true
	}
	if sequence.Number == 2 && soundTriggerControls.SendSoundToSequence2 {
		sequence.CurrentSpeed = time.Duration(12 * time.Hour)
		sequence.MusicTrigger = true
	}
	if sequence.Number == 3 && soundTriggerControls.SendSoundToSequence3 {
		sequence.CurrentSpeed = time.Duration(12 * time.Hour)
		sequence.MusicTrigger = true
	}
	if sequence.Number == 4 && soundTriggerControls.SendSoundToSequence4 {
		sequence.CurrentSpeed = time.Duration(12 * time.Hour)
		sequence.MusicTrigger = true
	}

	select {
	case command = <-soundTriggerChannel:
		break
	case command = <-commandChannel:
		break
	case <-time.After(sequence.CurrentSpeed):
		break
	}

	if command.UpdateSpeed {
		fmt.Printf("Received update speed command %d\n", command.Speed)
		sequence = currentSequence
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		sequence.Run = true
		return sequence
	}

	if command.UpdatePatten {
		savePattenName := command.Patten.Name
		fmt.Printf("Received update pattten %s\n", savePattenName)
		sequence = currentSequence
		sequence.Patten.Name = savePattenName
		return sequence
	}

	if command.UpdateFade {
		fadeTime := command.FadeTime
		fmt.Printf("Received new fade time of %v\n", fadeTime)
		sequence = currentSequence
		sequence.FadeTime = fadeTime
		return sequence
	}

	if command.Start {
		fmt.Printf("Received Start Command\n")
		sequence = currentSequence
		sequence.CurrentSpeed = SetSpeed(command.Speed)
		sequence.Run = true
		return sequence
	}

	if command.Stop {
		fmt.Printf("Received Stop Command\n")
		sequence.Run = false
		return sequence
	}

	// If we are being asked for our config we must replay with
	// the sequence inside our command.
	if command.ReadConfig {
		replayCommand := common.Command{}
		fmt.Printf("Sending Reply on %d\n", replayCommand.Number)
		replayCommand.X = command.X
		replayCommand.Y = command.Y
		replayCommand.Sequence = sequence
		replyChannel <- replayCommand
		return sequence
	}

	if command.LoadConfig {
		X := command.X
		Y := command.Y
		fmt.Printf("LoadConfig: Seq No %d, Load Config %d.%d.json\n", sequence.Number, X, Y)
		config := config.ReadConfig(fmt.Sprintf("config%d.%d.json", X, Y))

		for _, seq := range config {
			if seq.Number == sequence.Number {
				if seq.Number == 1 {
					soundTriggerControls.SendSoundToSequence1 = seq.MusicTrigger
				}
				if seq.Number == 2 {
					soundTriggerControls.SendSoundToSequence2 = seq.MusicTrigger
				}
				if seq.Number == 3 {
					soundTriggerControls.SendSoundToSequence3 = seq.MusicTrigger
				}
				if seq.Number == 4 {
					soundTriggerControls.SendSoundToSequence4 = seq.MusicTrigger
				}

				sequence = seq
			}
		}
		return sequence
	}
	return sequence
}

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
		Speed = 1000
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
		Speed = 100
	}
	if commandSpeed == 11 {
		Speed = 50
	}
	if commandSpeed == 12 {
		Speed = 25
	}
	return Speed * time.Millisecond
}

func SetFade(commandSpeed int) (Speed time.Duration) {
	if commandSpeed == 0 {
		Speed = 1000
	}
	if commandSpeed == 1 {
		Speed = 900
	}
	if commandSpeed == 2 {
		Speed = 800
	}
	if commandSpeed == 3 {
		Speed = 700
	}
	if commandSpeed == 4 {
		Speed = 600
	}
	if commandSpeed == 5 {
		Speed = 500
	}
	if commandSpeed == 6 {
		Speed = 400
	}
	if commandSpeed == 7 {
		Speed = 300
	}
	if commandSpeed == 8 {
		Speed = 200
	}
	if commandSpeed == 9 {
		Speed = 150
	}
	if commandSpeed == 10 {
		Speed = 100
	}
	if commandSpeed == 11 {
		Speed = 50
	}
	if commandSpeed == 12 {
		Speed = 25
	}
	return Speed * time.Millisecond
}
