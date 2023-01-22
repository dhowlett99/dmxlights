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
	SettingsPanel     *widget.Table
	SettingsList      []fixture.Setting
	SettingsOptions   []string
	ChannelOptions    []string
	CurrentChannel    int
	UpdateThisChannel int
	UpdateSettings    bool
}

const SETTING_NUMBER int = 0
const SETTING_NAME int = 1
const SETTING_CHANNEL int = 2
const SETTING_VALUE int = 3
const SETTING_DELETE int = 4
const SETTING_ADD int = 5

func NewSettingsPanel(SettingsList []fixture.Setting, channelFieldDisabled bool) *SettingsPanel {

	if debug {
		fmt.Printf("NewSettingsPanel\n")
	}

	var data = [][]string{}

	st := SettingsPanel{}
	st.SettingsList = SettingsList
	st.SettingsOptions = []string{"Off", "On", "Red", "Green", "Blue", "Soft", "Sharp", "Sound", "Rotate"}
	st.ChannelOptions = []string{"None"}

	// Settingses Selection Panel.
	st.SettingsPanel = widget.NewTable(

		func() (int, int) {
			height := len(data)
			width := 6
			return height, width
		},
		// Function to create table.
		func() (o fyne.CanvasObject) {

			// Load the fixtures into the array used by the table.
			data = makeSettingsArray(st.SettingsList)

			return container.NewMax(
				widget.NewLabel("template"), // Setting Number.
				widget.NewEntry(),           // Setting Name.
				widget.NewSelect(st.ChannelOptions, func(value string) {}), // Setting Value.// Channel Number.
				widget.NewEntry(),                // Setting Value.
				widget.NewButton("-", func() {}), // Delete this Setting.
				widget.NewButton("+", func() {}), // Add a new Setting below.
			)
		},

		// Function to update item in this table.
		func(i widget.TableCellID, o fyne.CanvasObject) {

			// Hide all field types.
			hideAllSettingsFields(o)

			// Show the setting a number.
			if i.Col == SETTING_NUMBER {
				showSettingsField(SETTING_NUMBER, o)
				o.(*fyne.Container).Objects[SETTING_NUMBER].(*widget.Label).SetText(data[i.Row][i.Col])
			}

			// Show and Edit the Name.
			if i.Col == SETTING_NAME {
				showSettingsField(SETTING_NAME, o)
				o.(*fyne.Container).Objects[SETTING_NAME].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[SETTING_NAME].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_NAME].(*widget.Entry).OnChanged = func(value string) {
					newSetting := fixture.Setting{}
					newSetting.Label = st.SettingsList[i.Row].Label
					newSetting.Name = value
					newSetting.Number = st.SettingsList[i.Row].Number
					if !channelFieldDisabled {
						newSetting.Channel = st.SettingsList[i.Row].Channel
					}
					newSetting.Value = st.SettingsList[i.Row].Value
					st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
				}
			}

			// Channel value.
			if i.Col == SETTING_CHANNEL {
				showSettingsField(SETTING_CHANNEL, o)
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).OnChanged = nil
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).SetSelected(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).Hidden = channelFieldDisabled
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).OnChanged = func(value string) {
					newSetting := fixture.Setting{}
					newSetting.Label = st.SettingsList[i.Row].Label
					newSetting.Name = st.SettingsList[i.Row].Name
					newSetting.Number = st.SettingsList[i.Row].Number
					newSetting.Channel = value
					newSetting.Value = st.SettingsList[i.Row].Value
					st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
				}
			}

			// Show and Edit the Setting Value.
			if i.Col == SETTING_VALUE {
				showSettingsField(SETTING_VALUE, o)
				o.(*fyne.Container).Objects[SETTING_VALUE].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[SETTING_VALUE].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_VALUE].(*widget.Entry).OnChanged = func(value string) {
					newSetting := fixture.Setting{}
					newSetting.Label = st.SettingsList[i.Row].Label
					newSetting.Name = st.SettingsList[i.Row].Name
					newSetting.Number = st.SettingsList[i.Row].Number
					if !channelFieldDisabled {
						newSetting.Channel = st.SettingsList[i.Row].Channel
					}
					newSetting.Value = value
					st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
				}
			}

			// Show the Delete Setting Button.
			if i.Col == SETTING_DELETE {
				showSettingsField(SETTING_DELETE, o)
				o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).OnTapped = func() {
					st.SettingsList = deleteSettingsItem(st.SettingsList, st.SettingsList[i.Row].Number-1)
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
					st.SettingsPanel.Refresh()
				}
			}

			// Show the Add Setting Button.
			if i.Col == SETTING_ADD {
				showSettingsField(SETTING_ADD, o)
				o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).OnTapped = func() {
					st.SettingsList = addSettingsItem(st.SettingsList, st.SettingsList[i.Row].Number, st.SettingsOptions)
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
					st.SettingsPanel.Refresh()
				}
			}
		},
	)

	// Setup the columns of this table.
	st.SettingsPanel.SetColumnWidth(0, 40)  // Number
	st.SettingsPanel.SetColumnWidth(1, 100) // Name
	if channelFieldDisabled {
		st.SettingsPanel.SetColumnWidth(2, 0) // Channel
	} else {
		st.SettingsPanel.SetColumnWidth(2, 100) // Channel
	}
	st.SettingsPanel.SetColumnWidth(3, 50) // Value
	st.SettingsPanel.SetColumnWidth(4, 20) // Delete
	st.SettingsPanel.SetColumnWidth(5, 20) // Add

	return &st
}

