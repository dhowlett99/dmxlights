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
	"sort"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type FixturesPanel struct {
	FixturePanel      *widget.Table
	FixtureList       []fixture.Fixture
	UpdateThisFixture int

	UpdateChannels      bool
	UpdateStates        bool
	UpdatedChannelsList []fixture.Channel
	UpdatedStatesList   []fixture.State

	UpdateUseFixture bool
	UseFixture       string

	UpdateDescription bool
	Description       string

	GroupOptions  []string
	NumberOptions []string
	TypeOptions   []string
}

const debug bool = false

const FIXTURE_ID int = 0
const FIXTURE_TYPE int = 1
const FIXTURE_GROUP int = 2
const FIXTURE_NUMBER int = 3
const FIXTURE_NAME int = 4
const FIXTURE_LABEL int = 5
const FIXTURE_ADDRESS int = 6
const FIXTURE_DESCRIPTION int = 7
const FIXTURE_DELETE int = 8
const FIXTURE_ADD int = 9
const FIXTURE_CHANNELS int = 10

// Convert the list of fixtures to an array of strings containing and array of strings with
// the values from each fixture.
// This is done once when the fixture panel is loaded and the fixture info comes from the fixtures.yaml
func makeArray(fixtures *fixture.Fixtures) [][]string {

	if debug {
		fmt.Printf("makeArray\n")
	}
	var data = [][]string{}

	// scan the fixtures structure for the selected fixture.
	for _, fixture := range fixtures.Fixtures {
		newFixture := []string{}
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.ID))
		newFixture = append(newFixture, fixture.Type)
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.Group))
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.Number))
		newFixture = append(newFixture, fixture.Name)
		newFixture = append(newFixture, fixture.Label)
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.Address))
		newFixture = append(newFixture, fixture.Description)
		newFixture = append(newFixture, "-")
		newFixture = append(newFixture, "+")
		newFixture = append(newFixture, "Channels")

		data = append(data, newFixture)
	}

	return data
}

// updateArray is called by either add or delete fixture, and takes the fixture data and
// updates the data array used by the panel's table.
func updateArray(fixtures []fixture.Fixture) [][]string {

	if debug {
		fmt.Printf("updateArray\n")
	}

	var data = [][]string{}

	// scan the fixtures structure for the selected fixture.
	for _, fixture := range fixtures {
		newFixture := []string{}
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.ID))
		newFixture = append(newFixture, fixture.Type)
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.Group))
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.Number))
		newFixture = append(newFixture, fixture.Name)
		newFixture = append(newFixture, fixture.Label)
		newFixture = append(newFixture, fmt.Sprintf("%d", fixture.Address))
		newFixture = append(newFixture, fixture.Description)
		newFixture = append(newFixture, "-")
		newFixture = append(newFixture, "+")
		newFixture = append(newFixture, "Channels")

		data = append(data, newFixture)
	}

	return data
}

