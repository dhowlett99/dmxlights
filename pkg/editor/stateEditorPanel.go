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

	// Title.
	title := widget.NewLabel(fmt.Sprintf("Edit Config for Sequence %d Fixture %d", thisFixture.Group, thisFixture.Number))
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Name, description and which Fixture to use.
	nameInput := widget.NewEntry()
	nameInput.SetPlaceHolder(thisFixture.Name)
	descInput := widget.NewEntry()
	descInput.SetPlaceHolder(thisFixture.Description)
	addrInput := widget.NewEntry()
	addrInput.SetPlaceHolder(thisFixture.UseFixture)

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
	formTopItem3 := widget.NewFormItem("Use Fixture", name3)
	formTopItems = append(formTopItems, formTopItem3)
	formTop := &widget.Form{
		Items: formTopItems,
	}

	labelStates := widget.NewLabel("Switch States")
	labelStates.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	// Populate State settings.
	statesList := thisFixture.States

	// Create Actions Panel.
	ap := NewActionsPanel([]fixture.Action{})
	ap.ActionsPanel.Hide()

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
