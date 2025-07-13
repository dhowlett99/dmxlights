// Copyright (C) 2022, 2023 dhowlett99.
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
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

func TestClearNoneActions(t *testing.T) {
	type args struct {
		statesList []fixture.State
	}
	tests := []struct {
		name string
		args args
		want []fixture.State
	}{
		{
			name: "test removal of None action in states list",
			args: args{
				statesList: []fixture.State{
					{
						Actions: []fixture.Action{
							{
								Mode:   "Off",
								Number: 1,
							},
							{
								Mode:   "None",
								Number: 2,
							},
						},
					},
				},
			},
			want: []fixture.State{
				{
					Actions: []fixture.Action{
						{
							Mode:   "Off",
							Number: 1,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearNoneActions(tt.args.statesList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClearNoneActions() = %v, want %v", got, tt.want)
			}
		})
	}
}
