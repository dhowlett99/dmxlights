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
	"regexp"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/override"
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

	DMXAddressEntryError  map[int]bool
	NameEntryError        map[int]bool
	LabelEntryError       map[int]bool
	DescriptionEntryError map[int]bool

	Fixtures *fixture.Fixtures
}

const RECTANGLE = 0
const TEXT = 1

const (
	FIXTURE_ID int = iota
	FIXTURE_TYPE
	FIXTURE_GROUP
	FIXTURE_NUMBER
	FIXTURE_NAME
	FIXTURE_LABEL
	FIXTURE_ADDRESS
	FIXTURE_DESCRIPTION
	FIXTURE_DELETE
	FIXTURE_ADD
	FIXTURE_CHANNELS
)

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

func generateFixtureNumberOptions(totalNumberOptions int) []string {

	var options []string
	for x := 1; x <= totalNumberOptions; x++ {
		options = append(options, strconv.Itoa(x))
	}
	return options
}

func NewFixturePanel(this *buttons.CurrentState, sequences []*common.Sequence, w fyne.Window, groupConfig *fixture.Groups, fixtures *fixture.Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, switchOverrides *[][]common.Override) (popupFixturePanel *widget.PopUp, err error) {

	if debug {
		fmt.Printf("NewFixturesPanel\n")
	}

	fp := FixturesPanel{}
	fp.Fixtures = fixtures
	fp.FixtureList = []fixture.Fixture{}

	// Populate group options from the available sequence labels.
	fp.GroupOptions = getGroupOptions(groupConfig)
	fp.NumberOptions = generateFixtureNumberOptions(8)
	fp.TypeOptions = []string{"rgb", "scanner", "switch", "projector"}

	// Storage for error flags for each fixture.
	fp.DMXAddressEntryError = make(map[int]bool, len(fp.FixtureList))
	fp.NameEntryError = make(map[int]bool, len(fp.FixtureList))
	fp.LabelEntryError = make(map[int]bool, len(fp.FixtureList))
	fp.DescriptionEntryError = make(map[int]bool, len(fp.FixtureList))

	// Create the save widget.
	var buttonSave *widget.Button

	// Create a dialog for error messages.
	var reports []string
	popupErrorPanel := &widget.PopUp{}
	// Ok button.
	button := widget.NewButton("OK", func() {
		popupErrorPanel.Hide()
	})
	popupErrorPanel = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("Title"),
			widget.NewLabel("Error Message"),
			widget.NewLabel("Report"),
			container.NewHBox(layout.NewSpacer(), button),
		),
		w.Canvas(),
	)

	// Load the fixtures into the array used by the table.
	data := makeArray(fixtures)

	// Title.
	title := widget.NewLabel("Fixture List")
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

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
		newItem.MultiFixtureDevice = f.MultiFixtureDevice
		newItem.NumberSubFixtures = f.NumberSubFixtures
		newItem.UseFixture = f.UseFixture
		fp.FixtureList = append(fp.FixtureList, newItem)
	}

	// Create a new fixtures list.
	fp.FixturePanel = widget.NewTableWithHeaders(
		// Function to find length of this table.
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
				address := fixture.GetFadeValuesFixtureAddressByName(fp.UseFixture, fixtures)
				data[fp.UpdateThisFixture][FIXTURE_ADDRESS] = address
				fp.FixtureList[fp.UpdateThisFixture].UseFixture = fp.UseFixture
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

		// Function to create items in this table.
		func() (o fyne.CanvasObject) {

			return container.NewStack(
				widget.NewLabel("id"), // ID.
				widget.NewSelect(fp.TypeOptions, func(value string) {}),   // Type rgb, scanner or switch.
				widget.NewSelect(fp.GroupOptions, func(value string) {}),  // Group Number.
				widget.NewSelect(fp.NumberOptions, func(value string) {}), // Fixture Number.
				container.NewStack(
					canvas.NewRectangle(colors.White),
					widget.NewEntry(), // Name.
				),
				container.NewStack(
					canvas.NewRectangle(colors.White),
					widget.NewEntry(), // Label.
				),
				container.NewStack(
					canvas.NewRectangle(colors.White),
					widget.NewEntry(), // DMX Address.
				),
				container.NewStack(
					canvas.NewRectangle(colors.White),
					widget.NewEntry(), // Description.
				),
				widget.NewButton("-", func() {}),        // Fixture delete button.
				widget.NewButton("+", func() {}),        // Fixture add button
				widget.NewButton("Channels", func() {}), // Channel Button
			)
		},

		// Function to update items in this table.
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
				o.(*fyne.Container).Objects[FIXTURE_TYPE].(*widget.Select).Selected = data[i.Row][i.Col]
				o.(*fyne.Container).Objects[FIXTURE_TYPE].(*widget.Select).Refresh()
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
						data = updateArray(fp.FixtureList)
					}
				}
			}

			// Fixture Group.
			if i.Col == FIXTURE_GROUP {
				showField(FIXTURE_GROUP, o)
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).SetSelected(getGroupName(groupConfig, data[i.Row][i.Col]))
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).Options = fp.GroupOptions
				o.(*fyne.Container).Objects[FIXTURE_GROUP].(*widget.Select).OnChanged = func(groupName string) {
					newFixture := makeNewFixture(data, i, FIXTURE_GROUP, getGroupFromName(groupConfig, groupName), fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
					data = updateArray(fp.FixtureList)
				}
			}

			// Fixture Number.
			if i.Col == FIXTURE_NUMBER {
				showField(FIXTURE_NUMBER, o)
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).SetSelected(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).OnChanged = func(value string) {
					newFixture := makeNewFixture(data, i, FIXTURE_NUMBER, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
					data = updateArray(fp.FixtureList)
				}
				o.(*fyne.Container).Objects[FIXTURE_NUMBER].(*widget.Select).PlaceHolder = "X"
			}

			// Fixture Name.
			if i.Col == FIXTURE_NAME {
				showField(FIXTURE_NAME, o)
				if fp.NameEntryError[fp.FixtureList[i.Row].ID] {
					o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.Red
				} else {
					o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.White
				}
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = func(value string) {
					if value != "" {
						o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).FocusGained()
						newFixture := makeNewFixture(data, i, FIXTURE_NAME, value, fp.FixtureList)
						fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
						data = updateArray(fp.FixtureList)

						// Clear all errors in all rows.
						for row := 0; row < len(data); row++ {
							fp.NameEntryError[row] = false
						}

						// Check the text entered.
						err := checkTextEntry(value)
						if err != nil {
							fp.NameEntryError[fp.FixtureList[i.Row].ID] = true
							fp.FixturePanel.Refresh()
							popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "Name Entry Error"
							popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
							popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
							o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
							fp.FixtureList[i.Row].Name = data[i.Row][i.Col]
							popupErrorPanel.Show()
							// Disable the save button.
							buttonSave.Disable()

						} else {
							fp.NameEntryError[fp.FixtureList[i.Row].ID] = false
							// And make sure we refresh every row, when we update this field.
							// So all the red error rectangls will disappear
							fp.FixturePanel.Refresh()
							// Enable the save button.
							buttonSave.Enable()
						}
					}
				}
				o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).PlaceHolder = "Name"
			}

			// Fixture Label.
			if i.Col == FIXTURE_LABEL {
				showField(FIXTURE_LABEL, o)
				if fp.LabelEntryError[fp.FixtureList[i.Row].ID] {
					o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.Red
				} else {
					o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.White
				}
				o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = func(value string) {
					o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[TEXT].(*widget.Entry).FocusGained()
					newFixture := makeNewFixture(data, i, FIXTURE_LABEL, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
					data = updateArray(fp.FixtureList)

					// Clear all errors in all rows.
					for row := 0; row < len(data); row++ {
						fp.LabelEntryError[row] = false
					}

					// Check label entry is valid.
					err := checkTextEntry(value)
					if err != nil {
						fp.LabelEntryError[fp.FixtureList[i.Row].ID] = true
						fp.FixturePanel.Refresh()
						popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "Label Entry Error"
						popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
						popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
						o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
						fp.FixtureList[i.Row].Label = data[i.Row][i.Col]
						popupErrorPanel.Show()
						// Disable the save button.
						buttonSave.Disable()
					} else {
						fp.LabelEntryError[fp.FixtureList[i.Row].ID] = false
						// And make sure we refresh every row, when we update this field.
						// So all the red error rectangls will disappear
						fp.FixturePanel.Refresh()
						// Enable the save button.
						buttonSave.Enable()
					}
				}
			}

			// Fixture DMX Address.
			if i.Col == FIXTURE_ADDRESS {
				showField(FIXTURE_ADDRESS, o)
				if fp.DMXAddressEntryError[fp.FixtureList[i.Row].ID] {
					o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.Red
				} else {
					o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.White
				}
				o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = func(value string) {
					if value != "" {
						o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).FocusGained()
						o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.White
						newFixture := makeNewFixture(data, i, FIXTURE_ADDRESS, value, fp.FixtureList)
						fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
						data = updateArray(fp.FixtureList)

						// Clear all errors in all rows.
						for row := 0; row < len(data); row++ {
							fp.DMXAddressEntryError[row] = false
						}

						// Check DMX Address is valid.
						err := checkDMXAddress(value)
						if err != nil {
							fp.DMXAddressEntryError[fp.FixtureList[i.Row].ID] = true
							fp.FixturePanel.Refresh()
							popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "DMX Entry Error"
							popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
							popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
							o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
							address, _ := strconv.Atoi(data[i.Row][i.Col])
							fp.FixtureList[i.Row].Address = int16(address)
							popupErrorPanel.Show()
							// Disable the save button.
							buttonSave.Disable()
						} else {
							fp.DMXAddressEntryError[fp.FixtureList[i.Row].ID] = false
							// And make sure we refresh every row, when we update this field.
							// So all the red error rectangls will disappear
							fp.FixturePanel.Refresh()
							// Enable the save button.
							buttonSave.Enable()
						}
					}
				}

				// Switch addresses are the address of the fixture being used.
				// So this comes from the state panel's use fixture field.
				// So if your a switch you can't change the DMX address here.
				if data[i.Row][FIXTURE_TYPE] == "switch" {
					o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).Disable()
				} else {
					o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).Enable()
				}
			}

			// Fixture Description.
			if i.Col == FIXTURE_DESCRIPTION {
				showField(FIXTURE_DESCRIPTION, o)
				if fp.DescriptionEntryError[fp.FixtureList[i.Row].ID] {
					o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.Red
				} else {
					o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = colors.White
				}
				o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = func(value string) {
					o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[TEXT].(*widget.Entry).FocusGained()
					newFixture := makeNewFixture(data, i, FIXTURE_DESCRIPTION, value, fp.FixtureList)
					fp.FixtureList = UpdateFixture(fp.FixtureList, fp.FixtureList[i.Row].ID, newFixture)
					data = updateArray(fp.FixtureList)

					// Clear all errors in all rows.
					for row := 0; row < len(data); row++ {
						fp.DescriptionEntryError[row] = false
					}

					// Check DMX Address is valid.
					err := checkTextEntry(value)
					if err != nil {
						fp.DescriptionEntryError[fp.FixtureList[i.Row].ID] = true
						fp.FixturePanel.Refresh()
						popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "Description Entry Error"
						popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
						popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
						o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
						fp.FixtureList[i.Row].Description = data[i.Row][i.Col]
						popupErrorPanel.Show()
						// Disable the save button.
						buttonSave.Disable()
					} else {
						fp.DescriptionEntryError[fp.FixtureList[i.Row].ID] = false
						// And make sure we refresh every row, when we update this field.
						// So all the red error rectangls will disappear
						fp.FixturePanel.Refresh()
						// Enable the save button.
						buttonSave.Enable()
					}
				}
			}

			// Fixture Delete Button.
			if i.Col == FIXTURE_DELETE {
				showField(FIXTURE_DELETE, o)
				o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[FIXTURE_DELETE].(*widget.Button).OnTapped = func() {
					fp.FixtureList = deleteFixture(fp.FixtureList, fp.FixtureList[i.Row].ID)
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
					fp.FixtureList = addFixture(fp.FixtureList, fp.FixtureList[i.Row].ID)
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
						modal, err = NewStatesEditor(w, fp.FixtureList[i.Row].ID, fp.FixtureList[i.Row].UseFixture, &fp, fixtures)
						if err != nil {
							fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", fp.FixtureList[i.Row].Group, fp.FixtureList[i.Row].Number, err)
							return
						}
					} else {
						modal, err = NewChannelEditor(w, fp.FixtureList[i.Row].ID, fp.FixtureList[i.Row].Channels, &fp, groupConfig, fixtures)
						if err != nil {
							fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", fp.FixtureList[i.Row].Group, fp.FixtureList[i.Row].Number, err)
							return
						}
					}
					modal.Resize(fyne.NewSize(800, 600))
					modal.Show()
					modal.Refresh()
				}
			}
		},
	)
	// Add headers
	fp.FixturePanel.ShowHeaderColumn = false
	fp.FixturePanel.CreateHeader = headerCreate
	fp.FixturePanel.UpdateHeader = headerUpdate

	fp.FixturePanel.SetColumnWidth(0, 40)  // Id
	fp.FixturePanel.SetColumnWidth(1, 100) // Type
	fp.FixturePanel.SetColumnWidth(2, 100) // Sequence Number
	fp.FixturePanel.SetColumnWidth(3, 59)  // Fixture Number
	fp.FixturePanel.SetColumnWidth(4, 80)  // Name
	fp.FixturePanel.SetColumnWidth(5, 80)  // Label
	fp.FixturePanel.SetColumnWidth(6, 50)  // DMX Address
	fp.FixturePanel.SetColumnWidth(7, 140) // Description
	fp.FixturePanel.SetColumnWidth(8, 20)  // Delete Button
	fp.FixturePanel.SetColumnWidth(9, 20)  // Add Button
	fp.FixturePanel.SetColumnWidth(10, 40) // Channels Button

	// Save button.
	buttonSave = widget.NewButton("OK", func() {

		// Remove any empty "None" actions from fixture list.
		fp.FixtureList = removeEmptyActions(fp.FixtureList)

		// Insert updated fixture into fixtures.
		fixtures.Fixtures = fp.FixtureList

		// Stop any running sequences.
		for _, sequence := range sequences {
			cmd := common.Command{
				Action: common.Stop,
			}
			// Send a message to the switch sequence.
			common.SendCommandToSequence(sequence.Number, cmd, commandChannels)
		}

		// Clear the sequence buttons.
		for _, sequence := range sequences {
			this.Running[sequence.Number] = false
			common.ShowRunningStatus(this.Running[sequence.Number], eventsForLaunchpad, guiButtons)
			// Clear the pattern function keys
			common.ClearSelectedRowOfButtons(sequence.Number, eventsForLaunchpad, guiButtons)
			// Turn off any function mode.
			this.SelectedMode[this.SelectedSequence] = buttons.NORMAL
			this.SelectButtonPressed[sequence.Number] = false
			buttons.SavePresetOff(this, eventsForLaunchpad, guiButtons)
			if this.Flood { // Turn off flood.
				buttons.FloodOff(len(sequences), this, commandChannels, eventsForLaunchpad, guiButtons)
			}
		}

		// Count the number of fixtures for all sequences since the user may have added or deleted fixtures.
		// The chaser uses the fixtures from the scanner group.
		for _, sequence := range sequences {
			cmd := common.Command{
				Action: common.UpdateFixturesConfig,
				Args: []common.Arg{
					{Name: "Fixtures", Value: fixtures},
				},
			}
			// Send a message to the switch sequence.
			common.SendCommandToSequence(sequence.Number, cmd, commandChannels)
		}

		// Find the switch sequence number.
		var SwitchSequenceNumber int
		for sequenceNumber, sequence := range sequences {
			if sequence.Type == "switch" {
				SwitchSequenceNumber = sequenceNumber
			}
		}
		// When we add a new set of fixtues with a possible new switch states we also need to populate a new override for that switch state.
		// So we recreate the overrides from scratch by using the pointer to SwitchOverrides.
		override.CreateOverrides(SwitchSequenceNumber, fixtures, switchOverrides)

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

		// Check DMX addresses don't overlap.
		reports, err = checkForNoOverlap(fixtures, fp)
		if err != nil {
			fmt.Printf("DMX Address %s:%s \n", err, strings.Join(reports, "\n"))
			fp.FixturePanel.Refresh()
			popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "DMX Address"
			popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
			popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
			popupErrorPanel.Show()
			// Disable the save button.
			buttonSave.Disable()
		} else {

			// Enable the save button.
			buttonSave.Enable()

			// Check Name is not duplicated.
			reports, err = checkForDuplicateName(fixtures, fp)
			if err != nil {
				fmt.Printf("Name Duplicated %s:%s \n", err, strings.Join(reports, "\n"))
				fp.FixturePanel.Refresh()
				popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "Name Duplicated "
				popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
				popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
				popupErrorPanel.Show()
				// Disable the save button.
				buttonSave.Disable()
			} else {
				// Enable the save button.
				buttonSave.Enable()

				// Check Name is not duplicated.
				reports, err = checkForDuplicateLabel(fixtures, fp)
				if err != nil {
					fmt.Printf("Label Duplicated %s:%s \n", err, strings.Join(reports, "\n"))
					fp.FixturePanel.Refresh()
					popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "Label Duplicated "
					popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
					popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
					popupErrorPanel.Show()
					// Disable the save button.
					buttonSave.Disable()
				} else {
					// Enable the save button.
					buttonSave.Enable()
					popupFixturePanel.Hide()
				}
			}
		}
	})

	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		popupFixturePanel.Hide()
	})
	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)
	panel := container.New(layout.NewGridWrapLayout(fyne.Size{Height: 500, Width: 750}), fp.FixturePanel)

	content := fyne.Container{}
	main := container.NewBorder(title, nil, nil, nil, panel)
	content = *container.NewBorder(main, nil, nil, nil, saveCancel)

	// popup fixture panel.
	popupFixturePanel = widget.NewModalPopUp(
		&content,
		w.Canvas(),
	)
	return popupFixturePanel, nil
}

