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

type AddUserHealthBody struct {
	IdToken             string `json:"id_token"`
	HealthConditionName string `json:"health_condition_name"`
	DiagnosisDate       string `json:"diagnosis_date"`
	Treatment           string `json:"treatment"`
	DischargedDate      string `json:"discharged_date"`
}

// This function will create or update the user health data
func AddUserHealth(params *AddUserHealthBody) *common.Response {
	user, err := getUserByIDToken(params.IdToken)

	if err != nil {
		return common.GetErrorResponse(err.Error(), http.StatusInternalServerError)
	}

	healthCondition, err := repositories.GetHealthConditionByName(params.HealthConditionName)

	if err != nil {
		return common.GetErrorResponse("That health condition name does not exists", http.StatusBadRequest)
	}

	userHealthData := &repositories.UserHealth{
		DiagnosisDate:     params.DiagnosisDate,
		Treatment:         params.Treatment,
		DischargedDate:    params.DischargedDate,
		HealthConditionId: healthCondition.Id,
		UserId:            user.Id,
	}

	err = repositories.CreateUserHealth(userHealthData)

	if err != nil {
		return common.GetErrorResponse("Error while adding your health data", http.StatusInternalServerError)
	}

	return common.GetSuccessResponse(map[string]string{"Message": "User Health Set"})
}

// This function will return the user health data
func GetUserHealth(idToken string) *common.Response {
	user, err := getUserByIDToken(idToken)

	if err != nil {
		return common.GetErrorResponse(err.Error(), http.StatusInternalServerError)
	}

	userHealthData, err := repositories.GetUserHealthByUserId(user.Id)

	if err != nil {
		return common.GetErrorResponse("Error while getting your health data", http.StatusInternalServerError)
	}

	return common.GetSuccessResponse(userHealthData)
}
