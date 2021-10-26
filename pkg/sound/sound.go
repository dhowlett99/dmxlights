package sound

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100

func NewSoundTrigger(soundTriggers []*common.Trigger, channels common.Channels) {

	go func() {

		gain := []float32{0.05, 0.06, 0.07, 0.08, 0.09, .1, .11, .12, .13, .14}

		fmt.Printf("Starting Sound System Version %s\n", portaudio.VersionText())

		err := portaudio.Initialize()
		if err != nil {
			fmt.Printf("error: portaudio: failed to initialise portaudio\n")
		}

		defer portaudio.Terminate()
		// Making the buffer bigger makes the music trigger have less latency.
		in := make([]float32, 128)
		out := make([]float32, 128)
		stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(in), in)
		if err != nil {
			fmt.Printf("error: portaudio: failed to open default stream \n")
		}

		// Start listening on the microphone input.
		stream.Start()
		if err != nil {
			fmt.Printf("error: portaudio: failed to start stream\n")
		}

		defer stream.Close()

		numSamples := 100
		gainSelected := 4
		for {
			stream.Read()
			if err != nil {
				fmt.Printf("error: portaudio: failed to read audio stream\n")
			}

			// Implenent a 800Hz low pass filter.
			cutoff := float32(800)

			for i := 1; i < numSamples; i++ {
				out[i] = in[i-1] + filter(cutoff)*in[i] - in[i-1]

				// fmt.Printf("gain selected is %d\n", gainSelected)
				// fmt.Printf("sound: current gain is %f\n", gain[gainSelected])
				if out[i] > gain[gainSelected] {
					cmd := common.Command{}
					for index, trigger := range soundTriggers {
						gainSelected = trigger.Gain
						if trigger.SequenceNumber == index {
							if trigger.State {
								channels.SoundTriggerChannels[index] <- cmd
							}
						}
					}
				}
			}
		}
	}()
}

func filter(cutofFreq float32) float32 {
	M_PI := float32(3.14159265358979323846264338327950288)
	RC := float32(1.0 / (cutofFreq * 2 * M_PI))
	dt := float32(1.0 / sampleRate) // SAMPLE_RATE
	alpha := dt / (RC + dt)

	return alpha
}