func NewFixturePanel(sequences []*common.Sequence, w fyne.Window, group int, number int, fixtures *fixture.Fixtures, commandChannels []chan common.Command) (modal *widget.PopUp, err error) {

	if debug {
		fmt.Printf("NewFixturePanel\n")
	}

	fp := FixturesPanel{}
	fp.FixtureList = []fixture.Fixture{}

	fp.GroupOptions = []string{"1", "2", "3", "4", "100", "101", "102", "103", "104", "105", "106", "107", "108", "109", "110"}
	fp.NumberOptions = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	fp.TypeOptions = []string{"rgb", "scanner", "switch", "projector"}

	// Load the fixtures into the array used by the table.
	data := makeArray(fixtures)

	// Title.
	title := widget.NewLabel("Fixture List")
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	var headerText = [][]string{{"ID", "Type", "Group", "No", "Name", "Label", "DMX", "Description", "-", "+", "Channels"},
		{"ID", "Type", "Group", "No", "Name", "Label", "DMX", "Description", "-", "+", "Channels"}}

	header := widget.NewTable(

		func() (int, int) {
			return len(headerText), len(headerText[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("headerText items")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(headerText[i.Row][i.Col])
		})

	header.SetColumnWidth(0, 40)  // Id
	header.SetColumnWidth(1, 100) // Type
	header.SetColumnWidth(2, 60)  // Sequence Number
	header.SetColumnWidth(3, 60)  // Fixture Number
	header.SetColumnWidth(4, 100) // Name
	header.SetColumnWidth(5, 100) // Label
	header.SetColumnWidth(6, 50)  // DMX Address
	header.SetColumnWidth(7, 150) // Description
	header.SetColumnWidth(8, 20)  // Delete Button
	header.SetColumnWidth(9, 20)  // Add Button
	header.SetColumnWidth(10, 40) // Channels Button

	// Geneate the fixture list.
	for no, f := range fixtures.Fixtures {
		newItem := fixture.Fixture{}

		if f.ID == 0 { // We have a empty ID for this fixture.
			newItem.ID = no + 1
		} else {
			newItem.ID = f.ID
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
	fp.FixturePanel = widget.NewTable(
		func() (int, int) {
			if fp.UpdateChannels {
				fp.FixtureList[fp.UpdateThisFixture].Channels = fp.UpdatedChannelsList
				fp.UpdateChannels = false
			}
			if fp.UpdateStates {
				fp.FixtureList[fp.UpdateThisFixture].States = fp.UpdatedStatesList
				fp.UpdateChannels = false
			}
			if fp.UpdateUseFixture {
				address := fixture.FindFixtureAddressByName(fp.UseFixture, fixtures)
				data[fp.UpdateThisFixture][FIXTURE_ADDRESS] = address
				dmx, _ := strconv.Atoi(address)
				fp.FixtureList[fp.UpdateThisFixture].Address = int16(dmx)
				fp.UpdateUseFixture = false
			}
			if fp.UpdateDescription {
				fp.FixtureList[fp.UpdateThisFixture].Description = fp.Description
				fp.UpdateDescription = false
			}

			return len(data), len(data[0])
		},

		// Create Table
		func() (o fyne.CanvasObject) {
			return container.NewMax(
				widget.NewLabel("id"), // ID.
				widget.NewSelect(fp.TypeOptions, func(value string) {}),   // Type rgb, scanner or switch.
				widget.NewSelect(fp.GroupOptions, func(value string) {}),  // Group Number.
				widget.NewSelect(fp.NumberOptions, func(value string) {}), // Fixture Number.
				widget.NewEntry(),                       // Name.
				widget.NewEntry(),                       // Label.
				widget.NewEntry(),                       // DMX Address.
				widget.NewEntry(),                       // Description
				widget.NewButton("-", func() {}),        // Fixture delete button.
				widget.NewButton("+", func() {}),        // Fixture add button
				widget.NewButton("Channels", func() {}), // Channel Button
			)
		},
		// Function to update item in this list.
		func(i widget.TableCellID, o fyne.CanvasObject) {

			// Hide all field types.
			hideAllFields(o)

			// Fixture ID.
			if i.Col == FIXTURE_ID {
				showField(FIXTURE_ID, o)
				o.(*fyne.Container).Objects[FIXTURE_ID].(*widget.Label).SetText(data[i.Row][i.Col])
			}

			// Fixture Type.
			if i.Col == FIXTURE_TYPE {
				showField(FIXTURE_TYPE, o)
				o.(*fyne.Container).Objects[FIXTURE_TYPE].(*widget.Select).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_TYPE].(*widget.Select).SetSelected(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_TYPE].(*widget.Select).OnChanged = func(value string) {
					if data[i.Row][FIXTURE_ID] == fmt.Sprintf("%d", fp.FixtureList[i.Row].ID) {
						newFixture := makeNewFixture(data, i, FIXTURE_TYPE, value, fp.FixtureList)

						if newFixture.Type == "switch" && newFixture.States == nil {
							newFixture.Channels = []fixture.Channel{}
							// Create some default states
							newFixture.States = []fixture.State{}
							newState := fixture.State{
								Number: 1,
								Name:   "New",
							}
							newFixture.States = append(newFixture.States, newState)
						}

						fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
					}
				}
			}

			// Fixture Group.
			if i.Col == FIXTURE_GROUP {
				showField(FIXTURE_GROUP, o)
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).SetSelected(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).Options = fp.GroupOptions
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).OnChanged = func(value string) {
					newFixture := makeNewFixture(data, i, FIXTURE_GROUP, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
				}
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).PlaceHolder = "XXX"
			}

			// Fixture Number.
			if i.Col == FIXTURE_NUMBER {
				showField(FIXTURE_NUMBER, o)
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).SetSelected(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).OnChanged = func(value string) {
					newFixture := makeNewFixture(data, i, FIXTURE_NUMBER, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
				}
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).PlaceHolder = "X"
			}

			// Fixture Name.
			if i.Col == FIXTURE_NAME {
				showField(FIXTURE_NAME, o)
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*widget.Entry).OnChanged = func(value string) {
					o.(*fyne.Container).Objects[FIXTURE_NAME].(*widget.Entry).FocusGained()
					newFixture := makeNewFixture(data, i, FIXTURE_NAME, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
				}
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*widget.Entry).PlaceHolder = "XXXXXXXXXXXXX"
			}

			// Fixture Label.
			if i.Col == FIXTURE_LABEL {
				showField(FIXTURE_LABEL, o)
				o.(*fyne.Container).Objects[FIXTURE_LABEL].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_LABEL].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_LABEL].(*widget.Entry).OnChanged = func(value string) {
					o.(*fyne.Container).Objects[FIXTURE_LABEL].(*widget.Entry).FocusGained()
					newFixture := makeNewFixture(data, i, FIXTURE_LABEL, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
				}
			}

			// Fixture DMX Address.
			if i.Col == FIXTURE_ADDRESS {
				showField(FIXTURE_ADDRESS, o)
				o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*widget.Entry).OnChanged = func(value string) {
					o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*widget.Entry).FocusGained()
					newFixture := makeNewFixture(data, i, FIXTURE_ADDRESS, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
				}
				// Switch addresses are the address of the fixture being used.
				// So this comes from the state panel's usefixture field.
				if data[i.Row][FIXTURE_TYPE] == "switch" {
					o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*widget.Entry).Disable()
				}
			}

			// Fixture Description.
			if i.Col == FIXTURE_DESCRIPTION {
				showField(FIXTURE_DESCRIPTION, o)
				o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*widget.Entry).OnChanged = func(value string) {
					o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*widget.Entry).FocusGained()
					newFixture := makeNewFixture(data, i, FIXTURE_DESCRIPTION, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
				}
			}

			// Fixture Delete Button.
			if i.Col == FIXTURE_DELETE {
				showField(FIXTURE_DELETE, o)
				o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).OnTapped = func() {
					fp.FixtureList = DeleteFixture(fp.FixtureList, fp.FixtureList[i.Row].ID)
					data = updateArray(fp.FixtureList)
					fp.FixturePanel.Refresh()
				}
			}

			// Fixture Add Button.
			if i.Col == FIXTURE_ADD {
				showField(FIXTURE_ADD, o)
				o.(*fyne.Container).Objects[FIXTURE_ADD].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[FIXTURE_ADD].(*widget.Button).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_ADD].(*widget.Button).OnTapped = func() {
					fp.FixtureList = AddFixture(fp.FixtureList, fp.FixtureList[i.Row].ID)
					data = updateArray(fp.FixtureList)
					fp.FixturePanel.Refresh()
				}
			}

			// Fixture Channels or States Button.
			if i.Col == FIXTURE_CHANNELS {
				showField(FIXTURE_CHANNELS, o)
				o.(*fyne.Container).Objects[FIXTURE_CHANNELS].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[FIXTURE_CHANNELS].(*widget.Button).SetText("->")
				o.(*fyne.Container).Objects[FIXTURE_CHANNELS].(*widget.Button).OnTapped = func() {
					fixtures.Fixtures = fp.FixtureList
					var modal *widget.PopUp
					if fp.FixtureList[i.Row].Type == "switch" {
						modal, err = NewStateEditor(w, fp.FixtureList[i.Row].ID, &fp, fixtures)
						if err != nil {
							fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", fp.FixtureList[i.Row].Group, fp.FixtureList[i.Row].Number, err)
							return
						}
					} else {
						modal, err = NewChannelEditor(w, fp.FixtureList[i.Row].ID, fp.FixtureList[i.Row].Channels, &fp, fixtures)
						if err != nil {
							fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", fp.FixtureList[i.Row].Group, fp.FixtureList[i.Row].Number, err)
							return
						}
					}
					modal.Resize(fyne.NewSize(800, 600))
					modal.Show()
				}
			}
		},
	)

	fp.FixturePanel.SetColumnWidth(0, 40)  // Id
	fp.FixturePanel.SetColumnWidth(1, 100) // Type
	fp.FixturePanel.SetColumnWidth(2, 60)  // Sequence Number
	fp.FixturePanel.SetColumnWidth(3, 60)  // Fixture Number
	fp.FixturePanel.SetColumnWidth(4, 100) // Name
	fp.FixturePanel.SetColumnWidth(5, 100) // Label
	fp.FixturePanel.SetColumnWidth(6, 50)  // DMX Address
	fp.FixturePanel.SetColumnWidth(7, 150) // Description
	fp.FixturePanel.SetColumnWidth(8, 20)  // Delete Button
	fp.FixturePanel.SetColumnWidth(9, 20)  // Add Button
	fp.FixturePanel.SetColumnWidth(10, 40) // Channels Button

	// Save button.
	buttonSave := widget.NewButton("Save", func() {

		// Insert updated fixture into fixtures.
		fixtures.Fixtures = fp.FixtureList

		// Clear switch positions to their first positions.
		for _, seq := range sequences {
			if seq.Type == "switch" {
				cmd := common.Command{
					Action: common.ResetAllSwitchPositions,
					Args: []common.Arg{
						{Name: "Fixtures", Value: fixtures},
					},
				}
				// Send a message to the switch sequence.
				common.SendCommandToSequence(seq.Number, cmd, commandChannels)
			}
		}

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
	panel := container.New(layout.NewGridWrapLayout(fyne.Size{Height: 430, Width: 750}), fp.FixturePanel)

	content := fyne.Container{}
	main := container.NewBorder(title, nil, nil, nil, header)
	two := container.NewBorder(main, nil, nil, nil, panel)
	content = *container.NewBorder(two, nil, nil, nil, saveCancel)

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		&content,
		w.Canvas(),
	)
	return modal, nil
}

