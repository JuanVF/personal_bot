package services

import (
	"fmt"
	"time"

	paws "github.com/JuanVF/personal_bot/aws"
	"github.com/JuanVF/personal_bot/bank"
	"github.com/JuanVF/personal_bot/repositories"
	"github.com/aws/aws-sdk-go/aws"
)

func CheckBudgetsByUserId(userId int) error {
	user, err := repositories.GetUser(userId)

	if err != nil {
		return err
	}

	budgets, err := repositories.GetBudgetsByUserId(userId)

	if err != nil {
		return err
	}

	overBudgets := GetOverBudgets(budgets, userId)

	if len(overBudgets) <= 0 {
		return nil
	}

	paws.SendEmail([]*string{aws.String(user.GoogleMe)}, aws.String(config.AWS.SES.BudgetTemplate), map[string]*string{
		"name": aws.String(user.Name),
		"data": aws.String(overBudgets),
	})

	return nil
}

// Returns those budgets reaching the limit
func GetOverBudgets(budgets []*repositories.Budget, userId int) string {
	s, e := getMonthStartAndEnd(time.Now())

	startOfMonthISO := s.Format(time.RFC3339)
	endOfMonthISO := e.Format(time.RFC3339)

	overBudgets := ""

	for _, budget := range budgets {
		summary, err := repositories.GetSummaryOfCertainPaymentsByDatesAndUserId(userId, budget.Matcher, startOfMonthISO, endOfMonthISO)

		if err != nil {
			logger.Error("Budget Service", err.Error())

			continue
		}

		if len(summary.PaymentsSummary) <= 0 {
			continue
		}

		total := GetTotalFromSummary(summary)

		// We will count those who are reaching the budget already
		if total >= budget.Amount*0.80 {
			overBudgets += fmt.Sprintf("<br>Budget: %s<br>Monto Gastado: %f<br>Monto Limite: %f<br><br>", budget.Name, total, budget.Amount)
		}
	}

	return overBudgets
}

// Returns the total in CRC of a summary
func GetTotalFromSummary(summary *repositories.UserPaymentsSummary) float64 {
	bank := bank.Bank{
		CurrentBank: bank.GF_DS,
	}

	CRCSum := 0.0
	USDSum := 0.0

	for _, s := range summary.PaymentsSummary {
		if s.Currency.Name == "USD" {
			USDSum += s.TotalAmount
		}

		if s.Currency.Name == "CRC" {
			CRCSum += s.TotalAmount
		}
	}

	return CRCSum + bank.Convert(USDSum, &repositories.Currency{Name: "USD"}, &repositories.Currency{Name: "USD"})
}
