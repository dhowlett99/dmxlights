// Copyright (C) 2022,2023 dhowlett99.
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
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type ActionPanel struct {
	ActionsPanel              *widget.List
	ActionsList               []fixture.Action
	ActionNameOptions         []string
	ActionColorsOptions       []string
	ActionMapOptions          []string
	ActionModeOptions         []string
	ActionFadeOptions         []string
	ActionSizeOptions         []string
	ActionSpeedOptions        []string
	ActionRotateOptions       []string
	ActionRotateSpeedOptions  []string
	ActionMusicOptions        []string
	ActionProgramOptions      []string
	ActionProgramSpeedOptions []string
	ActionStrobeOptions       []string
	ActionGoboOptions         []string
	ActionGoboSpeedOptions    []string
	UpdateActions             bool
	UpdateThisAction          int
	CurrentState              int
	CurrentStateName          string
}

const (
	LABEL int = iota
	SELECT
	COLOR_SELECTION_BOX
	RADIO_BUTTON
)

const (
	COLOR1 int = iota
	COLOR2
	COLOR3
	COLOR4
	COLOR5
	COLOR6
	COLOR7
	COLOR8
	COLOR9
	COLOR10
)

const (
	ACTIONS_MODE int = iota
	ACTIONS_COLORS
	ACTIONS_FADE
	ACTIONS_SIZE
	ACTIONS_SPEED
	ACTIONS_ROTATE
	ACTIONS_ROTATESPEED
	ACTIONS_PROGRAM
	ACTIONS_PROGRAM_SPEED
	ACTIONS_STROBE
	ACTIONS_GOBO
	ACTIONS_GOBO_SPEED
)

