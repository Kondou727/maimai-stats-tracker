package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	songdatadb "github.com/Kondou727/maimai-stats-tracker/internal/database/songdata"
	"github.com/pressly/goose"
)

const MAIMAI_SONGS_JSON_LINK = "https://maimai.sega.jp/data/maimai_songs.json"

type maimai_songs []struct {
	Artist     string `json:"artist"`
	Catcode    string `json:"catcode"`
	ImageURL   string `json:"image_url"`
	Release    string `json:"release"`
	LevBas     string `json:"lev_bas,omitempty"`
	LevAdv     string `json:"lev_adv,omitempty"`
	LevExp     string `json:"lev_exp,omitempty"`
	LevMas     string `json:"lev_mas,omitempty"`
	Sort       string `json:"sort"`
	Title      string `json:"title"`
	TitleKana  string `json:"title_kana"`
	Version    string `json:"version"`
	LevRemas   string `json:"lev_remas,omitempty"`
	DxLevBas   string `json:"dx_lev_bas,omitempty"`
	DxLevAdv   string `json:"dx_lev_adv,omitempty"`
	DxLevExp   string `json:"dx_lev_exp,omitempty"`
	DxLevMas   string `json:"dx_lev_mas,omitempty"`
	Key        string `json:"key,omitempty"`
	DxLevRemas string `json:"dx_lev_remas,omitempty"`
	Date       string `json:"date,omitempty"`
	LevUtage   string `json:"lev_utage,omitempty"`
	Kanji      string `json:"kanji,omitempty"`
	Comment    string `json:"comment,omitempty"`
	Buddy      string `json:"buddy,omitempty"`
}

// Fills the songdata.db
func (cfg *apiConfig) PopulateSongData() error {
	songs, err := pullOfficialJson()
	if err != nil {
		log.Printf("failed to pull official json")
		return err
	}
	for _, s := range songs {
		params := songdatadb.CreateSongParams{
			Artist:    s.Artist,
			Catcode:   s.Catcode,
			ImageUrl:  s.ImageURL,
			Release:   s.Release,
			LevBas:    sql.NullString{String: s.LevBas, Valid: true},
			LevAdv:    sql.NullString{String: s.LevAdv, Valid: true},
			LevExp:    sql.NullString{String: s.LevExp, Valid: true},
			LevMas:    sql.NullString{String: s.LevMas, Valid: true},
			Sort:      s.Sort,
			Title:     s.Title,
			TitleKana: s.TitleKana,
			Version:   s.Version,
			LevRemas:  sql.NullString{String: s.LevRemas, Valid: true},
			DxLevBas:  sql.NullString{String: s.DxLevBas, Valid: true},
			DxLevExp:  sql.NullString{String: s.DxLevExp, Valid: true},
			DxLevMas:  sql.NullString{String: s.DxLevMas, Valid: true},
			Date:      sql.NullString{String: s.Key, Valid: true},
			LevUtage:  sql.NullString{String: s.LevUtage, Valid: true},
			Comment:   sql.NullString{String: s.Comment, Valid: true},
			Buddy:     sql.NullString{String: s.Buddy, Valid: true},
		}
		err := cfg.songdataDBQueries.CreateSong(context.Background(), params)
		if err != nil {
			log.Printf("CreateSong failed")
			return err
		}
	}
	return nil

}

// Pulls the maimai_songs.json from the official website and marshals into a struct
func pullOfficialJson() (maimai_songs, error) {
	res, err := http.Get(MAIMAI_SONGS_JSON_LINK)
	if err != nil {
		log.Printf("get request to official maimai server failed")
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read get request body")
		return nil, err
	}

	var songs maimai_songs
	if err := json.Unmarshal(body, &songs); err != nil {
		log.Printf("failed to unmarshal maimai_songs.json")
		return nil, err
	}
	return songs, nil

}

func LoadSongdataDB() (*sql.DB, error) {
	songdataDB, err := sql.Open("sqlite", "songdata.db")
	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}
	if err := goose.Up(songdataDB, "sql/songdata/schema"); err != nil {
		return nil, err
	}
	return songdataDB, nil

}
