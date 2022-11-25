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
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type SettingsPanel struct {
	SettingsPanel      *widget.List
	SettingsList       []fixture.Setting
	SettingsOptions    []string
	CurrentChannel     int
	Update             bool
	UpdatedChannelList []fixture.Channel
}

func NewSettingsPanel(thisFixture fixture.Fixture, currentChannel *int, SettingsList []fixture.Setting, channelList []fixture.Channel, ap *ActionPanel) *SettingsPanel {

	st := SettingsPanel{}
	st.SettingsList = SettingsList
	st.SettingsOptions = []string{"Off", "On", "Red", "Green", "Blue", "SoftChase", "SharpChase", "SoundChase", "Rotate"}

	// Settingses Selection Panel.
	st.SettingsPanel = widget.NewList(
		func() int {
			return len(st.SettingsList)
		},
		// Function to create item.
		func() fyne.CanvasObject {
			return container.NewGridWithColumns(3,
				widget.NewLabel("template"),
				widget.NewEntry(),
				widget.NewEntry(),
				//widget.NewButton("Select", nil),
			)
		},

		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// Give the setting a number.
			st.SettingsList[i].Number = i + 1 // +1 so we count from 1 instead of 0.
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", st.SettingsList[i].Number))
			o.(*fyne.Container).Objects[1].(*widget.Entry).SetText(st.SettingsList[i].Name)
			o.(*fyne.Container).Objects[1].(*widget.Entry).OnChanged = func(value string) {
				newSetting := fixture.Setting{}
				newSetting.Label = st.SettingsList[i].Label
				newSetting.Name = value
				newSetting.Number = st.SettingsList[i].Number
				newSetting.Setting = st.SettingsList[i].Setting
				st.SettingsList = UpdateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
				channelList[st.CurrentChannel-1].Settings = st.SettingsList
				st.Update = true
				st.UpdatedChannelList = channelList
			}

			o.(*fyne.Container).Objects[2].(*widget.Entry).SetText(fmt.Sprintf("%d", st.SettingsList[i].Setting))
			o.(*fyne.Container).Objects[2].(*widget.Entry).OnChanged = func(value string) {
				newSetting := fixture.Setting{}
				newSetting.Label = st.SettingsList[i].Label
				newSetting.Name = st.SettingsList[i].Name
				newSetting.Number = st.SettingsList[i].Number
				newSetting.Setting, _ = strconv.Atoi(value)
				st.SettingsList = UpdateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
				channelList[st.CurrentChannel-1].Settings = st.SettingsList
				st.Update = true
				st.UpdatedChannelList = channelList
			}
		},
	)

	return &st
}

func SettingItemAllreadyExists(number int, settingsList []fixture.Setting) bool {
	// look through the settings list for the id's
	for _, item := range settingsList {
		if item.Number == number {
			return true
		}
	}
	return false
}

func FindLargestsettingsNumber(items []fixture.Setting) int {
	var number int
	for _, item := range items {
		if item.Number > number {
			number = item.Number
		}
	}
	return number
}

func AddSettingsItem(items []fixture.Setting, id int, options []string) []fixture.Setting {
	newItems := []fixture.Setting{}
	newItem := fixture.Setting{}
	newItem.Number = int(id) + 1
	if SettingItemAllreadyExists(newItem.Number, items) {
		newItem.Number = FindLargestsettingsNumber(items) + 1
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

func UpdateSettingsItem(items []fixture.Setting, id int, newItem fixture.Setting) []fixture.Setting {
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
