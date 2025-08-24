// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture mapping code, it is called to find fixtures and
// then sends messages to fixtures using the usb dmx library.
// Implemented and depends on usbdmx.
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

package fixture

import (
	"testing"

	"github.com/oliread/usbdmx/ft232"
)

func TestSetColorByNumber(t *testing.T) {

	type args struct {
		channel             Channel
		fixture             Fixture
		channelNumber       int
		color               int
		dmxController       *ft232.DMXController
		dmxInterfacePresent bool
	}
	tests := []struct {
		name  string
		args  args
		want  int16
		want1 int
	}{
		{
			name: "Color is 0",
			args: args{
				channel: Channel{
					Name: "Color",
					Settings: []Setting{
						{
							Number: 1,
							Name:   "White",
							Value:  "0",
						},
						{
							Number: 2,
							Name:   "Red",
							Value:  "8",
						},
						{
							Number: 3,
							Name:   "Green",
							Value:  "16",
						},
						{
							Number: 4,
							Name:   "Blue",
							Value:  "32",
						},
					},
				},
				fixture: Fixture{
					Address: 100,
					Channels: []Channel{
						{
							Name: "Color",
						},
					},
				},
				channelNumber:       0,
				color:               0,
				dmxController:       nil,
				dmxInterfacePresent: false,
			},
			want:  100,
			want1: 0,
		},
		{
			name: "Color is 1 want Red",
			args: args{
				channel: Channel{
					Name: "Color",
					Settings: []Setting{
						{
							Number: 1,
							Name:   "White",
							Value:  "0",
						},
						{
							Number: 2,
							Name:   "Red",
							Value:  "8",
						},
						{
							Number: 3,
							Name:   "Green",
							Value:  "16",
						},
						{
							Number: 4,
							Name:   "Blue",
							Value:  "32",
						},
					},
				},
				fixture: Fixture{
					Address: 100,
				},
				channelNumber:       0,
				color:               1,
				dmxController:       nil,
				dmxInterfacePresent: false,
			},
			want:  100,
			want1: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SetColorByNumber(tt.args.channel, tt.args.fixture, tt.args.channelNumber, tt.args.color, tt.args.dmxController, tt.args.dmxInterfacePresent)
			if got != tt.want {
				t.Errorf("SetColorByNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SetColorByNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
