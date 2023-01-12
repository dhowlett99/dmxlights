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
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type StatePanel struct {
	StatePanel         *widget.List
	StatesList         []fixture.State
	ButtonColorOptions []string
	StateOptions       []string
}

const (
	STATE_ID int = iota
	STATE_NAME
	STATE_BUTTONCOLOR
	STATE_DELETE
	STATE_ADD
	STATE_ACTIONS
	STATE_SETTINGS
)

func NewStatePanel(statesList []fixture.State, ap *ActionPanel, st *SettingsPanel) *StatePanel {

	sp := StatePanel{}
	sp.ButtonColorOptions = []string{"Red", "Orange", "Yellow", "Green", "Cyan", "Blue", "Purple", "Pink", "White", "Black"}
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
			if st.UpdateSettings {
				sp.StatesList[st.UpdateThisChannel].Settings = st.SettingsList
				st.UpdateSettings = false
			}
			return len(sp.StatesList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {

			return container.NewHBox(

				// State Id.
				container.NewWithoutLayout(widget.NewLabel("template")),

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
				widget.NewButton("Settings", nil),
			)
		},
		// Function to update item in this list.
		func(thisState widget.ListItemID, o fyne.CanvasObject) {

			// Show the state Number.
			o.(*fyne.Container).Objects[STATE_ID].(*fyne.Container).Objects[0].Resize(fyne.Size{Height: 50, Width: 30})
			o.(*fyne.Container).Objects[STATE_ID].(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", sp.StatesList[thisState].Number))

			// Show the state name.
			o.(*fyne.Container).Objects[STATE_NAME].(*widget.Entry).SetText(sp.StatesList[thisState].Name)
			o.(*fyne.Container).Objects[STATE_NAME].(*widget.Entry).OnChanged = func(value string) {
				newState := fixture.State{}
				newState.Name = value
				newState.Number = sp.StatesList[thisState].Number
				newState.Master = sp.StatesList[thisState].Master
				newState.Label = value
				newState.ButtonColor = sp.StatesList[thisState].ButtonColor
				newState.Flash = sp.StatesList[thisState].Flash
				newState.Settings = sp.StatesList[thisState].Settings
				newState.Actions = sp.StatesList[thisState].Actions
				sp.StatesList = updateStateItem(sp.StatesList, sp.StatesList[thisState].Number, newState)
			}

			// Show the selection box for button color.
			for _, option := range sp.ButtonColorOptions {
				if option == sp.StatesList[thisState].ButtonColor {
					o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).SetSelected(option)
				}
			}

			o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).OnChanged = func(value string) {
				newState := fixture.State{}
				newState.Name = sp.StatesList[thisState].Name
				newState.Number = sp.StatesList[thisState].Number
				newState.Master = sp.StatesList[thisState].Master
				newState.Label = sp.StatesList[thisState].Label
				newState.ButtonColor = value
				newState.Flash = sp.StatesList[thisState].Flash
				newState.Settings = sp.StatesList[thisState].Settings
				newState.Actions = sp.StatesList[thisState].Actions
				sp.StatesList = updateStateItem(sp.StatesList, sp.StatesList[thisState].Number, newState)
			}
			o.(*fyne.Container).Objects[STATE_BUTTONCOLOR].(*widget.Select).PlaceHolder = "Select"

			// Channel Delete Button.
			o.(*fyne.Container).Objects[STATE_DELETE].(*widget.Button).OnTapped = func() {
				sp.StatesList = deleteState(sp.StatesList, sp.StatesList[thisState].Number)
				sp.StatePanel.Refresh()
			}

			// Channel Add Button.
			o.(*fyne.Container).Objects[STATE_ADD].(*widget.Button).OnTapped = func() {
				sp.StatesList = addState(sp.StatesList, sp.StatesList[thisState].Number)
				sp.StatePanel.Refresh()
			}

			// Actions Button.
			o.(*fyne.Container).Objects[STATE_ACTIONS].(*widget.Button).OnTapped = func() {
				// Highlight this channel
				sp.StatePanel.Select(thisState)
				if sp.StatesList != nil {
					// Get Existing Actions for this state.
					index := sp.StatesList[thisState].Number - 1
					ap.ActionsList = sp.StatesList[index].Actions

					// If the settings are empty create a new set of settings.
					if len(ap.ActionsList) == 0 {
						// Create new settings.
						ap.ActionsList = CreateActionsList(sp.StatesList, thisState)
					}
				}
				ap.CurrentState = int(sp.StatesList[thisState].Number - 1)
				ap.CurrentStateName = sp.StatesList[thisState].Name
				ap.ActionsPanel.Hidden = false
				st.SettingsPanel.Hidden = true
				ap.ActionsPanel.Refresh()
			}

			// Settings Button.
			o.(*fyne.Container).Objects[STATE_SETTINGS].(*widget.Button).OnTapped = func() {
				// Highlight this channel
				sp.StatePanel.Select(thisState)
				fmt.Printf("--> settings pannel with list %+v\n", sp.StatesList[thisState])
				if sp.StatesList != nil {
					// Get Existing Settings for channel.
					st.SettingsList = populateSettingList(sp.StatesList, sp.StatesList[thisState].Number)
					// If the settings are empty create a new set of settings.
					if len(st.SettingsList) == 0 {
						// Create new settings.
						st.SettingsList = createSettingList(sp.StatesList[thisState].Number)
						st.CurrentChannel = int(sp.StatesList[thisState].Number)

					} else {
						// Edit existing settings.
						st.CurrentChannel = int(sp.StatesList[thisState].Number)
						fmt.Printf("Refresh the settings pannel with list %+v\n", st.SettingsList[thisState])
					}
					ap.ActionsPanel.Hidden = true
					st.SettingsPanel.Hidden = false
					st.SettingsPanel.Refresh()
				}
			}
		})
	return &sp
}

