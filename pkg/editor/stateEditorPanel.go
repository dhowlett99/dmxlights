// Copyright (C) 2022, 2023 dhowlett99.
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
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

// Show a list of States.
func NewStateEditor(w fyne.Window, id int, fp *FixturesPanel, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	if debug {
		fmt.Printf("NewStateEditor\n")
	}

	// Create the save button early so we can pass the pointer to error checks.
	buttonSave := widget.NewButton("OK", func() {})

	// Get the details of this fixture.
	thisFixture, err := fixture.GetFixureDetailsById(id, fixtures)
	if err != nil {
		return nil, fmt.Errorf("GetFixureDetailsById %s", err.Error())
	}

	// If this is a pretend virtual fixture i.e a switch.
	// Find the original fixure's details so we can make decisions on
	// what options to put in the menus based on the fixtures capabilities.
	basedOnFixture, err := fixture.GetFixureDetailsByLabel(thisFixture.UseFixture, fixtures)
	if err != nil {
		return nil, fmt.Errorf("GetFixureDetailsByLabel %s", err.Error())
	}
	fixtureInfo := fixture.FindFixtureInfo(basedOnFixture)
	if debug {
		fmt.Printf("This fixture has Rotate Feature %+v\n", fixtureInfo)
	}

	// Generate a list of functions that switches can use.
	fixturesAvailable := GetFixtureLabelsForSwitches(fixtures)

	// Title.
	title := widget.NewLabel(fmt.Sprintf("ID:%d Edit Config for Sequence %d Fixture %d", thisFixture.ID, thisFixture.Group, thisFixture.Number))
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Name.
	nameInput := widget.NewLabel(thisFixture.Name)
	nameLabel := widget.NewLabel("Name")
	name := container.NewAdaptiveGrid(3, nameLabel, nameInput, layout.NewSpacer())

	// Description.
	descInput := widget.NewEntry()
	descInput.SetPlaceHolder(thisFixture.Description)
	descLabel := widget.NewLabel("Description")
	desc := container.NewAdaptiveGrid(3, descLabel, descInput, layout.NewSpacer())
	// Update the Description.
	descInput.OnChanged = func(value string) {
		fp.UpdateThisFixture = thisFixture.ID - 1
		fp.UpdateDescription = true
		fp.Description = value
	}

	// Create Actions Panel. fixtureInfo controls what options we see.
	ap := NewActionsPanel(w, []fixture.Action{}, fixtureInfo)
	ap.ActionsPanel.Hide()

	// Create Settings Panel.
	st := NewSettingsPanel(w, []fixture.Setting{}, false, buttonSave)
	st.ChannelOptions = populateChannelNames(thisFixture.Channels)
	st.SettingsPanel.Hide()

	// Use Fixture.
	useInput := widget.NewSelect(fixturesAvailable, func(value string) {})
	useLabel := widget.NewLabel("Use Fixture")
	use := container.NewAdaptiveGrid(3, useLabel, useInput, layout.NewSpacer())

	// Fixture Address.
	addressInput := widget.NewEntry()
	addressLabel := widget.NewLabel("DMX Address")
	address := container.NewAdaptiveGrid(3, addressLabel, addressInput, layout.NewSpacer())

	// Update Use Fixture.
	useInput.OnChanged = func(value string) {

		// Update the address from the use fixture field in the fixture panel.
		fp.UseFixture = useInput.Selected
		fp.UpdateThisFixture = thisFixture.ID - 1
		fp.UpdateUseFixture = true

		useFixture, err := fixture.GetFixureDetailsByLabel(useInput.Selected, fixtures)
		if err != nil {
			addressInput.SetText("Not Found")
		} else {
			addressInput.SetText(fmt.Sprintf("%d", useFixture.Address))
		}

		// Based on a new use fixure - Try again to populate the program and rotate options as available for this states action.
		fixture, err := findFixtureByLabel(value, fixtures)
		if err != nil {
			fmt.Printf("findFixtureByName: fixtureName: %s error %s\n", fixture.Name, err.Error())
			return
		}
		ap.ActionProgramOptions = populateOptions(fixture, "Program", fixtures)

		// Based on a new use fixure - Try again to populate the channel names in the settings panel.
		st.ChannelOptions = populateChannelNames(useFixture.Channels)

	}

	// Show the currently selected fixture option.
	for _, option := range fixturesAvailable {
		if option == thisFixture.UseFixture {
			useInput.SetSelected(option)
		}
	}

	formTop := container.NewVBox(name, desc, use, address)

	labelStates := widget.NewLabel("Switch States")
	labelStates.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Populate State settings.
	statesList := thisFixture.States

	// Create States Panel.
	var StatesPanel *widget.Table
	sp := NewStatePanel(statesList, ap, st)
	StatesPanel = sp.StatePanel

	// Setup forms.
	scrollableStateList := container.NewScroll(StatesPanel)
	scrollableStateList.SetMinSize(fyne.Size{Height: 400, Width: 300})

	scrollableActionsList := container.NewScroll(ap.ActionsPanel)
	scrollableActionsList.SetMinSize(fyne.Size{Height: 400, Width: 300})

	scrollableSettingsList := container.NewScroll(st.SettingsPanel)
	scrollableSettingsList.SetMinSize(fyne.Size{Height: 400, Width: 300})

	// Setup OK buttons action.
	buttonSave.OnTapped = func() {

		// Populate the fixture used by this switch.
		fp.UseFixture = useInput.Selected

		// Update the address from the use fixture field in the fixture panel.
		fp.UseFixture = useInput.Selected
		fp.UpdateThisFixture = thisFixture.ID - 1
		fp.UpdateUseFixture = true

		// Insert updated fixture into fixtures.
		newFixtures := fixture.Fixtures{}
		for fixtureNumber, fixture := range fixtures.Fixtures {
			if fixture.ID == id {
				// Insert new states into fixture above us, in the fixture selection panel.
				fp.UpdateStates = true
				fp.UpdateThisFixture = fixtureNumber
				// Update our copy of the state list.
				fp.UpdatedStatesList = sp.StatesList
				newFixtures.Fixtures = append(newFixtures.Fixtures, thisFixture)
			} else {
				newFixtures.Fixtures = append(newFixtures.Fixtures, fixture)
			}
		}

		modal.Hide()

		// Refresh the fixtures panel incase something has changed.
		fp.FixturePanel.Refresh()
	}

	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		modal.Hide()
	})

	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)

	content := fyne.Container{}

	top := container.NewBorder(title, nil, nil, nil, formTop)
	main := container.NewBorder(top, nil, nil, nil, labelStates)
	scrollableList := container.New(layout.NewStackLayout(), scrollableActionsList, scrollableSettingsList)
	forms := container.NewAdaptiveGrid(2, scrollableStateList, scrollableList)
	bottom := container.NewBorder(main, nil, nil, nil, forms)
	content = *container.NewBorder(bottom, nil, nil, nil, saveCancel)

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		&content,
		w.Canvas(),
	)
	return modal, nil

}

