// Copyright (C) 2022, 2023 dhowlett99.
// This implements the function keys, used by the buttons package.
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

package buttons

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func ShowFunctionButtons(this *CurrentState, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("ShowFunctionButtons sequence target %d display %d\n", this.TargetSequence, this.DisplaySequence)
	}
	// Loop through the available functions for this sequence
	for index, function := range this.Functions[this.TargetSequence] {
		if debug {
			fmt.Printf("ShowFunctionButtons: function %s state %t\n", function.Name, function.State)
		}
		if !function.State && this.SelectMode[this.DisplaySequence] != CHASER { // Cyan
			common.LightLamp(common.ALight{X: index, Y: this.DisplaySequence, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
		}
		if !function.State && this.SelectMode[this.DisplaySequence] == CHASER { // Yellow
			common.LightLamp(common.ALight{X: index, Y: this.DisplaySequence, Brightness: 255, Red: 255, Green: 255, Blue: 0}, eventsForLauchpad, guiButtons)
		}
		if function.State { // Purple
			common.LightLamp(common.ALight{X: index, Y: this.DisplaySequence, Brightness: 255, Red: 200, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
		}
		common.LabelButton(index, this.DisplaySequence, function.Label, guiButtons)
	}
}

func processFunctions(X int, Y int, sequences []*common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence) {

	debug := false

	if this.SelectMode[this.SelectedSequence] == CHASER {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}
	if debug {
		fmt.Printf("Function Key X:%d Y:%d\n", X, Y)
		fmt.Printf("this.TargetSequence %d\n", this.TargetSequence)
		fmt.Printf("this.DisplaySequence %d\n", this.DisplaySequence)
	}

	if debug {
		fmt.Printf("FUNCS: this.Type = %s \n", this.SelectedType)
		for functionNumber := 0; functionNumber < 8; functionNumber++ {
			state := this.Functions[this.TargetSequence][functionNumber].State
			fmt.Printf("FUNCS: function %d state %t\n", functionNumber, state)
		}
		fmt.Printf("FUNCS: this.ChaserRunning %t \n", this.ScannerChaser)

		fmt.Printf("================== WHAT SELECT MODE =================\n")
		fmt.Printf("FUNCS: this.SelectButtonPressed[%d] = %t \n", this.DisplaySequence, this.SelectButtonPressed[this.DisplaySequence])
		if this.SelectMode[this.TargetSequence] == NORMAL {
			fmt.Printf("FUNCS: this.SelectMode[%d] = NORMAL \n", this.DisplaySequence)
		}
		if this.SelectMode[this.TargetSequence] == FUNCTION {
			fmt.Printf("FUNCS: this.SelectMode[%d] = FUNCTION \n", this.SelectedSequence)
		}
		if this.SelectMode[this.TargetSequence] == CHASER {
			fmt.Printf("FUNCS: this.SelectMode[%d] = CHASER \n", this.SelectedSequence)
		}
		if this.SelectMode[this.TargetSequence] == STATUS {
			fmt.Printf("FUNCS: this.SelectMode[%d] = STATUS \n", this.SelectedSequence)
		}

		fmt.Printf("================== WHAT EDIT MODES =================\n")
		fmt.Printf("FUNCS: this.EditSequenceColorsMode[%d] = %t \n", this.DisplaySequence, this.EditSequenceColorsMode)
		fmt.Printf("FUNCS: this.EditStaticColorsMode[%d] = %t \n", this.DisplaySequence, this.EditStaticColorsMode)
		fmt.Printf("FUNCS: this.EditGoboSelectionMode[%d] = %t \n", this.DisplaySequence, this.EditGoboSelectionMode)
		fmt.Printf("FUNCS: this.EditPatternMode[%d] = %t \n", this.DisplaySequence, this.EditPatternMode)
		fmt.Printf("===============================================\n")
	}

	// Map Function 1 - Go straight into pattern select mode, don't wait for a another select press.
	if X == common.Function1_Pattern {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function1_Pattern \n", this.DisplaySequence, this.SelectMode[this.DisplaySequence])
		}

		this.EditPatternMode = true
		this.Functions[this.DisplaySequence][common.Function1_Pattern].State = true
		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)
		this.EditFixtureSelectionMode = false

		ShowPatternSelectionButtons(this, sequences[this.TargetSequence].Master, *sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons)

		return
	}

	// Function 2 Set Auto Color - Toggle the auto color feature.
	if X == common.Function2_Auto_Color && !this.Functions[this.TargetSequence][common.Function2_Auto_Color].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function2_Auto_Color On\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function2_Auto_Color].State = true

		cmd := common.Command{
			Action: common.UpdateAutoColor,
			Args: []common.Arg{
				{Name: "AutoColor", Value: true},
				{Name: "Type", Value: this.SelectedType},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}
	if X == common.Function2_Auto_Color && this.Functions[this.TargetSequence][common.Function2_Auto_Color].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function2_Auto_Color Off\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function2_Auto_Color].State = false

		cmd := common.Command{
			Action: common.UpdateAutoColor,
			Args: []common.Arg{
				{Name: "AutoColor", Value: false},
				{Name: "Type", Value: this.SelectedType},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}

	// Function 3 Set Auto Pattern - Toggle Auto Pattern.
	if X == common.Function3_Auto_Pattern && !this.Functions[this.TargetSequence][common.Function3_Auto_Pattern].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function3_Auto_Pattern On\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function3_Auto_Pattern].State = true

		cmd := common.Command{
			Action: common.UpdateAutoPattern,
			Args: []common.Arg{
				{Name: "AutoPattern", Value: true},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}
	if X == common.Function3_Auto_Pattern && this.Functions[this.TargetSequence][common.Function3_Auto_Pattern].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function3_Auto_Pattern Off\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function3_Auto_Pattern].State = false

		cmd := common.Command{
			Action: common.UpdateAutoPattern,
			Args: []common.Arg{
				{Name: "AutoPattern", Value: false},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}

	// Function 4 Bounce - Toggle bounce feature.
	if X == common.Function4_Bounce && !this.Functions[this.TargetSequence][common.Function4_Bounce].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function4_Bounce On \n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function4_Bounce].State = true

		cmd := common.Command{
			Action: common.UpdateBounce,
			Args: []common.Arg{
				{Name: "Bounce", Value: true},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}
	if X == common.Function4_Bounce && this.Functions[this.TargetSequence][common.Function4_Bounce].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function4_Bounce Off \n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function4_Bounce].State = false

		cmd := common.Command{
			Action: common.UpdateBounce,
			Args: []common.Arg{
				{Name: "Bounce", Value: false},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}

	// Map Function 5 RGB - Go straight into RGB color edit mode, don't wait for a another select press.
	if X == common.Function5_Color && !this.Functions[this.TargetSequence][common.Function5_Color].State &&
		sequences[this.TargetSequence].Type == "rgb" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function5_Color RGB Color Mode \n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.EditSequenceColorsMode = true
		this.Functions[this.TargetSequence][common.Function5_Color].State = true

		time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)

		// Set the colors.
		// Get an upto date copy of the sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		// Make sure the sequence colors holds the correct colors from the pattern steps.
		sequences[this.TargetSequence].SequenceColors = common.HowManyColorsInSteps(sequences[this.TargetSequence].Pattern.Steps)

		if debug {
			fmt.Printf("Map Function 5 RGB ====> sequences[%d].SequenceColors %+v\n", this.TargetSequence, sequences[this.TargetSequence].SequenceColors)
		}
		sequences[this.TargetSequence].CurrentColors = sequences[this.TargetSequence].SequenceColors

		// Also save the sequence colors in local represention, so if the user doesn't select any colors we can restore the existing colors.
		this.SavedSequenceColors[this.SelectedSequence] = sequences[this.SelectedSequence].SequenceColors

		// Remember which sequence we are editing
		this.EditWhichSequence = this.TargetSequence

		// We use the whole launchpad for choosing from 24 colors.
		common.HideAllSequences(commandChannels)

		// Show the colors
		ShowRGBColorPicker(this.MasterBrightness, *sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons)

		return
	}

	// Map Function 5 Scanner Color Selection - Go straight into scanner color edit mode via select fixture, don't wait for a another select press.
	if X == common.Function5_Color && !this.Functions[this.TargetSequence][common.Function5_Color].State &&
		sequences[this.TargetSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function5_Color Scanner Color Selection Mode \n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function5_Color].State = true
		this.EditScannerColorsMode = true

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)

		// Select a fixture.
		this.EditFixtureSelectionMode = true
		this.SelectMode[this.TargetSequence] = FUNCTION
		sequences[this.TargetSequence].StaticColors[X].FirstPress = false

		// Remember which sequence we are editing
		this.EditWhichSequence = this.TargetSequence

		this.FollowingAction = "ShowScannerColorSelectionButtons"
		this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.TargetSequence], this.DisplaySequence, this, eventsForLaunchpad, this.FollowingAction, guiButtons)
		return
	}

	// Function 6 RGB - Turn on edit static color mode.
	if X == common.Function6_Static_Gobo &&
		!this.Functions[this.TargetSequence][common.Function6_Static_Gobo].State &&
		sequences[this.TargetSequence].Type == "rgb" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function6_Static_Gobo RGB Static Color Mode On\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function6_Static_Gobo].State = true

		if this.ScannerChaser {
			// Turn on edit static color mode in the scanner sequence.
			this.EditStaticColorsMode[this.ScannerSequenceNumber] = true
		}

		// Starting a static sequence will turn off any running sequence, so turn off the start lamp
		common.LightLamp(common.ALight{X: X, Y: this.DisplaySequence, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
		//  and remember that this sequence is off.
		this.Running[this.TargetSequence] = false

		this.EditGoboSelectionMode = false                    // Turn off the other option for this function key.
		this.EditStaticColorsMode[this.TargetSequence] = true // Turn on edit static color mode.
		this.SelectMode[this.DisplaySequence] = NORMAL        // Turn off functions.

		// Go straight to static color selection mode, don't wait for a another select press.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		time.Sleep(250 * time.Millisecond) // But give the launchpad time to light the function key purple.
		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)

		// Switch on any static colors.
		cmd := common.Command{
			Action: common.UpdateStatic,
			Args: []common.Arg{
				{Name: "Static", Value: true},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		common.RevealSequence(this.TargetSequence, commandChannels)

		// Remember which sequence we are editing
		this.EditWhichSequence = this.TargetSequence

		return
	}
	// Function 6 RGB - Turn off edit static color mode.
	if X == common.Function6_Static_Gobo &&
		this.Functions[this.TargetSequence][common.Function6_Static_Gobo].State &&
		sequences[this.TargetSequence].Type == "rgb" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function6_Static_Gobo RGB Static Color Mode Off\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function6_Static_Gobo].State = false

		this.EditGoboSelectionMode = false                     // Turn off the other option for this function key.
		this.EditStaticColorsMode[this.TargetSequence] = false // Turn off edit static color mode.
		this.SelectMode[this.TargetSequence] = NORMAL          // Turn off function selection mode.

		if this.ScannerChaser {
			// Turn on edit static color mode in the scanner sequence.
			this.EditStaticColorsMode[this.ScannerSequenceNumber] = true
		}

		// Go straight to static color selection mode, don't wait for a another select press.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		time.Sleep(250 * time.Millisecond) // But give the launchpad time to light the function key purple.
		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)

		// Switch off static colors.
		cmd := common.Command{
			Action: common.UpdateStatic,
			Args: []common.Arg{
				{Name: "Static", Value: false},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		common.RevealSequence(this.TargetSequence, commandChannels)
		return
	}

	// Map Function 6 Scanner GOBO Selection - Go to select gobo mode if we are in scanner sequence.
	if X == common.Function6_Static_Gobo && !this.Functions[this.TargetSequence][common.Function6_Static_Gobo].State &&
		sequences[this.TargetSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function6_Static_Gobo RGB Scanner Gobo Selection Mode\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function6_Static_Gobo].State = true
		this.EditStaticColorsMode[this.TargetSequence] = false // Turn off the other option for this function key.
		this.EditGoboSelectionMode = true

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)

		// Select a fixture.
		this.EditFixtureSelectionMode = true
		this.SelectMode[this.TargetSequence] = FUNCTION
		sequences[this.TargetSequence].StaticColors[X].FirstPress = false

		this.FollowingAction = "ShowGoboSelectionButtons"
		this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.TargetSequence], this.DisplaySequence, this, eventsForLaunchpad, this.FollowingAction, guiButtons)
		return
	}

	// Function 7 - Turn on the RGB Invert mode.
	if X == common.Function7_Invert_Chase &&
		!this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State &&
		sequences[this.TargetSequence].Type == "rgb" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function7_Invert_Chase RGB Invert Mode On\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State = true

		cmd := common.Command{
			Action: common.UpdateRGBInvert,
			Args: []common.Arg{
				{Name: "RGBInvert", Value: true},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Turn on invert means all the fixture states should be inverted.
		for fixtureNumber := 0; fixtureNumber < sequences[this.TargetSequence].NumberFixtures; fixtureNumber++ {
			state := this.FixtureState[this.TargetSequence][fixtureNumber]
			state.RGBInverted = true
			this.FixtureState[this.TargetSequence][fixtureNumber] = state
		}

		if debug {
			for fixtureNumber := 0; fixtureNumber < sequences[this.TargetSequence].NumberFixtures; fixtureNumber++ {
				fmt.Printf("Seq%d: Fixture:%d This FixtureState: %+v\n", this.TargetSequence, fixtureNumber, this.FixtureState[this.TargetSequence][fixtureNumber])
			}
		}

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}

	// Function 7 - Turn off the RGB Invert mode.
	if X == common.Function7_Invert_Chase &&
		this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State &&
		sequences[this.TargetSequence].Type == "rgb" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function7_Invert_Chase RGB Invert Mode Off\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State = false

		cmd := common.Command{
			Action: common.UpdateRGBInvert,
			Args: []common.Arg{
				{Name: "RGBInvert", Value: false},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Turn off invert means all the fixture states should be not inverted.
		for fixtureNumber := 0; fixtureNumber < sequences[this.TargetSequence].NumberFixtures; fixtureNumber++ {
			state := this.FixtureState[this.TargetSequence][fixtureNumber]
			state.RGBInverted = false
			state.ScannerPatternReversed = false
			this.FixtureState[this.TargetSequence][fixtureNumber] = state
		}

		if debug {
			for fixtureNumber := 0; fixtureNumber < sequences[this.TargetSequence].NumberFixtures; fixtureNumber++ {
				fmt.Printf("Seq%d: Fixture:%d This FixtureState: %+v\n", this.TargetSequence, fixtureNumber, this.FixtureState[this.TargetSequence][fixtureNumber])
			}
		}

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		return
	}

	// Function 7 - Toggle the shutter chaser mode. Start the chaser.
	if X == common.Function7_Invert_Chase &&
		!this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State &&
		sequences[this.TargetSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function7_Invert_Chase Scanner Shutter Chaser Mode On\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.ScannerChaser = true
		this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State = true // Chaser

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

		// Tell the scannern & chaser sequences that the scanner shutter chase flag is on.
		cmd := common.Command{
			Action: common.UpdateScannerHasShutterChase,
			Args: []common.Arg{
				{Name: "ScannerHasShutterChase", Value: this.ScannerChaser},
			},
		}
		common.SendCommandToSequence(this.ScannerSequenceNumber, cmd, commandChannels)
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

		// Tell the chaser to start.
		cmd = common.Command{
			Action: common.StartChase,
		}
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

		// Update the buttons: speed
		common.LabelButton(0, 7, "Chase\nSpeed\nDown", guiButtons)
		common.LabelButton(1, 7, "Chase\nSpeed\nUp", guiButtons)

		common.LabelButton(2, 7, "Chase\nShift\nDown", guiButtons)
		common.LabelButton(3, 7, "Chase\nShift\nUp", guiButtons)

		common.LabelButton(4, 7, "Chase\nSize\nDown", guiButtons)
		common.LabelButton(5, 7, "Chase\nSize\nUp", guiButtons)

		common.LabelButton(6, 7, "Chase\nFase\nSoft", guiButtons)
		common.LabelButton(7, 7, "Chase\nFade\nSharp", guiButtons)

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		return
	}

	// Function 7 - Toggle the shutter chaser mode. Stop the chaser.
	if X == common.Function7_Invert_Chase &&
		this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State &&
		sequences[this.TargetSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function7_Invert_Chase Scanner Shutter Chaser Mode Off\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.ScannerChaser = false
		this.Functions[this.TargetSequence][common.Function7_Invert_Chase].State = false

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

		// Tell scanner & chaser sequence that the scanner shutter chase flag is off.
		this.Running[this.ChaserSequenceNumber] = false
		cmd := common.Command{
			Action: common.UpdateScannerHasShutterChase,
			Args: []common.Arg{
				{Name: "ScannerHasShutterChase", Value: this.ScannerChaser},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

		// Stop the chaser.
		cmd = common.Command{
			Action: common.StopChase,
		}
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

		// Update the buttons: speed
		common.LabelButton(0, 7, "Speed\nDown", guiButtons)
		common.LabelButton(1, 7, "Speed\nUp", guiButtons)

		common.LabelButton(2, 7, "Shift\nDown", guiButtons)
		common.LabelButton(3, 7, "Shift\nUp", guiButtons)

		common.LabelButton(4, 7, "Size\nDown", guiButtons)
		common.LabelButton(5, 7, "Size\nUp", guiButtons)

		common.LabelButton(6, 7, "Fase\nSoft", guiButtons)
		common.LabelButton(7, 7, "Fade\nSharp", guiButtons)

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		return
	}

	// Function 8 MUSIC TRIGGER  - Send start music trigger for scanner & rgb sequences.
	if X == common.Function8_Music_Trigger &&
		this.SelectMode[this.TargetSequence] != CHASER &&
		!this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function8_Music_Trigger Music Trigger Mode On\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State = true

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		time.Sleep(250 * time.Millisecond) // But give the launchpad time to light the function key purple.

		// Starting a music trigger will start the sequence, so turn on the start lamp
		// and remember that this sequence is on.
		this.Running[this.DisplaySequence] = true
		common.ShowRunningStatus(this.TargetSequence, this.Running, eventsForLaunchpad, guiButtons)

		// Start the music trigger for the target sequence.
		cmd := common.Command{
			Action: common.UpdateMusicTrigger,
			Args: []common.Arg{
				{Name: "MusicTriger", Value: true},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// We want to exit from functioms immediately so we call handle.
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		return
	}

	// Function 8 MUSIC TRIGGER  - Send stop music trigger for scanner and rgb sequences.
	if X == common.Function8_Music_Trigger &&
		this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function8_Music_Trigger Music Trigger Mode Off\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State = false

		this.Running[this.TargetSequence] = false
		common.ShowRunningStatus(this.TargetSequence, this.Running, eventsForLaunchpad, guiButtons)

		// Stop the music trigger for the target sequence.
		cmd := common.Command{
			Action: common.UpdateMusicTrigger,
			Args: []common.Arg{
				{Name: "MusicTriger", Value: false},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// We want to exit from functioms immediately so we call handle.
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		return
	}

	// Function 8 MUSIC TRIGGER  - Send stop music trigger chaser sequences.
	if X == common.Function8_Music_Trigger &&
		this.SelectMode[this.TargetSequence] != CHASER &&
		this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State {

		if debug {
			fmt.Printf("Seq%d: Mode:%d common.Function8_Music_Trigger Shutter Chaser Music Trigger Mode Off\n", this.TargetSequence, this.SelectMode[this.TargetSequence])
		}

		this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State = false
		this.ScannerChaser = false

		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)
		time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

		// Tell scanner & chaser sequence that the scanner shutter chase flag is off.
		this.Running[this.ChaserSequenceNumber] = this.ScannerChaser
		cmd := common.Command{
			Action: common.UpdateScannerHasShutterChase,
			Args: []common.Arg{
				{Name: "ScannerHasShutterChase", Value: this.ScannerChaser},
			},
		}
		common.SendCommandToSequence(this.DisplaySequence, cmd, commandChannels)
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Stop the music trigger for the chaser sequence.
		cmd = common.Command{
			Action: common.UpdateMusicTrigger,
			Args: []common.Arg{
				{Name: "MusicTriger", Value: this.ScannerChaser},
			},
		}
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

		// We want to exit from functioms immediately so we call handle.
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		return
	}
}
