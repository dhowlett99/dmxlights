// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file controls which labels are shown at the top and bottom of the window.
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
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_checkStateActions(t *testing.T) {
	type args struct {
		state common.State
	}
	tests := []struct {
		name string
		args args
		want common.ColorDisplayControl
	}{
		{
			name: "Check for Off Mode",
			args: args{
				state: common.State{
					Actions: []common.Action{
						{
							Mode: "Off",
						},
					},
				},
			},
		},
		{
			name: "Check for Static Mode",
			args: args{
				state: common.State{
					Actions: []common.Action{
						{
							Mode: "Static",
							Colors: []string{
								"Red",
								"Green",
								"Blue",
							},
						},
					},
				},
			},
			want: common.ColorDisplayControl{
				Red:   true,
				Green: true,
				Blue:  true,
			},
		},
		{
			name: "Check for Chase Mode",
			args: args{
				state: common.State{
					Actions: []common.Action{
						{
							Mode: "Chase",
							Colors: []string{
								"Red",
								"Green",
								"Blue",
							},
						},
					},
				},
			},
			want: common.ColorDisplayControl{
				Red:   true,
				Green: true,
				Blue:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStateActions(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkStateActions() = %v, want %v", got, tt.want)
			}
		})
	}
}
