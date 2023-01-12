// Copyright (C) 2022 dhowlett99.
// This is the dmxlights fixture editor it is attached to a fixture and
// describes the fixtures properties which is then saved in the fixtures.yaml
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
	"fmt"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type SettingsPanel struct {
	SettingsPanel     *widget.List
	SettingsList      []fixture.Setting
	SettingsOptions   []string
	CurrentChannel    int
	UpdateThisChannel int
	UpdateSettings    bool
}

func NewSettingsPanel(SettingsList []fixture.Setting, channelFieldDisabled bool) *SettingsPanel {

	var SETTING_NUMBER int
	var SETTING_NAME int
	var SETTING_CHANNEL int
	var SETTING_VALUE int
	var SETTING_DELETE int
	var SETTING_ADD int

	st := SettingsPanel{}
	st.SettingsList = SettingsList
	st.SettingsOptions = []string{"Off", "On", "Red", "Green", "Blue", "Soft", "Sharp", "Sound", "Rotate"}

	// Settingses Selection Panel.
	st.SettingsPanel = widget.NewList(
		func() int {
			return len(st.SettingsList)
		},
		// Function to create item.
		func() fyne.CanvasObject {

			if channelFieldDisabled {

				SETTING_NUMBER = 0
				SETTING_NAME = 1
				SETTING_VALUE = 2
				SETTING_DELETE = 3
				SETTING_ADD = 4

				return container.NewGridWithColumns(5,
					widget.NewLabel("template"),      // Setting Number.
					widget.NewEntry(),                // Setting Name.
					widget.NewEntry(),                // Setting Value.
					widget.NewButton("-", func() {}), // Delete this Setting.
					widget.NewButton("+", func() {}), // Add a new Setting below.
				)
			}

			SETTING_NUMBER = 0
			SETTING_NAME = 1
			SETTING_CHANNEL = 2
			SETTING_VALUE = 3
			SETTING_DELETE = 4
			SETTING_ADD = 5

			return container.NewGridWithColumns(5,
				widget.NewLabel("template"),      // Setting Number.
				widget.NewEntry(),                // Setting Name.
				widget.NewEntry(),                // Channel Number.
				widget.NewEntry(),                // Setting Value.
				widget.NewButton("-", func() {}), // Delete this Setting.
				widget.NewButton("+", func() {}), // Add a new Setting below.
			)
		},

		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// Show the setting a number.
			o.(*fyne.Container).Objects[SETTING_NUMBER].(*widget.Label).SetText(fmt.Sprintf("%d", st.SettingsList[i].Number))

			// Show and Edit the Name.
			o.(*fyne.Container).Objects[SETTING_NAME].(*widget.Entry).SetText(st.SettingsList[i].Name)
			o.(*fyne.Container).Objects[SETTING_NAME].(*widget.Entry).OnChanged = func(value string) {
				newSetting := fixture.Setting{}
				newSetting.Label = st.SettingsList[i].Label
				newSetting.Name = value
				newSetting.Number = st.SettingsList[i].Number
				if !channelFieldDisabled {
					newSetting.Channel = st.SettingsList[i].Channel
				}
				newSetting.Setting = st.SettingsList[i].Setting
				st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
				st.UpdateSettings = true
				st.UpdateThisChannel = st.CurrentChannel - 1
			}

			if !channelFieldDisabled {
				// Channel value.
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Entry).Hidden = channelFieldDisabled
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Entry).SetText(st.SettingsList[i].Channel)
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Entry).OnChanged = func(value string) {
					newSetting := fixture.Setting{}
					newSetting.Label = st.SettingsList[i].Label
					newSetting.Name = st.SettingsList[i].Name
					newSetting.Number = st.SettingsList[i].Number
					newSetting.Channel = value
					newSetting.Setting = st.SettingsList[i].Setting
					st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
				}
			}

			// Show and Edit the Setting Value.
			o.(*fyne.Container).Objects[SETTING_VALUE].(*widget.Entry).SetText(st.SettingsList[i].Setting)
			o.(*fyne.Container).Objects[2].(*widget.Entry).OnChanged = func(value string) {
				newSetting := fixture.Setting{}
				newSetting.Label = st.SettingsList[i].Label
				newSetting.Name = st.SettingsList[i].Name
				newSetting.Number = st.SettingsList[i].Number
				if !channelFieldDisabled {
					newSetting.Channel = st.SettingsList[i].Channel
				}
				newSetting.Setting = value
				st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
				st.UpdateSettings = true
				st.UpdateThisChannel = st.CurrentChannel - 1
			}

			// Show the Delete Setting Button.
			o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).OnTapped = func() {
				st.SettingsList = deleteSettingsItem(st.SettingsList, st.SettingsList[i].Number-1)
				st.UpdateSettings = true
				st.UpdateThisChannel = st.CurrentChannel - 1
				st.SettingsPanel.Refresh()
			}

			// Show the Add Setting Button.
			o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).OnTapped = func() {
				st.SettingsList = addSettingsItem(st.SettingsList, st.SettingsList[i].Number, st.SettingsOptions)
				st.UpdateSettings = true
				st.UpdateThisChannel = st.CurrentChannel - 1
				st.SettingsPanel.Refresh()
			}
		},
	)

	return &st
}

func settingItemAllreadyExists(number int, settingsList []fixture.Setting) bool {
	// look through the settings list for the id's
	for _, item := range settingsList {
		if item.Number == number {
			return true
		}
	}
	return false
}

func findLargestsettingsNumber(items []fixture.Setting) int {
	var number int
	for _, item := range items {
		if item.Number > number {
			number = item.Number
		}
	}
	return number
}

func addSettingsItem(items []fixture.Setting, id int, options []string) []fixture.Setting {
	newItems := []fixture.Setting{}
	newItem := fixture.Setting{}
	newItem.Number = int(id) + 1
	if settingItemAllreadyExists(newItem.Number, items) {
		newItem.Number = findLargestsettingsNumber(items) + 1
	}
	newItem.Name = "New"

	for _, item := range items {
		if item.Number == id {
			newItems = append(newItems, newItem)
		}
		newItems = append(newItems, item)
	}
	sort.Slice(newItems, func(i, j int) bool {
		return newItems[i].Number < newItems[j].Number
	})
	return newItems
}

func updateSettingsItem(items []fixture.Setting, id int, newItem fixture.Setting) []fixture.Setting {
	newItems := []fixture.Setting{}
	for _, item := range items {
		if item.Number == id {
			// update the settings information.
			newItems = append(newItems, newItem)
		} else {
			// just add what was there before.
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func deleteSettingsItem(settingsList []fixture.Setting, id int) []fixture.Setting {
	newSettings := []fixture.Setting{}
	for settingNumber, setting := range settingsList {
		if settingNumber != id {
			newSettings = append(newSettings, setting)
		}
	}
	return newSettings
}
