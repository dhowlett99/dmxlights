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

type ChannelPanel struct {
	ChannelPanel   *widget.List
	ChannelList    []fixture.Channel
	ChannelOptions []string
}

func NewChannelPanel(currentChannel *int, channelList []fixture.Channel, ap *ActionPanel, st *SettingsPanel) *ChannelPanel {

	cp := ChannelPanel{}
	cp.ChannelList = channelList
	cp.ChannelOptions = []string{"Rotate", "Red1", "Red2", "Red3", "Red4", "Red5", "Red6", "Red7", "Red8", "Green1", "Green2", "Green3", "Green4", "Green5", "Green6", "Green7", "Green8", "Blue1", "Blue2", "Blue3", "Blue4", "Blue5", "Blue6", "Blue7", "Blue8", "White1", "White2", "White3", "White4", "White5", "White6", "White7", "White8", "Master", "Dimmer", "Static", "Pan", "FinePan", "Tilt", "FineTilt", "Shutter", "Strobe", "Color", "Gobo", "Program", "ProgramSpeed", "Programs", "ColorMacros"}

	// Channel or Switch State Selection Panel.
	cp.ChannelPanel = widget.NewList(
		// Function to find length.
		func() int {
			return len(cp.ChannelList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {
			return container.NewHBox(
				widget.NewLabel("template"),

				widget.NewSelect(cp.ChannelOptions, func(value string) {
					lastChannel, _ := strconv.Atoi(o.(*fyne.Container).Objects[0].(*widget.Label).Text)
					//fmt.Printf("We just pressed channel %d and set it to %s\n", lastChannel, value)
					item := fixture.Channel{}
					item.Name = value
					item.Number = int16(lastChannel)
					itemNumber := item.Number - 1
					item.Comment = cp.ChannelList[itemNumber].Comment
					item.MaxDegrees = cp.ChannelList[itemNumber].MaxDegrees
					item.Offset = cp.ChannelList[itemNumber].Offset
					item.Settings = cp.ChannelList[itemNumber].Settings
					item.Value = cp.ChannelList[itemNumber].Value
					cp.ChannelList = UpdateItem(cp.ChannelList, item.Number, item)
				}),

				widget.NewButton("-", func() {
					//log.Println("Delete Button pressed for ")
				}),
				widget.NewButton("+", func() {
					//log.Println("Add Button pressed for ")
				}),

				widget.NewButton("settings", func() {
					//log.Println("Add Button pressed for ")
				}),
			)
		},
		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", cp.ChannelList[i].Number))

			// find the selected option in the options list.
			for _, option := range cp.ChannelOptions {
				if option == cp.ChannelList[i].Name {
					o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
				}
			}

			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				cp.ChannelList = DeleteChannelItem(cp.ChannelList, cp.ChannelList[i].Number)
				cp.ChannelPanel.Refresh()
			}

			o.(*fyne.Container).Objects[3].(*widget.Button).OnTapped = func() {
				cp.ChannelList = AddChannelItem(cp.ChannelList, cp.ChannelList[i].Number, cp.ChannelOptions)
				cp.ChannelPanel.Refresh()
			}

			o.(*fyne.Container).Objects[4].(*widget.Button).OnTapped = func() {
				if cp.ChannelList != nil {
					if st != nil {
						st.SettingsList = PopulateSettingList(cp.ChannelList, cp.ChannelList[i].Number)
						st.SettingsPanel.Refresh()
					}
				}
			}
		})

	return &cp
}

func PopulateSettingList(channelList []fixture.Channel, channelNumber int16) (settingsList []fixture.Setting) {
	for _, channel := range channelList {
		if channelNumber == channel.Number {
			return channel.Settings
		}
	}
	return settingsList
}

func ChannelItemAllreadyExists(number int16, channelList []fixture.Channel) bool {
	// look through the channel list for the id's
	for _, item := range channelList {
		if item.Number == number {
			return true
		}
	}
	return false
}

func FindLargestChannelNumber(items []fixture.Channel) int16 {
	var number int16
	for _, item := range items {
		if item.Number > number {
			number = item.Number
		}
	}
	return number
}

func AddChannelItem(items []fixture.Channel, id int16, options []string) []fixture.Channel {
	newItems := []fixture.Channel{}
	newItem := fixture.Channel{}
	newItem.Number = id + 1
	if ChannelItemAllreadyExists(newItem.Number, items) {
		newItem.Number = FindLargestChannelNumber(items) + 1
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

func UpdateItem(items []fixture.Channel, id int16, newItem fixture.Channel) []fixture.Channel {
	newItems := []fixture.Channel{}
	for _, item := range items {
		if item.Number == id {
			// update the channel information.
			newItems = append(newItems, newItem)
		} else {
			// just add what was there before.
			newItems = append(newItems, item)
		}
	}
	return newItems
}
