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
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type SwitchPanel struct {
	SwitchPanel   *widget.List
	SwitchesList  []itemSelect
	SwitchOptions []string
}

func NewSwitchPanel(switchesAvailable bool, switchesList []itemSelect, ap *ActionPanel) *SwitchPanel {

	sw := SwitchPanel{}
	sw.SwitchesList = switchesList
	sw.SwitchOptions = []string{"Off", "On", "Red", "Green", "Blue", "Soft", "Sharp", "Sound", "Rotate"}

	// Switches Selection Panel.
	if switchesAvailable {
		sw.SwitchPanel = widget.NewList(
			func() int {
				return len(sw.SwitchesList)
			},
			func() (o fyne.CanvasObject) {
				return container.NewHBox(
					widget.NewLabel("template"),

					widget.NewSelect(sw.SwitchOptions, func(value string) {
					}),

					widget.NewButton("Select", nil),
				)
			},
			// Function to update item in this list.
			func(i widget.ListItemID, o fyne.CanvasObject) {

				o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", sw.SwitchesList[i].Number))

				// find the selected option in the options list.
				for _, option := range sw.SwitchesList[i].Options {
					if option == sw.SwitchesList[i].Label {
						o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
					}
				}

				// new part
				o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
					ap.ActionsList = []fixture.Action{}
					for _, action := range sw.SwitchesList[i].Actions {
						newAction := fixture.Action{}
						newAction.Name = action.Name
						newAction.Colors = action.Colors
						newAction.Mode = action.Mode
						newAction.Fade = action.Fade
						newAction.Speed = action.Speed
						ap.ActionsList = append(ap.ActionsList, newAction)
					}
					if debug {
						fmt.Printf("I am button %d actions %+v\n", sw.SwitchesList[i].Number, ap.ActionsList)
					}
					ap.ActionsPanel.Hidden = false
					ap.ActionsPanel.Refresh()
				}
			})
	}
	return &sw
}

func PopulateSwitches(thisFixture fixture.Fixture) (switchesAvailable bool, actionsAvailable bool,
	actionsList []fixture.Action, switchesList []itemSelect) {

	switchOptions := []string{"Off", "On", "Red", "Green", "Blue", "Soft", "Sharp", "Sound", "Rotate"}

	// Populate switch state settings and actions.
	if thisFixture.Type == "switch" {
		//labelSwitch.Text = "Switch States"
		for _, state := range thisFixture.States {
			switchesAvailable = true
			newSelect := itemSelect{}
			newSelect.Number = state.Number
			newSelect.Label = state.Name
			newSelect.Options = switchOptions
			if state.Actions != nil {
				actionsAvailable = true
				actionsList = []fixture.Action{}
				for _, action := range state.Actions {
					newAction := fixture.Action{}
					newAction.Name = action.Name
					newAction.Colors = action.Colors
					newAction.Mode = action.Mode
					newAction.Fade = action.Fade
					if action.Speed != "" {
						newAction.Speed = action.Speed
					} else {
						newAction.Speed = "none"
					}
					newAction.Rotate = action.Rotate
					newAction.Program = action.Program
					newAction.Strobe = action.Strobe

					actionsList = append(actionsList, newAction)
				}
			}
			newSelect.Actions = actionsList
			switchesList = append(switchesList, newSelect)
		}
	}

	return switchesAvailable, actionsAvailable, actionsList, switchesList
}
