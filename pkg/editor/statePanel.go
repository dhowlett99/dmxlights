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
	StatePanel         *widget.Table
	StatesList         []fixture.State
	ButtonColorOptions []string
	StateOptions       []string
}

const (
	STATE_NUMBER int = iota
	STATE_NAME
	STATE_BUTTONCOLOR
	STATE_DELETE
	STATE_ADD
	STATE_ACTIONS
	STATE_SETTINGS
)

func NewStatePanel(statesList []fixture.State, ap *ActionPanel, st *SettingsPanel) *StatePanel {

	if debug {
		fmt.Printf("NewStatePanel\n")
	}

	var data = [][]string{}

	sp := StatePanel{}
	sp.ButtonColorOptions = []string{"Red", "Orange", "Yellow", "Green", "Cyan", "Blue", "Purple", "Pink", "White", "Black"}
	sp.StateOptions = []string{"Off", "On", "Red", "Green", "Blue", "Soft", "Sharp", "Sound", "Rotate"}
	sp.StatesList = statesList

	// statees Selection Panel.
	sp.StatePanel = widget.NewTable(
		// Function to find length.
		func() (int, int) {
			if ap.UpdateActions {
				sp.StatesList[ap.UpdateThisAction].Actions = ap.ActionsList
				ap.UpdateActions = false
			}
			if st.UpdateSettings {
				sp.StatesList[st.UpdateThisChannel].Settings = st.SettingsList
				st.UpdateSettings = false
			}

			height := len(data)
			width := 7
			return height, width

		},
		// Function to create table.
		func() (o fyne.CanvasObject) {

			// Load the fixtures into the array used by the table.
			data = updateStatesArray(sp.StatesList)

			return container.NewStack(

				// State Number.
				widget.NewLabel("template"),

				// State Name.
				widget.NewEntry(),

				// Button Color.
				widget.NewSelect(sp.ButtonColorOptions, func(value string) {}),

				// Chanell delete button.
				widget.NewButton("-", func() {}),

				// Channel add button
				widget.NewButton("+", func() {}),

				// Actions button.
				widget.NewButton("Actions", nil),

				// Settings button.
				widget.NewButton("Values", nil),
			)
		},
		// Function to update item in this table.
		func(thisState widget.TableCellID, o fyne.CanvasObject) {

			// Hide all field types.
			hideAllStatesFields(o)

			// Show the state Number.
			if thisState.Col == STATE_NUMBER {
				showStatesField(STATE_NUMBER, o)
				o.(*fyne.Container).Objects[STATE_NUMBER].(*widget.Label).SetText(fmt.Sprintf("%d", sp.StatesList[thisState.Row].Number))
			}

			// Show the state name.
			if thisState.Col == STATE_NAME {
				showStatesField(STATE_NAME, o)
				o.(*fyne.Container).Objects[STATE_NAME].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[STATE_NAME].(*widget.Entry).SetText(sp.StatesList[thisState.Row].Name)
				o.(*fyne.Container).Objects[STATE_NAME].(*widget.Entry).OnChanged = func(value string) {
					newState := fixture.State{}
					newState.Name = value
					newState.Number = sp.StatesList[thisState.Row].Number
					newState.Master = sp.StatesList[thisState.Row].Master
					newState.Label = value
					newState.ButtonColor = sp.StatesList[thisState.Row].ButtonColor
					newState.Flash = sp.StatesList[thisState.Row].Flash
					newState.Settings = sp.StatesList[thisState.Row].Settings
					newState.Actions = sp.StatesList[thisState.Row].Actions
					sp.StatesList = updateStateItem(sp.StatesList, sp.StatesList[thisState.Row].Number, newState)
					data = updateStatesArray(sp.StatesList)
				}
			}

			// Show the selection box for button color.
			if thisState.Col == STATE_BUTTONCOLOR {
				showStatesField(STATE_BUTTONCOLOR, o)

				o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).OnChanged = nil
				for _, option := range sp.ButtonColorOptions {
					if option == sp.StatesList[thisState.Row].ButtonColor {
						o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).SetSelected(option)
					}
				}
				o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).OnChanged = func(value string) {
					newState := fixture.State{}
					newState.Name = sp.StatesList[thisState.Row].Name
					newState.Number = sp.StatesList[thisState.Row].Number
					newState.Master = sp.StatesList[thisState.Row].Master
					newState.Label = sp.StatesList[thisState.Row].Label
					newState.ButtonColor = value
					newState.Flash = sp.StatesList[thisState.Row].Flash
					newState.Settings = sp.StatesList[thisState.Row].Settings
					newState.Actions = sp.StatesList[thisState.Row].Actions
					sp.StatesList = updateStateItem(sp.StatesList, sp.StatesList[thisState.Row].Number, newState)
					data = updateStatesArray(sp.StatesList)
				}
				o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).PlaceHolder = "Select"
			}

			// State delete button.
			if thisState.Col == STATE_DELETE {
				showStatesField(STATE_DELETE, o)
				o.(*fyne.Container).Objects[STATE_DELETE].(*widget.Button).OnTapped = func() {
					sp.StatesList = deleteState(sp.StatesList, sp.StatesList[thisState.Row].Number)
					data = updateStatesArray(sp.StatesList)
					sp.StatePanel.Refresh()
				}
			}

			// State add button.
			if thisState.Col == STATE_ADD {
				showStatesField(STATE_ADD, o)
				o.(*fyne.Container).Objects[STATE_ADD].(*widget.Button).OnTapped = func() {
					sp.StatesList = addState(sp.StatesList, sp.StatesList[thisState.Row].Number)
					data = updateStatesArray(sp.StatesList)
					sp.StatePanel.Refresh()
				}
			}

			// Actions button.
			if thisState.Col == STATE_ACTIONS {
				showStatesField(STATE_ACTIONS, o)
				o.(*fyne.Container).Objects[STATE_ACTIONS].(*widget.Button).OnTapped = func() {
					// Highlight this channel
					sp.StatePanel.Select(thisState)
					if sp.StatesList != nil {
						// Get Existing Actions for this state.
						index := sp.StatesList[thisState.Row].Number - 1
						data = updateStatesArray(sp.StatesList)
						ap.ActionsList = sp.StatesList[index].Actions

						// Remove any actions which are off from any previous selections.
						//ap.ActionsList = ClearOffActions(ap.ActionsList)

						// If the settings are empty create a new set of settings.
						if len(ap.ActionsList) == 0 {
							// Create new settings.
							ap.ActionsList = append(ap.ActionsList, CreateActionsList(sp.StatesList, thisState.Row))
						}
					}
					ap.CurrentState = int(sp.StatesList[thisState.Row].Number - 1)
					ap.CurrentStateName = sp.StatesList[thisState.Row].Name
					ap.ActionsPanel.Hidden = false
					st.SettingsPanel.Hidden = true
					ap.ActionsPanel.Refresh()
				}
			}

			// Settings Button.
			if thisState.Col == STATE_SETTINGS {
				showStatesField(STATE_SETTINGS, o)
				o.(*fyne.Container).Objects[STATE_SETTINGS].(*widget.Button).OnTapped = func() {
					// Highlight this channel
					sp.StatePanel.Select(thisState)
					if sp.StatesList != nil {
						// Get Existing Settings for channel.
						st.SettingsList = populateSettingList(sp.StatesList, sp.StatesList[thisState.Row].Number)
						data = updateStatesArray(sp.StatesList)
						// If the settings are empty create a new set of settings.
						if len(st.SettingsList) == 0 {
							// Create new settings.
							st.SettingsList = createSettingList(sp.StatesList[thisState.Row].Number)
							st.CurrentChannel = int(sp.StatesList[thisState.Row].Number)

						} else {
							// Edit existing settings.
							st.CurrentChannel = int(sp.StatesList[thisState.Row].Number)
						}
						ap.ActionsPanel.Hidden = true
						st.SettingsPanel.Hidden = false
						st.SettingsPanel.Refresh()
					}
				}
			}
		},
	)

	// Setup the columns of this table.
	sp.StatePanel.SetColumnWidth(0, 40)  // Number
	sp.StatePanel.SetColumnWidth(1, 80)  // Name
	sp.StatePanel.SetColumnWidth(2, 100) // Button Color
	sp.StatePanel.SetColumnWidth(3, 20)  // Delete
	sp.StatePanel.SetColumnWidth(4, 20)  // Add
	sp.StatePanel.SetColumnWidth(5, 60)  // Actions
	sp.StatePanel.SetColumnWidth(6, 60)  // Settings

	return &sp
}

