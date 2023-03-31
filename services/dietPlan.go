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

type DietPlanService struct {
}

// This function will read the latest Walmart Invoice and generate a diet plan based
// on the items purchased
func (s DietPlanService) GenerateDietPlan(idToken, accessToken string) *common.Response {
	repo := repositories.NewWalmartRepository(common.GetDB().GetConnection())

	user, err := getUserByIDToken(idToken)

	if err != nil {
		return common.GetErrorResponse("Invalid ID Token", http.StatusBadRequest)
	}

	userHealth, err := repositories.GetUserHealthByUserId(user.Id)

	if err != nil {
		return common.GetErrorResponse("Error getting user health data", http.StatusBadRequest)
	}

	fitnessGoals, err := repositories.GetFitnessGoalsByUser(user.Id)

	if err != nil {
		return common.GetErrorResponse("Error getting fitness goals", http.StatusBadRequest)
	}

	walmartData, err := s.getWalmartData(user, accessToken)

	if err != nil {
		return common.GetErrorResponse("Error getting Walmart data", http.StatusBadRequest)
	}

	walmartId, err := repo.CreateWalmartInvoice(&walmartData.WalmartInvoice)

	if err != nil {
		return common.GetErrorResponse("Error saving Walmart invoice", http.StatusBadRequest)
	}

	err = repo.BulkInsertIngredients(walmartId, walmartData.Ingredients)

	if err != nil {
		return common.GetErrorResponse("Error saving ingredients", http.StatusBadRequest)
	}

	dietPlan, err := s.getDietPlan(walmartData, fitnessGoals, userHealth, user)

	if err != nil {
		return common.GetErrorResponse("Error getting diet plan", http.StatusBadRequest)
	}

	id, err := repo.CreateDietPlan(dietPlan)

	if err != nil {
		return common.GetErrorResponse("Error saving diet plan", http.StatusBadRequest)
	}

	dietPlan.Id = id

	return common.GetSuccessResponse(dietPlan)
}
