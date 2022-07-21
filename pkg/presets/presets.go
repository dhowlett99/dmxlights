package presets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/rakyll/launchpad/mk3"
)

type Preset struct {
	State    bool   `json:"state"`
	Selected bool   `json:"-"`
	Label    string `json:"label"`
}

func InitPresets(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, presets map[string]Preset) {
	for y := 4; y < 7; y++ {
		for x := 0; x < 8; x++ {
			// Set to Preset Yellow.
			common.LightLamp(common.ALight{X: x, Y: y, Red: 150, Green: 150, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons)
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].State {
				if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].Selected {
					common.FlashLight(x, y, common.Red, common.PresetYellow, eventsForLauchpad, guiButtons)
				} else {
					common.LightLamp(common.ALight{X: x, Y: y, Red: 255, Green: 0, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons)
				}
			}
			common.LabelButton(x, y, presets[fmt.Sprint(x)+","+fmt.Sprint(y)].Label, guiButtons)
		}
	}
}

func SavePresets(presets map[string]Preset) {
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

func LoadPresets() map[string]Preset {

	presets := map[string]Preset{}

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

func ClearAll(pad *mk3.Launchpad, presets map[string]Preset, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, sequences []chan common.Command) {
	command := common.Command{
		Action: common.Stop,
	}

	for _, sequence := range sequences {
		sequence <- command
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].State {
				label := presets[fmt.Sprint(x)+","+fmt.Sprint(y)].Label
				presets[fmt.Sprint(x)+","+fmt.Sprint(y)] = Preset{State: true, Selected: false, Label: label}
				common.LightLamp(common.ALight{X: x, Y: y, Red: 255, Green: 0, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons)
			}
		}
	}
}
