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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type itemSelect struct {
	Number  int16
	Label   string
	Options []string
	Actions []actionItems
}

func NewEditor(w fyne.Window, group int, number int, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	fixture, err := fixture.GetFixureDetails(group, number, fixtures)
	if err != nil {
		return nil, fmt.Errorf("GetFixureDetails %s", err.Error())
	}

	// Title.
	title := widget.NewLabel(fmt.Sprintf("Edit Config for Sequence %d Fixture %d", fixture.Group, fixture.Number))
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Name, description and DMX address
	nameInput := widget.NewEntry()
	nameInput.SetPlaceHolder(fixture.Name)
	descInput := widget.NewEntry()
	descInput.SetPlaceHolder(fixture.Description)
	addrInput := widget.NewEntry()
	addrInput.SetPlaceHolder(fmt.Sprintf("%d", fixture.Address))

	// Top Form.
	var formTopItems []*widget.FormItem
	name1 := widget.NewEntry()
	name1.SetText(fixture.Name)
	formTopItem := widget.NewFormItem("Name", name1)
	formTopItems = append(formTopItems, formTopItem)
	name2 := widget.NewEntry()
	name2.SetText(fixture.Description)
	formTopItem2 := widget.NewFormItem("Description", name2)
	formTopItems = append(formTopItems, formTopItem2)
	name3 := widget.NewEntry()
	name3.SetText(fmt.Sprintf("%d", fixture.Address))
	formTopItem3 := widget.NewFormItem("DMX Address", name3)
	formTopItems = append(formTopItems, formTopItem3)
	formTop := &widget.Form{
		Items: formTopItems,
	}

	labelChannels := widget.NewLabel("Channels")
	// Channel or Switch label.
	labelChannels.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	labelSwitch := widget.NewLabel("Switch States")

	// Describe the options.
	channelOptions := []string{"Rotate", "Red1", "Red2", "Red3", "Red4", "Red5", "Red6", "Red7", "Red8", "Green1", "Green2", "Green3", "Green4", "Green5", "Green6", "Green7", "Green8", "Blue1", "Blue2", "Blue3", "Blue4", "Blue5", "Blue6", "Blue7", "Blue8", "White1", "White2", "White3", "White4", "White5", "White6", "White7", "White8", "Master", "Dimmer", "Static", "Pan", "FinePan", "Tilt", "FineTilt", "Shutter", "Strobe", "Color", "Gobo", "Program", "ProgramSpeed", "Programs", "ColorMacros"}
	switchOptions := []string{"Off", "On", "Red", "Green", "Blue", "SoftChase", "SharpChase", "SoundChase", "Rotate"}
	//actionOptions := []string{"Colors", "Fade", "Mode", "Music", "Program", "Rotate", "Size", "Speed"}
	actionNameOptions := switchOptions
	actionModeOptions := []string{"Off", "Chase", "Static"}
	actionColorOptions := []string{"Red", "Green", "Blue"}
	actionFadeOptions := []string{"Off", "Soft", "Sharp"}
	actionSpeedOptions := []string{"Slow", "Medium", "Fast", "VeryFast", "Music"}

	// Populate fixture channels form.
	channelList := PopulateChannels(fixture, channelOptions)

	// Create Channel Panel.
	cp := NewChannelPanel(channelList, channelOptions)

	// Populate switch state settings and actions.
	switchesAvailable, actionsAvailable, actionsList, switchesList := PopulateSwitches(switchOptions, fixture)

	// Create Actions Panel.
	var ap *ActionPanel
	if actionsAvailable {
		ap = NewActionsPanel(actionsAvailable, actionsList, actionNameOptions, actionColorOptions, actionModeOptions, actionFadeOptions, actionSpeedOptions)
		ap.ActionsPanel.Hide()
	}

	// Create Switches Panel.
	var switchesPanel *widget.List
	if switchesAvailable {
		sw := NewSwitchPanel(switchesAvailable, switchesList, switchOptions, ap)
		switchesPanel = sw.SwitchPanel
	}

	// Setup forms.
	scrollableChannelList := container.NewScroll(cp.ChannelPanel)
	scrollableDevicesList := container.NewScroll(switchesPanel)

	// Size accordingly
	scrollableChannelList.SetMinSize(fyne.Size{Height: 400, Width: 300})
	scrollableDevicesList.SetMinSize(fyne.Size{Height: 0, Width: 0})
	if actionsAvailable {
		scrollableChannelList.SetMinSize(fyne.Size{Height: 250, Width: 300})
		scrollableDevicesList.SetMinSize(fyne.Size{Height: 250, Width: 300})
	}

	// Save button.
	buttonSave := widget.NewButton("Save", func() {
		for _, channel := range cp.ChannelList {
			fmt.Printf("---> channel \n")
			fmt.Printf("\t number %d\n", channel.Number)
			fmt.Printf("\t name   %s\n", channel.Label)
		}
		modal.Hide()
	})

	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		modal.Hide()
	})

	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)

	content := fyne.Container{}
	top := container.NewBorder(formTop, nil, nil, nil, labelChannels)
	middle := container.NewBorder(top, nil, nil, nil, scrollableChannelList)
	if switchesAvailable {
		var states *fyne.Container
		if actionsAvailable {
			scrollableActionsList := container.NewScroll(ap.ActionsPanel)
			scrollableActionsList.SetMinSize(fyne.Size{Height: 0, Width: 0})
			scrollableActionsList.SetMinSize(fyne.Size{Height: 250, Width: 300})
			states = container.NewHBox(scrollableDevicesList, scrollableActionsList)
		}
		main := container.NewBorder(middle, nil, nil, nil, labelSwitch)
		if actionsAvailable {
			bottom := container.NewBorder(main, nil, nil, nil, states)
			content = *container.NewBorder(bottom, nil, nil, nil, saveCancel)
		} else {
			content = *container.NewBorder(main, nil, nil, nil, saveCancel)
		}
	} else {
		content = *container.NewBorder(middle, nil, nil, nil, saveCancel)
	}

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		&content,
		w.Canvas(),
	)
	return modal, nil
}

func DeleteItem(channelList []itemSelect, id int16) []itemSelect {
	newItems := []itemSelect{}
	for _, item := range channelList {

		if item.Number != id {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func ItemAllreadyExists(number int16, channelList []itemSelect) bool {
	// look through the channel list for the id's
	for _, item := range channelList {
		if item.Number == number {
			return true
		}
	}
	return false
}

func FindLargest(items []itemSelect) int16 {
	var number int16
	for _, item := range items {
		if item.Number > number {
			number = item.Number
		}
	}
	return number
}

func AddItem(items []itemSelect, id int16, options []string) []itemSelect {
	newItems := []itemSelect{}
	newItem := itemSelect{}
	newItem.Number = id + 1
	if ItemAllreadyExists(newItem.Number, items) {
		newItem.Number = FindLargest(items) + 1
	}
	newItem.Label = "New"
	newItem.Options = options

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

func UpdateItem(items []itemSelect, id int16, newItem itemSelect) []itemSelect {
	newItems := []itemSelect{}
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
