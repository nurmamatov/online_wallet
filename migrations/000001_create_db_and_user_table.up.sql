CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    phone VARCHAR(20) NOT NULL,
    is_approved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);