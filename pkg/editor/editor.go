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
	"os"

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
	Actions []fixture.Action
}

func NewChannelEditor(w fyne.Window, id string, group int, number int, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	var currentChannel int
	thisFixture, err := fixture.GetFixureDetails(id, fixtures)
	if err != nil {
		return nil, fmt.Errorf("GetFixureDetails %s", err.Error())
	}

	// Title.
	title := widget.NewLabel(fmt.Sprintf("Edit Config for Sequence %d Fixture %d", thisFixture.Group, thisFixture.Number))
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

	labelSwitch := widget.NewLabel("Switch States")
	labelSwitch.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Populate fixture channels form.
	channelList := thisFixture.Channels
	settingsList := []fixture.Setting{}
	var st *SettingsPanel

	// Populate switch state settings and actions.
	switchesAvailable, actionsAvailable, actionsList, switchesList := PopulateSwitches(thisFixture)

	// Create Actions Panel.
	var ap *ActionPanel
	if actionsAvailable {
		ap = NewActionsPanel(actionsAvailable, actionsList)
		ap.ActionsPanel.Hide()
	}

	// Create Switches Panel.
	var switchesPanel *widget.List
	if switchesAvailable {
		sw := NewSwitchPanel(switchesAvailable, switchesList, ap)
		switchesPanel = sw.SwitchPanel
	}

	// Create Settings Panel
	var settingsPanel *widget.List
	st = NewSettingsPanel(thisFixture, &currentChannel, settingsList, channelList, ap)
	settingsPanel = st.SettingsPanel

	// Create Channel Panel.
	cp := NewChannelPanel(thisFixture, &currentChannel, channelList, ap, st)

	// Setup forms.
	scrollableChannelList := container.NewScroll(cp.ChannelPanel)
	scrollableDevicesList := container.NewScroll(switchesPanel)

	// Size accordingly
	scrollableChannelList.SetMinSize(fyne.Size{Height: 400, Width: 300})
	scrollableDevicesList.SetMinSize(fyne.Size{Height: 0, Width: 0})
	if actionsAvailable {
		scrollableChannelList.SetMinSize(fyne.Size{Height: 250, Width: 400})
		scrollableDevicesList.SetMinSize(fyne.Size{Height: 250, Width: 400})
	}

	// Save button.
	buttonSave := widget.NewButton("Save", func() {

		// Insert updated fixture into fixtures.
		newFixtures := fixture.Fixtures{}
		for _, fixture := range fixtures.Fixtures {
			if fixture.Group == group && fixture.Number == number+1 {
				// Insert new channels into fixture.
				thisFixture.Channels = channelList
				newFixtures.Fixtures = append(newFixtures.Fixtures, thisFixture)
			} else {
				newFixtures.Fixtures = append(newFixtures.Fixtures, fixture)
			}
		}

		// Save the new fixtures file.
		err := fixture.SaveFixtures("newfixtures.yaml", &newFixtures)
		if err != nil {
			fmt.Printf("error saving fixtures %s\n", err.Error())
			os.Exit(1)
		}

		modal.Hide()
	})

	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		modal.Hide()
	})

	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)

	content := fyne.Container{}
	t := container.NewBorder(title, nil, nil, nil, formTop)
	top := container.NewBorder(t, nil, nil, nil, labelChannels)
	scrollableSettingsList := container.NewScroll(settingsPanel)
	scrollableSettingsList.SetMinSize(fyne.Size{Height: 250, Width: 400})
	box := container.NewAdaptiveGrid(2, scrollableChannelList, scrollableSettingsList)
	middle := container.NewBorder(top, nil, nil, nil, box)
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