func getGroupOptions(groupConfig *fixture.Groups) []string {
	if debug {
		fmt.Printf("getGroupOptions\n")
	}
	var groupOptions []string
	for _, group := range groupConfig.Groups {
		groupOptions = append(groupOptions, group.Name)
	}
	return groupOptions
}

func getGroupFromName(groupConfig *fixture.Groups, groupName string) string {
	if debug {
		fmt.Printf("getGroupFromName %s\n", groupName)
	}
	for _, group := range groupConfig.Groups {
		if group.Name == groupName {
			return group.Number
		}
	}
	return groupName
}

func getGroupName(groupConfig *fixture.Groups, groupNumber string) string {
	if debug {
		fmt.Printf("getGroupName %s\n", groupNumber)
	}
	for _, group := range groupConfig.Groups {
		if group.Number == groupNumber {
			return group.Name
		}
	}
	return groupNumber
}

func checkForDuplicateName(fixtures *fixture.Fixtures, fp FixturesPanel) ([]string, error) {

	if debug {
		fmt.Printf("checkForDuplicateName\n")
	}

	var err error
	var reports []string

	for _, fixture := range fixtures.Fixtures {
		for _, testfixture := range fixtures.Fixtures {
			if fixture.Type != "switch" && fixture.ID != testfixture.ID {
				if len(strings.TrimSpace(fixture.Name)) == 0 {
					fp.NameEntryError[fixture.ID] = true
					// We have an empty name
					err = fmt.Errorf("empty name")
					reports = append(reports, fmt.Sprintf("empty name on fixture ID %d and ID %d with name %s", fixture.ID, testfixture.ID, fixture.Name))
					return reports, err
				}
				if fixture.Name == testfixture.Name {
					fp.NameEntryError[fixture.ID] = true
					// We have an duplicate name
					err = fmt.Errorf("duplicate names")
					reports = append(reports, fmt.Sprintf("duplicate names on fixture ID %d and ID %d with name %s", fixture.ID, testfixture.ID, fixture.Name))
				}
			}
		}
	}
	return reports, err
}

