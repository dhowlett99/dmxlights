// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights main sequencer responsible for controlling all
// of the fixtures in a group.
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
