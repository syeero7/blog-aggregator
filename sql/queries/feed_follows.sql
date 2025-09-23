-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
INSERT INTO feed_follows (feed_id, user_id, created_at, updated_at) 
VALUES ($1, $2, $3, $4) RETURNING *)
SELECT
  ff.*,
  feeds.name AS feed_name,
  users.name AS username
FROM inserted_feed_follow ff
INNER JOIN feeds ON ff.feed_id = feeds.id
INNER JOIN users ON ff.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT
  ff.*,
  feeds.name AS feed_name,
  (SELECT name FROM users WHERE feeds.user_id = id) AS creator
FROM feed_follows ff
INNER JOIN feeds ON ff.feed_id = feeds.id
INNER JOIN users ON ff.user_id = users.id
WHERE users.name = $1;

