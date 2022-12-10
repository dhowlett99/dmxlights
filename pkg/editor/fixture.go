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
	"os"
	"strconv"

	"github.com/google/uuid"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type FixturesPanel struct {
	FixturePanel      *widget.List
	FixtureList       []fixture.Fixture
	UpdateFixture     bool
	UpdateChannelList []fixture.Channel
	UpdateThisFixture int
}

func NewFixturePanel(w fyne.Window, group int, number int, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	fp := FixturesPanel{}
	fp.FixtureList = []fixture.Fixture{}

	groupOptions := []string{"1", "2", "3", "4", "100", "101", "102", "103", "104", "105", "106", "107", "108", "109", "110"}
	// Title.
	title := widget.NewLabel("Fixture List")
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	label := container.NewGridWithColumns(7, widget.NewLabel("Group"), widget.NewLabel("Number"), widget.NewLabel("Name"), widget.NewLabel("Label"), widget.NewLabel("DMX Address"), widget.NewLabel("Description"), widget.NewLabel("Channels"))

	for _, f := range fixtures.Fixtures {
		newItem := fixture.Fixture{}
		if len(f.UUID) == 0 { // We have a empty UUID for this fixture.
			if debug {
				fmt.Printf("Generating UUID for Fixture: %s\n", f.Name)
			}
			newItem.UUID = uuid.New().String()[:7]
		} else {
			newItem.UUID = f.UUID
		}
		newItem.Name = f.Name
		newItem.Label = f.Label
		newItem.Group = f.Group
		newItem.Number = f.Number
		newItem.Address = f.Address
		newItem.Description = f.Description
		newItem.Type = f.Type
		newItem.Channels = f.Channels
		newItem.States = f.States
		newItem.NumberChannels = f.NumberChannels
		newItem.UseFixture = f.UseFixture
		fp.FixtureList = append(fp.FixtureList, newItem)
	}

	// Create a new list.
	fp.FixturePanel = widget.NewList(
		func() int {
			if fp.UpdateFixture {
				fp.FixtureList[fp.UpdateThisFixture].Channels = fp.UpdateChannelList
				fp.UpdateFixture = false
			}
			return len(fp.FixtureList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {
			return container.NewGridWithColumns(8,
				widget.NewEntry(), // UUID.
				widget.NewSelect(groupOptions, func(value string) {}), // Group Number.
				widget.NewEntry(),                       // Number.
				widget.NewEntry(),                       // Name.
				widget.NewEntry(),                       // Label.
				widget.NewEntry(),                       // DMX Address.
				widget.NewEntry(),                       // Description
				widget.NewButton("Channels", func() {}), // Channel Button
			)
		},
		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// Show the UUID.
			o.(*fyne.Container).Objects[0].(*widget.Entry).SetText(fp.FixtureList[i].UUID)
			o.(*fyne.Container).Objects[0].(*widget.Entry).Hidden = true

			// find the selected group in the options list.
			for _, option := range groupOptions {
				if option == strconv.Itoa(fp.FixtureList[i].Group) {
					o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
				}
			}
			o.(*fyne.Container).Objects[1].(*widget.Select).OnChanged = func(value string) {
				// if value isn't what we expect it to be ignore.
				if o.(*fyne.Container).Objects[0].(*widget.Entry).Text == fp.FixtureList[i].UUID {
					//o.(*fyne.Container).Objects[1].(*widget.Select).FocusGained()
					newSetting := fixture.Fixture{}
					newSetting.UUID = fp.FixtureList[i].UUID
					newSetting.Label = fp.FixtureList[i].Label
					newSetting.Name = fp.FixtureList[i].Name
					newSetting.Number = fp.FixtureList[i].Number
					v, _ := strconv.Atoi(value)
					newSetting.Group = v
					newSetting.Description = fp.FixtureList[i].Description
					newSetting.Type = fp.FixtureList[i].Type
					newSetting.Channels = fp.FixtureList[i].Channels
					newSetting.States = fp.FixtureList[i].States
					newSetting.NumberChannels = fp.FixtureList[i].NumberChannels
					newSetting.UseFixture = fp.FixtureList[i].UseFixture
					newSetting.Address = fp.FixtureList[i].Address
					fp.FixtureList = UpdateListItem(fp.FixtureList, fp.FixtureList[i].UUID, newSetting)
				}
			}

			// Show the Fixture Number.
			o.(*fyne.Container).Objects[2].(*widget.Entry).SetText(strconv.Itoa(fp.FixtureList[i].Number))

			// Show the Fixture Name.
			o.(*fyne.Container).Objects[3].(*widget.Entry).SetText(fp.FixtureList[i].Name)

			// Show the Fixture Label.
			o.(*fyne.Container).Objects[4].(*widget.Entry).SetText(fp.FixtureList[i].Label)

			// Show and Edit the Fixture DMX label.
			o.(*fyne.Container).Objects[5].(*widget.Entry).SetText(strconv.Itoa(int(fp.FixtureList[i].Address)))
			o.(*fyne.Container).Objects[5].(*widget.Entry).OnChanged = func(value string) {
				// if value isn't what we expect it to be ignore.
				if o.(*fyne.Container).Objects[0].(*widget.Entry).Text == fp.FixtureList[i].UUID {
					o.(*fyne.Container).Objects[5].(*widget.Entry).FocusGained()
					newSetting := fixture.Fixture{}
					newSetting.UUID = fp.FixtureList[i].UUID
					newSetting.Label = fp.FixtureList[i].Label
					newSetting.Name = fp.FixtureList[i].Name
					newSetting.Number = fp.FixtureList[i].Number
					newSetting.Group = fp.FixtureList[i].Group
					newSetting.Description = fp.FixtureList[i].Description
					newSetting.Type = fp.FixtureList[i].Type
					newSetting.Channels = fp.FixtureList[i].Channels
					newSetting.States = fp.FixtureList[i].States
					newSetting.NumberChannels = fp.FixtureList[i].NumberChannels
					newSetting.UseFixture = fp.FixtureList[i].UseFixture
					v, _ := strconv.Atoi(value)
					newSetting.Address = int16(v)
					fp.FixtureList = UpdateListItem(fp.FixtureList, fp.FixtureList[i].UUID, newSetting)
				}
			}

			// Show Fixture Description.
			o.(*fyne.Container).Objects[6].(*widget.Entry).SetText(fp.FixtureList[i].Description)

			// Show and Edit Channel Definitions using the Channel Editor.
			o.(*fyne.Container).Objects[7].(*widget.Button).OnTapped = func() {
				modal, err := NewChannelEditor(w, fp.FixtureList[i].UUID, fp.FixtureList[i].Channels, &fp, fixtures)
				if err != nil {
					fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", fp.FixtureList[i].Group, fp.FixtureList[i].Number, err)
					return
				}
				modal.Resize(fyne.NewSize(800, 600))
				modal.Show()
			}
		},
	)

	// Save button.
	buttonSave := widget.NewButton("Save", func() {

		// Insert updated fixture into fixtures.
		copy(fixtures.Fixtures, fp.FixtureList)

		// Save the new fixtures file.
		err := fixture.SaveFixtures("fixtures.yaml", fixtures)
		if err != nil {
			fmt.Printf("error saving fixtures %s\n", err.Error())
			os.Exit(1)
		}

		modal.Hide()
	})
	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		modal.Hide()
	})
	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)
	scrollableList := container.NewVScroll(fp.FixturePanel)
	scrollableList.SetMinSize(fyne.Size{Height: 430, Width: 600})

	content := fyne.Container{}
	main := container.NewBorder(title, nil, nil, nil, label)
	two := container.NewBorder(main, nil, nil, nil, scrollableList)
	content = *container.NewBorder(two, nil, nil, nil, saveCancel)

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		&content,
		w.Canvas(),
	)
	return modal, nil
}

func UpdateListItem(items []fixture.Fixture, id string, newItem fixture.Fixture) []fixture.Fixture {
	newItems := []fixture.Fixture{}
	for _, item := range items {
		if item.UUID == id {
			// update the settings information.
			newItems = append(newItems, newItem)
		} else {
			// just add what was there before.
			newItems = append(newItems, item)
		}
	}
	return newItems
}
