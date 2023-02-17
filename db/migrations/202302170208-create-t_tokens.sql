CREATE TABLE personal_bot.t_tokens (
    id SERIAL,
    refresh_token TEXT NOT NULL,
    last_token TEXT NOT NULL,
    expires_in TEXT NOT NULL,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    user_id INT NOT NULL
);

ALTER TABLE ONLY personal_bot.t_tokens ADD CONSTRAINT t_tokens_pkey PRIMARY KEY(id);
ALTER TABLE ONLY personal_bot.t_tokens ADD CONSTRAINT t_tokens_userid_fkey FOREIGN KEY (user_id) REFERENCES personal_bot.t_users(id);