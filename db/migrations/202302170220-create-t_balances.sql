CREATE TABLE personal_bot.t_balances (
    id SERIAL,
    amount NUMERIC(15, 2) NOT NULL,
    expenses NUMERIC(15, 2) NOT NULL,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    user_id INT NOT NULL,
    currency_id INT NOT NULL
);

ALTER TABLE ONLY personal_bot.t_balances ADD CONSTRAINT t_balances_pkey PRIMARY KEY(id);
ALTER TABLE ONLY personal_bot.t_balances ADD CONSTRAINT t_balances_userid_fkey FOREIGN KEY (user_id) REFERENCES personal_bot.t_users(id);
ALTER TABLE ONLY personal_bot.t_balances ADD CONSTRAINT t_balances_currencyid_fkey FOREIGN KEY (currency_id) REFERENCES personal_bot.t_currencies(id);