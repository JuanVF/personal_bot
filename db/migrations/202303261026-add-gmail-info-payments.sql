ALTER TABLE ONLY personal_bot.t_payments ADD COLUMN gmail_id TEXT;
ALTER TABLE ONLY personal_bot.t_payments ADD COLUMN description TEXT;

CREATE INDEX t_payments_gmail_index ON personal_bot.t_payments USING btree (gmail_id);