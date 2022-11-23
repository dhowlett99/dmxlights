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

func NewChannelPanel(channelList []fixture.Channel, channelOptions []string) *ChannelPanel {

	cp := ChannelPanel{}
	cp.ChannelList = channelList
	cp.ChannelOptions = channelOptions

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
					item.Number = int16(lastChannel)
					item.Name = value
					cp.ChannelList = UpdateItem(cp.ChannelList, item.Number, item)
				}),

				widget.NewButton("-", func() {
					//log.Println("Delete Button pressed for ")
				}),
				widget.NewButton("+", func() {
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

		})

	return &cp
}

func PopulateChannels(thisFixture fixture.Fixture, channelOptions []string) []fixture.Channel {
	channelList := []fixture.Channel{}
	// Populate fixture channels form.
	for _, channel := range thisFixture.Channels {
		newSelect := fixture.Channel{}
		newSelect.Number = channel.Number
		newSelect.Name = channel.Name
		newSelect.Offset = channel.Offset
		newSelect.MaxDegrees = channel.MaxDegrees
		newSelect.Settings = channel.Settings
		newSelect.Value = channel.Value
		newSelect.Comment = channel.Name
		channelList = append(channelList, newSelect)
	}
	return channelList
}
