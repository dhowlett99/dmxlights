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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

// Show a list of Statees
func NewStateEditor(w fyne.Window, id int, fp *FixturesPanel, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	thisFixture, err := fixture.GetFixureDetails(id, fixtures)
	if err != nil {
		return nil, fmt.Errorf("GetFixureDetails %s", err.Error())
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

	// Create Actions Panel.
	ap := NewActionsPanel(w, []fixture.Action{})
	ap.ActionsPanel.Hide()

	// Use Fixture.
	useInput := widget.NewSelect(fixturesAvailable, func(value string) {})
	useLabel := widget.NewLabel("Use Fixture")
	use := container.NewAdaptiveGrid(3, useLabel, useInput, layout.NewSpacer())

	// Update Use Fixture.
	useInput.OnChanged = func(value string) {

		// Update the address from the use fixture field in the fixture panel.
		fp.UseFixture = useInput.Selected
		fp.UpdateThisFixture = thisFixture.ID - 1
		fp.UpdateUseFixture = true

		// Try again to populate the program options as available for this states action.
		ap.ActionProgramOptions = populateProgramOptions(value, fixtures)
	}

	// Show the currently selected fixture option.
	for _, option := range fixturesAvailable {
		if option == thisFixture.UseFixture {
			useInput.SetSelected(option)
		}
	}

	formTop := container.NewVBox(name, desc, use)

	labelStates := widget.NewLabel("Switch States")
	labelStates.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Populate State settings.
	statesList := thisFixture.States

	// Create States Panel.
	var StatesPanel *widget.List
	sp := NewStatePanel(statesList, ap)
	StatesPanel = sp.StatePanel

	// Setup forms.
	scrollableStateList := container.NewScroll(StatesPanel)
	scrollableStateList.SetMinSize(fyne.Size{Height: 400, Width: 300})
	scrollableActionsList := container.NewScroll(ap.ActionsPanel)
	scrollableActionsList.SetMinSize(fyne.Size{Height: 400, Width: 300})

	// OK button.
	buttonSave := widget.NewButton("OK", func() {

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
	})

	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		modal.Hide()
	})

	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)

	content := fyne.Container{}

	top := container.NewBorder(title, nil, nil, nil, formTop)
	main := container.NewBorder(top, nil, nil, nil, labelStates)
	forms := container.NewAdaptiveGrid(2, scrollableStateList, scrollableActionsList)
	bottom := container.NewBorder(main, nil, nil, nil, forms)
	content = *container.NewBorder(bottom, nil, nil, nil, saveCancel)

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		&content,
		w.Canvas(),
	)
	return modal, nil

}

func populateProgramOptions(fixtureName string, fixtures *fixture.Fixtures) []string {

	programOptions := []string{}
	programSettings, err := fixture.GetProgramSettins(fixtureName, fixtures)
	if err != nil {
		fmt.Printf("populateProgramOptions: no program settings found for fixture %s\n", fixtureName)
	} else {
		for _, setting := range programSettings {
			programOptions = append(programOptions, setting.Name)
		}
	}
	if len(programOptions) == 0 {
		programOptions = append(programOptions, "None")
	}

	return programOptions
}
