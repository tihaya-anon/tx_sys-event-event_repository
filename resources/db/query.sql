-- name: UpdateEventStatus :one
UPDATE events SET status = $2 WHERE event_id = $1
RETURNING *;