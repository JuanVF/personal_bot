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
    discharged_date timestamp with time zone NOT NULL,
    activity_level_id int REFERENCES personal_bot.t_activity_levels(id)
);

CREATE INDEX idx_t_health_conditions_name ON personal_bot.t_health_conditions (name);
CREATE INDEX idx_t_activity_levels_name ON personal_bot.t_activity_levels (name);

CREATE INDEX idx_t_user_health_user_id ON personal_bot.t_user_health (user_id);
CREATE INDEX idx_t_user_health_health_condition_id ON personal_bot.t_user_health (health_condition_id);
CREATE INDEX idx_t_user_health_activity_level_id ON personal_bot.t_user_health (activity_level_id);