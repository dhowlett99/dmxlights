// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights sound trigger system.
// Implemented by and depends on portaudio.
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
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/gordonklaus/portaudio"
)

const debug = false

const sampleRate = 44100

var gainSelected = 4
var gainCounters = make([]int, 10)

type SoundConfig struct {
	deviceName      string
	availableInputs []string
	stream          *portaudio.Stream
	SoundTriggers   []*common.Trigger
	gainSelected    int
	gainCounters    []int
	inputChannels   []*portaudio.HostApiInfo
	stopChannel     chan bool
}

func NewSoundTrigger(channels common.Channels, guiButtons chan common.ALight, eventsForLaunchpad chan common.ALight) *SoundConfig {

	soundConfig := SoundConfig{}
	soundConfig.stopChannel = make(chan bool)
	soundConfig.gainSelected = gainSelected
	soundConfig.gainCounters = gainCounters
	soundConfig.SoundTriggers = channels.SoundTriggers

	soundConfig.getAvailableInputs()
	soundConfig.StartSoundConfig("Built-in Microphone", guiButtons, eventsForLaunchpad)

	return &soundConfig

}

func (soundConfig *SoundConfig) StartSoundConfig(deviceName string, guiButtons chan common.ALight, eventsForLaunchpad chan common.ALight) {

	fmt.Printf("Starting Sound System Version %s\n", portaudio.VersionText())

	soundConfig.deviceName = deviceName

	go func() {

		err := portaudio.Initialize()
		if err != nil {
			fmt.Printf("error: portaudio: failed to initialise portaudio\n")
		}

		defer portaudio.Terminate()

		in := make([]float32, 128) // Making the buffer bigger makes the music trigger have less latency.
		out := make([]float32, 128)
		gain := []float32{0.05, 0.06, 0.07, 0.08, 0.09, 0.1, 0.11, 0.12, 0.13, 0.14, 0.15}

		if deviceName == "Built-in Microphone" {
			// Open the default input stream.
			soundConfig.stream, err = portaudio.OpenDefaultStream(1, 0, sampleRate, len(in), in)
			if err != nil {
				fmt.Printf("error: portaudio: failed to open default stream \n")
			}
		} else {
			inputChannels, err := portaudio.HostApis()
			if err != nil {
				fmt.Printf("error: portaudio: failed to list input channels \n")
			}
			for _, inputChannel := range inputChannels {
				fmt.Printf("New Input Channel %v\n", *inputChannel)

				for _, device := range inputChannel.Devices {
					if device.MaxInputChannels > 0 {
						if device.Name == deviceName {
							fmt.Printf("Found device %s\n", device.Name)
							p := portaudio.HighLatencyParameters(device, nil)
							fmt.Printf("Input.Channels %d\n", device.MaxInputChannels)
							p.Input.Channels = device.MaxInputChannels
							fmt.Printf("Output.Channels %d\n", device.MaxOutputChannels)
							p.Output.Channels = device.MaxOutputChannels
							fmt.Printf("SampleRate %f\n", device.DefaultSampleRate)
							p.SampleRate = device.DefaultSampleRate
							p.FramesPerBuffer = len(in)
							soundConfig.stream, err = portaudio.OpenStream(p, in)
							if err != nil {
								fmt.Printf("error: portaudio: failed to open stream %s\n", device.Name)
							}
						}
					}
				}
			}
		}

		// Start listening on the microphone input.
		soundConfig.stream.Start()
		if err != nil {
			fmt.Printf("error: portaudio: failed to start stream\n")
		}

		defer soundConfig.stream.Close()

		numSamples := 10

		// Start the thread that reports the gain.
		go soundConfig.gainChecker()

		for {
			// We need a way to shutdown the sound trigger subsystem when we switch
			// audio inputs in the settings dialog box.
			select {
			case <-soundConfig.stopChannel:
				return
			case <-time.After(1 * time.Millisecond):
			}

			// Read from the input stream.
			soundConfig.stream.Read()

			// Implenent a 800Hz low pass filter.
			cutoff := float32(800)

			// Now loop getting beats from portaudio.
			for i := 1; i < numSamples; i++ {

				out[i] = in[i-1] + soundConfig.filter(cutoff)*in[i] - in[i-1]

				// Tell the automatic gain control what level we're at.
				soundConfig.reportLevels(out[i])

				// Allow fine adjustment.
				actualGain := gain[gainSelected] + soundConfig.SoundTriggers[0].Gain

				if out[i] > actualGain {

					cmd := common.Command{}

					for triggerNumber, trigger := range soundConfig.SoundTriggers {
						if trigger.State {
							if debug {
								fmt.Printf("SOUND Trying to send to %s %d\n", trigger.Name, triggerNumber)
							}
							// Update status bar.
							common.LightLamp(common.Button{X: 0, Y: -1}, common.Magenta, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

							select {
							case soundConfig.SoundTriggers[triggerNumber].Channel <- cmd:

							case <-time.After(1000 * time.Millisecond):
								continue
							}

						}
					}
					// A short delay stop a sequnece being overwhelmed by trigger events.
					time.Sleep(time.Millisecond * 10)
					common.LightLamp(common.Button{X: 0, Y: -1}, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
				}
			}
		}
	}()
}

func (soundConfig *SoundConfig) GetDeviceName() string {
	return soundConfig.deviceName
}

// EnableSoundTrigger  - Register the Trigger.
func (soundConfig *SoundConfig) EnableSoundTrigger(name string) error {
	// Step through the existing sound triggers and find the one we want to enable.
	for triggerNumber, trigger := range soundConfig.SoundTriggers {
		if trigger.Name == name {
			//fmt.Printf("Enable trigger number %d name %s\n", triggerNumber, trigger.Name)
			soundConfig.SoundTriggers[triggerNumber].State = true
			return nil
		}
	}
	return fmt.Errorf("sound trigger %s not found", name)
}

// DisableSoundTrigger  - Disable the Trigger.
func (soundConfig *SoundConfig) DisableSoundTrigger(name string) error {

	// Step through the existing sound triggers and find the one we want to disable.
	for triggerNumber, trigger := range soundConfig.SoundTriggers {
		if trigger.Name == name {
			soundConfig.SoundTriggers[triggerNumber].State = false
			return nil
		}
	}
	return fmt.Errorf("sound trigger %s not found", name)
}

// GetSoundTriggerState  - What state is this trigger in ?
func (soundConfig *SoundConfig) GetSoundTriggerState(name string) bool {

	// Step through the existing sound triggers and find the one we want to find the status of.
	for triggerNumber, trigger := range soundConfig.SoundTriggers {
		if trigger.Name == name {
			if soundConfig.SoundTriggers[triggerNumber].State {
				return soundConfig.SoundTriggers[triggerNumber].State
			}
		}
	}
	return false
}

func (soundConfig *SoundConfig) getAvailableInputs() {
	// Fire up the audio subsystem just to find the number of audio inputs.
	err := portaudio.Initialize()
	if err != nil {
		fmt.Printf("error: portaudio: failed to initialise portaudio\n")
	}

	soundConfig.inputChannels, err = portaudio.HostApis()
	if err != nil {
		fmt.Printf("error: portaudio: failed to list input channels \n")
	}

	for _, inputChannel := range soundConfig.inputChannels {
		for _, device := range inputChannel.Devices {
			if device.MaxInputChannels > 0 {
				fmt.Printf("device %s\n", device.Name)
				soundConfig.availableInputs = append(soundConfig.availableInputs, device.Name)
			}
		}
	}

	portaudio.Terminate()
}

func (soundConfig *SoundConfig) StopSoundConfig() {

	if debug {
		fmt.Printf("Stop sound config\n")
	}
	// Send a signal for the current sound triggers to stop.
	soundConfig.stopChannel <- true

}

func (soundConfig *SoundConfig) GetSoundConfig() []string {
	if debug {
		fmt.Printf("sound config avail ins %s\n", soundConfig.availableInputs)
	}
	return soundConfig.availableInputs
}

func (soundConfig *SoundConfig) gainChecker() {
	for {
		timer1 := time.NewTimer(3 * time.Second)
		<-timer1.C

		if debug {
			fmt.Printf(">>>> I AM CHECKING THE GAIN \n")
		}
		// Calculate and the gain.
		gain := soundConfig.findGain(gainCounters)

		// Reset the counters.
		for index := range gainCounters {
			gainCounters[index] = 0
		}
		gainSelected = gain
	}
}

// findGain determine which counter has the largest value
// and returns the element number i.e. what gain.
func (soundConfig *SoundConfig) findGain(values []int) int {
	// Find minimum
	min := values[0]
	for _, v := range values {
		if v == 0 { // exlcude the empty counters to find peak.
			continue
		}
		if v < min {
			min = v
		}
	}

	// Find element
	for i, v := range values {
		if v == min {
			return i
		}
	}
	return 0
}

func (soundConfig *SoundConfig) reportLevels(level float32) {

	gain := []float32{
		// Peak  Gain Set.
		0.11, // 0.05
		0.12, // 0.06
		0.13, // 0.07
		0.14, // 0.08
		0.15, // 0.09
		0.16, // 0.10
		0.17, // 0.11
		0.18, // 0.12
		0.19, // 0.13
		0.20, // 0.14
		0.21, // 0.15
	}

	if level > gain[9] {
		gainCounters[9]++
	}
	if level > gain[8] {
		gainCounters[8]++
	}
	if level > gain[7] {
		gainCounters[7]++
	}
	if level > gain[6] {
		gainCounters[6]++
	}
	if level > gain[5] {
		gainCounters[5]++
	}
	if level > gain[4] {
		gainCounters[4]++
	}
	if level > gain[3] {
		gainCounters[3]++
	}
	if level > gain[2] {
		gainCounters[2]++
	}
	if level > gain[1] {
		gainCounters[1]++
	}
	if level > gain[0] {
		gainCounters[0]++
	}
}

func (soundConfig *SoundConfig) filter(cutofFreq float32) float32 {
	M_PI := float32(3.14159265358979323846264338327950288)
	RC := float32(1.0 / (cutofFreq * 2 * M_PI))
	dt := float32(1.0 / sampleRate) // SAMPLE_RATE
	alpha := dt / (RC + dt)

	return alpha
}
