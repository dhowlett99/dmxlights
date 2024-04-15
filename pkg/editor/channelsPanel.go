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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type ChannelPanel struct {
	ChannelPanel   *widget.Table
	ChannelList    []fixture.Channel
	ChannelOptions []string
}

const (
	CHANNEL_NUMBER int = iota
	CHANNEL_NAME
	CHANNEL_DELETE
	CHANNEL_ADD
	CHANNEL_SETTINGS
)

func NewChannelPanel(thisFixture fixture.Fixture, channels []fixture.Channel, st *SettingsPanel) *ChannelPanel {

	if debug {
		fmt.Printf("NewChannelPanel\n")
	}

	var data = [][]string{}

	cp := ChannelPanel{}
	cp.ChannelOptions = []string{"Rotate", "Macro",
		"Red1", "Red2", "Red3", "Red4", "Red5", "Red6", "Red7", "Red8",
		"Green1", "Green2", "Green3", "Green4", "Green5", "Green6", "Green7", "Green8",
		"Blue1", "Blue2", "Blue3", "Blue4", "Blue5", "Blue6", "Blue7", "Blue8",
		"White1", "White2", "White3", "White4", "White5", "White6", "White7", "White8",
		"Master", "Master Reverse", "Dimmer", "Static",
		"Pan", "FinePan", "Tilt", "FineTilt", "Shutter", "Strobe",
		"Color", "Gobo", "Program", "ProgramSpeed", "Programs",
		"ColorMacros", "SoundActive", "DimmerCurve", "Speed"}
	cp.ChannelList = channels

	// Channel or Switch State Selection Panel.
	cp.ChannelPanel = widget.NewTable(

		// Function to find length.
		func() (int, int) {
			if st.UpdateSettings {
				if st.UpdateThisChannel < 0 {
					st.UpdateThisChannel = 0
				}
				cp.ChannelList[st.UpdateThisChannel].Settings = st.SettingsList
				cp.ChannelList[st.UpdateThisChannel].MaxDegrees = st.SettingMaxDegrees
				st.UpdateSettings = false
			}
			height := len(data)
			width := 5
			return height, width
		},

		// Function to create table.
		func() (o fyne.CanvasObject) {

			// Load the settings into the array used by the table.
			data = makeChannelsArray(cp.ChannelList)

			return container.NewStack(

				// Channel Number.
				widget.NewLabel("template"),

				// Channel Value as a selectable dialog box.
				widget.NewSelect(cp.ChannelOptions, func(value string) {}),

				// Chanell delete button.
				widget.NewButton("-", func() {}),

				// Channel add button
				widget.NewButton("+", func() {}),

				// Settings for this channel button.
				widget.NewButton("settings", func() {}),
			)
		},

		// Function to update items in this table.
		func(i widget.TableCellID, o fyne.CanvasObject) {

			// Hide all field types.
			hideAllChannelsFields(o)

			// Show the Channel Number.
			if i.Col == CHANNEL_NUMBER {
				showChannelsField(CHANNEL_NUMBER, o)
				o.(*fyne.Container).Objects[CHANNEL_NUMBER].(*widget.Label).SetText(fmt.Sprintf("%d", cp.ChannelList[i.Row].Number))
			}

			// Show the currently selected Channel option.
			if i.Col == CHANNEL_NAME {
				showChannelsField(CHANNEL_NAME, o)
				o.(*fyne.Container).Objects[CHANNEL_NAME].(*widget.Select).OnChanged = nil
				o.(*fyne.Container).Objects[CHANNEL_NAME].(*widget.Select).Selected = cp.ChannelList[i.Row].Name
				o.(*fyne.Container).Objects[CHANNEL_NAME].(*widget.Select).Refresh()
				// Edit the channel Value.
				o.(*fyne.Container).Objects[CHANNEL_NAME].(*widget.Select).OnChanged = func(value string) {
					newChannel := fixture.Channel{}
					newChannel.Name = value
					newChannel.Number = cp.ChannelList[i.Row].Number
					newChannel.Value = cp.ChannelList[i.Row].Value
					newChannel.Settings = cp.ChannelList[i.Row].Settings
					newChannel.Comment = cp.ChannelList[i.Row].Comment
					newChannel.MaxDegrees = cp.ChannelList[i.Row].MaxDegrees
					newChannel.Offset = cp.ChannelList[i.Row].Offset
					cp.ChannelList = updateChannelItem(cp.ChannelList, cp.ChannelList[i.Row].Number, newChannel)
					data = makeChannelsArray(cp.ChannelList)
				}
			}

			// Channel Delete Button.
			if i.Col == CHANNEL_DELETE {
				showChannelsField(CHANNEL_DELETE, o)
				o.(*fyne.Container).Objects[CHANNEL_DELETE].(*widget.Button).OnTapped = func() {
					cp.ChannelList = deleteChannelItem(cp.ChannelList, cp.ChannelList[i.Row].Number)
					data = makeChannelsArray(cp.ChannelList)
					cp.ChannelPanel.Refresh()
				}
			}

			// Channel Add Button.
			if i.Col == CHANNEL_ADD {
				showChannelsField(CHANNEL_ADD, o)
				o.(*fyne.Container).Objects[CHANNEL_ADD].(*widget.Button).OnTapped = func() {
					cp.ChannelList = addChannelItem(cp.ChannelList, cp.ChannelList[i.Row].Number)
					data = makeChannelsArray(cp.ChannelList)
					cp.ChannelPanel.Refresh()
				}
			}

			// Settings Button.
			if i.Col == CHANNEL_SETTINGS {
				showChannelsField(CHANNEL_SETTINGS, o)
				o.(*fyne.Container).Objects[CHANNEL_SETTINGS].(*widget.Button).OnTapped = func() {
					// Highlight this channel
					cp.ChannelPanel.Select(i)
					if cp.ChannelList != nil {
						// Get Existing Settings for channel.
						st.SettingsList = populateChannelSettingList(cp.ChannelList, cp.ChannelList[i.Row].Number)
						// If the settings are empty create a new set of settings.
						if len(st.SettingsList) == 0 {
							// Create new settings.
							st.SettingsList = createChannelSettingList()
							st.CurrentChannel = int(cp.ChannelList[i.Row].Number)
							st.SettingsPanel.Hidden = false
							st.SettingsPanel.Refresh()
						} else {
							// Edit existing settings.
							st.CurrentChannel = int(cp.ChannelList[i.Row].Number)
							st.SettingsPanel.Refresh()
						}
					}
				}
			}
		},
	)

	// Setup the columns of this table.
	cp.ChannelPanel.SetColumnWidth(0, 40)  // Number
	cp.ChannelPanel.SetColumnWidth(1, 160) // Name
	cp.ChannelPanel.SetColumnWidth(2, 20)  // Delete
	cp.ChannelPanel.SetColumnWidth(3, 20)  // Add
	cp.ChannelPanel.SetColumnWidth(4, 100) // Settings

	return &cp
}

