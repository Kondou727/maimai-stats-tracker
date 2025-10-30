package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Kondou727/maimai-stats-tracker/internal/database"
)

type RawScores struct {
	SongName    string  `json:"songName"`
	ChartType   int     `json:"chartType"`
	Difficulty  int     `json:"difficulty"`
	Achievement float64 `json:"achievement"`
	Genre       string  `json:"genre"`
	Level       float64 `json:"level"`
}

func ScoresJSONToStruct(jsonPath string) ([]RawScores, error) {
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}
	var rs []RawScores
	if err := json.Unmarshal(jsonData, &rs); err != nil {
		return nil, err
	}

	return rs, nil
}

func (cfg *apiConfig) ImportScoresToDB(jsonPath string) error {
	scores, err := ScoresJSONToStruct(jsonPath)
	if err != nil {
		return err
	}
	for _, s := range scores {
		_, err = cfg.scoresDBQueries.CreateScore(context.Background(), database.CreateScoreParams{
			SongName:    s.SongName,
			ChartType:   int64(s.ChartType),
			Difficulty:  int64(s.Difficulty),
			Achievement: s.Achievement,
			Genre:       s.Genre,
			Level:       s.Level,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
