package main

import (
	"database/sql"
	"fmt"
	"log"

	scoresdb "github.com/Kondou727/maimai-stats-tracker/internal/database/scores"
	songdatadb "github.com/Kondou727/maimai-stats-tracker/internal/database/songdata"

	_ "modernc.org/sqlite"
)

type apiConfig struct {
	scoresDB          *sql.DB
	scoresDBQueries   *scoresdb.Queries
	songdataDB        *sql.DB
	songdataDBQueries *songdatadb.Queries
}

func main() {
	log.Printf("Starting maimai-stats-tracker")

	scoresDB, err := LoadScoresDB()
	if err != nil {
		log.Fatalf("Failed loading scores DB: %s", err)
	}
	defer scoresDB.Close()

	songdataDB, err := LoadSongdataDB()
	if err != nil {
		log.Fatalf("Failed loading song data DB: %s", err)
	}
	defer songdataDB.Close()

	scoresDBQueries := scoresdb.New(scoresDB)
	songdataDBQueries := songdatadb.New(songdataDB)

	cfg := apiConfig{
		scoresDB:          scoresDB,
		scoresDBQueries:   scoresDBQueries,
		songdataDB:        songdataDB,
		songdataDBQueries: songdataDBQueries,
	}
	/*
		err = cfg.loadTSV()
		if err != nil {
			log.Fatal(err)
		}

		err = cfg.PopulateSongData()
		if err != nil {
			log.Fatal(err)
		}

	*/
	err = cfg.pullJackets()
	if err != nil {
		log.Fatal(err)
	}

}

func (cfg *apiConfig) loadTSV() error {
	fmt.Print("Path to TSV file: ")
	var tsvPath string
	fmt.Scan(&tsvPath)
	if err := cfg.ImportScoresToDB(tsvPath, "tsv_file"); err != nil {
		return err
	}
	return nil
}
