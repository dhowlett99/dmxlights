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

type SettingsPanel struct {
	SettingsPanel   *widget.List
	SettingsesList  []fixture.Setting
	SettingsOptions []string
}

func NewSettingsPanel(SettingsesAvailable bool, SettingsesList []fixture.Setting, ap *ActionPanel) *SettingsPanel {

	st := SettingsPanel{}
	st.SettingsesList = SettingsesList
	st.SettingsOptions = []string{"Off", "On", "Red", "Green", "Blue", "SoftChase", "SharpChase", "SoundChase", "Rotate"}

	// Settingses Selection Panel.
	if SettingsesAvailable {
		st.SettingsPanel = widget.NewList(
			func() int {
				return len(st.SettingsesList)
			},
			// Function to create item.
			func() fyne.CanvasObject {
				return container.NewHBox(
					widget.NewLabel("template"),
					widget.NewEntry(),
					widget.NewButton("Select", nil),
				)
			},

			// Function to update item in this list.
			func(i widget.ListItemID, o fyne.CanvasObject) {
				//fmt.Printf("Settings ID is %d   Settings Setting is %s\n", i, SettingsOptions[i])
				o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", st.SettingsesList[i].Number))

				// find the selected option in the options list.
				o.(*fyne.Container).Objects[1].(*widget.Entry).Text = st.SettingsesList[i].Name

				// new part
				o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
					fmt.Printf("I am button %d actions %+v\n", st.SettingsesList[i].Number, ap.ActionsList)
				}
			},
		)
	}
	return &st
}
