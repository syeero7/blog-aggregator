-- name: CreatePost :exec
INSERT INTO posts (title, description, url, created_at, updated_at, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetPostsForUser :many
SELECT p.* FROM posts p
INNER JOIN feed_follows ff ON p.feed_id = ff.feed_id
INNER JOIN users u ON ff.user_id = u.id 
WHERE u.name = $1
ORDER BY p.published_at DESC
LIMIT $2;
