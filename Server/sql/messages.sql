CREATE TABLE IF NOT EXISTS messages (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
    where_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    status TEXT NOT NULL,
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    is_changed BOOLEAN NOT NULL DEFAULT FALSE,
    is_forwarded BOOLEAN NOT NULL DEFAULT FALSE,
    type TEXT NOT NULL,
    file_link TEXT NOT NULL
)
