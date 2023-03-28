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
