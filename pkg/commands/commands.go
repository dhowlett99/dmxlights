package commands

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
		CurrentSpeed = SetSpeed(command.Speed)
		command = currentCommand
		command.CurrentSpeed = CurrentSpeed
		command.Speed = saveSpeed
	}

	if command.UpdatePatten {
		savePattenName := command.Patten.Name
		fmt.Printf("Received update pattten %s\n", savePattenName)
		command = currentCommand
		command.Patten.Name = savePattenName
		command.UpdatePatten = false
	}

	if command.UpdateFade {
		fadeTime := command.FadeTime
		fmt.Printf("Received new fade time of %v\n", fadeTime)
		command = currentCommand
		command.FadeTime = fadeTime
		command.UpdateFade = true
	}

	if command.Start {
		fmt.Printf("Received Start Command\n")
		command = currentCommand
		command.Run = true
		command.Start = false
	}

	if command.Stop {
		fmt.Printf("Received Stop Command\n")
		command = currentCommand
		command.Stop = false
		command.Start = false
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
		//command.LoadConfig = true
	}
	return command
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
	return Speed * time.Millisecond
}
