package sound

import (
	"fmt"
	"reflect"
	"testing"
	"time"

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

func TestSoundConfig_RegisterSoundTrigger(t *testing.T) {

	type fields struct {
		deviceName       string
		availableInputs  []string
		stream           *portaudio.Stream
		BPMChannel       chan bool
		SoundTriggers    map[int]*common.Trigger
		gainSelected     int
		gainCounters     []int
		inputChannels    []*portaudio.HostApiInfo
		stopChannel      chan bool
		BPMtimer         *time.Timer
		BPMcounter       int
		BPMactualCounter int
		BPMsecondUp      bool
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
		want1  map[int]common.Trigger
	}{
		{
			name: "add first trigger",
			args: args{
				name:         "switch1",
				channel:      nil,
				switchNumber: 1,
			},
			fields: fields{
				SoundTriggers: map[int]*common.Trigger{
					0: {
						Name:  "sequence0",
						State: true,
					},
					1: {
						Name:  "sequence1",
						State: true,
					},
					2: {
						Name:  "sequence2",
						State: true,
					},
				},
			},
			want1: map[int]common.Trigger{
				0:  {Name: "sequence0", State: true, Gain: 0, BPM: 0},
				1:  {Name: "sequence1", State: true, Gain: 0, BPM: 0},
				2:  {Name: "sequence2", State: true, Gain: 0, BPM: 0},
				11: {Name: "switch1", State: true, Gain: 0, BPM: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			soundConfig := &SoundConfig{
				deviceName:       tt.fields.deviceName,
				availableInputs:  tt.fields.availableInputs,
				stream:           tt.fields.stream,
				BPMChannel:       tt.fields.BPMChannel,
				SoundTriggers:    tt.fields.SoundTriggers,
				gainSelected:     tt.fields.gainSelected,
				gainCounters:     tt.fields.gainCounters,
				inputChannels:    tt.fields.inputChannels,
				stopChannel:      tt.fields.stopChannel,
				BPMtimer:         tt.fields.BPMtimer,
				BPMcounter:       tt.fields.BPMcounter,
				BPMactualCounter: tt.fields.BPMactualCounter,
				BPMsecondUp:      tt.fields.BPMsecondUp,
			}

			soundConfig.RegisterSoundTrigger(tt.args.name, tt.args.channel, tt.args.switchNumber)

			if !reflect.DeepEqual(soundConfig.SoundTriggers, tt.want1) {
				t.Errorf("SoundConfig.RegisterSoundTrigger() got = %+v, want %+v", soundConfig.SoundTriggers, tt.want1)
			}
		})
	}
}

func TestSoundConfig_DeRegisterSoundTrigger(t *testing.T) {

	type fields struct {
		deviceName       string
		availableInputs  []string
		stream           *portaudio.Stream
		BPMChannel       chan bool
		SoundTriggers    map[int]*common.Trigger
		gainSelected     int
		gainCounters     []int
		inputChannels    []*portaudio.HostApiInfo
		stopChannel      chan bool
		BPMtimer         *time.Timer
		BPMcounter       int
		BPMactualCounter int
		BPMsecondUp      bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  []common.Trigger
	}{
		{
			name: "add first trigger",
			args: args{
				name: "switch1",
			},
			fields: fields{
				SoundTriggers: map[int]*common.Trigger{
					0: {
						Name:  "sequence0",
						State: true,
					},
					1: {
						Name:  "sequence1",
						State: true,
					},
					2: {
						Name:  "sequence2",
						State: true,
					},
					3: {
						Name:  "switch1",
						State: true,
					},
				},
			},
			want1: []common.Trigger{
				{Name: "sequence0", State: true, Gain: 0, BPM: 0},
				{Name: "sequence1", State: true, Gain: 0, BPM: 0},
				{Name: "sequence2", State: true, Gain: 0, BPM: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			soundConfig := &SoundConfig{
				deviceName:       tt.fields.deviceName,
				availableInputs:  tt.fields.availableInputs,
				stream:           tt.fields.stream,
				BPMChannel:       tt.fields.BPMChannel,
				SoundTriggers:    tt.fields.SoundTriggers,
				gainSelected:     tt.fields.gainSelected,
				gainCounters:     tt.fields.gainCounters,
				inputChannels:    tt.fields.inputChannels,
				stopChannel:      tt.fields.stopChannel,
				BPMtimer:         tt.fields.BPMtimer,
				BPMcounter:       tt.fields.BPMcounter,
				BPMactualCounter: tt.fields.BPMactualCounter,
				BPMsecondUp:      tt.fields.BPMsecondUp,
			}

			soundConfig.DeRegisterSoundTrigger(tt.args.name)

			triggers := []common.Trigger{}
			for _, trigger := range soundConfig.SoundTriggers {
				fmt.Printf("Trigger %+v\n", trigger)
				triggers = append(triggers, *trigger)
			}

			if !reflect.DeepEqual(triggers, tt.want1) {
				t.Errorf("SoundConfig.RegisterSoundTrigger() got = %+v, want %+v", triggers, tt.want1)
			}

		})
	}
}
