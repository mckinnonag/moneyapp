ALTER TABLE "transactions" DROP COLUMN "iso_currency_code";
ALTER TABLE "transactions" ADD COLUMN "iso_currency_code" int REFERENCES currency_codes (currency_id);