-- name: CreateEvent :one
INSERT INTO events (
  event_id,
  event_topic,
  event_type,
  source,
  created_at,
  expires_at,
  status,
  retry_count,
  dedup_key,
  metadata,
  payload,
  target_service,
  correlation_id
) VALUES (
  $1, $2, $3, $4, 
  $5, $6, $7, $8, 
  $9, $10, $11, $12, $13
)
RETURNING *;

-- name: UpdateEventStatus :exec
UPDATE events SET status = $2 WHERE event_id = $1;

-- name: ReadEventByEventId :one
SELECT * FROM events WHERE event_id = $1;