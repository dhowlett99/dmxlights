package sound

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/gordonklaus/portaudio"
)

func NewSoundTrigger(soundTriggers []*common.Trigger, channels common.Channels) {

	go func() {

		fmt.Printf("Starting Sound System Version %s\n", portaudio.VersionText())

		err := portaudio.Initialize()
		if err != nil {
			fmt.Printf("error: portaudio: failed to initialise portaudio\n")
		}

		defer portaudio.Terminate()
		// Making the buffer bigger makes the music trigger have less latency.
		in := make([]int32, 128)
		stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
		if err != nil {
			fmt.Printf("error: portaudio: failed to open default stream \n")
		}

		// Start listening on the microphone input.
		stream.Start()
		if err != nil {
			fmt.Printf("error: portaudio: failed to start stream\n")
		}

		defer stream.Close()

		for {
			stream.Read()
			if err != nil {
				fmt.Printf("error: portaudio: failed to read audio stream\n")
			}

			//if in[0] > 1000000000 {
			if in[0] > 10000000 {
				// Trigger
				time.Sleep(10 * time.Millisecond)
				cmd := common.Command{}
				for index, trigger := range soundTriggers {
					if trigger.SequenceNumber == index {
						if trigger.State {
							channels.SoundTriggerChannels[index] <- cmd
						}
					}
				}
			}
		}
	}()
}