func GetFixtureLabelsForSwitches(fixtures *fixture.Fixtures) []string {

	if debug {
		fmt.Printf("GetFixtureLabelsForSwitches\n")
	}

	fixturesAvailable := []string{}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Type != "switch" {
			fixturesAvailable = append(fixturesAvailable, fixture.Label)
		}
	}
	return fixturesAvailable
}

func UpdateFixture(fixtures []fixture.Fixture, id int, newItem fixture.Fixture) []fixture.Fixture {

	if debug {
		fmt.Printf("UpdateFixture\n")
	}

	newFixtures := []fixture.Fixture{}
	for _, fixture := range fixtures {
		if fixture.ID == id {
			// update the settings information.
			newFixtures = append(newFixtures, newItem)
		} else {
			// just add what was there before.
			newFixtures = append(newFixtures, fixture)
		}
	}
	return newFixtures
}

func AddFixture(fixtures []fixture.Fixture, id int) []fixture.Fixture {

	if debug {
		fmt.Printf("AddFixture\n")
	}

	newFixtures := []fixture.Fixture{}
	newFixture := fixture.Fixture{}
	newFixture.ID = id + 1
	if FixtureItemAllreadyExists(newFixture.ID, fixtures) {
		newFixture.ID = FindLargestFixtureNumber(fixtures) + 1
	}
	newFixture.Name = "New"
	newFixture.Type = "rgb"

	// Create a empty channel for this fixture.
	newChannels := []fixture.Channel{}
	newChannel := fixture.Channel{
		Number: 1,
	}
	newChannels = append(newChannels, newChannel)
	newFixture.Channels = newChannels

	for _, fixture := range fixtures {
		if fixture.ID == id {
			newFixtures = append(newFixtures, newFixture)
		}
		newFixtures = append(newFixtures, fixture)
	}
	sort.Slice(newFixtures, func(i, j int) bool {
		return newFixtures[i].ID < newFixtures[j].ID
	})
	return newFixtures
}