func checkForDuplicateLabel(fixtures *fixture.Fixtures, fp FixturesPanel) ([]string, error) {

	if debug {
		fmt.Printf("checkForDuplicateLabel\n")
	}

	var err error
	var reports []string

	for _, fixture := range fixtures.Fixtures {
		for _, testfixture := range fixtures.Fixtures {
			if fixture.Type != "switch" && fixture.ID != testfixture.ID {
				if fixture.Label == testfixture.Label {
					fp.LabelEntryError[fixture.ID] = true
					// We have an duplicate label
					err = fmt.Errorf("duplicate labels")
					reports = append(reports, fmt.Sprintf("duplicate label on fixture ID %d and ID %d with label %s", fixture.ID, testfixture.ID, fixture.Label))
				}
			}
		}
	}
	return reports, err
}

func checkForNoOverlap(fixtures *fixture.Fixtures, fp FixturesPanel) ([]string, error) {

	if debug {
		fmt.Printf("checkForNoOverlap\n")
	}

	var err error
	var reports []string

	for _, fixture := range fixtures.Fixtures {
		for _, testfixture := range fixtures.Fixtures {
			if fixture.Type != "switch" && fixture.ID != testfixture.ID {
				if checkOverlap(int(fixture.Address), int(fixture.Address)+len(fixture.Channels), int(testfixture.Address), int(testfixture.Address)+len(testfixture.Channels)) {
					fp.DMXAddressEntryError[fixture.ID] = true
					// We have an overlapping DMX address.
					err = fmt.Errorf("overlapping DMX Address")
					reports = append(reports, fmt.Sprintf("overlapping DMX Address on fixture %s with fixture %s", fixture.Name, testfixture.Name))
					return reports, err
				}
			}
		}
	}
	return reports, err
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

// checkOverlap
func checkOverlap(aStart int, aEnd int, bStart int, bEnd int) bool {
	if debug {
		fmt.Printf("checkOverlap\n")
	}
	return (aStart >= bEnd) != (aEnd > bStart)
}

func checkDMXAddress(value string) error {

	if len(strings.TrimSpace(value)) == 0 || len(value) == 0 {
		return fmt.Errorf("DMX error, value is empty")
	}

	address, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("DMX Address error, must only contain numbers")
	}
	if address == 0 {
		return fmt.Errorf("DMX Address error, cannot be zero")
	}
	if address > common.MAX_DMX_ADDRESS {
		return fmt.Errorf("DMX Address error, cannot be greater than %d", common.MAX_DMX_ADDRESS)
	}
	return nil
}

