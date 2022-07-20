package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func SaveConfig(config []common.Sequence, filename string) {

	// Marshall the config into a json object.
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatalf("error: marshalling config: %v", err)
	}
	// Write to file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("error: writing config: %v to file:%s", err, filename)
	}
}

func LoadConfig(filename string) []common.Sequence {

	config := []common.Sequence{}

	// Read the file.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: reading config: %v from file:%s", err, filename)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error:  reading config: %v from file:%s", err, filename)
	}
	return config
}

func DeleteConfig(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		log.Fatalf("error:  deleting config: %v from file:%s", err, filename)
	}
	return nil
}

func AskToLoadConfig(commandChannels []chan common.Command, X int, Y int) {
	command := common.Command{
		Action: common.LoadConfig,
		Args: []common.Arg{
			{Name: "X", Value: X},
			{Name: "Y", Value: Y},
		},
	}

	for _, seq := range commandChannels {
		seq <- command
	}
}

func AskToSaveConfig(sequences []chan common.Command, replyChannel []chan common.Sequence, X int, Y int) {

	config := []common.Sequence{}

	go func() {
		// Wait for responses from sequences.
		time.Sleep(100 * time.Millisecond)
		for _, replyChannel := range replyChannel {
			config = append(config, WaitForConfig(replyChannel))
		}

		// Write to config file.
		SaveConfig(config, fmt.Sprintf("config%d.%d.json", X, Y))
	}()

	// Ask for all the sequencers for their config.
	command := common.Command{
		Action: common.ReadConfig,
		Args: []common.Arg{
			{Name: "X", Value: X},
			{Name: "Y", Value: Y},
		},
	}

	for _, seq := range sequences {
		seq <- command
	}
}

func WaitForConfig(replyChannel chan common.Sequence) common.Sequence {
	sequence := common.Sequence{}
	select {
	case sequence = <-replyChannel:
		break
	case <-time.After(500 * time.Millisecond):
		break
	}
	return sequence
}
