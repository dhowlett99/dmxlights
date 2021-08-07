package sound

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/gordonklaus/portaudio"
)

func NewSoundTrigger(trigger []chan common.Command) {

	go func() {
		fmt.Println("Starting Sound System")

		portaudio.Initialize()
		defer portaudio.Terminate()
		in := make([]int32, 64)
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
			//fmt.Printf("Read Sound\n")
			if err != nil {
				fmt.Printf("failed to read audio stream\n")
			}

			if in[0] > 1000000000 {
				if !fired {
					fmt.Printf("TRIGGER\n")
					time.Sleep(100 * time.Millisecond)
					cmd := common.Command{
						// Start: true,
					}
					for seq := range trigger {
						trigger[seq] <- cmd
					}

					fired = false
				}
				// fired = true
			}
		}
	}()
}
