CREATE TABLE IF NOT EXISTS "users" (
  "user_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "email" varchar UNIQUE NOT NULL,
  "password" bytea NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "plaid_items" DROP COLUMN "user_id";
ALTER TABLE "plaid_items" ADD COLUMN "user_id" int NOT NULL REFERENCES users (user_id),;

ALTER TABLE "transactions" DROP COLUMN "user_id";
ALTER TABLE "transactions" ADD COLUMN "user_id" int REFERENCES users (user_id);

ALTER TABLE "friends" DROP COLUMN "user_id_1";
ALTER TABLE "friends" ADD COLUMN "user_id_1" int NOT NULL REFERENCES users (user_id);

ALTER TABLE "friends" DROP COLUMN "user_id_2";
ALTER TABLE "friends" ADD COLUMN "user_id_2" int NOT NULL REFERENCES users (user_id);

ALTER TABLE "shared_transactions" DROP COLUMN "user_id";
ALTER TABLE "shared_transactions" ADD COLUMN "user_id" int NOT NULL REFERENCES users (user_id);

ALTER TABLE "shared_transactions" DROP COLUMN "shared_with";
ALTER TABLE "shared_transactions" ADD COLUMN "shared_with" int NOT NULL REFERENCES users (user_id);