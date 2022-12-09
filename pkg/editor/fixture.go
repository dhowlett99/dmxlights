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

func NewFixturePanel(w fyne.Window, group int, number int, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	groupOptions := []string{"1", "2", "3", "4", "100", "101", "102", "103", "104", "105", "106", "107", "108", "109", "110"}
	// Title.
	title := widget.NewLabel("Fixture List")
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	label := container.NewGridWithColumns(7, widget.NewLabel("Group"), widget.NewLabel("Number"), widget.NewLabel("Name"), widget.NewLabel("Label"), widget.NewLabel("DMX Address"), widget.NewLabel("Description"), widget.NewLabel("Channels"))

	fixtureList := []fixture.Fixture{}

	for _, f := range fixtures.Fixtures {
		newItem := fixture.Fixture{}
		if len(f.UUID) == 0 { // We have a empty UUID for this fixture.
			fmt.Printf("Generating UUID for Fixture: %s\n", f.Name)
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
		fixtureList = append(fixtureList, newItem)
	}

	// Create a new list.
	listPanel := widget.NewList(
		func() int {
			return len(fixtureList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {
			return container.NewGridWithColumns(7,
				widget.NewSelect(groupOptions, func(value string) {}),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewButton("Channels", func() {}),
			)
		},
		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {

			// find the selected group in the options list.
			for _, option := range groupOptions {
				if option == strconv.Itoa(fixtureList[i].Group) {
					o.(*fyne.Container).Objects[0].(*widget.Select).SetSelected(option)

				}
			}

			o.(*fyne.Container).Objects[1].(*widget.Entry).SetText(strconv.Itoa(fixtureList[i].Number))
			o.(*fyne.Container).Objects[2].(*widget.Entry).SetText(fixtureList[i].Name)
			o.(*fyne.Container).Objects[3].(*widget.Entry).SetText(fixtureList[i].Label)
			o.(*fyne.Container).Objects[4].(*widget.Entry).SetText(strconv.Itoa(int(fixtureList[i].Address)))
			o.(*fyne.Container).Objects[4].(*widget.Entry).OnChanged = func(value string) {
				newSetting := fixture.Fixture{}
				newSetting.UUID = fixtureList[i].UUID
				newSetting.Label = fixtureList[i].Label
				newSetting.Name = fixtureList[i].Name
				newSetting.Number = fixtureList[i].Number
				newSetting.Group = fixtureList[i].Group
				newSetting.Description = fixtureList[i].Description
				newSetting.Type = fixtureList[i].Type
				newSetting.Channels = fixtureList[i].Channels
				newSetting.States = fixtureList[i].States
				newSetting.NumberChannels = fixtureList[i].NumberChannels
				newSetting.UseFixture = fixtureList[i].UseFixture
				i, _ := strconv.Atoi(value)
				newSetting.Address = int16(i)
				fixtureList = UpdateListItem(fixtureList, newSetting.UUID, newSetting)
				// channelList[st.CurrentChannel-1].Settings = st.fixtureList
				// st.Update = true
				// st.UpdatedChannelList = channelList
			}
			o.(*fyne.Container).Objects[5].(*widget.Entry).SetText(fixtureList[i].Description)
			o.(*fyne.Container).Objects[6].(*widget.Button).OnTapped = func() {
				modal, err := NewChannelEditor(w, fixtureList[i].UUID, fixtureList[i].Group, fixtureList[i].Number, fixtures)
				if err != nil {
					fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", fixtureList[i].Group, fixtureList[i].Number, err)
					return
				}
				modal.Resize(fyne.NewSize(800, 600))
				modal.Show()
			}
		},
	)

	// Save button.
	buttonSave := widget.NewButton("Save", func() {

		// for fixtureNumber := range fixtureList {
		// 	fixtures.Fixtures[fixtureNumber] = fixtureList[fixtureNumber]
		// }

		// Insert updated fixture into fixtures.
		copy(fixtures.Fixtures, fixtureList)

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
	scrollableList := container.NewVScroll(listPanel)
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
