package labels

import (
	"errors"
	"os"

	"github.com/go-yaml/yaml"
)

type Label struct {
	Number int    `yaml:"number"`
	Name   string `yaml:"name"`
	Label  string `yaml:"label"`
}

type LabelEntry struct {
	Name   string  `yaml:"name"`
	Labels []Label `yaml:"labels"`
}

type LabelData struct {
	LabelEntrys []LabelEntry `yaml:"entrys"`
}

func LoadLabels() (labelData *LabelData, err error) {

	filename := "labels.yaml"

	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error: loading labels.yaml file: " + err.Error())
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: reading labels.yaml file: " + err.Error())
	}

	labelData = &LabelData{}
	err = yaml.Unmarshal(data, labelData)
	if err != nil {
		return nil, errors.New("error: unmarshalling labels config: " + err.Error())
	}
	return labelData, nil
}

func GetLabel(data *LabelData, name string, labelname string) string {

	for _, labelEntry := range data.LabelEntrys {
		if labelEntry.Name == name {
			// Found the right label entry. Now look for the matching text for this label.
			for _, label := range labelEntry.Labels {
				if label.Name == labelname {
					return label.Label
				}
			}
		}

	}

	return labelname + "\nNot\nFound"

}
