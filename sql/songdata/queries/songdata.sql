-- name: CreateSong :exec
INSERT INTO songdata (artist, catcode, image_url, release, lev_bas, lev_adv, lev_exp, lev_mas, sort, title, title_kana, version, lev_remas, dx_lev_bas, dx_lev_adv, dx_lev_exp, dx_lev_mas, dx_lev_remas, date, lev_utage, kanji, comment, buddy)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
ON CONFLICT (artist, title) DO NOTHING
;
