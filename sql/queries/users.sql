-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: DeleteFeedFollowByUserAndURL :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.feed_id = feeds.id
  AND feed_follows.user_id = $1
  AND feeds.url = $2;