func settingItemAllreadyExists(number int, settingsList []fixture.Setting) bool {

	if debug {
		fmt.Printf("settingItemAllreadyExists\n")
	}

	// look through the settings list for the id's
	for _, item := range settingsList {
		if item.Number == number {
			return true
		}
	}
	return false
}

func findLargestsettingsNumber(items []fixture.Setting) int {

	if debug {
		fmt.Printf("findLargestsettingsNumber\n")
	}

	var number int
	for _, item := range items {
		if item.Number > number {
			number = item.Number
		}
	}
	return number
}

func addSettingsItem(items []fixture.Setting, id int, options []string) []fixture.Setting {

	if debug {
		fmt.Printf("addSettingsItem\n")
	}

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

	if debug {
		fmt.Printf("updateSettingsItem\n")
	}

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

// makeSettingsArray - Convert the list of settings to an array of strings containing and array of strings with
// the values from each fixture.
// This is done once when the settings panel is loaded.
func makeSettingsArray(settings []fixture.Setting) [][]string {

	if debug {
		fmt.Printf("makeSettingsArray\n")
	}

	var data = [][]string{}

	for _, setting := range settings {
		newSetting := []string{}
		newSetting = append(newSetting, fmt.Sprintf("%d", setting.Number))
		newSetting = append(newSetting, setting.Name)
		newSetting = append(newSetting, setting.Channel)
		newSetting = append(newSetting, setting.Value)
		newSetting = append(newSetting, "-")
		newSetting = append(newSetting, "+")
		newSetting = append(newSetting, "Channels")

		data = append(data, newSetting)
	}

	return data
}

func showSettingsField(field int, o fyne.CanvasObject) {
	if debug {
		fmt.Printf("showField\n")
	}
	// Now show the selected field.
	switch {
	case field == SETTING_NUMBER:
		o.(*fyne.Container).Objects[SETTING_NUMBER].(*widget.Label).Hidden = false
	case field == SETTING_NAME:
		o.(*fyne.Container).Objects[SETTING_NAME].(*widget.Entry).Hidden = false
	case field == SETTING_CHANNEL:
		o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).Hidden = false
	case field == SETTING_VALUE:
		o.(*fyne.Container).Objects[SETTING_VALUE].(*widget.Entry).Hidden = false
	case field == SETTING_DELETE:
		o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).Hidden = false
	case field == SETTING_ADD:
		o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).Hidden = false
	}
}

func hideAllSettingsFields(o fyne.CanvasObject) {
	if debug {
		fmt.Printf("hideAllSettingsFields\n")
	}
	o.(*fyne.Container).Objects[SETTING_NUMBER].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[SETTING_NAME].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).Hidden = true
	o.(*fyne.Container).Objects[SETTING_VALUE].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).Hidden = true
}
