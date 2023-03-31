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
package repositories

import (
	"database/sql"
	"fmt"
	"strings"
)

type WalmartInvoice struct {
	Id             int
	UserId         int
	Amount         float64
	InvoiceDate    string
	ItemsPurchased int
	GmailId        string
}

type Ingredient struct {
	Id               int
	WalmartInvoiceId int
	Name             string
}

type DietPlan struct {
	Id           int
	UserId       int
	Name         string
	Description  string
	MealPlan     string
	Warning      string
	CreationDate string
}

type WalmartRepository struct {
	db *sql.DB
}

func NewWalmartRepository(db *sql.DB) *WalmartRepository {
	return &WalmartRepository{db}
}

// Create a new Diet Plan
func (r *WalmartRepository) CreateDietPlan(dietPlan *DietPlan) (int, error) {
	statement := `INSERT INTO personal_bot.t_diet_plans 
                    (user_id, name, description, meal_plan, warning, creation_date)
                VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`

	var id int

	err := r.db.QueryRow(statement, dietPlan.UserId, dietPlan.Name, dietPlan.Description, dietPlan.MealPlan, dietPlan.Warning).Scan(&id)

	if err != nil {
		logger.Error("Walmart Repository - Create Diet Plan", err.Error())

		return 0, err
	}

	return id, nil
}

// Create a new Walmart Invoice
func (r *WalmartRepository) CreateWalmartInvoice(invoice *WalmartInvoice) (int, error) {
	statement := `INSERT INTO personal_bot.t_walmart_invoice 
                    (user_id, amount, invoice_date, items_purchased, gmail_id)
                VALUES ($1, $2, NOW(), $3, $4) RETURNING id`

	var id int

	err := r.db.QueryRow(statement, invoice.UserId, invoice.Amount, invoice.ItemsPurchased, invoice.GmailId).Scan(&id)

	if err != nil {
		logger.Error("Walmart Repository - Create Walmart Invoice", err.Error())

		return 0, err
	}

	return id, nil
}

// Bulk insert ingredients into the database
func (r *WalmartRepository) BulkInsertIngredients(walmartInvoiceId int, ingredients []Ingredient) error {
	if len(ingredients) == 0 {
		return nil
	}

	// Prepare the SQL statement with placeholders for bulk insertion
	placeholders := make([]string, len(ingredients))
	values := make([]interface{}, 0, len(ingredients)*2)

	for i, ingredient := range ingredients {
		placeholders[i] = fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2)

		values = append(values, walmartInvoiceId, ingredient.Name)
	}

	statement := `INSERT INTO personal_bot.t_ingredients
                    (walmart_invoice_id, name)
                  VALUES ` + strings.Join(placeholders, ", ")

	// Execute the SQL statement
	_, err := r.db.Exec(statement, values...)

	if err != nil {
		logger.Error("Walmart Repository - Bulk Insert Ingredients", err.Error())
		return err
	}

	return nil
}
