package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func SaveConfig(config []common.Sequence, filename string) {

	// Marshall the config into a json object.
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling config: %v", err)
	}
	// Write to file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Error writing config: %v to file:%s", err, filename)
	}
}

func LoadConfig(filename string) []common.Sequence {

	config := []common.Sequence{}

	// Read the file.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading config: %v from file:%s", err, filename)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error reading config: %v from file:%s", err, filename)
	}
	return config
}

func AskToLoadConfig(commandChannels []chan common.Command, X int, Y int) {
	fmt.Printf("AskToLoadConfig\n")
	cmd := common.Command{
		LoadConfig: true,
		X:          X,
		Y:          Y,
	}
	for _, seq := range commandChannels {
		seq <- cmd
	}
}

func AskToSaveConfig(sequences []chan common.Command, replyChannel []chan common.Sequence, X int, Y int) {

	fmt.Printf("askToSaveConfig: Save Preset in X:%d Y:%d \n", X, Y)
	config := []common.Sequence{}

	go func() {
		// wait for responses from sequences.
		time.Sleep(100 * time.Millisecond)
		for _, replyChannel := range replyChannel {
			config = append(config, WaitForConfig(replyChannel))
		}

		// write to config file.
		SaveConfig(config, fmt.Sprintf("config%d.%d.json", X, Y))
	}()

	// ask for all the sequencers for their config.
	cmd := common.Command{
		ReadConfig: true,
		X:          X,
		Y:          Y,
	}
	for _, seq := range sequences {
		seq <- cmd
	}
}

// waitForConfig
func WaitForConfig(replyChannel chan common.Sequence) common.Sequence {
	sequence := common.Sequence{}
	select {
	case sequence = <-replyChannel:
		fmt.Printf("Config Received for seq: %s\n", sequence.Name)
		break
	case <-time.After(500 * time.Millisecond):
		break
	}
	return sequence
}
