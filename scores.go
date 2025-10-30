package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Kondou727/maimai-stats-tracker/internal/database"
	"github.com/gocarina/gocsv"
)

type RawScores struct {
	SongName    string `csv:"Song"`
	ChartType   string `csv:"Chart"`
	Difficulty  string `csv:"Difficulty"`
	Achievement string `csv:"Achv"`
	FCAP        string `csv:"FC/AP"`
	Sync        string `csv:"Sync"`
	DXStar      int    `csv:"DX ✦"`
	DXPercent   string `csv:"DX %"`
}

func ScoresTSVToStruct(tsvPath string) ([]RawScores, error) {
	tsvFile, err := os.Open(tsvPath)
	if err != nil {
		return nil, err
	}
	csvreader := csv.NewReader(tsvFile)
	csvreader.Comma = '\t'
	csvreader.LazyQuotes = true
	var rs []RawScores
	if err := gocsv.UnmarshalCSV(csvreader, &rs); err != nil {
		return nil, err
	}

	return rs, nil
}

func (cfg *apiConfig) ImportScoresToDB(input string, format string) error {
	log.Print("processing tsv..")
	scores, err := ScoresTSVToStruct(input)
	if err != nil {
		return err
	}

	for _, s := range scores {
		achievementInt, err := percentToInt(s.Achievement)
		if err != nil {
			return err
		}

		dxpercentInt, err := percentToInt(s.DXPercent)
		if err != nil {
			return err
		}

		_, err = cfg.scoresDBQueries.CreateScore(context.Background(), database.CreateScoreParams{
			SongName:    s.SongName,
			ChartType:   s.ChartType,
			Difficulty:  s.Difficulty,
			Achievement: int64(achievementInt),
			FcAp:        s.FCAP,
			Sync:        s.Sync,
			DxStar:      int64(s.DXStar),
			DxPercent:   int64(dxpercentInt),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func percentToInt(s string) (int, error) {
	return strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(s, "%", ""), ".", ""))
}
