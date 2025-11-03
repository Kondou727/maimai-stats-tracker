-- name: CreateSong :exec
INSERT INTO songdata (
    id, title, artist, genre, img, release, version, is_dx, diff, level, const, is_utage, is_buddy
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
ON CONFLICT (id, diff, is_dx, is_utage) DO NOTHING;


-- name: ReturnAllJackets :many
SELECT img FROM songdata;