func DeleteFixture(fixtureList []fixture.Fixture, id int) []fixture.Fixture {

	if debug {
		fmt.Printf("DeleteFixture\n")
	}

	newFixtures := []fixture.Fixture{}
	if id == 1 {
		return fixtureList
	}
	for _, fixture := range fixtureList {
		if fixture.ID != id {
			newFixtures = append(newFixtures, fixture)
		}
	}
	return newFixtures
}

func FixtureItemAllreadyExists(id int, fixtureList []fixture.Fixture) bool {

	if debug {
		fmt.Printf("FixtureItemAllreadyExists\n")
	}

	// look through the fixture list for the id's
	for _, fixture := range fixtureList {
		if fixture.ID == id {
			return true
		}
	}
	return false
}

func FindLargestFixtureNumber(fixtures []fixture.Fixture) int {

	if debug {
		fmt.Printf("FindLargestFixtureNumber\n")
	}

	var number int
	for _, fixture := range fixtures {
		if fixture.ID > number {
			number = fixture.ID
		}
	}

	if debug {
		fmt.Printf("Largest %d\n", number)
	}

	return number
}

func hideAllFields(o fyne.CanvasObject) {

	if debug {
		fmt.Printf("hideAllFields\n")
	}

	// Hide everything.
	o.(*fyne.Container).Objects[FIXTURE_ID].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_TYPE].(*widget.Select).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_NAME].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_LABEL].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_ADD].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_CHANNELS].(*widget.Button).Hidden = true

}

