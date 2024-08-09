package sequence

import (
	"errors"
	"os"

	"github.com/go-yaml/yaml"
)

// LoadSequences loads sequence configuration information.
// Each sequence has a :-
//
//	name: sequence name,  a singe word.
//	description: free text describing the sequence.
//	group: assignes to one of the top 4 rows of the launchpad. 1-4
//	type:  rgb, scanner or switch
func LoadSequences() (sequences *SequencesConfig, err error) {
	filename := "sequences.yaml"

	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error: loading sequences.yaml file: " + err.Error())
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: reading sequences.yaml file: " + err.Error())
	}

	sequences = &SequencesConfig{}
	err = yaml.Unmarshal(data, sequences)
	if err != nil {
		return nil, errors.New("error: unmarshalling sequences config: " + err.Error())
	}
	return sequences, nil
}
