CREATE TABLE IF NOT EXISTS wallet (
    wallet_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    balance REAL,
    is_approved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS wallet_fall (
    fall_id UUID PRIMARY KEY,
    wallet_id UUID REFERENCES wallet(wallet_id),
    balance REAL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);
