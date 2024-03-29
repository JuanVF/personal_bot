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
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/JuanVF/personal_bot/google"
	"github.com/JuanVF/personal_bot/openai"
	"github.com/JuanVF/personal_bot/repositories"
	"github.com/dvsekhvalnov/jose2go/base64url"
)

type WalmartData struct {
	WalmartInvoice repositories.WalmartInvoice
	Ingredients    []repositories.Ingredient
}

type WalmartInvoiceXML struct {
	DetalleServicio struct {
		LineaDetalle []struct {
			Detalle         string  `xml:"Detalle"`
			MontoTotalLinea float64 `xml:"MontoTotalLinea"`
		} `xml:"LineaDetalle"`
	} `xml:"DetalleServicio"`
	ResumenFactura struct {
		TotalComprobante float64 `xml:"TotalComprobante"`
	} `xml:"ResumenFactura"`
}

// This function will read the latest Walmart Invoice and generate a diet plan based
func (s DietPlanService) getDietPlan(
	walmartData *WalmartData,
	fitnessGoals []repositories.FitnessGoal,
	healthProblems []*repositories.UserHealth,
	user *repositories.User) (*repositories.DietPlan, error) {
	// Organize the ingredients into a string to be used in the prompt
	ingredientsInterface := make([]interface{}, len(walmartData.Ingredients))

	for i, ingredient := range walmartData.Ingredients {
		ingredientsInterface[i] = ingredient
	}

	ingredients := s.organizeItems(ingredientsInterface, func(item interface{}) string {
		return item.(repositories.Ingredient).Name
	})

	// Organize the fitness goals into a string to be used in the prompt
	fitnessGoalsInterface := make([]interface{}, len(fitnessGoals))

	for i, fitnessGoal := range fitnessGoals {
		fitnessGoalsInterface[i] = fitnessGoal
	}

	goals := s.organizeItems(fitnessGoalsInterface, func(item interface{}) string {
		return item.(repositories.FitnessGoal).Description
	})

	// Organize the user health problems into a string to be used in the prompt
	userHealthInterface := make([]interface{}, len(healthProblems))

	for i, healthProblem := range healthProblems {
		userHealthInterface[i] = *healthProblem
	}

	healthProblemsPrompt := s.organizeItems(userHealthInterface, func(item interface{}) string {
		healthProblem := item.(repositories.UserHealth)

		return fmt.Sprintf("%s, which is %s and is being treated with %s",
			healthProblem.HealthCondition.Name, healthProblem.HealthCondition.Description, healthProblem.Treatment)
	})

	ChatPrompts := []openai.GPT3Chat{
		{
			Role:    "user",
			Content: fmt.Sprintf("I bought this ingredients yesterday: %s", ingredients),
		},
		{
			Role:    "system",
			Content: "That's pretty nice, now tell me, what do you want me to do with those ingredients?",
		},
		{
			Role: "user",
			Content: fmt.Sprintf(`Create a personalized diet plan for me. My goals are %s. 
								I have a history of %s. 
								Also my weight is %.2f kg and my height is %.2f cm, and I have an activity level of %s.
								The user recently made a Walmart purchase, and the ingredients are provided in the data. 
								Please create a list of possible recipes that can help him to achieve its goals, organize it in breakfast, lunch and dinner. 
								The diet should contain macronutrient and micronutrient information, and any necessary food restrictions or notes. 
								Don't ask anything else, and give me the diet plan please.`, goals, healthProblemsPrompt, user.Weight, user.Height, user.ActivityLevel.Name),
		},
	}

	resp, err := openai.Chat(&ChatPrompts)

	if err != nil {
		return nil, fmt.Errorf("Error getting response from OpenAI")
	}

	dietPlan := &repositories.DietPlan{
		MealPlan: resp.Choices[0].Message.Content,
		UserId:   user.Id,
		Description: fmt.Sprintf("Diet plan for user %s, with goals %s, and health problems %s",
			user.Name, goals, healthProblemsPrompt),
		Name:    fmt.Sprintf("Diet plan for %s", user.Name),
		Warning: "This diet plan was generated by an AI, and it's not guaranteed to be accurate. Consult a doctor or nutritionist before following this diet plan.",
	}

	return dietPlan, err
}

func (s DietPlanService) organizeItems(items []interface{}, format func(item interface{}) string) string {
	organizedItems := ""

	for i, item := range items {
		if i == len(items)-1 && len(items) > 1 {
			organizedItems += fmt.Sprintf("and %s.", format(item))
		}

		organizedItems += format(item) + ", "
	}

	return organizedItems
}

// Will read the latest Walmart Invoice and structure the data
func (s DietPlanService) getWalmartData(user *repositories.User, accessToken string) (*WalmartData, error) {
	mails, err := google.GetGmailMessageList(user.GoogleMe, walmartQuery, accessToken)

	if err != nil {
		return nil, err
	}

	if len(mails.Messages) == 0 {
		return nil, fmt.Errorf("No Walmart Mails")
	}

	latestWalmartMail := mails.Messages[0]

	mail, err := google.GetGmailMessage(user.GoogleMe, latestWalmartMail.Id, accessToken)

	if err != nil {
		return nil, err
	}

	xmlPart := s.getXMLPart(mail.Payload.Parts)

	if xmlPart == nil {
		return nil, fmt.Errorf("No XML Part")
	}

	xml, err := google.RetrieveFile(user.GoogleMe, latestWalmartMail.Id, xmlPart.Body.AttachmentId, accessToken)

	if err != nil {
		return nil, err
	}

	decodedXML, err := base64url.Decode(xml.Data)

	if err != nil {
		return nil, err
	}

	data := s.getWalmartInvoice(decodedXML)

	if data == nil {
		return nil, fmt.Errorf("No Walmart Invoice")
	}

	data.WalmartInvoice.UserId = user.Id
	data.WalmartInvoice.GmailId = latestWalmartMail.Id

	return data, nil
}

// Will extract the data from the XML invoice
func (s DietPlanService) getWalmartInvoiceXML(data []byte) *WalmartInvoiceXML {
	invoice := &WalmartInvoiceXML{}

	err := xml.Unmarshal(data, &invoice)

	if err != nil {
		return nil
	}

	return invoice
}

// Will extract the data from the XML invoice
func (s DietPlanService) getWalmartInvoice(data []byte) *WalmartData {
	invoice := s.getWalmartInvoiceXML(data)

	if invoice == nil {
		return nil
	}

	// Extract data for WalmartInvoice struct
	walmartInvoice := &repositories.WalmartInvoice{}
	ingredients := []repositories.Ingredient{}

	for _, linea := range invoice.DetalleServicio.LineaDetalle {
		ingredient := &repositories.Ingredient{
			Name: linea.Detalle,
		}

		ingredients = append(ingredients, *ingredient)

		walmartInvoice.ItemsPurchased++
	}

	walmartInvoice.Amount = invoice.ResumenFactura.TotalComprobante

	// Create a WalmartData struct to hold the extracted data
	walmartData := &WalmartData{
		WalmartInvoice: *walmartInvoice,
		Ingredients:    ingredients,
	}

	return walmartData
}

// Will get the XML invoice of the mail
func (s DietPlanService) getXMLPart(parts []google.GmailMessagePayload) *google.GmailMessagePayload {
	for _, payload := range parts {
		if strings.HasSuffix(payload.Filename, ".xml") && strings.HasPrefix(payload.MimeType, "application/") && payload.Filename[0] == 'T' {
			return &payload
		} else if payload.Parts != nil {
			if pdf := s.getXMLPart(payload.Parts); pdf != nil {
				return pdf
			}
		}
	}

	return nil
}
