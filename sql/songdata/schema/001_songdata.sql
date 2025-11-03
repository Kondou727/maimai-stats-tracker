-- +goose Up
CREATE TABLE IF NOT EXISTS songdata(
    id TEXT NOT NULL,
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    genre TEXT NOT NULL,
    img TEXT NOT NULL,
    release TEXT NOT NULL,
    version TEXT NOT NULL,
    is_dx BOOLEAN NOT NULL,
    diff TEXT NOT NULL,
    level TEXT NOT NULL,
    const STRING NOT NULL,
    is_utage BOOLEAN NOT NULL,
    is_buddy TEXT,
	PRIMARY KEY(id, diff, is_dx, is_utage)
);

-- +goose Down
DROP TABLE songdata;
