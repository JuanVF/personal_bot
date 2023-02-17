CREATE TABLE personal_bot.t_currencies (
    id SERIAL,
    name TEXT NOT NULL
);

ALTER TABLE ONLY personal_bot.t_currencies ADD CONSTRAINT t_currencies_pkey PRIMARY KEY(id);