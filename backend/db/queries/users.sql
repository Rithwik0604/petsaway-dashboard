-- name: GetUserByName :one
SELECT
    *
FROM
    users
WHERE
    username = ?;

-- name: InsertUserByName :one
INSERT INTO
    users (id, username, password_hash)
VALUES
    (?, ?, ?) RETURNING *;

-- name: UpdateUserByName :one
UPDATE users
SET
    id = ?,
    username = ?,
    password_hash = ? RETURNING *;