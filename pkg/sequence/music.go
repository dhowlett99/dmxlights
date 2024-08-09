package sequence

import (
	"fmt"
	"os"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/sound"
)

func enableMusicTrigger(sequence *common.Sequence, soundConfig *sound.SoundConfig) {

	sequence.CurrentSpeed = time.Duration(12 * time.Hour)
	err := soundConfig.EnableSoundTrigger(sequence.Name)
	if err != nil {
		fmt.Printf("Error while trying to enable sound trigger %s\n", err.Error())
		os.Exit(1)
	}
	if debug {
		fmt.Printf("Sound trigger %s enabled \n", sequence.Name)
	}
	sequence.ChangeMusicTrigger = false
}

func disableMusicTrigger(sequence *common.Sequence, soundConfig *sound.SoundConfig) {

	err := soundConfig.DisableSoundTrigger(sequence.Name)
	if err != nil {
		fmt.Printf("Error while trying to disable sound trigger %s\n", err.Error())
		os.Exit(1)
	}
	if debug {
		fmt.Printf("Sound trigger %s disabled\n", sequence.Name)
	}
	sequence.CurrentSpeed = common.SetSpeed(sequence.Speed)
	sequence.ChangeMusicTrigger = false
}