// UpdateItem replaces the selected item by id with newItem.
func updateStateItem(states []fixture.State, id int16, newState fixture.State) []fixture.State {

	if debug {
		fmt.Printf("updateStateItem\n")
	}

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

func addState(states []fixture.State, id int16) (outItems []fixture.State) {

	if debug {
		fmt.Printf("addState\n")
	}

	newStates := []fixture.State{}
	newItem := fixture.State{}
	newItem.Number = id + 1
	newItem.Name = "New"

	var added bool // Only add once.

	for no, item := range states {
		// Add at the start of an empty list.
		if len(states) == 0 && !added {
			newStates = append(newStates, newItem)
			added = true
		}
		// Insert at this position.
		if item.Number == id+1 && !added {
			newStates = append(newStates, newItem)
			added = true
		}
		newStates = append(newStates, item)
		// Append an item at the very end.
		if no == len(states)-1 && !added {
			newStates = append(newStates, newItem)
			added = true
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newStates {
		indexedItem.Number = int16(number + 1)
		outItems = append(outItems, indexedItem)
	}

	return outItems
}

func deleteState(stateList []fixture.State, id int16) (outItems []fixture.State) {

	if debug {
		fmt.Printf("deleteState\n")
	}

	newStates := []fixture.State{}
	for _, channel := range stateList {
		if channel.Number != id {
			newStates = append(newStates, channel)
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newStates {
		indexedItem.Number = int16(number + 1)
		outItems = append(outItems, indexedItem)
	}

	if len(outItems) == 0 {
		// Create a default State
		newItem := fixture.State{}
		newItem.Number = 1
		newItem.Name = "NewState"
		outItems = append(outItems, newItem)
	}

	return outItems
}

func populateSettingList(statesList []fixture.State, stateNumber int16) (settingsList []fixture.Setting) {

	if debug {
		fmt.Printf("populateSettingList\n")
	}

	for _, state := range statesList {
		if stateNumber == state.Number {
			return state.Settings
		}
	}
	return settingsList
}

func createSettingList(channelNumber int16) (settingsList []fixture.Setting) {

	if debug {
		fmt.Printf("createSettingList\n")
	}

	newItem := fixture.Setting{}
	newItem.Name = "New Setting"
	newItem.Number = 1
	newItem.Value = "0"
	settingsList = append(settingsList, newItem)
	return settingsList
}

// makeStatesArray - Convert the list of states to an array of strings containing and array of strings with
// the values from each state.
// This is done once when the state panel is loaded.
func updateStatesArray(states []fixture.State) [][]string {

	if debug {
		fmt.Printf("makeSettingsArray\n")
	}

	var data = [][]string{}

	for _, state := range states {
		newState := []string{}
		newState = append(newState, fmt.Sprintf("%d", state.Number))
		newState = append(newState, state.Name)
		newState = append(newState, state.ButtonColor)
		newState = append(newState, "-")
		newState = append(newState, "+")
		newState = append(newState, "Actions")
		newState = append(newState, "Values")
		data = append(data, newState)
	}

	return data
}

func showStatesField(field int, o fyne.CanvasObject) {
	if debug {
		fmt.Printf("showField\n")
	}
	// Now show the selected field.
	switch {
	case field == STATE_NUMBER:
		o.(*fyne.Container).Objects[STATE_NUMBER].(*widget.Label).Hidden = false
	case field == STATE_NAME:
		o.(*fyne.Container).Objects[STATE_NAME].(*widget.Entry).Hidden = false
	case field == STATE_BUTTONCOLOR:
		o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).Hidden = false
	case field == STATE_DELETE:
		o.(*fyne.Container).Objects[STATE_DELETE].(*widget.Button).Hidden = false
	case field == STATE_ADD:
		o.(*fyne.Container).Objects[STATE_ADD].(*widget.Button).Hidden = false
	case field == STATE_ACTIONS:
		o.(*fyne.Container).Objects[STATE_ACTIONS].(*widget.Button).Hidden = false
	case field == STATE_SETTINGS:
		o.(*fyne.Container).Objects[STATE_SETTINGS].(*widget.Button).Hidden = false
	}
}

func hideAllStatesFields(o fyne.CanvasObject) {
	if debug {
		fmt.Printf("hideAllSettingsFields\n")
	}
	o.(*fyne.Container).Objects[STATE_NUMBER].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[STATE_NAME].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).Hidden = true
	o.(*fyne.Container).Objects[STATE_DELETE].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[STATE_ADD].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[STATE_ACTIONS].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[STATE_SETTINGS].(*widget.Button).Hidden = true

}
