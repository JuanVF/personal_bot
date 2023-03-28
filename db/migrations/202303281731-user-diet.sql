CREATE TABLE personal_bot.t_walmart_invoice (
    id serial PRIMARY KEY,
    user_id int REFERENCES personal_bot.t_users(id),
    amount numeric(10,2),
    invoice_date timestamp with time zone NOT NULL,
    items_purchased int
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
    macronutrients text NOT NULL,
    micronutrients text NOT NULL,
    meal_plan text NOT NULL,
    food_restriction text,
    notes text NOT NULL,
    creation_date timestamp with time zone NOT NULL
);

CREATE INDEX idx_t_walmart_invoice_user_id ON personal_bot.t_walmart_invoice (user_id);
CREATE INDEX idx_t_ingredients_walmart_invoice_id ON personal_bot.t_ingredients (walmart_invoice_id);