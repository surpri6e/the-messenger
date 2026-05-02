CREATE TABLE IF NOT EXISTS users (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    username TEXT NOT NULL,
    theme TEXT NOT NULL,
    info TEXT NOT NULL,
    avatar_link TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    is_online BOOLEAN DEFAULT FALSE,
    last_seen TIMESTAMP NOT NULL,
    is_email_accepted BOOLEAN DEFAULT FALSE,
    contacts TEXT[],
    is_muted_chats_id INTEGER[],
    is_pinned_chats INTEGER[] NOT NULL
)