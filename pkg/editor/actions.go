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
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type ActionPanel struct {
	ActionsPanel         *widget.List
	ActionsList          []fixture.Action
	ActionNameOptions    []string
	ActionColorsOptions  []string
	ActionModeOptions    []string
	ActionFadeOptions    []string
	ActionSizeOptions    []string
	ActionSpeedOptions   []string
	ActionRotateOptions  []string
	ActionRotateSpeed    []string
	ActionMusicOptions   []string
	ActionProgramOptions []string
	ActionStrobeOptions  []string
	UpdateActions        bool
	UpdateThisAction     int
	CurrentState         int
	CurrentStateName     string
}

const LABEL = 0
const SELECT = 1

const ACTIONS_MODE = 0
const ACTIONS_COLORS = 1
const ACTIONS_FADE = 2
const ACTIONS_SIZE = 3
const ACTIONS_SPEED = 4
const ACTIONS_ROTATE = 5
const ACTIONS_ROTATESPEED = 6
const ACTIONS_PROGRAM = 7
const ACTIONS_STROBE = 8

func NewActionsPanel(w fyne.Window, actionsList []fixture.Action) *ActionPanel {

	ap := ActionPanel{}
	ap.ActionsList = actionsList
	ap.ActionModeOptions = []string{"Off", "Static", "Chase", "Control"}
	ap.ActionSizeOptions = []string{"Off", "Short", "Medium", "Long"}
	ap.ActionFadeOptions = []string{"Off", "Soft", "Sharp"}
	ap.ActionSpeedOptions = []string{"Off", "Slow", "Medium", "Fast", "VeryFast", "Music"}
	ap.ActionRotateOptions = []string{"Off", "Clockwise", "Anti Clockwise", "Auto"}
	ap.ActionRotateSpeed = []string{"Slow", "Medium", "Fast"}
	ap.ActionMusicOptions = []string{"Off", "On"}
	ap.ActionStrobeOptions = []string{"Off", "Slow", "Medium", "Fast"}

	cp := NewColorPickerPanel(w)

	// Actions Selection Panel.
	ap.ActionsPanel = widget.NewList(
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
				newAction.Program = ap.ActionsList[cp.UpdateThisAction].Program
				newAction.Strobe = ap.ActionsList[cp.UpdateThisAction].Strobe
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[cp.UpdateThisAction].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
				cp.UpdateColors = false
			}

			return len(ap.ActionsList)
		},
		func() fyne.CanvasObject {
			return container.NewVBox(

				container.NewHBox(
					widget.NewLabel("Mode"),
					widget.NewSelect(ap.ActionModeOptions, func(value string) {}),
				),

				container.NewHBox(
					widget.NewLabel("Colors"),
					widget.NewButton("Select", func() {}),
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
					widget.NewSelect(ap.ActionRotateSpeed, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Program"),
					widget.NewSelect(ap.ActionProgramOptions, func(value string) {}),
				),
				container.NewHBox(
					widget.NewLabel("Strobe"),
					widget.NewSelect(ap.ActionStrobeOptions, func(value string) {}),
				),
			)
		},

		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// Mode
			o.(*fyne.Container).Objects[ACTIONS_MODE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Mode)
			o.(*fyne.Container).Objects[ACTIONS_MODE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := fixture.Action{}
				if value == "Off" || value == "" {
					newAction.Name = ap.ActionsList[i].Name
					newAction.Number = ap.ActionsList[i].Number
					newAction.Colors = []string{}
					newAction.Mode = value
					newAction.Fade = ""
					newAction.Size = ""
					newAction.Speed = ""
					newAction.Rotate = ""
					newAction.RotateSpeed = ""
					newAction.Program = ""
					newAction.Strobe = ""
					ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState

					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[2].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[3].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[4].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[5].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[6].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[7].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[8].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[9].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[10].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[11].(*canvas.Rectangle).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
				}

				if value == "Static" {
					newAction.Name = ap.ActionsList[i].Name
					newAction.Number = ap.ActionsList[i].Number
					newAction.Colors = ap.ActionsList[i].Colors
					newAction.Mode = value
					newAction.Fade = ""
					newAction.Size = ""
					newAction.Speed = ""
					newAction.Rotate = ""
					newAction.RotateSpeed = ""
					newAction.Program = ""
					newAction.Strobe = ""
					ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[2].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[3].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[4].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[5].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[6].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[7].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[8].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[9].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[10].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[11].(*canvas.Rectangle).Hidden = false

					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true

					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
				}

				if value == "Chase" {
					newAction.Name = ap.ActionsList[i].Name
					newAction.Number = ap.ActionsList[i].Number
					newAction.Colors = ap.ActionsList[i].Colors
					newAction.Mode = value
					newAction.Fade = ap.ActionsList[i].Fade
					newAction.Size = ap.ActionsList[i].Size
					newAction.Speed = ap.ActionsList[i].Speed
					newAction.Rotate = ap.ActionsList[i].Rotate
					newAction.RotateSpeed = ap.ActionsList[i].RotateSpeed
					newAction.Program = ap.ActionsList[i].Program
					newAction.Strobe = ap.ActionsList[i].Strobe
					ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState
					// Color Selection.
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).Hidden = false
					// Color Display Boxes.
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[2].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[3].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[4].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[5].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[6].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[7].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[8].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[9].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[10].(*canvas.Rectangle).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[11].(*canvas.Rectangle).Hidden = false
					// Fade
					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
					// Size
					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
					// Speed
					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
					// Rotate
					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
					// Rotate RotateSpeed
					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
					// Program
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
					// Strobe
					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
				}

				if value == "Control" {
					newAction.Name = ap.ActionsList[i].Name
					newAction.Number = ap.ActionsList[i].Number
					newAction.Colors = []string{}
					newAction.Mode = value
					newAction.Fade = ""
					newAction.Size = ""
					newAction.Speed = ""
					newAction.Rotate = ""
					newAction.RotateSpeed = ""
					newAction.Program = ap.ActionsList[i].Program
					newAction.Strobe = ""
					ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
					ap.UpdateActions = true
					ap.UpdateThisAction = ap.CurrentState

					// Color Selection
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[SELECT].(*widget.Button).Hidden = true
					// Color Display Boxes
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[2].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[3].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[4].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[5].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[6].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[7].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[8].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[9].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[10].(*canvas.Rectangle).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[11].(*canvas.Rectangle).Hidden = true
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
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = false
					o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = false
					// Strobe
					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[LABEL].(*widget.Label).Hidden = true
					o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).Hidden = true
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

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[2].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[2].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[3].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[3].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[4].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[4].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[5].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[5].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[6].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[6].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[7].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[7].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[8].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[8].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[9].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[9].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[10].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[10].(*canvas.Rectangle))

			o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[11].(*canvas.Rectangle).SetMinSize(fyne.Size{Height: 5, Width: 8})
			cp.Rectanges = append(cp.Rectanges, o.(*fyne.Container).Objects[ACTIONS_COLORS].(*fyne.Container).Objects[11].(*canvas.Rectangle))

			SetRectangleColorsFromString(cp, ap.ActionsList[i].Colors)

			// Fade
			o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Fade)
			o.(*fyne.Container).Objects[ACTIONS_FADE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := fixture.Action{}
				newAction.Name = ap.ActionsList[i].Name
				newAction.Number = ap.ActionsList[i].Number
				newAction.Colors = ap.ActionsList[i].Colors
				newAction.Mode = ap.ActionsList[i].Mode
				newAction.Fade = value
				newAction.Size = ap.ActionsList[i].Size
				newAction.Speed = ap.ActionsList[i].Speed
				newAction.Rotate = ap.ActionsList[i].Rotate
				newAction.RotateSpeed = ap.ActionsList[i].RotateSpeed
				newAction.Program = ap.ActionsList[i].Program
				newAction.Strobe = ap.ActionsList[i].Strobe
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Size
			o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Size)
			o.(*fyne.Container).Objects[ACTIONS_SIZE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := fixture.Action{}
				newAction.Name = ap.ActionsList[i].Name
				newAction.Number = ap.ActionsList[i].Number
				newAction.Colors = ap.ActionsList[i].Colors
				newAction.Mode = ap.ActionsList[i].Mode
				newAction.Fade = ap.ActionsList[i].Fade
				newAction.Size = value
				newAction.Speed = ap.ActionsList[i].Speed
				newAction.Rotate = ap.ActionsList[i].Rotate
				newAction.RotateSpeed = ap.ActionsList[i].RotateSpeed
				newAction.Program = ap.ActionsList[i].Program
				newAction.Strobe = ap.ActionsList[i].Strobe
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Speed)
			o.(*fyne.Container).Objects[ACTIONS_SPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := fixture.Action{}
				newAction.Name = ap.ActionsList[i].Name
				newAction.Number = ap.ActionsList[i].Number
				newAction.Colors = ap.ActionsList[i].Colors
				newAction.Mode = ap.ActionsList[i].Mode
				newAction.Fade = ap.ActionsList[i].Fade
				newAction.Size = ap.ActionsList[i].Size
				newAction.Speed = value
				newAction.Rotate = ap.ActionsList[i].Rotate
				newAction.RotateSpeed = ap.ActionsList[i].RotateSpeed
				newAction.Program = ap.ActionsList[i].Program
				newAction.Strobe = ap.ActionsList[i].Strobe
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}
			// Rotate
			o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Rotate)
			o.(*fyne.Container).Objects[ACTIONS_ROTATE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := fixture.Action{}
				newAction.Name = ap.ActionsList[i].Name
				newAction.Number = ap.ActionsList[i].Number
				newAction.Colors = ap.ActionsList[i].Colors
				newAction.Mode = ap.ActionsList[i].Mode
				newAction.Fade = ap.ActionsList[i].Fade
				newAction.Size = ap.ActionsList[i].Size
				newAction.Speed = ap.ActionsList[i].Speed
				newAction.Rotate = value
				newAction.RotateSpeed = ap.ActionsList[i].RotateSpeed
				newAction.Program = ap.ActionsList[i].Program
				newAction.Strobe = ap.ActionsList[i].Strobe
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// RotateSpeed
			o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].RotateSpeed)
			o.(*fyne.Container).Objects[ACTIONS_ROTATESPEED].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
				newAction := fixture.Action{}
				newAction.Name = ap.ActionsList[i].Name
				newAction.Number = ap.ActionsList[i].Number
				newAction.Colors = ap.ActionsList[i].Colors
				newAction.Mode = ap.ActionsList[i].Mode
				newAction.Fade = ap.ActionsList[i].Fade
				newAction.Size = ap.ActionsList[i].Size
				newAction.Speed = ap.ActionsList[i].Speed
				newAction.Rotate = ap.ActionsList[i].Rotate
				newAction.RotateSpeed = value
				newAction.Program = ap.ActionsList[i].Program
				newAction.Strobe = ap.ActionsList[i].Strobe
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			// Program
			o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Program)
			o.(*fyne.Container).Objects[ACTIONS_PROGRAM].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
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
				newAction.Program = value
				newAction.Strobe = ap.ActionsList[i].Strobe
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

			o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).SetSelected(ap.ActionsList[i].Strobe)
			o.(*fyne.Container).Objects[ACTIONS_STROBE].(*fyne.Container).Objects[SELECT].(*widget.Select).OnChanged = func(value string) {
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
				newAction.Program = ap.ActionsList[i].Strobe
				newAction.Strobe = value
				ap.ActionsList = UpdateAction(ap.CurrentStateName, ap.ActionsList, ap.ActionsList[i].Number, newAction)
				ap.UpdateActions = true
				ap.UpdateThisAction = ap.CurrentState
			}

		})
	return &ap
}

// UpdateItem replaces the selected item by id with newItem.
func UpdateAction(currentStateName string, actions []fixture.Action, id int, newAction fixture.Action) []fixture.Action {
	newActions := []fixture.Action{}
	for _, action := range actions {
		if action.Number == id {
			// update the channel information.
			newAction.Name = currentStateName
			newActions = append(newActions, newAction)
		} else {
			// just add what was there before.
			newActions = append(newActions, action)
		}
	}
	return newActions
}

func CreateActionsList(stateList []fixture.State, selectedState int) (actionsList []fixture.Action) {

	newAction := fixture.Action{}
	newAction.Name = stateList[selectedState].Name // Action Name has the same name as the state.
	newAction.Number = 1
	newAction.Size = "Off"
	newAction.Rotate = "Off"
	newAction.RotateSpeed = "Off"
	newAction.Program = "Off"
	newAction.Strobe = "Off"
	newAction.Colors = []string{"Off"}
	newAction.Mode = "Off"
	newAction.Fade = "Off"
	newAction.Speed = "Off"

	actionsList = append(actionsList, newAction)
	return actionsList
}
