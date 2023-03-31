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
CREATE TABLE personal_bot.t_walmart_invoice (
    id serial PRIMARY KEY,
    user_id int REFERENCES personal_bot.t_users(id),
    amount numeric(10,2),
    invoice_date timestamp with time zone NOT NULL,
    items_purchased int,
    gmail_id text NOT NULL,
);

CREATE TABLE personal_bot.t_ingredients (
    id serial PRIMARY KEY,
    walmart_invoice_id int REFERENCES personal_bot.t_walmart_invoice(id),
    name text NOT NULL
);

CREATE TABLE personal_bot.t_diet_plans (
    id serial PRIMARY KEY,
    user_id int REFERENCES personal_bot.t_users(id),
    name text NOT NULL,
    description text NOT NULL,
    meal_plan text NOT NULL,
    warning text NOT NULL,
    creation_date timestamp with time zone NOT NULL
);

CREATE INDEX idx_t_walmart_invoice_user_id ON personal_bot.t_walmart_invoice (user_id);
CREATE INDEX idx_t_ingredients_walmart_invoice_id ON personal_bot.t_ingredients (walmart_invoice_id);
CREATE INDEX idx_t_walmart_invoice_gmail_id ON personal_bot.t_walmart_invoice (gmail_id);