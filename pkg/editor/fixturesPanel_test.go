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

import "testing"

func Test_checkDMXnumber(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name: "some text which is an error",
			args: args{
				value: "Hello",
			},
			wantErr: true,
		},
		{
			name: "single number",
			args: args{
				value: "100",
			},
			wantErr: false,
		},
		{
			name: "single number with text is an error",
			args: args{
				value: "1A00",
			},
			wantErr: true,
		},
		{
			name: "range of numbers",
			args: args{
				value: "100-200",
			},
			wantErr: false,
		},
		{
			name: "range of numbers with text in first nunber is an error",
			args: args{
				value: "10A0-200",
			},
			wantErr: true,
		},
		{
			name: "range of numbers with text in seconde number is an error",
			args: args{
				value: "100-20B0",
			},
			wantErr: true,
		},
		{
			name: "range of numbers with text in both numbers is an error",
			args: args{
				value: "10AA0-20B0",
			},
			wantErr: true,
		},
		{
			name: "range of numbers second number is less than first is an error",
			args: args{
				value: "100-2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDMXValue(tt.args.value)
			if err != nil && !tt.wantErr {
				t.Errorf("checkDMXValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
