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

package sound

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/gordonklaus/portaudio"
)

func Test_findLargest(t *testing.T) {

	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				values: []int{119, 7, 0, 0},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				values: []int{12342, 7293, 4930, 3378, 2364, 1661, 1124, 732, 489, 309},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			soundConfig := SoundConfig{}
			if got := soundConfig.findGain(tt.args.values); got != tt.want {
				t.Errorf("findGain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoundConfig_EnableSoundTrigger(t *testing.T) {

	type fields struct {
		deviceName      string
		availableInputs []string
		stream          *portaudio.Stream
		SoundTriggers   []*common.Trigger
		gainSelected    int
		gainCounters    []int
		inputChannels   []*portaudio.HostApiInfo
		stopChannel     chan bool
	}
	type args struct {
		name         string
		channel      chan common.Command
		switchNumber int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []common.Trigger
	}{
		{
			name: "enable first trigger",
			args: args{
				name:         "switch1",
				channel:      nil,
				switchNumber: 1,
			},
			fields: fields{
				SoundTriggers: []*common.Trigger{
					{
						Name:  "sequence0",
						State: true,
					},
					{
						Name:  "sequence1",
						State: true,
					},
					{
						Name:  "sequence2",
						State: true,
					},
					{
						Name:  "switch0",
						State: false,
					},
					{
						Name:  "switch1",
						State: false,
					},
				},
			},
			want: []common.Trigger{
				{Name: "sequence0", State: true, Gain: 0, Channel: nil},
				{Name: "sequence1", State: true, Gain: 0, Channel: nil},
				{Name: "sequence2", State: true, Gain: 0, Channel: nil},
				{Name: "switch0", State: false, Gain: 0, Channel: nil},
				{Name: "switch1", State: true, Gain: 0, Channel: nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			soundConfig := &SoundConfig{
				deviceName:      tt.fields.deviceName,
				availableInputs: tt.fields.availableInputs,
				stream:          tt.fields.stream,
				SoundTriggers:   tt.fields.SoundTriggers,
				gainSelected:    tt.fields.gainSelected,
				gainCounters:    tt.fields.gainCounters,
				inputChannels:   tt.fields.inputChannels,
				stopChannel:     tt.fields.stopChannel,
			}

			soundConfig.EnableSoundTrigger(tt.args.name)

			// Resolve the pointers.
			triggers := []common.Trigger{}
			for _, trigger := range soundConfig.SoundTriggers {
				triggers = append(triggers, *trigger)
			}

			if !reflect.DeepEqual(triggers, tt.want) {
				t.Errorf("SoundConfig.RegisterSoundTrigger() got = %+v, want %+v", triggers, tt.want)
			}
		})
	}
}

func TestSoundConfig_DisableSoundTrigger(t *testing.T) {

	type fields struct {
		deviceName      string
		availableInputs []string
		stream          *portaudio.Stream
		SoundTriggers   []*common.Trigger
		gainSelected    int
		gainCounters    []int
		inputChannels   []*portaudio.HostApiInfo
		stopChannel     chan bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []common.Trigger
	}{
		{
			name: "disable switch1 trigger",
			args: args{
				name: "switch1",
			},
			fields: fields{
				SoundTriggers: []*common.Trigger{
					{
						Name:  "sequence0",
						State: true,
					},
					{
						Name:  "sequence1",
						State: true,
					},
					{
						Name:  "sequence2",
						State: true,
					},
					{
						Name:  "switch0",
						State: false,
					},
					{
						Name:  "switch1",
						State: true,
					},
				},
			},
			want: []common.Trigger{
				{Name: "sequence0", State: true, Gain: 0},
				{Name: "sequence1", State: true, Gain: 0},
				{Name: "sequence2", State: true, Gain: 0},
				{Name: "switch0", State: false, Gain: 0},
				{Name: "switch1", State: false, Gain: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			soundConfig := &SoundConfig{
				deviceName:      tt.fields.deviceName,
				availableInputs: tt.fields.availableInputs,
				stream:          tt.fields.stream,
				SoundTriggers:   tt.fields.SoundTriggers,
				gainSelected:    tt.fields.gainSelected,
				gainCounters:    tt.fields.gainCounters,
				inputChannels:   tt.fields.inputChannels,
				stopChannel:     tt.fields.stopChannel,
			}

			soundConfig.DisableSoundTrigger(tt.args.name)

			// Resolve the pointers.
			triggers := []common.Trigger{}
			for _, trigger := range soundConfig.SoundTriggers {
				triggers = append(triggers, *trigger)
			}

			if !reflect.DeepEqual(triggers, tt.want) {
				t.Errorf("SoundConfig.RegisterSoundTrigger() got = %+v, want %+v", triggers, tt.want)
			}

		})
	}
}