func createChannelSettingList() (settingsList []fixture.Setting) {

	if debug {
		fmt.Printf("createChannelSettingList\n")
	}

	newItem := fixture.Setting{}
	newItem.Name = "New"
	newItem.Label = "New"
	newItem.Channel = "New"
	newItem.Number = 1
	newItem.Value = "0"
	settingsList = append(settingsList, newItem)
	return settingsList
}

func populateChannelSettingList(channelList []fixture.Channel, channelNumber int16) (settingsList []fixture.Setting) {

	if debug {
		fmt.Printf("populateChannelSettingList\n")
	}

	for _, channel := range channelList {
		if channelNumber == channel.Number {
			return channel.Settings
		}
	}
	return settingsList
}

func addChannelItem(channels []fixture.Channel, number int16) (outItems []fixture.Channel) {

	if debug {
		fmt.Printf("addChannelItem\n")
	}

	newChannels := []fixture.Channel{}
	newItem := fixture.Channel{}
	newItem.Number = number + 1
	newItem.Name = "(Select one)"
	emptyValue := int16(0)
	newItem.Value = &emptyValue

	var added bool // Only add once.

	for no, item := range channels {
		// Add at the start of an empty list.
		if len(channels) == 0 && !added {
			newChannels = append(newChannels, newItem)
			added = true
		}
		// Insert at this position.
		if item.Number == number+1 && !added {
			newChannels = append(newChannels, newItem)
			added = true
		}
		newChannels = append(newChannels, item)
		// Append an item at the very end.
		if no == len(channels)-1 && !added {
			newChannels = append(newChannels, newItem)
			added = true
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newChannels {
		indexedItem.Number = int16(number + 1)
		outItems = append(outItems, indexedItem)
	}

	return outItems
}

func deleteChannelItem(channelList []fixture.Channel, id int16) (outItems []fixture.Channel) {

	if debug {
		fmt.Printf("deleteChannelItem\n")
	}

	newChannels := []fixture.Channel{}
	for _, channel := range channelList {
		if channel.Number != id {
			newChannels = append(newChannels, channel)
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newChannels {
		indexedItem.Number = int16(number + 1)
		outItems = append(outItems, indexedItem)
	}

	// If we have no items create a default one.
	if len(outItems) == 0 {
		// Create a default Channel
		newItem := fixture.Channel{}
		newItem.Number = 1
		newItem.Name = "(Select one)"
		outItems = append(outItems, newItem)
	}

	return outItems
}

// UpdateItem replaces the selected item by id with newItem.
func updateChannelItem(channels []fixture.Channel, id int16, newChannel fixture.Channel) []fixture.Channel {

	if debug {
		fmt.Printf("updateChannelItem\n")
	}

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

// makeChannelsArray - Convert the list of channels to an array of strings containing and array of strings with
// the values from each channel.
// This is done once when the channels panel is loaded.
func makeChannelsArray(channels []fixture.Channel) [][]string {

	if debug {
		fmt.Printf("makeChannelsArray\n")
	}

	var data = [][]string{}

	for _, channel := range channels {
		newChannel := []string{}
		newChannel = append(newChannel, fmt.Sprintf("%d", channel.Number))
		newChannel = append(newChannel, channel.Name)
		newChannel = append(newChannel, "-")
		newChannel = append(newChannel, "+")
		newChannel = append(newChannel, "Settings")

		data = append(data, newChannel)
	}

	return data
}

func showChannelsField(field int, o fyne.CanvasObject) {
	if debug {
		fmt.Printf("showChannelsField\n")
	}
	// Now show the selected field.
	switch {
	case field == CHANNEL_NUMBER:
		o.(*fyne.Container).Objects[CHANNEL_NUMBER].(*widget.Label).Hidden = false
	case field == CHANNEL_NAME:
		o.(*fyne.Container).Objects[CHANNEL_NAME].(*widget.Select).Hidden = false
	case field == CHANNEL_DELETE:
		o.(*fyne.Container).Objects[CHANNEL_DELETE].(*widget.Button).Hidden = false
	case field == CHANNEL_ADD:
		o.(*fyne.Container).Objects[CHANNEL_ADD].(*widget.Button).Hidden = false
	case field == CHANNEL_SETTINGS:
		o.(*fyne.Container).Objects[CHANNEL_SETTINGS].(*widget.Button).Hidden = false
	}
}

func hideAllChannelsFields(o fyne.CanvasObject) {
	if debug {
		fmt.Printf("hideAllChannelsFields\n")
	}
	o.(*fyne.Container).Objects[CHANNEL_NUMBER].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[CHANNEL_NAME].(*widget.Select).Hidden = true
	o.(*fyne.Container).Objects[CHANNEL_DELETE].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[CHANNEL_ADD].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[CHANNEL_SETTINGS].(*widget.Button).Hidden = true
}
