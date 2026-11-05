// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is configuration for the context sensitive bottom
// buttons and their lables on the bottom status bar.
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
	"image/color"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/labels"
)

type bottonButton struct {
	Label string
	Color color.RGBA
}

type ButtonConfig struct {
	BlankButtons          []bottonButton
	RGBButtons            []bottonButton
	ScannerButtons        []bottonButton
	ChaserButtons         []bottonButton
	DerbyButtons          []bottonButton
	ProjectorButtons      []bottonButton
	SwitchSettingsButtons []bottonButton
	SwitchStaticsButtons  []bottonButton
	SwitchControlButtons  []bottonButton
	SwitchChaseButtons    []bottonButton
}

func ConfigureButtons(this *CurrentState) ButtonConfig {

	buttonConfig := ButtonConfig{}

	// Storage for the rgb buttons on the bottom row.
	var RGBButtons []bottonButton
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Down"), Color: colors.Cyan})
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Up"), Color: colors.Cyan})
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shift", "Down"), Color: colors.Cyan})
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shift", "Up"), Color: colors.Cyan})
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Size", "Down"), Color: colors.Cyan})
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Size", "Up"), Color: colors.Cyan})
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Fade", "Soft"), Color: colors.Cyan})
	RGBButtons = append(RGBButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Fade", "Sharp"), Color: colors.Cyan})
	buttonConfig.RGBButtons = RGBButtons

	// Storage for the scanner buttons on the bottom row.
	var ScannerButtons []bottonButton
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Down"), Color: colors.Cyan})
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Up"), Color: colors.Cyan})
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shift", "Down"), Color: colors.Cyan})
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shift", "Up"), Color: colors.Cyan})
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Size", "Down"), Color: colors.Cyan})
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Size", "Up"), Color: colors.Cyan})
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Coord", "Down"), Color: colors.Cyan})
	ScannerButtons = append(ScannerButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Coord", "Up"), Color: colors.Cyan})
	buttonConfig.ScannerButtons = ScannerButtons

	// Storage for chaser buttons on the bottom row.
	var ChaserButtons []bottonButton
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Speed", "Down"), Color: colors.Cyan})
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Speed", "Up"), Color: colors.Cyan})
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Shift", "Down"), Color: colors.Cyan})
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Shift", "Up"), Color: colors.Cyan})
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Size", "Down"), Color: colors.Cyan})
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Size", "Up"), Color: colors.Cyan})
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Fade", "Soft"), Color: colors.Cyan})
	ChaserButtons = append(ChaserButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Chase Fade", "Sharp"), Color: colors.Cyan})
	buttonConfig.ChaserButtons = ChaserButtons

	// Storage for derby buttons on the bottom row.
	var DerbyButtons []bottonButton
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Down"), Color: colors.Cyan})
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Up"), Color: colors.Cyan})
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Rotate", "Down"), Color: colors.Cyan})
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Rotate", "Up"), Color: colors.Cyan})
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Down"), Color: colors.Cyan})
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Up"), Color: colors.Cyan})
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Gobo", "Down"), Color: colors.Cyan})
	DerbyButtons = append(DerbyButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Gobo", "Up"), Color: colors.Cyan})
	buttonConfig.DerbyButtons = DerbyButtons

	// Storage for projector buttons on the bottom row.
	var ProjectorButtons []bottonButton
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Down"), Color: colors.Cyan})
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Up"), Color: colors.Cyan})
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Rotate", "Down"), Color: colors.Cyan})
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Rotate", "Up"), Color: colors.Cyan})
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Down"), Color: colors.Cyan})
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Up"), Color: colors.Cyan})
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Gobo", "Down"), Color: colors.Cyan})
	ProjectorButtons = append(ProjectorButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Gobo", "Up"), Color: colors.Cyan})
	buttonConfig.ProjectorButtons = ProjectorButtons

	// Storage for the rgb labels on the bottom row.
	var SwitchSettingsButtons []bottonButton
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Down"), Color: colors.Cyan})
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Speed", "Up"), Color: colors.Cyan})
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shift", "Down"), Color: colors.Cyan})
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shift", "Up"), Color: colors.Cyan})
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Size", "Down"), Color: colors.Cyan})
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Size", "Up"), Color: colors.Cyan})
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Fade", "Soft"), Color: colors.Cyan})
	SwitchSettingsButtons = append(SwitchSettingsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Fade", "Sharp"), Color: colors.Cyan})
	buttonConfig.SwitchSettingsButtons = SwitchSettingsButtons

	// Storage for the switch label buttons in off mode on the bottom row.
	var BlankButtons []bottonButton
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	BlankButtons = append(BlankButtons, bottonButton{Label: "", Color: colors.Cyan})
	buttonConfig.SwitchChaseButtons = BlankButtons

	// Storage for the switch buttons in static mode on the bottom row.
	var SwitchStaticsButtons []bottonButton
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Down"), Color: colors.Cyan})
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Up"), Color: colors.Cyan})
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchStaticsButtons = append(SwitchStaticsButtons, bottonButton{Label: "", Color: colors.Cyan})
	buttonConfig.SwitchStaticsButtons = SwitchStaticsButtons

	// Storage for the switch buttons in chase mode on the bottom row.
	var SwitchChaseButtons []bottonButton
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shutter Speed", "Down"), Color: colors.Cyan})
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Shutter Speed", "Up"), Color: colors.Cyan})
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Rotate Speed", "Down"), Color: colors.Cyan})
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Rotate Speed", "Up"), Color: colors.Cyan})
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Down"), Color: colors.Cyan})
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Color", "Up"), Color: colors.Cyan})
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Gobo", "Down"), Color: colors.Cyan})
	SwitchChaseButtons = append(SwitchChaseButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Gobo", "Up"), Color: colors.Cyan})
	buttonConfig.SwitchChaseButtons = SwitchChaseButtons

	// Storage for the switch buttons on the bottom row.
	var SwitchControlButtons []bottonButton
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Program Speed", "Down"), Color: colors.Cyan})
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Program Speed", "Up"), Color: colors.Cyan})
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Program Number", "Down"), Color: colors.Cyan})
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: labels.GetLabel(this.Labels, "Program Number", "Up"), Color: colors.Cyan})
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: "", Color: colors.Cyan})
	SwitchControlButtons = append(SwitchControlButtons, bottonButton{Label: "", Color: colors.Cyan})
	buttonConfig.SwitchControlButtons = SwitchControlButtons

	return buttonConfig
}
