package sound

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/gordonklaus/portaudio"
)

type Sound struct {
	SendSoundToSequence1 bool
	SendSoundToSequence2 bool
	SendSoundToSequence3 bool
	SendSoundToSequence4 bool
}

func NewSoundTrigger(trigger []chan common.Command) *Sound {

	s := Sound{}
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

		var fired bool
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
				if !fired {
					// Trigger
					time.Sleep(10 * time.Millisecond)
					cmd := common.Command{
						// Start: true,
					}
					if s.SendSoundToSequence1 {
						trigger[0] <- cmd
					}
					if s.SendSoundToSequence2 {
						trigger[1] <- cmd
					}
					if s.SendSoundToSequence3 {
						trigger[2] <- cmd
					}
					if s.SendSoundToSequence4 {
						trigger[3] <- cmd
					}
					fired = false
				}
			}
		}
	}()
	return &s
}

func (s *Sound) SetSoundTrigger(seq int) {
	if seq == 1 {
		s.SendSoundToSequence1 = true
	}
	if seq == 1 {
		s.SendSoundToSequence2 = true
	}
	if seq == 1 {
		s.SendSoundToSequence3 = true
	}
	if seq == 1 {
		s.SendSoundToSequence4 = true
	}
}
