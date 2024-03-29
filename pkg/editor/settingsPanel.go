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
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

type SettingsPanel struct {
	SettingsPanel     *widget.Table
	SettingsList      []fixture.Setting
	SettingMaxDegrees *int
	SettingsOptions   []string
	ChannelOptions    []string
	CurrentChannel    int
	UpdateThisChannel int
	UpdateSettings    bool

	NameEntryError     map[int]bool
	DMXValueEntryError map[int]bool
}

const (
	SETTING_NUMBER int = iota
	SETTING_NAME
	SETTING_CHANNEL
	SETTING_VALUE
	SETTING_DELETE
	SETTING_ADD
)

const (
	TITLE int = iota
	MESSAGE
	REPORT
)

func NewSettingsPanel(w fyne.Window, SettingsList []fixture.Setting, channelFieldDisabled bool, buttonSave *widget.Button) *SettingsPanel {

	if debug {
		fmt.Printf("NewSettingsPanel\n")
	}

	var data = [][]string{}

	st := SettingsPanel{}
	st.SettingsList = SettingsList
	st.SettingsOptions = []string{"Off", "On", "Red", "Green", "Blue", "Soft", "Sharp", "Sound", "Rotate"}
	st.ChannelOptions = []string{"None"}

	Red := color.RGBA{}
	Red.R = uint8(255)
	Red.G = uint8(0)
	Red.B = uint8(0)
	Red.A = 255

	White := color.White

	// Storage for error flags for each fixture.
	st.NameEntryError = make(map[int]bool, len(st.SettingsList))
	st.DMXValueEntryError = make(map[int]bool, len(st.SettingsList))

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

	// Settingses Selection Panel.
	st.SettingsPanel = widget.NewTable(

		// Function to find length.
		func() (int, int) {
			height := len(data)
			width := 6
			return height, width
		},

		// Function to create table.
		func() (o fyne.CanvasObject) {

			// Load the settings into the array used by the table.
			data = makeSettingsArray(st.SettingsList)

			return container.NewStack(
				// SETTING_NUMBER
				widget.NewLabel("template"),

				// SETTING_NAME
				container.NewStack(
					canvas.NewRectangle(color.White),
					widget.NewEntry(),
				),

				// SETTING_CHANNEL
				widget.NewSelect(st.ChannelOptions, func(value string) {}),

				// SETTING_VALUE
				container.NewStack(
					canvas.NewRectangle(color.White),
					widget.NewEntry(),
				),

				// SETTING_DELETE
				widget.NewButton("-", func() {}),

				// SETTING_ADD
				widget.NewButton("+", func() {}),
			)
		},

		// Function to update item in this table.
		func(i widget.TableCellID, o fyne.CanvasObject) {

			// Hide all field types.
			hideAllSettingsFields(o)

			// Show the setting a number.
			if i.Col == SETTING_NUMBER {
				showSettingsField(SETTING_NUMBER, o)
				o.(*fyne.Container).Objects[SETTING_NUMBER].(*widget.Label).SetText(data[i.Row][i.Col])
			}

			// Show and Edit the Setting Name.
			if i.Col == SETTING_NAME {
				showSettingsField(SETTING_NAME, o)
				if st.NameEntryError[st.SettingsList[i.Row].Number] {
					o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = Red
				} else {
					o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = White
				}
				o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = func(settingName string) {
					if settingName != "" {
						newSetting := fixture.Setting{}
						newSetting.Label = settingName // Label same as name.
						newSetting.Name = settingName
						newSetting.Number = st.SettingsList[i.Row].Number
						if !channelFieldDisabled {
							newSetting.Channel = st.SettingsList[i.Row].Channel
						}
						newSetting.Value = st.SettingsList[i.Row].Value
						st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
						data = makeSettingsArray(st.SettingsList)
						st.UpdateSettings = true
						st.UpdateThisChannel = st.CurrentChannel - 1

						// Clear all errors in all rows.
						for row := 0; row < len(data); row++ {
							st.NameEntryError[row] = false
						}

						// Check the text entered.
						err := checkTextEntry(settingName)
						if err != nil {
							st.NameEntryError[st.SettingsList[i.Row].Number] = true
							st.SettingsPanel.Refresh()
							// Populate error message panel.
							popupErrorPanel.Content.(*fyne.Container).Objects[TITLE].(*widget.Label).Text = "Name Entry Error"
							popupErrorPanel.Content.(*fyne.Container).Objects[MESSAGE].(*widget.Label).Text = err.Error()
							popupErrorPanel.Content.(*fyne.Container).Objects[REPORT].(*widget.Label).Text = strings.Join(reports, "\n")
							o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
							st.SettingsList[i.Row].Name = data[i.Row][i.Col]
							popupErrorPanel.Show()
							// Disable the save button.
							buttonSave.Disable()

						} else {
							st.NameEntryError[st.SettingsList[i.Row].Number] = false
							// And make sure we refresh every row, when we update this field.
							// So all the red error rectangls will disappear
							st.SettingsPanel.Refresh()
							// Enable the save button.
							buttonSave.Enable()
						}
					}
				}
			}

			// Channel number.
			if i.Col == SETTING_CHANNEL {
				showSettingsField(SETTING_CHANNEL, o)
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).OnChanged = nil
				// Update the options to include any thing that might specified in the config file.
				st.ChannelOptions = addChannelOption(st.ChannelOptions, data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).Options = st.ChannelOptions
				// Match the options to the data in the field and display in anyway.
				for _, option := range st.ChannelOptions {
					if option == data[i.Row][i.Col] {
						o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).SetSelected(option)
					}
				}
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).Hidden = channelFieldDisabled
				o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).OnChanged = func(settingChannel string) {
					newSetting := fixture.Setting{}
					newSetting.Label = st.SettingsList[i.Row].Label
					newSetting.Name = st.SettingsList[i.Row].Name
					newSetting.Number = st.SettingsList[i.Row].Number
					newSetting.Channel = settingChannel
					newSetting.Value = st.SettingsList[i.Row].Value
					st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
				}
			}

			// Show and Edit the Setting Value.
			if i.Col == SETTING_VALUE {
				showSettingsField(SETTING_VALUE, o)
				if st.DMXValueEntryError[st.SettingsList[i.Row].Number] {
					o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = Red
				} else {
					o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).FillColor = White
				}
				o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = nil
				o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[TEXT].(*widget.Entry).OnChanged = func(settingValue string) {
					if settingValue != "" {
						newSetting := fixture.Setting{}
						newSetting.Label = st.SettingsList[i.Row].Label
						newSetting.Name = st.SettingsList[i.Row].Name
						newSetting.Number = st.SettingsList[i.Row].Number
						if !channelFieldDisabled {
							newSetting.Channel = st.SettingsList[i.Row].Channel
						}
						newSetting.Value = settingValue
						st.SettingsList = updateSettingsItem(st.SettingsList, newSetting.Number, newSetting)
						data = makeSettingsArray(st.SettingsList)
						st.UpdateSettings = true
						st.UpdateThisChannel = st.CurrentChannel - 1

						// Clear all errors in all rows.
						for row := 0; row < len(data); row++ {
							st.NameEntryError[row] = false
						}

						if newSetting.Name == "MaxDegrees" {
							// We're adding a MaxDegrees setting to the channel.
							// Populate MaxDegrees - If we are a channel name of Pan or Tilt and any of the settings contain the name MaxDegrees
							// add the value to newChannel.MaxDegrees
							maxDegrees, _ := strconv.Atoi(newSetting.Value)
							st.SettingMaxDegrees = &maxDegrees
							st.UpdateSettings = true
							st.UpdateThisChannel = st.CurrentChannel - 1
						} else {
							// Check the text entered.
							err := checkDMXValue(settingValue)
							if err != nil {
								st.DMXValueEntryError[st.SettingsList[i.Row].Number] = true
								st.SettingsPanel.Refresh()
								popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "Value Entry Error"
								popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = err.Error()
								popupErrorPanel.Content.(*fyne.Container).Objects[2].(*widget.Label).Text = strings.Join(reports, "\n")
								o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[TEXT].(*widget.Entry).SetText(data[i.Row][i.Col])
								st.SettingsList[i.Row].Value = data[i.Row][i.Col]
								popupErrorPanel.Show()
								// Disable the save button.
								buttonSave.Disable()

							} else {
								st.DMXValueEntryError[st.SettingsList[i.Row].Number] = false
								// And make sure we refresh every row, when we update this field.
								// So all the red error rectangls will disappear
								st.SettingsPanel.Refresh()
								// Enable the save button.
								buttonSave.Enable()
							}
						}
					}
				}
			}

			// Show the Delete Setting Button.
			if i.Col == SETTING_DELETE {
				showSettingsField(SETTING_DELETE, o)
				o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).OnTapped = func() {
					if len(st.SettingsList) != 0 {
						st.SettingsList = deleteSettingsItem(st.SettingsList, st.SettingsList[i.Row].Number-1)
					}
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1

					if len(st.SettingsList) == 0 {
						st.SettingsPanel.Hide()
					} else {
						st.SettingsPanel.Refresh()
					}
				}
			}

			// Show the Add Setting Button.
			if i.Col == SETTING_ADD {
				showSettingsField(SETTING_ADD, o)
				o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).OnTapped = nil
				o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).SetText(data[i.Row][i.Col])
				o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).OnTapped = func() {
					if len(st.SettingsList) != 0 {
						st.SettingsList = addSettingsItem(st.SettingsList, st.SettingsList[i.Row].Number)
					} else {
						st.SettingsList = addSettingsItem(st.SettingsList, 0)
					}
					data = makeSettingsArray(st.SettingsList)
					st.UpdateSettings = true
					st.UpdateThisChannel = st.CurrentChannel - 1
					st.SettingsPanel.Refresh()
				}
			}
		},
	)

	// Setup the columns of this table.
	st.SettingsPanel.SetColumnWidth(0, 40)  // Number
	st.SettingsPanel.SetColumnWidth(1, 100) // Name
	if channelFieldDisabled {
		st.SettingsPanel.SetColumnWidth(2, 0) // Channel
	} else {
		st.SettingsPanel.SetColumnWidth(2, 100) // Channel
	}
	st.SettingsPanel.SetColumnWidth(3, 80) // Value
	st.SettingsPanel.SetColumnWidth(4, 20) // Delete
	st.SettingsPanel.SetColumnWidth(5, 20) // Add

	return &st
}

