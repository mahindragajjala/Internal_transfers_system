CREATE TABLE IF NOT EXISTS accounts (
    account_id BIGINT PRIMARY KEY,
    balance NUMERIC(20, 5) NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    source_account_id BIGINT REFERENCES accounts(account_id),
    destination_account_id BIGINT REFERENCES accounts(account_id),
    amount NUMERIC(20, 5) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
