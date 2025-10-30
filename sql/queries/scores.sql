-- name: CreateScore :one
INSERT INTO scores (song_name, chart_type, difficulty, achievement, genre, level)
VALUES (
    ?, ?, ?, ?, ?, ?
)
ON CONFLICT(song_name, difficulty) DO UPDATE SET
    achievement = excluded.achievement
RETURNING *;
