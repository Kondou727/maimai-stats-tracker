-- +goose Up
CREATE TABLE IF NOT EXISTS songdata(
	artist TEXT NOT NULL,
	catcode TEXT NOT NULL,
	image_url TEXT NOT NULL,
	release TEXT NOT NULL,
	lev_bas TEXT,
	lev_adv TEXT,
	lev_exp TEXT,
	lev_mas TEXT,
	sort TEXT NOT NULL,
	title TEXT NOT NULL,
	title_kana TEXT NOT NULL,
	version TEXT NOT NULL,
	lev_remas TEXT,
	dx_lev_bas TEXT,
	dx_lev_adv TEXT,
	dx_lev_exp TEXT,
	dx_lev_mas TEXT,
	dx_lev_remas TEXT,
	date TEXT,
	lev_utage TEXT,
	kanji TEXT,
	comment TEXT,
	buddy TEXT,
	PRIMARY KEY(artist, title)
);

-- +goose Down
DROP TABLE songdata;
