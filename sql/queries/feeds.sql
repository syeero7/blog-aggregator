-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetALLFeeds :many
SELECT
  f.name AS name,
  f.url AS url,
  u.name AS username,
  f.created_at as created_at,
  f.updated_at as updated_at
FROM feeds f
INNER JOIN users u ON f.user_id = u.id
ORDER BY f.created_at DESC;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = $2, updated_at = $3
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY last_fetched_at NULLS FIRST 
LIMIT 1;
