// Copyright (C) 2022 dhowlett99.
// This is the dmxlights channel editor it is attached to a fixture and
// describes the fixtures channel properties which is then saved in the fixtures.yaml
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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

const debug bool = false

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

func NewChannelEditor(w fyne.Window, id int, channels []fixture.Channel, fp *FixturesPanel, groupConfig *fixture.Groups, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	if debug {
		fmt.Printf("NewChannelEditor\n")
	}

	// Store the current channels incase we cancel changes.
	savedChannels := append(channels[:0:0], channels...)

	// Create the save button early so we can pass the pointer to error checks.
	buttonSave := widget.NewButton("OK", func() {})

	thisFixture, err := fixture.GetFixtureDetailsById(id, fixtures)
	if err != nil {
		return nil, fmt.Errorf("GetFixtureDetailsById %s", err.Error())
	}

	// Resolve the group number to a name.
	groupName := getGroupName(groupConfig, strconv.Itoa(thisFixture.Group))

	// Title.
	title := widget.NewLabel(fmt.Sprintf("Edit Channel Config for Sequence %s Fixture %d", groupName, thisFixture.Number))
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Name, description and DMX address
	nameInput := widget.NewEntry()
	nameInput.SetPlaceHolder(thisFixture.Name)
	descInput := widget.NewEntry()
	descInput.SetPlaceHolder(thisFixture.Description)
	addrInput := widget.NewEntry()
	addrInput.SetPlaceHolder(fmt.Sprintf("%d", thisFixture.Address))

	// Top Form.
	var formTopItems []*widget.FormItem
	name1 := widget.NewEntry()
	name1.SetText(thisFixture.Name)
	formTopItem := widget.NewFormItem("Name", name1)
	formTopItems = append(formTopItems, formTopItem)
	name2 := widget.NewEntry()
	name2.SetText(thisFixture.Description)
	formTopItem2 := widget.NewFormItem("Description", name2)
	formTopItems = append(formTopItems, formTopItem2)
	name3 := widget.NewEntry()
	name3.SetText(fmt.Sprintf("%d", thisFixture.Address))
	formTopItem3 := widget.NewFormItem("DMX Address", name3)
	formTopItems = append(formTopItems, formTopItem3)
	formTop := &widget.Form{
		Items: formTopItems,
	}

	labelChannels := widget.NewLabel("Channels")
	labelChannels.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	settingsList := []fixture.Setting{}
	var st *SettingsPanel

	// Create Channel Settings Panel
	var settingsPanel *widget.Table
	channelPanel := true
	// You can have a setting for every channel on a fixture.
	// So if your creating a fixture channel setting
	// technically you could occupy the whole of the DMX universe.
	st = NewSettingsPanel(w, channelPanel, settingsList, MAX_NUMBER_SETTINGS, buttonSave)
	settingsPanel = st.SettingsPanel

	// Create Channel Panel.
	cp := NewChannelPanel(thisFixture, channels, st)

	// Setup forms.
	scrollableChannelList := container.NewVScroll(cp.ChannelPanel)
	scrollableChannelList.SetMinSize(fyne.Size{Height: 400, Width: 250})
	scrollableSettingsList := container.NewVScroll(settingsPanel)
	scrollableSettingsList.SetMinSize(fyne.Size{Height: 400, Width: 250})

	// Setup OK buttons action.
	buttonSave.OnTapped = func() {

		// Insert updated fixture into fixtures.
		newFixtures := fixture.Fixtures{}
		for fixtureNumber, fixture := range fixtures.Fixtures {
			if fixture.ID == id {
				// Insert new channels into fixture above us, in the fixture selection panel.
				fp.UpdateChannels = true
				fp.UpdateThisFixture = fixtureNumber
				fp.UpdatedChannelsList = cp.ChannelList
				// Update our copy of the channel list.
				thisFixture.Channels = cp.ChannelList
				newFixtures.Fixtures = append(newFixtures.Fixtures, thisFixture)
			} else {
				newFixtures.Fixtures = append(newFixtures.Fixtures, fixture)
			}
		}

		modal.Hide()
	}

	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		// Restore any changed channels settings.
		for channelNumber := range savedChannels {
			cp.ChannelList[channelNumber].Settings = append(cp.ChannelList[channelNumber].Settings[:0:0], savedChannels[channelNumber].Settings...)
		}
		modal.Hide()
	})

	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)

	content := fyne.Container{}
	t := container.NewBorder(title, nil, nil, nil, formTop)
	top := container.NewBorder(t, nil, nil, nil, labelChannels)
	box := container.NewAdaptiveGrid(2, scrollableChannelList, scrollableSettingsList)
	middle := container.NewBorder(top, nil, nil, nil, box)
	content = *container.NewBorder(middle, nil, nil, nil, saveCancel)

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		&content,
		w.Canvas(),
	)
	return modal, nil
}

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
							st.SettingsList = createNewChannelSettingList()
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

func createNewChannelSettingList() (settingsList []fixture.Setting) {

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
