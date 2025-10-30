-- +goose Up
CREATE TABLE IF NOT EXISTS scores (
	song_name TEXT NOT NULL,
	chart_type TEXT NOT NULL,
	difficulty TEXT NOT NULL,
	achievement INTEGER NOT NULL,
	fc_ap STRING NOT NULL,
	sync STRING NOT NULL,
	dx_star INTEGER NOT NULL,
	dx_percent INTEGER NOT NULL,
	PRIMARY KEY (song_name, chart_type, difficulty)
);

-- +goose Down
DROP TABLE scores;
