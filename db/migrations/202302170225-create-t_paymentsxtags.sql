CREATE TABLE personal_bot.t_payments_per_tags (
    id_payment INT NOT NULL,
    id_tag INT NOT NULL
);

ALTER TABLE ONLY personal_bot.t_payments_per_tags ADD CONSTRAINT t_payments_per_tags_userid_fkey FOREIGN KEY (id_payment) REFERENCES personal_bot.t_payments(id);
ALTER TABLE ONLY personal_bot.t_payments_per_tags ADD CONSTRAINT t_payments_per_tags_tagid_fkey FOREIGN KEY (id_tag) REFERENCES personal_bot.t_tags(id);