func NewActionsPanel(w fyne.Window, actionsList []fixture.Action, fixtureInfo fixture.FixtureInfo) *ActionPanel {

	ap := ActionPanel{}
	ap.ActionsList = actionsList
	ap.ActionModeOptions = []string{"None", "Off", "Static", "Chase", "Control"}
	ap.ActionSizeOptions = []string{"Off", "Short", "Medium", "Long"}
	ap.ActionFadeOptions = []string{"Off", "Soft", "Normal", "Sharp"}
	ap.ActionSpeedOptions = []string{"Off", "Slow", "Medium", "Fast", "VeryFast", "Music"}
	ap.ActionRotateOptions = []string{"Off", "Clockwise", "Anti Clockwise", "Auto"}
	ap.ActionRotateSpeedOptions = []string{"Slow", "Medium", "Fast"}
	ap.ActionProgramSpeedOptions = []string{"Slow", "Medium", "Fast"}
	ap.ActionMusicOptions = []string{"Off", "On"}
	ap.ActionStrobeOptions = []string{"Off", "Slow", "Medium", "Fast"}
	ap.ActionMapOptions = []string{"Off", "On"}
	// ap.ActionGoboOptions are setup in the StatePanel that calls this func.
	ap.ActionGoboSpeedOptions = []string{"Slow", "Medium", "Fast"}

	cp := NewColorPickerPanel()

	// Actions Selection Panel.
	ap.ActionsPanel = widget.NewList(
		// Function to find length.
		func() int {
			if cp.UpdateColors {
				newAction := fixture.Action{}
				newAction.Name = ap.ActionsList[cp.UpdateThisAction].Name
				newAction.Number = ap.ActionsList[cp.UpdateThisAction].Number
				newAction.Colors = strings.Split(cp.ColorSelection, ",")
				newAction.Mode = ap.ActionsList[cp.UpdateThisAction].Mode
				newAction.Fade = ap.ActionsList[cp.UpdateThisAction].Fade
				newAction.Size = ap.ActionsList[cp.UpdateThisAction].Size
				newAction.Speed = ap.ActionsList[cp.UpdateThisAction].Speed
				newAction.Rotate = ap.ActionsList[cp.UpdateThisAction].Rotate
				newAction.RotateSpeed = ap.ActionsList[cp.UpdateThisAction].RotateSpeed
				newAction.Program = ap.ActionsList[cp.UpdateThisAction].Program
				newAction.Strobe = ap.ActionsList[cp.UpdateThisAction].Strobe
				newAction.Map = ap.ActionsList[cp.UpdateThisAction].Map
				newAction.Gobo = ap.ActionsList[cp.UpdateThisAction].Gobo
				newAction.GoboSpeed = ap.ActionsList[cp.UpdateThisAction].GoboSpeed
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[cp.UpdateThisAction].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
				cp.UpdateColors = false
			}

			return len(ap.ActionsList)
		},

		// Function to create list.
		func() fyne.CanvasObject {
			return container.NewVBox(

				container.NewHBox(
					widget.NewLabel("Mode"),
					widget.NewSelect(ap.ActionModeOptions, func(value string) {}),
				),

				container.NewHBox(
					widget.NewLabel("Colors"),
					widget.NewButton("Select", func() {}),

					container.NewHBox(
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
						canvas.NewRectangle(color.White),
					),

					widget.NewRadioGroup(ap.ActionMapOptions, nil),
				),

				container.NewHBox(
					widget.NewLabel("Fade"),
					widget.NewSelect(ap.ActionFadeOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Size"),
					widget.NewSelect(ap.ActionSizeOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Speed"),
					widget.NewSelect(ap.ActionSpeedOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Rotate"),
					widget.NewSelect(ap.ActionRotateOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Rotate Speed"),
					widget.NewSelect(ap.ActionRotateSpeedOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Program"),
					widget.NewSelect(ap.ActionProgramOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Program Speed"),
					widget.NewSelect(ap.ActionProgramSpeedOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Strobe"),
					widget.NewSelect(ap.ActionStrobeOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Gobo"),
					widget.NewSelect(ap.ActionGoboOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Gobo Speed"),
					widget.NewSelect(ap.ActionGoboSpeedOptions, func(value string) {}),
				),
			)
		},

		// Function to update items in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {
			hideAllActionFields(o.(*fyne.Container))

			// Mode
			o.(*fyne.Container).Objects[ACTIONS_MODE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Mode)
			o.(*fyne.Container).Objects[ACTIONS_MODE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {

				if value == "None" || value == "" {
					hideAllActionFields(o.(*fyne.Container))
					newAction := createBlankAction(ap, i)
					newAction.Mode = value
					ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState
				}

				if value == "Off" || value == "" {
					hideAllActionFields(o.(*fyne.Container))
					newAction := createBlankAction(ap, i)
					newAction.Mode = value
					ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState
				}

				if value == "Static" {
					hideAllActionFields(o.(*fyne.Container))
					newAction := createBlankAction(ap, i)
					newAction.Mode = value
					ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR1].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR2].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR3].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR4].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR5].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR6].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR7].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR8].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR9].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR10].(*canvas.Rectangle).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false

				}

				if value == "Chase" {
					hideAllActionFields(o.(*fyne.Container))

					newAction := createCopyOfAction(ap, i)
					newAction.Mode = value
					ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState

					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR1].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR2].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR3].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR4].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR5].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR6].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR7].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR8].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR9].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR10].(*canvas.Rectangle).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[RADIO_BUTTON].(*widget.RadioGroup).Horizontal = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[RADIO_BUTTON].(*widget.RadioGroup).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = !fixtureInfo.HasRotate
					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = !fixtureInfo.HasRotate

					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = !fixtureInfo.HasRotate
					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = !fixtureInfo.HasRotate

					o.(*fyne.Container).Objects[ACTIONS_GOBO].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = !fixtureInfo.HasGobo
					o.(*fyne.Container).Objects[ACTIONS_GOBO].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = !fixtureInfo.HasGobo

					o.(*fyne.Container).Objects[ACTIONS_GOBO_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = !fixtureInfo.HasGobo
					o.(*fyne.Container).Objects[ACTIONS_GOBO_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = !fixtureInfo.HasGobo
				}

				if value == "Control" {
					newAction := createBlankAction(ap, i)
					newAction.Name = ap.ActionsList[i].Name
					newAction.Number = ap.ActionsList[i].Number
					newAction.Colors = []string{}
					newAction.Mode = value
					newAction.Program = ap.ActionsList[i].Program
					newAction.ProgramSpeed = ap.ActionsList[i].ProgramSpeed
					ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState

					// Program
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
					// Program Speed
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
				}
			}

			// Button for Color Selection.
			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).OnTapped = func() {
				cp.ActionNumber = i
				cp.Modal = widget.NewModalPopUp(
					cp.Panel,
					w.Canvas(),
				)
				cp.Modal.Show()
			}

			// Setup the color selection box.
			setColorBoxSizes(o, cp)
			SetRectangleColorsFromString(cp, ap.ActionsList[i].Colors)

			// Map
			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[RADIO_BUTTON].(*widget.RadioGroup).SetSelected(ap.ActionsList[i].Map)
			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[RADIO_BUTTON].(*widget.RadioGroup).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Map = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Fade
			o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Fade)
			o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Fade = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Size
			o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Size)
			o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Size = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Speed
			o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Speed)
			o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Speed = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Rotate
			o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Rotate)
			o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Rotate = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// RotateSpeed
			o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].RotateSpeed)
			o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.RotateSpeed = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Program
			o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Program)
			o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Program = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Program Speed
			o.(*fyne.Container).Objects[ACTIONS_PROGRAM_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].ProgramSpeed)
			o.(*fyne.Container).Objects[ACTIONS_PROGRAM_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.ProgramSpeed = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Strobe
			o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Strobe)
			o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Strobe = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Gobo
			o.(*fyne.Container).Objects[ACTIONS_GOBO].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Gobo)
			o.(*fyne.Container).Objects[ACTIONS_GOBO].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.Gobo = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Gobo
			o.(*fyne.Container).Objects[ACTIONS_GOBO_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].GoboSpeed)
			o.(*fyne.Container).Objects[ACTIONS_GOBO_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := createCopyOfAction(ap, i)
				newAction.GoboSpeed = value
				ap.ActionsList = updateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}
		})

	return &ap
}

