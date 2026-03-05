-- +goose Up
CREATE TABLE accounts(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    balance NUMERIC(15, 2) DEFAULT 0.00,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE accounts;
