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

type StatePanel struct {
	StatePanel   *widget.List
	StatesList   []fixture.State
	StateOptions []string
}

const STATE_ID int = 0
const STATE_NAME int = 1
const STATE_ACTIONS int = 2

func NewStatePanel(statesList []fixture.State, ap *ActionPanel) *StatePanel {

	sp := StatePanel{}
	sp.StateOptions = []string{"Off", "On", "Red", "Green", "Blue", "Soft", "Sharp", "Sound", "Rotate"}
	sp.StatesList = statesList

	// statees Selection Panel.
	sp.StatePanel = widget.NewList(
		// Function to find length.
		func() int {
			if ap.UpdateActions {
				sp.StatesList[ap.UpdateThisAction].Actions = ap.ActionsList
				ap.UpdateActions = false
			}
			return len(sp.StatesList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {
			return container.NewHBox(

				// State Number.
				widget.NewLabel("template"),

				// State Name.
				widget.NewSelect(sp.StateOptions, func(value string) {
				}),

				// Select Actions.
				widget.NewButton("Select", nil),
			)
		},
		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// Show the Channel Number.
			o.(*fyne.Container).Objects[STATE_ID].(*widget.Label).SetText(fmt.Sprintf("%d", sp.StatesList[i].Number))

			// Show the currently selected state from the options list.
			for _, option := range sp.StateOptions {
				if option == sp.StatesList[i].Label {
					o.(*fyne.Container).Objects[STATE_NAME].(*widget.Select).SetSelected(option)
				}
			}
			o.(*fyne.Container).Objects[STATE_NAME].(*widget.Select).OnChanged = func(value string) {
				newState := fixture.State{}
				newState.Name = value
				newState.Number = sp.StatesList[i].Number
				newState.Master = sp.StatesList[i].Master
				newState.Label = sp.StatesList[i].Label
				newState.ButtonColor = sp.StatesList[i].ButtonColor
				newState.Flash = sp.StatesList[i].Flash
				newState.Values = sp.StatesList[i].Values
				newState.Actions = sp.StatesList[i].Actions
				sp.StatesList = UpdateStateItem(sp.StatesList, sp.StatesList[i].Number, newState)
			}

			// Show the Actions Button.
			o.(*fyne.Container).Objects[STATE_ACTIONS].(*widget.Button).OnTapped = func() {
				ap.ActionsList = []fixture.Action{}
				for _, action := range sp.StatesList[i].Actions {
					newAction := fixture.Action{}
					newAction.Name = action.Name
					newAction.Colors = action.Colors
					newAction.Mode = action.Mode
					newAction.Fade = action.Fade
					newAction.Speed = action.Speed
					ap.ActionsList = append(ap.ActionsList, newAction)
				}
				if debug {
					fmt.Printf("I am button %d actions %+v\n", sp.StatesList[i].Number, ap.ActionsList)
				}
				ap.ActionsPanel.Hidden = false
				ap.CurrentState = int(sp.StatesList[i].Number)
				ap.ActionsPanel.Refresh()
			}
		})
	return &sp
}

func populateStates(thisFixture fixture.Fixture) (actionsList []fixture.Action, statesList []fixture.State) {

	// Populate state state settings and actions.
	for _, state := range thisFixture.States {
		newState := fixture.State{}
		newState.Name = state.Name
		newState.Number = state.Number
		newState.Label = state.Label
		newState.Values = state.Values
		newState.ButtonColor = state.ButtonColor
		newState.Master = state.Master
		newState.Actions = state.Actions
		newState.Flash = state.Flash
		if state.Actions != nil {
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
				newAction.Music = action.Music
				newAction.Program = action.Program
				newAction.Strobe = action.Strobe

				actionsList = append(actionsList, newAction)
			}
		}
		newState.Actions = actionsList
		statesList = append(statesList, newState)
	}

	return actionsList, statesList
}

// UpdateItem replaces the selected item by id with newItem.
func UpdateStateItem(states []fixture.State, id int16, newState fixture.State) []fixture.State {
	newStates := []fixture.State{}
	for _, state := range states {
		if state.Number == id {
			// update the channel information.
			newStates = append(newStates, newState)
		} else {
			// just add what was there before.
			newStates = append(newStates, state)
		}
	}
	return newStates
}