func createCopyOfAction(ap ActionPanel, i int) fixture.Action {
	newAction := fixture.Action{}
	newAction.Name = ap.ActionsList[i].Name
	newAction.Number = ap.ActionsList[i].Number
	newAction.Colors = ap.ActionsList[i].Colors
	newAction.Mode = ap.ActionsList[i].Mode
	newAction.Fade = ap.ActionsList[i].Fade
	newAction.Size = ap.ActionsList[i].Size
	newAction.Speed = ap.ActionsList[i].Speed
	newAction.Rotate = ap.ActionsList[i].Rotate
	newAction.RotateSpeed = ap.ActionsList[i].RotateSpeed
	newAction.Program = ap.ActionsList[i].Program
	newAction.ProgramSpeed = ap.ActionsList[i].ProgramSpeed
	newAction.Strobe = ap.ActionsList[i].Strobe
	newAction.Gobo = ap.ActionsList[i].Gobo
	newAction.GoboSpeed = ap.ActionsList[i].GoboSpeed
	newAction.Map = ap.ActionsList[i].Map
	return newAction
}

func createBlankAction(ap ActionPanel, i int) fixture.Action {
	newAction := fixture.Action{}
	newAction.Name = ap.ActionsList[i].Name
	newAction.Number = ap.ActionsList[i].Number
	newAction.Colors = []string{}
	newAction.Mode = "None"
	newAction.Fade = ""
	newAction.Size = ""
	newAction.Speed = ""
	newAction.Rotate = ""
	newAction.RotateSpeed = ""
	newAction.Program = ""
	newAction.ProgramSpeed = ""
	newAction.Strobe = ""
	newAction.Gobo = ""
	newAction.GoboSpeed = ""
	newAction.Map = ""
	return newAction
}

func setColorBoxSizes(o fyne.CanvasObject, cp *ColorPanel) {

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR1].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR1].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR2].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR2].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR3].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR3].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR4].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR4].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR5].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR5].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR6].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR6].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR7].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR7].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR8].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR8].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR9].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR9].(*canvas.Rectangle))

	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR10].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
	cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR10].(*canvas.Rectangle))
}

func hideAllActionFields(o fyne.CanvasObject) {

	// Color Selection.
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).Hidden = true
	// Color Display Boxes.
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR1].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR2].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR3].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR4].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR5].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR6].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR7].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR8].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR9].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[COLOR_SELECTION_BOX].(*fyne.Container).Objects[COLOR10].(*canvas.Rectangle).Hidden = true
	// Map Brightness
	o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[RADIO_BUTTON].(*widget.RadioGroup).Hidden = true
	// Fade
	o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Size
	o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Speed
	o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Rotate
	o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Rotate RotateSpeed
	o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Program
	o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Program Speed
	o.(*fyne.Container).Objects[ACTIONS_PROGRAM_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_PROGRAM_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Strobe
	o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Gobo
	o.(*fyne.Container).Objects[ACTIONS_GOBO].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_GOBO].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
	// Gobo Speed
	o.(*fyne.Container).Objects[ACTIONS_GOBO_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[ACTIONS_GOBO_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

}

// UpdateItem replaces the selected item by id with newItem.
func updateAction(currentStateName string, actions []fixture.Action, id int, newAction fixture.Action) []fixture.Action {
	newActions := []fixture.Action{}
	for _, action := range actions {
		if debug {
			fmt.Printf("updateAction: Name %s Mode %s\n", newAction.Name, newAction.Mode)
		}
		if action.Number == id {
			// update the channel information.
			newAction.Name = currentStateName
			newActions = append(newActions, newAction)
		} else {
			// just add what was there before.
			// Unless it's still set to None.
			if action.Mode != "None" {
				newActions = append(newActions, action)
			}
		}
	}
	return newActions
}

func CreateActionsList(stateList []fixture.State, selectedState int) fixture.Action {

	if debug {
		fmt.Printf("createActionList with Name %s\n", stateList[selectedState].Name)
	}

	newAction := fixture.Action{}
	newAction.Name = stateList[selectedState].Name // Action Name has the same name as the state.
	newAction.Number = 1
	newAction.Size = "Off"
	newAction.Rotate = "Off"
	newAction.RotateSpeed = "Off"
	newAction.Program = "Off"
	newAction.ProgramSpeed = "Off"
	newAction.Strobe = "Off"
	newAction.Map = "Off"
	newAction.Colors = []string{"Off"}
	newAction.Mode = "None"
	newAction.Fade = "Off"
	newAction.Speed = "Off"
	newAction.Gobo = "Default"
	newAction.GoboSpeed = "Slow"

	return newAction
}
