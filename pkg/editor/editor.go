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
	Actions []itemSelect
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

	var actionsAvailable bool
	var switchesAvailable bool

	labelChannels := widget.NewLabel("Channels")
	// Channel or Switch label.
	labelChannels.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	labelSwitch := widget.NewLabel("switch")
	channelList := []itemSelect{}
	switchesList := []itemSelect{}
	actionsList := []itemSelect{}

	// Describe the options.
	channelOptions := []string{"Red1", "Red2", "Red3", "Red4", "Red5", "Red6", "Red7", "Red8", "Green1", "Green2", "Green3", "Green4", "Green5", "Green6", "Green7", "Green8", "Blue1", "Blue2", "Blue3", "Blue4", "Blue5", "Blue6", "Blue7", "Blue8", "Master", "Dimmer", "Static", "Pan", "FinePan", "Tilt", "FineTilt", "Shutter", "Strobe", "Color", "Gobo", "Programs", "ColorMacros"}
	switchOptions := []string{"Off", "On", "Red", "Green", "Blue", "Softchase", "Hardchase", "Soundchase", "Rotate"}
	//actionOptions := []string{"Colors", "Fade", "Mode", "Music", "Program", "Rotate", "Size", "Speed"}
	actionOptions := []string{"Off", "Chase", "Static"}
	// Populate fixture channels form.
	for _, channel := range fixture.Channels {
		newSelect := itemSelect{}
		newSelect.Number = channel.Number
		newSelect.Label = channel.Name
		newSelect.Options = channelOptions
		channelList = append(channelList, newSelect)
	}

	channelPanel := NewChannelPanel(channelList, channelOptions)

	// Populate switch state settings and actions.
	if fixture.Type == "switch" {
		labelSwitch.Text = "Switch States"
		for _, state := range fixture.States {
			switchesAvailable = true
			newSelect := itemSelect{}
			newSelect.Number = state.Number
			newSelect.Label = state.Name
			newSelect.Options = switchOptions
			if state.Actions != nil {
				actionsAvailable = true
				actionsList = []itemSelect{}
				for actionNumber, action := range state.Actions {
					fmt.Printf("----->Add action %+v\n", action)
					newAction := itemSelect{}
					newAction.Number = int16(actionNumber)
					newAction.Label = action.Mode
					newAction.Options = actionOptions
					actionsList = append(actionsList, newAction)
				}
			}
			newSelect.Actions = actionsList
			fmt.Printf("----->Actions %+v\n", newSelect.Actions)
			switchesList = append(switchesList, newSelect)
		}
	}

	ap := NewActionsPanel(actionsAvailable, actionsList, actionOptions)
	switchesPanel := NewSwitchPanel(switchesAvailable, switchesList, switchOptions, ap)

	// Setup forms.
	scrollableChannelList := container.NewScroll(channelPanel)
	scrollableDevicesList := container.NewScroll(switchesPanel)
	scrollableActionsList := container.NewScroll(ap.ActionsPanel)
	scrollableChannelList.SetMinSize(fyne.Size{Height: 400, Width: 300})
	scrollableDevicesList.SetMinSize(fyne.Size{Height: 0, Width: 0})
	scrollableActionsList.SetMinSize(fyne.Size{Height: 0, Width: 0})

	// Size accordingly
	if actionsAvailable {
		scrollableChannelList.SetMinSize(fyne.Size{Height: 250, Width: 300})
		scrollableDevicesList.SetMinSize(fyne.Size{Height: 250, Width: 300})
		scrollableActionsList.SetMinSize(fyne.Size{Height: 250, Width: 300})
	}

	// Save button.
	buttonSave := widget.NewButton("Save", func() {
		for _, channel := range channelList {
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
		main := container.NewBorder(middle, nil, nil, nil, labelSwitch)
		states := container.NewHBox(scrollableDevicesList, scrollableActionsList)
		bottom := container.NewBorder(main, nil, nil, nil, states)
		content = *container.NewBorder(bottom, nil, nil, nil, saveCancel)
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
