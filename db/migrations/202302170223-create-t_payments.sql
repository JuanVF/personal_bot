CREATE TABLE personal_bot.t_payments (
    id SERIAL,
    amount NUMERIC(15, 2) NOT NULL,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    user_id INT NOT NULL,
    currency_id INT NOT NULL,
    dolar_price NUMERIC(15, 2) NOT NULL,
    tags JSON NOT NULL
);

ALTER TABLE ONLY personal_bot.t_payments ADD CONSTRAINT t_payments_pkey PRIMARY KEY(id);
ALTER TABLE ONLY personal_bot.t_payments ADD CONSTRAINT t_payments_userid_fkey FOREIGN KEY (user_id) REFERENCES personal_bot.t_users(id);
ALTER TABLE ONLY personal_bot.t_payments ADD CONSTRAINT t_payments_currencyid_fkey FOREIGN KEY (currency_id) REFERENCES personal_bot.t_currencies(id);