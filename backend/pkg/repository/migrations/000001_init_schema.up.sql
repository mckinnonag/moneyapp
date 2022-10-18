CREATE TABLE IF NOT EXISTS "users" (
  "user_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "email" varchar UNIQUE NOT NULL,
  "password" bytea NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "plaid_items" (
    "item_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "user_id" int NOT NULL REFERENCES users (user_id),
    "access_token" varchar NOT NULL,
    "plaid_item_id" varchar NOT NULL, 
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "currency_codes" (
    "currency_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "code" varchar (3) UNIQUE,
    "currency_name" varchar
);

CREATE TABLE IF NOT EXISTS "transactions" (
  "tx_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "plaid_id" int NOT NULL,
  "plaid_item_id" int REFERENCES plaid_items (item_id),
  "user_id" int REFERENCES users (user_id),
  "category" text [],
  "location" varchar,
  "tx_name" varchar,
  "amount" float,
  "iso_currency_code" int REFERENCES currency_codes (currency_id),
  "tx_date" timestamptz,
  "pending" boolean,
  "merchant_name" varchar,
  "payment_channel" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "friends" (
    "user_id_1" int NOT NULL REFERENCES users (user_id),
    "user_id_2" int NOT NULL REFERENCES users (user_id),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_currency_code_code ON currency_codes(code);
CREATE INDEX IF NOT EXISTS idx_friends_user1 ON friends(user_id_1);
CREATE INDEX IF NOT EXISTS idx_friends_user2 ON friends(user_id_2);