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
package services

import (
	"net/http"

	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/repositories"
)

type CreateUserFitnessBody struct {
	IdToken               string `json:"id_token"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	StartDate             string `json:"start_date"`
	FitnessGoalStatusName string `json:"fitness_goal_status_name"`
	FitnessTargetName     string `json:"fitness_target_name"`
	MeasureName           string `json:"measure_name"`
}

// This function will create new fitness goal for the user
func AddUserFitnessGoal(params *CreateUserFitnessBody) *common.Response {
	user, err := getUserByIDToken(params.IdToken)

	if err != nil {
		return common.GetErrorResponse(err.Error(), http.StatusInternalServerError)
	}

	fitnessGoalStatus, err := repositories.GetFitnessGoalStatusByName(params.FitnessGoalStatusName)

	if err != nil {
		return common.GetErrorResponse("That fitness goal status name does not exists", http.StatusBadRequest)
	}

	fitnessTarget, err := repositories.GetFitnessTargetByName(params.FitnessTargetName)

	if err != nil {
		return common.GetErrorResponse("That fitness target name does not exists", http.StatusBadRequest)
	}

	measure, err := repositories.GetMeasureByName(params.MeasureName)

	if err != nil {
		return common.GetErrorResponse("That measure name does not exists", http.StatusBadRequest)
	}

	fitnessGoal := &repositories.FitnessGoal{
		UserId:              user.Id,
		Name:                params.Name,
		Description:         params.Description,
		StartDate:           params.StartDate,
		FitnessGoalStatusId: fitnessGoalStatus.Id,
		FitnessTargetId:     fitnessTarget.Id,
		MeasureId:           measure.Id,
	}

	err = repositories.CreateFitnessGoal(fitnessGoal)

	if err != nil {
		return common.GetErrorResponse("Error creating fitness goal", http.StatusInternalServerError)
	}

	return common.GetSuccessResponse(map[string]string{"message": "Fitness goal created successfully"})
}
