-- +goose Up
CREATE TABLE IF NOT EXISTS scores (
	song_name TEXT NOT NULL,
	chart_type INTEGER NOT NULL,
	difficulty INTEGER NOT NULL,
	achievement REAL NOT NULL,
	genre TEXT NOT NULL,
	level REAL NOT NULL,
	PRIMARY KEY (song_name, difficulty)
);

-- +goose Down
DROP TABLE scores;
