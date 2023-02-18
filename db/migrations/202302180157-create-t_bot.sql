CREATE TABLE personal_bot.t_bots (
    id SERIAL,
    user_id INT NOT NULL,
    last_thread_id TEXT
);

ALTER TABLE ONLY personal_bot.t_bots ADD CONSTRAINT t_bots_pkey PRIMARY KEY(id);
ALTER TABLE ONLY personal_bot.t_bots ADD CONSTRAINT t_bots_userid_fkey FOREIGN KEY (user_id) REFERENCES personal_bot.t_users(id);