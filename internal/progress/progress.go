package progress

import (
	"encoding/json"
	"os"
)

type Progress struct {
	Language  string   `json:"language"`
	Chapters  []string `json:"chapters_completed"`
	Exercises []string `json:"exercises_completed"`
}

const saveFile = "progress.json"

func LoadProgress() (Progress, error) {
	var p Progress
	file, err := os.ReadFile(saveFile)
	if err != nil {
		if os.IsNotExist(err) {
			return Progress{}, nil
		}
		return p, err
	}

	err = json.Unmarshal(file, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SaveProgress(p Progress) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(saveFile, data, 0644)
}

func (p *Progress) AddChapter(chapter string) {
	for _, c := range p.Chapters {
		if c == chapter {
			return
		}
	}
	p.Chapters = append(p.Chapters, chapter)
}

func (p *Progress) AddExercise(ex string) {
	for _, e := range p.Exercises {
		if e == ex {
			return
		}
	}
	p.Exercises = append(p.Exercises, ex)
}

