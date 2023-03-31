/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
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