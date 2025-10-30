package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Kondou727/maimai-stats-tracker/internal/database"
	"github.com/pressly/goose"
	_ "modernc.org/sqlite"
)

const DBFILE = "scores.db"

type apiConfig struct {
	scoresDB        *sql.DB
	scoresDBQueries *database.Queries
}

func main() {
	scoresDB, err := LoadScoresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer scoresDB.Close()

	scoresDBQueries := database.New(scoresDB)
	cfg := apiConfig{
		scoresDB:        scoresDB,
		scoresDBQueries: scoresDBQueries,
	}
	fmt.Print("Enter path to json file: ")
	var jsonPath string
	fmt.Scan(&jsonPath)
	if err = cfg.ImportScoresToDB(jsonPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Score import success!")
}

func LoadScoresDB() (*sql.DB, error) {
	scoresDB, err := sql.Open("sqlite", DBFILE)
	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}
	if err := goose.Up(scoresDB, "sql/schema"); err != nil {
		return nil, err
	}
	return scoresDB, nil

}