func populateOptions(thisFixture *fixture.Fixture, name string, fixtures *fixture.Fixtures) []string {

	if debug {
		fmt.Printf("populateOptions\n")
	}

	options := []string{}
	programSettings, err := fixture.GetChannelSettinsByName(thisFixture, name, fixtures)
	if err == nil {
		for _, setting := range programSettings {
			options = append(options, setting.Name)
		}
	}
	if len(options) == 0 {
		options = append(options, "None")
	}

	return options
}

func populateChannelNames(channels []fixture.Channel) []string {

	if debug {
		fmt.Printf("Channels available are %v\n", channels)
	}

	options := []string{}

	for _, channel := range channels {
		options = append(options, channel.Name)
	}
	if len(options) == 0 {
		options = append(options, "None")
	}

	return options
}

func findFixtureByLabel(label string, fixtures *fixture.Fixtures) (*fixture.Fixture, error) {

	if debug {
		fmt.Printf("Look for fixture by Label %s\n", label)
	}

	if label == "" {
		return nil, fmt.Errorf("findFixtureByName: fixture name is empty")
	}

	for _, fixture := range fixtures.Fixtures {
		if strings.Contains(fixture.Label, label) {
			if debug {
				fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
			}
			return &fixture, nil
		}
	}
	return nil, fmt.Errorf("findFixtureByLabel: failed to find fixture by label %s", label)
}
