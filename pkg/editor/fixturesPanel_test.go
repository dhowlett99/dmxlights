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
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

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
			name: "string with a - in is an error",
			args: args{
				value: "1A-00",
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

func Test_removeEmptyActions(t *testing.T) {

	one := int16(1)
	two := int(2)

	type args struct {
		fixtureList []fixture.Fixture
	}
	tests := []struct {
		name string
		args args
		want []fixture.Fixture
	}{
		{
			name: "Test1 delete action named none.",
			args: args{
				fixtureList: []fixture.Fixture{
					{
						ID:          1,
						Name:        "Fixture1",
						Number:      1,
						Label:       "Label1",
						Description: "Description1",
						Type:        "Type1",
						Group:       1,
						Address:     1,
						Channels: []fixture.Channel{
							{
								Number:     1,
								Name:       "Name",
								Value:      &one,
								MaxDegrees: &two,
								Offset:     &two,
								Comment:    "Comment",
								Settings: []fixture.Setting{
									{
										Name:          "Name",
										Label:         "Label",
										Number:        1,
										Channel:       "Channel",
										Value:         "Value",
										SelectedValue: "SelectedValue",
									},
								},
							},
						},
						States: []fixture.State{
							{
								Name:        "State1",
								Number:      1,
								Label:       "Label1",
								ButtonColor: "ButtonColor",
								Master:      1,
								Actions: []fixture.Action{
									{
										Name:         "Action1",
										Number:       1,
										Mode:         "Chase",
										Colors:       []string{"Red", "Green", "Blue"},
										Map:          "Map",
										Fade:         "Fade",
										Size:         "Size",
										Speed:        "Speed",
										Rotate:       "Rotate",
										RotateSpeed:  "RotateSpeed",
										Program:      "Progran",
										ProgramSpeed: "ProgramSpeed",
										Strobe:       "Strobe",
										Gobo:         "Gobo",
										GoboSpeed:    "GoboSpeed",
									},
									{
										Name:         "Action2",
										Number:       2,
										Mode:         "None",
										Colors:       []string{"Red", "Green", "Blue"},
										Map:          "Map",
										Fade:         "Fade",
										Size:         "Size",
										Speed:        "Speed",
										Rotate:       "Rotate",
										RotateSpeed:  "RotateSpeed",
										Program:      "Progran",
										ProgramSpeed: "ProgramSpeed",
										Strobe:       "Strobe",
										Gobo:         "Gobo",
										GoboSpeed:    "GoboSpeed",
									},
								},
								Settings: []fixture.Setting{
									{
										Name:          "Name",
										Label:         "Label",
										Number:        1,
										Channel:       "Channel",
										Value:         "Value",
										SelectedValue: "SelectedValue",
									},
								},
								Flash: false,
							},
						},
						MultiFixtureDevice: false,
						NumberSubFixtures:  8,
						UseFixture:         "PAR1",
					},
				},
			},
			want: []fixture.Fixture{
				{
					Name:        "Fixture1",
					ID:          1,
					Label:       "Label1",
					Number:      1,
					Description: "Description1",
					Type:        "Type1",
					Group:       1,
					Address:     1,
					Channels: []fixture.Channel{
						{
							Number:     1,
							Name:       "Name",
							Value:      &one,
							MaxDegrees: &two,
							Offset:     &two,
							Comment:    "Comment",
							Settings: []fixture.Setting{
								{
									Name:          "Name",
									Label:         "Label",
									Number:        1,
									Channel:       "Channel",
									Value:         "Value",
									SelectedValue: "SelectedValue",
								},
							},
						},
					},
					States: []fixture.State{
						{
							Name:        "State1",
							Number:      1,
							Label:       "Label1",
							ButtonColor: "ButtonColor",
							Master:      1,
							Actions: []fixture.Action{
								{
									Name:         "Action1",
									Number:       1,
									Mode:         "Chase",
									Colors:       []string{"Red", "Green", "Blue"},
									Map:          "Map",
									Fade:         "Fade",
									Size:         "Size",
									Speed:        "Speed",
									Rotate:       "Rotate",
									RotateSpeed:  "RotateSpeed",
									Program:      "Progran",
									ProgramSpeed: "ProgramSpeed",
									Strobe:       "Strobe",
									Gobo:         "Gobo",
									GoboSpeed:    "GoboSpeed",
								},
							},
							Settings: []fixture.Setting{
								{
									Name:          "Name",
									Label:         "Label",
									Number:        1,
									Channel:       "Channel",
									Value:         "Value",
									SelectedValue: "SelectedValue",
								},
							},
							Flash: false,
						},
					},
					MultiFixtureDevice: false,
					NumberSubFixtures:  8,
					UseFixture:         "PAR1",
				},
			},
		},

		{
			name: "Test2 nothing to delete, should pass all through",
			args: args{
				fixtureList: []fixture.Fixture{
					{
						ID:          1,
						Name:        "Fixture1",
						Number:      1,
						Label:       "Label1",
						Description: "Description1",
						Type:        "Type1",
						Group:       1,
						Address:     1,
						Channels: []fixture.Channel{
							{
								Number:     1,
								Name:       "Name",
								Value:      &one,
								MaxDegrees: &two,
								Offset:     &two,
								Comment:    "Comment",
								Settings: []fixture.Setting{
									{
										Name:          "Name",
										Label:         "Label",
										Number:        1,
										Channel:       "Channel",
										Value:         "Value",
										SelectedValue: "SelectedValue",
									},
								},
							},
						},
						States: []fixture.State{
							{
								Name:        "State1",
								Number:      1,
								Label:       "Label1",
								ButtonColor: "ButtonColor",
								Master:      1,
								Actions: []fixture.Action{
									{
										Name:         "Action1",
										Number:       1,
										Mode:         "Chase",
										Colors:       []string{"Red", "Green", "Blue"},
										Map:          "Map",
										Fade:         "Fade",
										Size:         "Size",
										Speed:        "Speed",
										Rotate:       "Rotate",
										RotateSpeed:  "RotateSpeed",
										Program:      "Progran",
										ProgramSpeed: "ProgramSpeed",
										Strobe:       "Strobe",
										Gobo:         "Gobo",
										GoboSpeed:    "GoboSpeed",
									},
									{
										Name:         "Action2",
										Number:       2,
										Mode:         "Static",
										Colors:       []string{"Red", "Green", "Blue"},
										Map:          "Map",
										Fade:         "Fade",
										Size:         "Size",
										Speed:        "Speed",
										Rotate:       "Rotate",
										RotateSpeed:  "RotateSpeed",
										Program:      "Progran",
										ProgramSpeed: "ProgramSpeed",
										Strobe:       "Strobe",
										Gobo:         "Gobo",
										GoboSpeed:    "GoboSpeed",
									},
								},
								Settings: []fixture.Setting{
									{
										Name:          "Name",
										Label:         "Label",
										Number:        1,
										Channel:       "Channel",
										Value:         "Value",
										SelectedValue: "SelectedValue",
									},
								},
								Flash: false,
							},
						},
						MultiFixtureDevice: false,
						NumberSubFixtures:  8,
						UseFixture:         "PAR1",
					},
				},
			},
			want: []fixture.Fixture{
				{
					Name:        "Fixture1",
					ID:          1,
					Label:       "Label1",
					Number:      1,
					Description: "Description1",
					Type:        "Type1",
					Group:       1,
					Address:     1,
					Channels: []fixture.Channel{
						{
							Number:     1,
							Name:       "Name",
							Value:      &one,
							MaxDegrees: &two,
							Offset:     &two,
							Comment:    "Comment",
							Settings: []fixture.Setting{
								{
									Name:          "Name",
									Label:         "Label",
									Number:        1,
									Channel:       "Channel",
									Value:         "Value",
									SelectedValue: "SelectedValue",
								},
							},
						},
					},
					States: []fixture.State{
						{
							Name:        "State1",
							Number:      1,
							Label:       "Label1",
							ButtonColor: "ButtonColor",
							Master:      1,
							Actions: []fixture.Action{
								{
									Name:         "Action1",
									Number:       1,
									Mode:         "Chase",
									Colors:       []string{"Red", "Green", "Blue"},
									Map:          "Map",
									Fade:         "Fade",
									Size:         "Size",
									Speed:        "Speed",
									Rotate:       "Rotate",
									RotateSpeed:  "RotateSpeed",
									Program:      "Progran",
									ProgramSpeed: "ProgramSpeed",
									Strobe:       "Strobe",
									Gobo:         "Gobo",
									GoboSpeed:    "GoboSpeed",
								},
								{
									Name:         "Action2",
									Number:       2,
									Mode:         "Static",
									Colors:       []string{"Red", "Green", "Blue"},
									Map:          "Map",
									Fade:         "Fade",
									Size:         "Size",
									Speed:        "Speed",
									Rotate:       "Rotate",
									RotateSpeed:  "RotateSpeed",
									Program:      "Progran",
									ProgramSpeed: "ProgramSpeed",
									Strobe:       "Strobe",
									Gobo:         "Gobo",
									GoboSpeed:    "GoboSpeed",
								},
							},
							Settings: []fixture.Setting{
								{
									Name:          "Name",
									Label:         "Label",
									Number:        1,
									Channel:       "Channel",
									Value:         "Value",
									SelectedValue: "SelectedValue",
								},
							},
							Flash: false,
						},
					},
					MultiFixtureDevice: false,
					NumberSubFixtures:  8,
					UseFixture:         "PAR1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeEmptyActions(tt.args.fixtureList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeEmptyActions() got = %+v\n", got)
				t.Errorf("removeEmptyActions() want %+v\n", tt.want)
			}
		})
	}
}
