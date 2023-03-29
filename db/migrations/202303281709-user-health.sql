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