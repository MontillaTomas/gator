-- name: CreatePost :one
INSERT INTO posts (feed_id, title, description, url, published_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT
    posts.id,
    posts.title,
    posts.description,
    posts.url,
    posts.published_at
FROM posts
JOIN feeds ON posts.feed_id = feeds.id
JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;