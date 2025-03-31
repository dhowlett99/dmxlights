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

	"github.com/dhowlett99/dmxlights/pkg/colors"
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
						common.FlashLight(common.Button{X: x, Y: y}, colors.Red, colors.PresetYellow, eventsForLauchpad, guiButtons)
					} else {
						// There is a color in the presets datatbase so set the color
						color := common.GetRGBColorByName(presets[fmt.Sprint(x)+","+fmt.Sprint(y)].ButtonColor)
						common.FlashLight(common.Button{X: x, Y: y}, color, colors.PresetYellow, eventsForLauchpad, guiButtons)
					}
				} else {
					// Not Selected and there's no button color defined so just light the lamp red.
					if presets[fmt.Sprint(x)+","+fmt.Sprint(y)].ButtonColor == "" {
						color := colors.Red
						common.LightLamp(common.Button{X: x, Y: y}, color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)

					} else {
						// We're not selected and there is a button color specified so set that color.
						color := common.GetRGBColorByName(presets[fmt.Sprint(x)+","+fmt.Sprint(y)].ButtonColor)
						common.LightLamp(common.Button{X: x, Y: y}, color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
					}
				}
			} else {
				// Unused preset is set to yellow.
				common.LightLamp(common.Button{X: x, Y: y}, colors.PresetYellow, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
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

func RemovePreset(presets map[string]Preset, x int, y int) {
	newPreset := presets[fmt.Sprint(x)+","+fmt.Sprint(y)]
	newPreset.Selected = false
	presets[fmt.Sprint(x)+","+fmt.Sprint(y)] = newPreset
}

func SavePresets(presets map[string]Preset, projectName string) {
	// Marshall the config into a json object.
	data, err := json.MarshalIndent(presets, "", " ")
	if err != nil {
		log.Fatalf("error: marshalling config: %v", err)
	}

	// Get the preset filename for this project.
	path := GetPresetNamePath(projectName)

	// Write to file
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		log.Fatalf("error: writing config: %v to file:%s", err, "presets.json")
	}
}

func GetPresetNamePath(projectName string) string {
	path := "projects/." + projectName + "/presets.json"
	return path
}

func LoadPresets(projectName string) (map[string]Preset, error) {

	path := GetPresetNamePath(projectName)

	presets := map[string]Preset{}

	// Read the file.
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading presets: %v from file:%s\n", err, "presets.json")
		return presets, err
	}

	// Unmarshal into presets map.
	err = json.Unmarshal(data, &presets)
	if err != nil {
		log.Fatalf("error unmashalling presets: %v from file:%s", err, "presets.json")
		return presets, err
	}

	return presets, nil
}

func GetPresetNumber(X int, Y int) string {

	switch Y {
	case 5:
		switch X {
		case 0:
			return "Preset1"
		case 1:
			return "Preset2"
		case 2:
			return "Preset3"
		case 3:
			return "Preset4"
		case 4:
			return "Preset5"
		case 5:
			return "Preset6"
		case 6:
			return "Preset7"
		case 7:
			return "Preset8"
		}

	case 6:
		switch X {
		case 0:
			return "Preset9"
		case 1:
			return "Preset10"
		case 2:
			return "Preset11"
		case 3:
			return "Preset12"
		case 4:
			return "Preset13"
		case 5:
			return "Preset14"
		case 6:
			return "Preset15"
		case 7:
			return "Preset16"
		}

	case 7:
		switch X {
		case 0:
			return "Preset17"
		case 1:
			return "Preset18"
		case 2:
			return "Preset19"
		case 3:
			return "Preset20"
		case 4:
			return "Preset21"
		case 5:
			return "Preset22"
		case 6:
			return "Preset23"
		case 7:
			return "Preset24"
		}
	}
	return ""
}
