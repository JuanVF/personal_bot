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
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/bank"
	"github.com/JuanVF/personal_bot/classifier"
	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/google"
	"github.com/JuanVF/personal_bot/repositories"
)

type GeneratePaymentsBody struct {
	BearerToken string `json:"bearer_token"`
	IDToken     string `json:"id_token"`
}

type GetPaymentsBody struct {
	IDToken string `json:"id_token"`
}

type ProcessPaymentPayload struct {
	Mails *google.GmailThreads
	Bot   *repositories.Bot
	User  *repositories.User
	Token string
}

// Returns the payments for an specific user
func GetPaymentsByTokenId(body *GetPaymentsBody) *common.Response {
	user, err := getUserByIDToken(body.IDToken)

	if err != nil {
		return &common.Response{
			Status: http.StatusBadRequest,
			Body: &common.ErrorResponse{
				Message: err.Error(),
			},
		}
	}

	payments, err := repositories.GetPaymentsByUserId(user.Id)

	if err != nil {
		return &common.Response{
			Status: http.StatusInternalServerError,
			Body: &common.ErrorResponse{
				Message: "An error ocurred. Please try again",
			},
		}
	}

	return &common.Response{
		Status: http.StatusOK,
		Body:   payments,
	}
}

// Generates the payments of an user using its ID Token
func GeneratePayments(body *GeneratePaymentsBody) *common.Response {
	var response *common.Response = &common.Response{
		Status: http.StatusOK,
	}

	user, err := getUserByIDToken(body.IDToken)

	if err != nil {
		return common.GetErrorResponse(err.Error(), http.StatusBadRequest)
	}

	bot, err := repositories.GetBotByUserId(user.Id)

	if err != nil {
		return common.GetErrorResponse("This user des not have a bot, please request a bot.", http.StatusBadRequest)
	}

	mails, err := google.GetGmailMessageList(user.GoogleMe, googleQuery, body.BearerToken)

	if err != nil {
		return common.GetErrorResponse("An error has ocurred while reading Gmail Mails. Please try again.", http.StatusBadRequest)
	}

	if mails.ResultSizeEstimate == 0 {
		return common.GetErrorResponse("There are not new payments to process.", http.StatusBadRequest)
	}

	payments, err := ProcessPayments(&ProcessPaymentPayload{
		Mails: mails,
		Bot:   bot,
		User:  user,
		Token: body.BearerToken,
	})

	if err != nil {
		return common.GetErrorResponse("An error has ocurred while processing the payments. Please check your data and verify.", http.StatusInternalServerError)
	}

	repositories.InsertPayments(payments)
	repositories.UpdateBot(bot)

	response.Body = payments

	return response
}

// Process all the payments
func ProcessPayments(payload *ProcessPaymentPayload) ([]*repositories.CreatePayment, error) {
	messages := payload.Mails.Messages

	if payload.Bot.LastGmailId == nil {
		return nil, fmt.Errorf("Configure the Bot ID[%d] for start reading the last thread id", payload.Bot.Id)
	}

	newLastGmailId := ""
	payments := make([]*repositories.CreatePayment, 0)

	bank := bank.Bank{
		CurrentBank: bank.GF_DS,
	}

	dollarPrice := bank.Convert(1, &repositories.Currency{Name: "USD"}, &repositories.Currency{Name: "CRC"})

	for i, message := range messages {
		if i == 0 {
			newLastGmailId = message.Id
		}

		if message.Id == *payload.Bot.LastGmailId {
			break
		}

		payment, err := ProcessPayment(message.Id, payload.Token, payload.User)

		if err != nil {
			continue
		}

		payment.DolarPrice = dollarPrice

		payments = append(payments, payment)
	}

	payload.Bot.LastGmailId = &newLastGmailId

	return payments, nil
}

// Process and return a specific payment
func ProcessPayment(threadId, token string, user *repositories.User) (*repositories.CreatePayment, error) {
	thread, err := google.GetGmailMessage(user.GoogleMe, threadId, token)

	if err != nil {
		return nil, err
	}

	body, err := GetBodyFromPayload(&thread.Payload)

	if err != nil {
		common.GetLogger().Error("Payment Service", err.Error())
		return nil, err
	}

	header, err := GetHeader(&thread.Payload.Headers, "From")

	if err != nil {
		return nil, err
	}

	date, err := GetHeader(&thread.Payload.Headers, "Date")

	if err != nil {
		return nil, err
	}

	bank := bank.Bank{}

	paymentData := bank.GetPaymentData(body, header)

	if paymentData == nil {
		common.GetLogger().Error("Payment Service", fmt.Sprintf("Body for Message ID[%s] didn't match requisites", threadId))
		return nil, fmt.Errorf("Body didn't match requisites")
	}

	classifier := classifier.Classifier{
		Model: classifier.OPEN_AI,
	}

	tags, err := classifier.GetClassifier().Classify(paymentData.Body)

	if err != nil {
		return nil, fmt.Errorf("Error while getting tags")
	}

	currency := repositories.GetCurrencyByName(paymentData.Currency.Name)

	dbBank, err := repositories.GetBankByName(paymentData.BankName)

	if err != nil {
		common.GetLogger().Error("Payment Service", fmt.Sprintf("Bank[%s] not found", paymentData.BankName))

		return nil, fmt.Errorf("Error while getting bank")
	}

	payment := &repositories.CreatePayment{
		Amount:      paymentData.Amount,
		CurrencyId:  currency.Id,
		Tags:        tags,
		UserId:      user.Id,
		LastUpdated: ConvertFROMRFC822toISOString(date),
		Description: paymentData.Body,
		GmailId:     thread.Id,
		IdBank:      dbBank.Id,
	}

	return payment, nil
}
