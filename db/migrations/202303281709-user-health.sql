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
CREATE TABLE personal_bot.t_health_conditions (
    id serial PRIMARY KEY,
    name text,
    description text
);

CREATE TABLE personal_bot.t_activity_levels (
    id serial PRIMARY KEY,
    name text,
    description text
);

CREATE TABLE personal_bot.t_user_health (
    id serial PRIMARY KEY,
    user_id int REFERENCES personal_bot.t_users(id),
    health_condition_id int REFERENCES personal_bot.t_health_conditions(id),
    diagnosis_date timestamp with time zone NOT NULL,
    treatment text,
    discharged_date timestamp with time zone NOT NULL
);

CREATE INDEX idx_t_health_conditions_name ON personal_bot.t_health_conditions (name);
CREATE INDEX idx_t_activity_levels_name ON personal_bot.t_activity_levels (name);

CREATE INDEX idx_t_user_health_user_id ON personal_bot.t_user_health (user_id);
CREATE INDEX idx_t_user_health_health_condition_id ON personal_bot.t_user_health (health_condition_id);

-- User Health Data in User Table
ALTER TABLE personal_bot.t_users ADD COLUMN weight NUMERIC(10,2);
ALTER TABLE personal_bot.t_users ADD COLUMN height NUMERIC(10,2);
ALTER TABLE personal_bot.t_users ADD COLUMN activity_level_id int REFERENCES personal_bot.t_activity_levels(id);

CREATE INDEX idx_t_users_activity_level_id ON personal_bot.t_users (activity_level_id);