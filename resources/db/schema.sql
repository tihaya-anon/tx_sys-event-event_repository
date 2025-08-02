-- dialogue: PostgreSQL
CREATE TYPE delivery_status AS ENUM (
    'DELIVERY_STATUS_UNSPECIFIED',
    'PENDING',
    'RETRYING',
    'DELIVERED',
    'DEAD'
);

CREATE TABLE events (
    event_id TEXT PRIMARY KEY,
    event_topic TEXT NOT NULL,
    event_type TEXT NOT NULL,
    source TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    expires_at BIGINT,
    status delivery_status NOT NULL DEFAULT 'DELIVERY_STATUS_UNSPECIFIED',
    retry_count INTEGER DEFAULT 0,
    dedup_key TEXT,
    metadata JSONB,
    payload TEXT, -- Base64 encoded
    target_service TEXT,
    correlation_id TEXT
);