func showField(field int, o fyne.CanvasObject) {

	if debug {
		fmt.Printf("showField\n")
	}

	// Now show the selected field.
	switch {
	case field == FIXTURE_ID:
		o.(*fyne.Container).Objects[FIXTURE_ID].(*widget.Label).Hidden = false
	case field == FIXTURE_TYPE:
		o.(*fyne.Container).Objects[FIXTURE_TYPE].(*widget.Select).Hidden = false
	case field == FIXTURE_GROUP:
		o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).Hidden = false
	case field == FIXTURE_NUMBER:
		o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).Hidden = false
	case field == FIXTURE_NAME:
		o.(*fyne.Container).Objects[FIXTURE_NAME].(*widget.Entry).Hidden = false
	case field == FIXTURE_LABEL:
		o.(*fyne.Container).Objects[FIXTURE_LABEL].(*widget.Entry).Hidden = false
	case field == FIXTURE_ADDRESS:
		o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*widget.Entry).Hidden = false
	case field == FIXTURE_DESCRIPTION:
		o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*widget.Entry).Hidden = false
	case field == FIXTURE_DELETE:
		o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).Hidden = false
	case field == FIXTURE_ADD:
		o.(*fyne.Container).Objects[FIXTURE_ADD].(*widget.Button).Hidden = false
	case field == FIXTURE_CHANNELS:
		o.(*fyne.Container).Objects[FIXTURE_CHANNELS].(*widget.Button).Hidden = false
	}
}

func makeNewFixture(data [][]string, i widget.TableCellID, field int, value string, fixtureList []fixture.Fixture) fixture.Fixture {

	if debug {
		fmt.Printf("makeNewFixture\n")
	}

	// Set up all the default values.
	newFixture := fixture.Fixture{}
	newFixture.ID, _ = strconv.Atoi(data[i.Row][FIXTURE_ID])

	newFixture.Type = data[i.Row][FIXTURE_TYPE]
	newFixture.Label = data[i.Row][FIXTURE_LABEL]
	newFixture.Name = data[i.Row][FIXTURE_NAME]
	number, _ := strconv.Atoi(data[i.Row][FIXTURE_NUMBER])
	newFixture.Number = number
	group, _ := strconv.Atoi(data[i.Row][FIXTURE_GROUP])
	newFixture.Group = group
	newFixture.Description = data[i.Row][FIXTURE_DESCRIPTION]
	address, _ := strconv.Atoi(data[i.Row][FIXTURE_ADDRESS])
	newFixture.Address = int16(address)

	// Set up the pointers to further data.
	newFixture.Channels = fixtureList[i.Row].Channels
	newFixture.States = fixtureList[i.Row].States
	newFixture.NumberChannels = fixtureList[i.Row].NumberChannels
	newFixture.UseFixture = fixtureList[i.Row].UseFixture

	// Now setup the new selected value.
	switch {
	case field == FIXTURE_ID:
		id, _ := strconv.Atoi(value)
		newFixture.ID = id

	case field == FIXTURE_TYPE:
		newFixture.Type = value

	case field == FIXTURE_GROUP:
		group, _ := strconv.Atoi(value)
		newFixture.Group = group

	case field == FIXTURE_NUMBER:
		number, _ := strconv.Atoi(value)
		newFixture.Number = number

	case field == FIXTURE_NAME:
		newFixture.Name = value

	case field == FIXTURE_LABEL:
		newFixture.Label = value

	case field == FIXTURE_ADDRESS:
		address, _ := strconv.Atoi(value)
		newFixture.Address = int16(address)

	case field == FIXTURE_DESCRIPTION:
		newFixture.Description = value

	}
	return newFixture
}