func checkDMXValue(value string) error {

	if len(strings.TrimSpace(value)) == 0 || len(value) == 0 {
		return fmt.Errorf("DMX Value error, value is empty")
	}

	// Filter out any characters and special characters.
	var IsRangeOrNumber = regexp.MustCompile(`^[0-9\-]+$`).MatchString
	if !IsRangeOrNumber(value) {
		return fmt.Errorf("DMX Value error, must only numbers, ranges can be specified e.g. 10-20")
	}

	// No check the numbers in a range.
	if strings.Contains(value, "-") {
		// We've found a range of values.
		// Find the start value
		numbers := strings.Split(value, "-")

		// Now apply the range depending on the speed
		// Check the start of the range.
		err := checkDMXnumber(numbers[0])
		if err != nil {
			return err
		}
		// Check the stop value of the range.
		if numbers[1] != "" {
			err = checkDMXnumber(numbers[1])
			if err != nil {
				return err
			}
		}

		// Check the range makes sense.
		if numbers[1] < numbers[0] && numbers[1] != "" {
			return fmt.Errorf("second value in range must be greater than first")
		}
	} else {
		// Check the single value (no range case).
		err := checkDMXnumber(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkDMXnumber(value string) error {

	if debug {
		fmt.Printf("checkDMXnumber: checking value %s\n", value)
	}

	if value == "" {
		return nil
	}
	address, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("DMX Number error, must only contain numbers")
	}
	if address < 0 {
		return fmt.Errorf("DMX Number error, cannot be less than zero")
	}
	if address > common.MAX_DMX_ADDRESS {
		return fmt.Errorf("DMX Number error, cannot be greater than %d", common.MAX_DMX_ADDRESS)
	}

	return nil
}

func checkTextEntry(value string) error {

	if debug {
		fmt.Printf("checkTextEntry %s\n", value)
	}

	if len(strings.TrimSpace(value)) == 0 || len(value) == 0 {
		return fmt.Errorf("cannot be an empty string")
	}

	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9\ \.\_]+$`).MatchString
	if !IsLetter(value) {
		return fmt.Errorf("must only contain letters and numbers, underscore and dots are allowed")
	}

	if len(value) > common.MAX_TEXT_ENTRY_LENGTH {
		return fmt.Errorf("cannot be greater than %d characters", common.MAX_TEXT_ENTRY_LENGTH)
	}
	return nil
}

func addFixture(fixtures []fixture.Fixture, id int) (outItems []fixture.Fixture) {

	if debug {
		fmt.Printf("AddFixture\n")
	}

	newFixtures := []fixture.Fixture{}
	newFixture := fixture.Fixture{}
	newFixture.ID = id + 1
	newFixture.Group = 1
	newFixture.Number = 1
	newFixture.Name = fmt.Sprintf("Fixture %d", newFixture.ID)
	newFixture.Label = fmt.Sprintf("Label %d", newFixture.ID)
	newFixture.Description = fmt.Sprintf("Desc %d", newFixture.ID)
	newFixture.Type = "(Select One)"

	// Create a empty channel list for this fixture.
	newChannels := []fixture.Channel{}
	newChannel := fixture.Channel{
		Number: 1,
	}
	newChannels = append(newChannels, newChannel)
	newFixture.Channels = newChannels

	var added bool // Only add once.

	for no, fixture := range fixtures {
		// Add at the start of an empty list.
		if len(fixtures) == 0 && !added {
			newFixtures = append(newFixtures, newFixture)
			added = true
		}
		// Insert at this position.
		if fixture.ID == id+1 && !added {
			newFixtures = append(newFixtures, newFixture)
			added = true
		}
		newFixtures = append(newFixtures, fixture)
		// Append an item at the very end.
		if no == len(fixtures)-1 && !added {
			newFixtures = append(newFixtures, newFixture)
			added = true
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newFixtures {
		indexedItem.ID = number + 1
		outItems = append(outItems, indexedItem)
	}

	return outItems
}

func deleteFixture(fixtureList []fixture.Fixture, id int) (outItems []fixture.Fixture) {

	if debug {
		fmt.Printf("DeleteFixture\n")
	}

	newFixtures := []fixture.Fixture{}
	// if id == 1 {
	// 	return fixtureList
	// }
	for _, fixture := range fixtureList {
		if fixture.ID != id {
			newFixtures = append(newFixtures, fixture)
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newFixtures {
		indexedItem.ID = number + 1
		outItems = append(outItems, indexedItem)
	}

	if len(outItems) == 0 {
		// Create a default fixture with one default channel.
		newFixture := fixture.Fixture{}
		newFixture.ID = 1
		newFixture.Group = 1
		newFixture.Number = 1
		newFixture.Name = fmt.Sprintf("Fixture %d", newFixture.ID)
		newFixture.Label = fmt.Sprintf("Label %d", newFixture.ID)
		newFixture.Description = fmt.Sprintf("Desc %d", newFixture.ID)
		newFixture.Type = "(Select One)"

		emptyValue := int16(0)
		newChannel := fixture.Channel{
			Number:  0,
			Name:    "New",
			Value:   &emptyValue,
			Comment: "New",
			//Settings
		}
		newFixture.Channels = append(newFixture.Channels, newChannel)
		outItems = append(outItems, newFixture)
	}

	return outItems
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
	o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = true
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
		o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = false
		o.(*fyne.Container).Objects[FIXTURE_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = false
	case field == FIXTURE_LABEL:
		o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = false
		o.(*fyne.Container).Objects[FIXTURE_LABEL].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = false
	case field == FIXTURE_ADDRESS:
		o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = false
		o.(*fyne.Container).Objects[FIXTURE_ADDRESS].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = false
	case field == FIXTURE_DESCRIPTION:
		o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = false
		o.(*fyne.Container).Objects[FIXTURE_DESCRIPTION].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = false
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
	newFixture.MultiFixtureDevice = fixtureList[i.Row].MultiFixtureDevice
	newFixture.NumberSubFixtures = fixtureList[i.Row].NumberSubFixtures
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

type ActiveHeader struct {
	widget.Label
	OnTapped func()
}

func headerCreate() fyne.CanvasObject {
	h := &ActiveHeader{}
	h.ExtendBaseWidget(h)
	h.SetText("000")
	return h
}

func headerUpdate(id widget.TableCellID, o fyne.CanvasObject) {
	header := o.(*ActiveHeader)
	header.TextStyle.Bold = true
	switch id.Col {
	case -1:
		header.SetText(strconv.Itoa(id.Row + 1))
	case 0:
		header.SetText("ID")
	case 1:
		header.SetText("Type")
	case 2:
		header.SetText("Group")
	case 3:
		header.SetText("No")
	case 4:
		header.SetText("Name")
	case 5:
		header.SetText("Label")
	case 6:
		header.SetText("DMX")
	case 7:
		header.SetText("Description")
	case 8:
		header.SetText("-")
	case 9:
		header.SetText("+")
	case 10:
		header.SetText("Select")
	}

	// header.OnTapped = func() {
	// 	fmt.Printf("Header %d tapped\n", id.Col)
	// }
}

func (h *ActiveHeader) Tapped(_ *fyne.PointEvent) {
	if h.OnTapped != nil {
		h.OnTapped()
	}
}

func (h *ActiveHeader) TappedSecondary(_ *fyne.PointEvent) {
}

func removeEmptyActions(fixtureList []fixture.Fixture) []fixture.Fixture {

	var newFixtureList []fixture.Fixture

	for _, f := range fixtureList {

		newFixture := fixture.Fixture{}

		newFixture.ID = f.ID
		newFixture.Name = f.Name
		newFixture.Label = f.Label
		newFixture.Number = f.Number
		newFixture.Description = f.Description
		newFixture.Type = f.Type
		newFixture.Group = f.Group
		newFixture.Address = f.Address
		newFixture.Channels = f.Channels

		newFixture.MultiFixtureDevice = f.MultiFixtureDevice
		newFixture.NumberSubFixtures = f.NumberSubFixtures
		newFixture.UseFixture = f.UseFixture

		newStates := []fixture.State{}

		for _, state := range f.States {

			newState := fixture.State{}
			newActions := []fixture.Action{}
			for _, action := range state.Actions {
				if action.Mode != "None" {
					newActions = append(newActions, action)
				}
			}

			newState.Name = state.Name
			newState.Number = state.Number
			newState.Label = state.Label
			newState.ButtonColor = state.ButtonColor
			newState.Master = state.Master
			newState.Actions = newActions
			newState.Settings = state.Settings
			newState.Flash = state.Flash

			newStates = append(newStates, newState)

		}
		newFixture.States = newStates
		newFixtureList = append(newFixtureList, newFixture)

	}
	return newFixtureList
}
