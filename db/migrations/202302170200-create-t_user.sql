CREATE TABLE personal_bot.t_users (
    id SERIAL,
    name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    google_me TEXT NOT NULL,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

ALTER TABLE ONLY personal_bot.t_users ADD CONSTRAINT t_users_pkey PRIMARY KEY (id);