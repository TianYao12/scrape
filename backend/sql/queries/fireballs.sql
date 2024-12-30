-- name: CreateFireball :one

INSERT INTO fireballs (
    id,
    created_at,
    updated_at,
    total_radiated_energy, 
    feed_id
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFireballsForUser :many
SELECT fireballs.* 
FROM fireballs
JOIN feed_follows ON fireballs.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY fireballs.created_at DESC
LIMIT $2;
