package command

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
)

// listenCommandChannelAndWait listens on channel for instructions or timeout and go to next step of sequence.
func ListenCommandChannelAndWait(command common.Sequence, commandChannel chan common.Sequence, replyChannel chan common.Sequence, CurrentSpeed time.Duration, mySequenceNumber int) common.Sequence {

	currentCommand := command
	select {
	case command = <-commandChannel:
		//fmt.Printf("COMMAND\n")
		break
	case <-time.After(CurrentSpeed):
		//fmt.Printf("TIMEOUT\n")
		break
	}
	if command.UpdateSpeed {
		saveSpeed := command.Speed
		fmt.Printf("Received update speed %d\n", saveSpeed)
		CurrentSpeed = setSpeed(command.Speed)
		command = currentCommand
		command.CurrentSpeed = CurrentSpeed
		command.Speed = saveSpeed
	}

	if command.UpdatePatten {
		savePattenName := command.Patten.Name
		fmt.Printf("Received update pattten %s\n", savePattenName)
		command = currentCommand
		command.Patten.Name = savePattenName
		command.UpdatePatten = true
	}

	if command.UpdateFade {
		fadeTime := command.FadeTime
		fmt.Printf("Received new fade time of %v\n", fadeTime)
		command = currentCommand
		command.FadeTime = fadeTime
		command.UpdateFade = true
	}

	if command.Start {
		fmt.Printf("Received Start Seq \n")
		command = currentCommand
		command.Run = true
	}

	if command.Stop {
		fmt.Printf("Received Stop Seq \n")
		command = currentCommand
		command.Run = false
	}

	if command.ReadConfig {
		fmt.Printf("Sending Reply on %d\n", currentCommand.Number)
		currentCommand.X = command.X
		currentCommand.Y = command.Y
		replyChannel <- currentCommand
		command = currentCommand
	}

	if command.LoadConfig {
		X := command.X
		Y := command.Y
		config := config.ReadConfig(fmt.Sprintf("config%d.%d.json", X, Y))

		for _, seq := range config {
			if seq.Number == mySequenceNumber {
				command = seq
			}
		}
		command.LoadConfig = true
	}
	return command
}
