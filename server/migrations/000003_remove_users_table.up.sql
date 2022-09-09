ALTER TABLE "plaid_items" DROP COLUMN "user_id";
ALTER TABLE "plaid_items" ADD COLUMN "user_id" varchar NOT NULL;

ALTER TABLE "transactions" DROP COLUMN "user_id";
ALTER TABLE "transactions" ADD COLUMN "user_id" varchar NOT NULL;

ALTER TABLE "friends" DROP COLUMN "user_id_1";
ALTER TABLE "friends" ADD COLUMN "user_id_1" varchar NOT NULL;

ALTER TABLE "friends" DROP COLUMN "user_id_2";
ALTER TABLE "friends" ADD COLUMN "user_id_2" varchar NOT NULL;

ALTER TABLE "shared_transactions" DROP COLUMN "user_id";
ALTER TABLE "shared_transactions" ADD COLUMN "user_id" varchar NOT NULL;

ALTER TABLE "shared_transactions" DROP COLUMN "shared_with";
ALTER TABLE "shared_transactions" ADD COLUMN "shared_with" varchar NOT NULL;

DROP TABLE IF EXISTS users;