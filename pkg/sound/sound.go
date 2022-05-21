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

func NewSoundTrigger(soundTriggers []*common.Trigger, channels common.Channels) {

	go func() {

		gain := []float32{0.05, 0.06, 0.07, 0.08, 0.09, 0.1, 0.11, 0.12, 0.13, 0.14, 0.15}

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

		numSamples := 10

		// Start the thread that reports the gain.
		go gainChecker()

		for {
			stream.Read()
			if err != nil {
				fmt.Printf("error: portaudio: failed to read audio stream\n")
			}

			// Implenent a 800Hz low pass filter.
			cutoff := float32(800)

			// Now loop getting beats from portaudio.
			for i := 1; i < numSamples; i++ {
				out[i] = in[i-1] + filter(cutoff)*in[i] - in[i-1]

				// Tell the automatic gain control what level we're at.
				reportLevels(out[i], soundTriggers)

				// Allow fine adjustment.
				actualGain := gain[gainSelected] + soundTriggers[0].Gain

				if out[i] > actualGain {
					cmd := common.Command{}
					for index, trigger := range soundTriggers {
						if trigger.SequenceNumber == index {
							if trigger.State {
								if debug {
									fmt.Printf("%d: Index %d Gain %f   State %t  \n", trigger.SequenceNumber, index, trigger.Gain, trigger.State)
								}
								channels.SoundTriggerChannels[index] <- cmd
							}
						}
					}
					// A short delay stop a sequnece being overwhelmed by trigger events.
					time.Sleep(time.Millisecond * 10)
				}
			}
		}
	}()
}

func gainChecker() {
	for {
		timer1 := time.NewTimer(3 * time.Second)
		<-timer1.C

		if debug {
			fmt.Printf(">>>> I AM CHECKING THE GAIN \n")
		}
		// Calculate and the gain.
		gain := findGain(gainCounters)

		// Reset the counters.
		for index := range gainCounters {
			gainCounters[index] = 0
		}
		gainSelected = gain
	}
}

// findGain determine which counter has the largest value
// and returns the element number i.e. what gain.
func findGain(values []int) int {
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

func reportLevels(level float32, soundTriggers []*common.Trigger) {

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

func filter(cutofFreq float32) float32 {
	M_PI := float32(3.14159265358979323846264338327950288)
	RC := float32(1.0 / (cutofFreq * 2 * M_PI))
	dt := float32(1.0 / sampleRate) // SAMPLE_RATE
	alpha := dt / (RC + dt)

	return alpha
}
