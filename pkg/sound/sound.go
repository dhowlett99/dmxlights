package sound

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/gordonklaus/portaudio"
)

func NewSoundTrigger(sequences []*common.Sequence, channels common.Channels) {

	go func() {

		fmt.Println("Starting Sound System")

		portaudio.Initialize()
		defer portaudio.Terminate()
		// Making the buffer bigger makes the music trigger have less latency.
		in := make([]int32, 128)
		stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
		if err != nil {
			fmt.Printf("failed to initialise portaudio\n")
		}

		// Start listening on the microphone input.
		stream.Start()
		if err != nil {
			fmt.Printf("failed to start stream\n")
		}

		defer stream.Close()

		for {
			stream.Read()
			if err != nil {
				fmt.Printf("failed to read audio stream\n")
			}

			if in[0] > 1000000000 {
				// Trigger
				time.Sleep(10 * time.Millisecond)
				cmd := common.Command{}
				channels.SoundTriggerChannels[0] <- cmd
				channels.SoundTriggerChannels[1] <- cmd
				channels.SoundTriggerChannels[2] <- cmd
				channels.SoundTriggerChannels[3] <- cmd
			}
		}
	}()
}
