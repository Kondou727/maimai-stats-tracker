package config

import (
	"database/sql"

	scoresdb "github.com/Kondou727/maimai-stats-tracker/internal/database/scores"
	songdatadb "github.com/Kondou727/maimai-stats-tracker/internal/database/songdata"
)

type ApiConfig struct {
	ScoresDB          *sql.DB
	ScoresDBQueries   *scoresdb.Queries
	SongdataDB        *sql.DB
	SongdataDBQueries *songdatadb.Queries
}
