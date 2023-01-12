// Copyright (C) 2022,2023 dhowlett99.
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

type ChannelPanel struct {
	ChannelPanel   *widget.List
	ChannelList    []fixture.Channel
	ChannelOptions []string
}

func NewChannelPanel(thisFixture fixture.Fixture, channels []fixture.Channel, st *SettingsPanel) *ChannelPanel {

	cp := ChannelPanel{}
	cp.ChannelOptions = []string{"Rotate", "Red1", "Red2", "Red3", "Red4", "Red5", "Red6", "Red7", "Red8", "Green1", "Green2", "Green3", "Green4", "Green5", "Green6", "Green7", "Green8", "Blue1", "Blue2", "Blue3", "Blue4", "Blue5", "Blue6", "Blue7", "Blue8", "White1", "White2", "White3", "White4", "White5", "White6", "White7", "White8", "Master", "Dimmer", "Static", "Pan", "FinePan", "Tilt", "FineTilt", "Shutter", "Strobe", "Color", "Gobo", "Program", "ProgramSpeed", "Programs", "ColorMacros", "SoundActive", "DimmerCurve", "Speed"}
	cp.ChannelList = channels

	// Channel or Switch State Selection Panel.
	cp.ChannelPanel = widget.NewList(
		// Function to find length.
		func() int {
			if st.UpdateSettings {
				cp.ChannelList[st.UpdateThisChannel].Settings = st.SettingsList
				st.UpdateSettings = false
			}
			return len(cp.ChannelList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {
			return container.NewHBox(

				// Channel Number.
				widget.NewLabel("template"),

				// Channel Value as a selectable dialog box.
				widget.NewSelect(cp.ChannelOptions, func(value string) {}),

				// Chanell delete button.
				widget.NewButton("-", func() {}),

				// Channel add button
				widget.NewButton("+", func() {}),

				// Channel access settings for this channel button.
				widget.NewButton("settings", func() {}),
			)
		},
		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// Show the Channel Number.
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", cp.ChannelList[i].Number))

			// Show the currently selected Channel option.
			for _, option := range cp.ChannelOptions {
				if option == cp.ChannelList[i].Name {
					o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
				}
			}
			// Edit the channel Value.
			o.(*fyne.Container).Objects[1].(*widget.Select).OnChanged = func(value string) {
				newChannel := fixture.Channel{}
				newChannel.Name = value
				newChannel.Number = cp.ChannelList[i].Number
				newChannel.Value = cp.ChannelList[i].Value
				newChannel.Settings = cp.ChannelList[i].Settings
				newChannel.Comment = cp.ChannelList[i].Comment
				newChannel.MaxDegrees = cp.ChannelList[i].MaxDegrees
				newChannel.Offset = cp.ChannelList[i].Offset
				cp.ChannelList = updateChannelItem(cp.ChannelList, cp.ChannelList[i].Number, newChannel)
			}

			// Channel Delete Button.
			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				cp.ChannelList = deleteChannelItem(cp.ChannelList, cp.ChannelList[i].Number)
				cp.ChannelPanel.Refresh()
			}

			// Channel Add Button.
			o.(*fyne.Container).Objects[3].(*widget.Button).OnTapped = func() {
				cp.ChannelList = addChannelItem(cp.ChannelList, cp.ChannelList[i].Number, cp.ChannelOptions)
				cp.ChannelPanel.Refresh()
			}

			// Settings Button.
			o.(*fyne.Container).Objects[4].(*widget.Button).OnTapped = func() {
				// Highlight this channel
				cp.ChannelPanel.Select(i)
				if cp.ChannelList != nil {
					// Get Existing Settings for channel.
					st.SettingsList = populateChannelSettingList(cp.ChannelList, cp.ChannelList[i].Number)
					// If the settings are empty create a new set of settings.
					if len(st.SettingsList) == 0 {
						// Create new settings.
						st.SettingsList = createChannelSettingList(cp.ChannelList[i].Number)
						st.CurrentChannel = int(cp.ChannelList[i].Number)
						st.SettingsPanel.Hidden = false
						st.SettingsPanel.Refresh()
					} else {
						// Edit existing settings.
						st.CurrentChannel = int(cp.ChannelList[i].Number)
						st.SettingsPanel.Refresh()
					}
				}
			}
		})

	return &cp
}

func createChannelSettingList(channelNumber int16) (settingsList []fixture.Setting) {
	newItem := fixture.Setting{}
	newItem.Name = "New Setting"
	newItem.Number = 1
	newItem.Setting = "0"
	settingsList = append(settingsList, newItem)
	return settingsList
}

func populateChannelSettingList(channelList []fixture.Channel, channelNumber int16) (settingsList []fixture.Setting) {
	for _, channel := range channelList {
		if channelNumber == channel.Number {
			return channel.Settings
		}
	}
	return settingsList
}

func channelItemAllreadyExists(number int16, channelList []fixture.Channel) bool {
	// look through the channel list for the id's
	for _, item := range channelList {
		if item.Number == number {
			return true
		}
	}
	return false
}

func findLargestChannelNumber(items []fixture.Channel) int16 {
	var number int16
	for _, item := range items {
		if item.Number > number {
			number = item.Number
		}
	}
	return number
}

func addChannelItem(channels []fixture.Channel, id int16, options []string) []fixture.Channel {
	newChannels := []fixture.Channel{}
	newItem := fixture.Channel{}
	newItem.Number = id + 1
	if channelItemAllreadyExists(newItem.Number, channels) {
		newItem.Number = findLargestChannelNumber(channels) + 1
	}
	newItem.Name = "New"

	for _, item := range channels {
		if item.Number == id {
			newChannels = append(newChannels, newItem)
		}
		newChannels = append(newChannels, item)
	}
	sort.Slice(newChannels, func(i, j int) bool {
		return newChannels[i].Number < newChannels[j].Number
	})
	return newChannels
}

func deleteChannelItem(channelList []fixture.Channel, id int16) []fixture.Channel {
	newChannels := []fixture.Channel{}
	if id == 1 {
		return channelList
	}
	for _, channel := range channelList {
		if channel.Number != id {
			newChannels = append(newChannels, channel)
		}
	}
	return newChannels
}

// UpdateItem replaces the selected item by id with newItem.
func updateChannelItem(channels []fixture.Channel, id int16, newChannel fixture.Channel) []fixture.Channel {
	newChannels := []fixture.Channel{}
	for _, channel := range channels {
		if channel.Number == id {
			// update the channel information.
			newChannels = append(newChannels, newChannel)
		} else {
			// just add what was there before.
			newChannels = append(newChannels, channel)
		}
	}
	return newChannels
}
