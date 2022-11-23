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
	"strings"

	"fyne.io/fyne/v2"
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
	ActionSpeedOptions   []string
	ActionRotateOptions  []string
	ActionMusicOptions   []string
	ActionProgramOptions []string
	ActionStrobeOptions  []string
}

func NewActionsPanel(actionsAvailable bool, actionsList []fixture.Action) *ActionPanel {

	ap := ActionPanel{}
	ap.ActionsList = actionsList
	ap.ActionNameOptions = []string{"Off", "On", "Red", "Green", "Blue", "SoftChase", "SharpChase", "SoundChase", "Rotate"}
	ap.ActionModeOptions = []string{"Off", "Chase", "Static"}
	ap.ActionColorsOptions = []string{"None", "Red", "Green", "Blue"}
	ap.ActionFadeOptions = []string{"Off", "Soft", "Sharp"}
	ap.ActionSpeedOptions = []string{"Off", "Slow", "Medium", "Fast", "VeryFast", "Music"}
	ap.ActionRotateOptions = []string{"Off", "Forward", "Reverse", "Auto"}
	ap.ActionMusicOptions = []string{"Off", "On"}
	ap.ActionProgramOptions = []string{"Off", "Slow", "Medium", "Fast"}
	ap.ActionStrobeOptions = []string{"Off", "Slow", "Medium", "Fast"}

	// Actions Selection Panel.
	if actionsAvailable {
		ap.ActionsPanel = widget.NewList(
			func() int {
				return len(ap.ActionsList)
			},
			func() fyne.CanvasObject {
				return container.NewVBox(
					container.NewHBox(
						widget.NewLabel("Name"),
						widget.NewSelect(ap.ActionNameOptions, func(value string) {}),
					),
					container.NewHBox(
						widget.NewLabel("Colors"),
						widget.NewSelect(ap.ActionColorsOptions, func(value string) {}),
					),
					container.NewHBox(
						widget.NewLabel("Mode"),
						widget.NewSelect(ap.ActionModeOptions, func(value string) {}),
					),
					container.NewHBox(
						widget.NewLabel("Fade"),
						widget.NewSelect(ap.ActionFadeOptions, func(value string) {}),
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
						widget.NewLabel("Music"),
						widget.NewSelect(ap.ActionMusicOptions, func(value string) {}),
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
				o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Name)
				o.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(strings.Join(ap.ActionsList[i].Colors[:], ","))
				o.(*fyne.Container).Objects[2].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Mode)
				o.(*fyne.Container).Objects[3].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Fade)
				o.(*fyne.Container).Objects[4].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Speed)
				o.(*fyne.Container).Objects[5].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Rotate)
				o.(*fyne.Container).Objects[6].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Music)
				o.(*fyne.Container).Objects[7].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Program)
				o.(*fyne.Container).Objects[8].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Strobe)
			})
	}
	return &ap
}

func (ap *ActionPanel) UpdateActionList(actionList []fixture.Action) {
	ap.ActionsList = actionList
}
