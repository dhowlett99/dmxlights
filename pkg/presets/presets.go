package presets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func InitPresets(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, presets map[string]bool, GuiFlashButtons [][]bool) {
	for y := 4; y < 7; y++ {
		for x := 0; x < 8; x++ {
			common.LightLamp(common.ALight{X: x, Y: y, Red: 100, Green: 100, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons, GuiFlashButtons)
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				common.LightLamp(common.ALight{X: x, Y: y, Red: 255, Green: 0, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons, GuiFlashButtons)
			}
		}
	}
}

func ClearPresets(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, presets map[string]bool, GuiFlashButtons [][]bool) {
	for y := 4; y < 7; y++ {
		for x := 0; x < 8; x++ {
			common.LightLamp(common.ALight{X: x, Y: y, Red: 100, Green: 100, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons, GuiFlashButtons)
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				common.LightLamp(common.ALight{X: x, Y: y, Red: 100, Green: 100, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons, GuiFlashButtons)
			}
		}
	}
}

func SavePresets(presets map[string]bool) {
	// Marshall the config into a json object.
	data, err := json.MarshalIndent(presets, "", " ")
	if err != nil {
		log.Fatalf("error: marshalling config: %v", err)
	}

	// Write to file
	err = ioutil.WriteFile("presets.json", data, 0644)
	if err != nil {
		log.Fatalf("error: writing config: %v to file:%s", err, "presets.json")
	}
}

func LoadPresets() map[string]bool {

	presets := map[string]bool{}

	// Read the file.
	data, err := ioutil.ReadFile("presets.json")
	if err != nil {
		fmt.Printf("error reading presets: %v from file:%s\n", err, "presets.json")
		return presets
	}

	err = json.Unmarshal(data, &presets)
	if err != nil {
		log.Fatalf("error unmashalling presets: %v from file:%s", err, "presets.json")
	}

	return presets
}
