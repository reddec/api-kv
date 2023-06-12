-- name: Put :exec
INSERT INTO value (namespace, key, value, expire_at, content_type)
VALUES (?, ?, ?, ?, ?) ON CONFLICT (namespace, key) DO
UPDATE SET expire_at = excluded.expire_at, value = excluded.value, content_type = excluded.content_type;

-- name: Get :one
SELECT value, expire_at, content_type
FROM value
WHERE namespace = ? AND key = ? AND (expire_at IS NULL OR expire_at > unixepoch()*1000);

-- TODO: unixepoch to unixepoch('subsec') once modernc/sqlite will switch to 3.42

-- name: ListStart :many
SELECT *
FROM value_key
WHERE namespace = ?
  AND (expire_at IS NULL OR expire_at > unixepoch() * 1000)
ORDER BY id LIMIT ?;

-- name: ListNext :many
SELECT *
FROM value_key
WHERE namespace = ?
  AND id > ?
  AND (expire_at IS NULL OR expire_at > unixepoch() * 1000)
ORDER BY id LIMIT ?;

-- name: Delete :exec
DELETE
FROM value
WHERE namespace = ? AND key = ?;

-- name: DeleteNamespace :exec
DELETE
FROM value
WHERE namespace = ?;

-- name: DeleteExpired :exec
DELETE
FROM value
WHERE expire_at < unixepoch() * 1000;