// UpdateItem replaces the selected item by id with newItem.
func updateStateItem(states []fixture.State, id int16, newState fixture.State) []fixture.State {
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

func addState(states []fixture.State, id int16) []fixture.State {
	newStates := []fixture.State{}
	newItem := fixture.State{}
	newItem.Number = id + 1
	if stateItemAllreadyExists(newItem.Number, states) {
		newItem.Number = findLargestStateNumber(states) + 1
	}
	newItem.Name = "New"

	for _, item := range states {
		if item.Number == id {
			newStates = append(newStates, newItem)
		}
		newStates = append(newStates, item)
	}
	sort.Slice(newStates, func(i, j int) bool {
		return newStates[i].Number < newStates[j].Number
	})
	return newStates
}

func deleteState(stateList []fixture.State, id int16) []fixture.State {
	newStates := []fixture.State{}
	if id == 1 {
		return stateList
	}
	for _, channel := range stateList {
		if channel.Number != id {
			newStates = append(newStates, channel)
		}
	}
	return newStates
}

func stateItemAllreadyExists(number int16, stateList []fixture.State) bool {
	// look through the state list for the id's
	for _, item := range stateList {
		if item.Number == number {
			return true
		}
	}
	return false
}

func findLargestStateNumber(items []fixture.State) int16 {
	var number int16
	for _, item := range items {
		if item.Number > number {
			number = item.Number
		}
	}
	return number
}

func populateSettingList(statesList []fixture.State, stateNumber int16) (settingsList []fixture.Setting) {
	for _, state := range statesList {
		if stateNumber == state.Number {
			return state.Settings
		}
	}
	return settingsList
}

func createSettingList(channelNumber int16) (settingsList []fixture.Setting) {
	newItem := fixture.Setting{}
	newItem.Name = "New Setting"
	newItem.Number = 1
	newItem.Setting = "0"
	settingsList = append(settingsList, newItem)
	return settingsList
}
