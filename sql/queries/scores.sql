-- name: CreateScore :one
INSERT INTO scores (song_name, chart_type, difficulty, achievement, fc_ap, sync, dx_star, dx_percent)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
ON CONFLICT(song_name, chart_type, difficulty) DO UPDATE SET
    achievement = excluded.achievement,
    fc_ap = excluded.fc_ap,
    sync = excluded.sync,
    dx_star = excluded.dx_star,
    dx_percent = excluded.dx_percent

RETURNING *;
