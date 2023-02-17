CREATE TABLE personal_bot.t_tags (
    id SERIAL,
    name TEXT NOT NULL
);

ALTER TABLE ONLY personal_bot.t_tags ADD CONSTRAINT t_tags_pkey PRIMARY KEY(id);