// Copyright (C) 2022-2026 dhowlett99.
// This is part of the dmxlights channel editor it is used to load and save fixture configuration from a file.
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

package editor

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// showDirectoryPopup reads dirPath and displays its contents in a modal popup.
// When an item is clicked, onSelect is called with the full path and the popup closes.
func showDirectoryPopup(dirPath string, parent fyne.Window, onSelect func(selectedPath string)) dialog.Dialog {
	var popup dialog.Dialog

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		dialog.ShowError(err, parent)
		return nil
	}

	list := widget.NewList(
		func() int {
			return len(entries)
		},
		func() fyne.CanvasObject {
			return widget.NewLabelWithStyle("template", fyne.TextAlignLeading, fyne.TextStyle{})
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			name := entries[id].Name()
			if entries[id].IsDir() {
				name += "/"
			}
			label.SetText(name)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		fullPath := filepath.Join(dirPath, entries[id].Name())
		if popup != nil {
			popup.Hide()
		}
		if onSelect != nil {
			onSelect(fullPath)
		}
	}

	popup = dialog.NewCustom("Fixture Library", "Close", list, parent)

	return popup
}

func WriteFixtureConfigFile(cp *ChannelPanel, configFileName string) {
	// Marshall the config into a json object.
	data, err := json.MarshalIndent(cp, "", " ")
	if err != nil {
		log.Fatalf("error: marshalling config: %v", err)
	}

	// Get the preset filename for this project.
	path := GetPresetNamePath(configFileName)

	// Write to file
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		log.Fatalf("error: writing fixture config: %v to file:%s", err, path)
	}
}

func GetPresetNamePath(configFileName string) string {
	path := "./fixturesLibrary/" + configFileName + ".json"
	return path
}
