package repositories

import (
	"database/sql"
	"fmt"
)

type FitnessGoalStatus struct {
	Id   int
	Name string
}

type FitnessTarget struct {
	Id   int
	Name string
}

type Measure struct {
	Id   int
	Name string
}

type FitnessGoal struct {
	Id                  int
	UserId              int
	Name                string
	Description         string
	StartDate           string
	FitnessGoalStatusId int
	FitnessGoalStatus   FitnessGoalStatus
	FitnessTargetId     int
	FitnessTarget       FitnessTarget
	MeasureId           int
	Measure             Measure
	CreationDate        string
}

// CreateFitnessGoal creates a new fitness goal for the user
func CreateFitnessGoal(goal *FitnessGoal) error {
	statement := `
		INSERT INTO personal_bot.t_fitness_goals (
			user_id, name, description, start_date,
			fitness_goal_status_id, fitness_target_id, measure_id, creation_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`

	_, err := db.GetConnection().Exec(
		statement,
		goal.UserId,
		goal.Name,
		goal.Description,
		goal.StartDate,
		goal.FitnessGoalStatusId,
		goal.FitnessTargetId,
		goal.MeasureId,
	)

	if err != nil {
		logger.Error("Fitness Goal Repository - Create Fitness Goal", err.Error())

		return err
	}

	return nil
}

// Return a measure by its name
func GetMeasureByName(name string) (*Measure, error) {
	statement := "SELECT id, name FROM personal_bot.t_measures WHERE name = $1"

	var measure Measure

	err := db.GetConnection().QueryRow(statement, name).Scan(&measure.Id, &measure.Name)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("measure not found")
	}

	if err != nil {
		logger.Error("Fitness Repository - Get Measure By Name", err.Error())

		return nil, err
	}

	return &measure, nil
}

// Return a fitness target by its name
func GetFitnessTargetByName(name string) (*FitnessTarget, error) {
	statement := "SELECT id, name FROM personal_bot.t_fitness_targets WHERE name = $1"

	var target FitnessTarget

	err := db.GetConnection().QueryRow(statement, name).Scan(&target.Id, &target.Name)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("fitness target not found")
	}

	if err != nil {
		logger.Error("Fitness Repository - Get Fitness Target By Name", err.Error())

		return nil, err
	}

	return &target, nil
}

// Return a fitness goal status by its name
func GetFitnessGoalStatusByName(name string) (*FitnessGoalStatus, error) {
	statement := "SELECT id, name FROM personal_bot.t_fitness_goal_statuses WHERE name = $1"

	var status FitnessGoalStatus

	err := db.GetConnection().QueryRow(statement, name).Scan(&status.Id, &status.Name)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("fitness goal status not found")
	}

	if err != nil {
		logger.Error("Fitness Repository - Get Fitness Goal Status By Name", err.Error())

		return nil, err
	}

	return &status, nil
}
