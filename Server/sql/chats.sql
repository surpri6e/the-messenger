CREATE TABLE IF NOT EXISTS chat (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    first_person_id INTEGER NOT NULL,
    second_person_id INTEGER NOT NULL,
    created_at TEXT
)
