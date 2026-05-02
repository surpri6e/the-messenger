CREATE TABLE IF NOT EXISTS groups (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    owner_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    info TEXT,
    users_id INTEGER[] NOT NULL,
    admins_id INTEGER[] NOT NULL,
    enemies_id INTEGER[],
    created_at TEXT    
)