// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights presets mechanism, controlling the saving and
// recalling of sequence configurations.
// All presets are saved in configX,Y.json files.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package presets

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

type Preset struct {
	State       bool   `json:"state"`
	Selected    bool   `json:"-"`
	Label       string `json:"label"`
	ButtonColor string `json:"buttoncolor"`
}

// RefeshPresets is used to refresh the view of presets.
func RefreshPresets(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, presets map[string]Preset) {
	for y := 4; y < 7; y++ {
		for x := 0; x < 8; x++ {
			// State true is a preset which is being used and has a saved config.
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].State {
				// Selected preset is set to it's flashing color.
				if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].Selected {
					// Selected.
					if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].ButtonColor == "" {
						// There's no color defined so flash red & yellow.
						common.FlashLight(x, y, common.Red, common.PresetYellow, eventsForLauchpad, guiButtons)
					} else {
						// There is a color in the presets datatbase so set the color
						color, _ := common.GetRGBColorByName(presets[fmt.Sprint(x)+","+fmt.Sprint(y)].ButtonColor)
						common.FlashLight(x, y, color, common.PresetYellow, eventsForLauchpad, guiButtons)
					}
				} else {
					// Not Selected and there's no button color defined so just light the lamp red.
					if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].ButtonColor == "" {
						color := common.Red
						common.LightLamp(common.ALight{X: x, Y: y, Flash: false, Red: color.R, Green: color.G, Blue: color.B, Brightness: 255}, eventsForLauchpad, guiButtons)

					} else {
						// We're not selected and there is a button color specified so set that color.
						color, _ := common.GetRGBColorByName(presets[fmt.Sprint(x)+","+fmt.Sprint(y)].ButtonColor)
						common.LightLamp(common.ALight{X: x, Y: y, Flash: false, Red: color.R, Green: color.G, Blue: color.B, Brightness: 255}, eventsForLauchpad, guiButtons)
					}
				}
			} else {
				// Unused preset is set to yellow.
				common.LightLamp(common.ALight{X: x, Y: y, Red: 150, Green: 150, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons)
			}
			common.LabelButton(x, y, presets[fmt.Sprint(x)+","+fmt.Sprint(y)].Label, guiButtons)
		}
	}
}

// ClearPresets is used to un-select all presets.
func ClearPresets(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, presets map[string]Preset) {
	for y := 4; y < 7; y++ {
		for x := 0; x < 8; x++ {
			newPreset := presets[fmt.Sprint(x)+","+fmt.Sprint(y)]
			newPreset.Selected = false
			presets[fmt.Sprint(x)+","+fmt.Sprint(y)] = newPreset
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
	err = os.WriteFile("presets.json", data, 0644)
	if err != nil {
		log.Fatalf("error: writing config: %v to file:%s", err, "presets.json")
	}
}

func LoadPresets() map[string]Preset {

	presets := map[string]Preset{}

	// Read the file.
	data, err := os.ReadFile("presets.json")
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
