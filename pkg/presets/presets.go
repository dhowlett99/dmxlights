package presets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func InitPresets(eventsForLauchpad chan common.ALight, presets map[string]bool) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				common.LightOn(eventsForLauchpad, common.ALight{X: x, Y: y, Brightness: 3, Red: 3, Green: 0, Blue: 0})
			}
		}
	}
}

func SavePresets(presets map[string]bool) {
	// Marshall the config into a json object.
	data, err := json.MarshalIndent(presets, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling config: %v", err)
	}

	// Write to file
	err = ioutil.WriteFile("presets.json", data, 0644)
	if err != nil {
		log.Fatalf("Error writing config: %v to file:%s", err, "presets.json")
	}
}

func LoadPresets() map[string]bool {

	presets := map[string]bool{}

	// Read the file.
	data, err := ioutil.ReadFile("presets.json")
	if err != nil {
		fmt.Printf("Error reading prests: %v from file:%s", err, "presets.json")
		return presets
	}

	err = json.Unmarshal(data, &presets)
	if err != nil {
		log.Fatalf("Error unmashalling presets: %v from file:%s", err, "presets.json")
	}

	return presets
}
