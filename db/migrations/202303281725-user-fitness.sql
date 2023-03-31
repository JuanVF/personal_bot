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
CREATE TABLE personal_bot.t_fitness_goal_statuses (
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE personal_bot.t_fitness_targets (
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE personal_bot.t_measures (
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE personal_bot.t_fitness_goals (
    id serial PRIMARY KEY,
    user_id int REFERENCES personal_bot.t_users(id),
    name text NOT NULL,
    description text,
    start_date timestamp with time zone NOT NULL,
    fitness_goal_status_id int REFERENCES personal_bot.t_fitness_goal_statuses(id),
    fitness_target_id int REFERENCES personal_bot.t_fitness_targets(id),
    measure_id int REFERENCES personal_bot.t_measures(id),
    creation_date timestamp with time zone NOT NULL DEFAULT now()
);

CREATE INDEX idx_t_fitness_goals_user_id ON personal_bot.t_fitness_goals (user_id);
CREATE INDEX idx_t_fitness_goals_fitness_goal_status_id ON personal_bot.t_fitness_goals (fitness_goal_status_id);
CREATE INDEX idx_t_fitness_goals_fitness_target_id ON personal_bot.t_fitness_goals (fitness_target_id);
CREATE INDEX idx_t_fitness_goals_measure_id ON personal_bot.t_fitness_goals (measure_id);
