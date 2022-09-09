CREATE TABLE IF NOT EXISTS "shared_transactions" (
  "shared_tx_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "tx_id" int NOT NULL REFERENCES transactions (tx_id),
  "user_id" int NOT NULL REFERENCES users (user_id),
  "shared_with" int NOT NULL REFERENCES users (user_id),
  "amount" float,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);