func addChannelOption(options []string, newOption string) []string {
	newOptions := []string{}
	for _, option := range options {
		if option != newOption {
			newOptions = append(newOptions, option)
		}
	}
	// now add the new option.
	newOptions = append(newOptions, newOption)
	return newOptions
}

func addSettingsItem(items []fixture.Setting, number int) (outItems []fixture.Setting) {

	if debug {
		fmt.Printf("addSettingsItem\n")
	}

	newItems := []fixture.Setting{}
	newItem := fixture.Setting{}
	newItem.Number = int(number) + 1
	newItem.Name = "New"
	newItem.Label = "New"
	newItem.Channel = "(Select one)"
	newItem.Value = "0"

	var added bool // Only add once.

	if number == 0 {
		items = append(items, newItem)
	}

	for no, item := range items {
		// Add at the start of an empty list.
		if len(items) == 0 && !added {
			newItems = append(newItems, newItem)
			added = true
		}
		// Insert at this position.
		if item.Number == number+1 && !added {
			fmt.Printf("Insert at this position %+v\n", newItem)
			newItems = append(newItems, newItem)
			added = true
		}
		newItems = append(newItems, item)
		// Append an item at the very end.
		if no == len(items)-1 && !added {
			newItems = append(newItems, newItem)
			added = true
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newItems {
		indexedItem.Number = number + 1
		outItems = append(outItems, indexedItem)
	}

	return outItems
}

func deleteSettingsItem(settingsList []fixture.Setting, id int) (outItems []fixture.Setting) {

	if debug {
		fmt.Printf("deleteSettingsItem\n")
	}

	newSettings := []fixture.Setting{}
	for settingNumber, setting := range settingsList {
		if settingNumber != id {
			newSettings = append(newSettings, setting)
		}
	}

	// Now fix the item numbers
	for number, indexedItem := range newSettings {
		indexedItem.Number = number + 1
		outItems = append(outItems, indexedItem)
	}

	return outItems
}

func updateSettingsItem(items []fixture.Setting, id int, newItem fixture.Setting) []fixture.Setting {

	if debug {
		fmt.Printf("updateSettingsItem\n")
	}

	newItems := []fixture.Setting{}
	for _, item := range items {
		if item.Number == id {
			// update the settings information.
			newItems = append(newItems, newItem)
		} else {
			// just add what was there before.
			newItems = append(newItems, item)
		}
	}
	return newItems
}

// makeSettingsArray - Convert the list of settings to an array of strings containing and array of strings with
// the values from each fixture.
// This is done once when the settings panel is loaded.
func makeSettingsArray(settings []fixture.Setting) [][]string {

	if debug {
		fmt.Printf("makeSettingsArray\n")
	}

	var data = [][]string{}

	for _, setting := range settings {
		newSetting := []string{}
		newSetting = append(newSetting, fmt.Sprintf("%d", setting.Number))
		newSetting = append(newSetting, setting.Name)
		newSetting = append(newSetting, setting.Channel)
		newSetting = append(newSetting, setting.Value)
		newSetting = append(newSetting, "-")
		newSetting = append(newSetting, "+")
		newSetting = append(newSetting, "Channels")

		data = append(data, newSetting)
	}

	return data
}

func showSettingsField(field int, o fyne.CanvasObject) {
	if debug {
		fmt.Printf("showField\n")
	}
	// Now show the selected field.
	switch {
	case field == SETTING_NUMBER:
		o.(*fyne.Container).Objects[SETTING_NUMBER].(*widget.Label).Hidden = false
	case field == SETTING_NAME:
		o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = false
		o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = false
	case field == SETTING_CHANNEL:
		o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).Hidden = false
	case field == SETTING_VALUE:
		o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = false
		o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = false
	case field == SETTING_DELETE:
		o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).Hidden = false
	case field == SETTING_ADD:
		o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).Hidden = false
	}
}

func hideAllSettingsFields(o fyne.CanvasObject) {
	if debug {
		fmt.Printf("hideAllSettingsFields\n")
	}
	o.(*fyne.Container).Objects[SETTING_NUMBER].(*widget.Label).Hidden = true
	o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[SETTING_NAME].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[SETTING_CHANNEL].(*widget.Select).Hidden = true
	o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[TEXT].(*widget.Entry).Hidden = true
	o.(*fyne.Container).Objects[SETTING_VALUE].(*fyne.Container).Objects[RECTANGLE].(*canvas.Rectangle).Hidden = true
	o.(*fyne.Container).Objects[SETTING_DELETE].(*widget.Button).Hidden = true
	o.(*fyne.Container).Objects[SETTING_ADD].(*widget.Button).Hidden = true
}
