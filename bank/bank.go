package bank

import (
	"regexp"

	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/repositories"
)

var logger *common.Logger = common.GetLogger()

var BlankSpaces *regexp.Regexp = regexp.MustCompile("\\s{2,}")
var BlankSpace *regexp.Regexp = regexp.MustCompile("\\s")
var CurrenciesRegex *regexp.Regexp = regexp.MustCompile("USD|CRC")

// For now we are just using Google Finance Data since
// we need to investigate a way to get the Local Banks data
const (
	GF_DS int = iota
)

type Bank struct {
	CurrentBank int
}

// It will convert a currency to another using the defined datasource
func (b *Bank) Convert(value float64, from, to *repositories.Currency) float64 {
	ds := b.getDataSource()

	CRCValue, _ := ds.getCRCValue()

	if from.Name == "USD" && to.Name == "CRC" {
		return value * CRCValue
	}

	if from.Name == "CRC" && to.Name == "USD" {
		return value / CRCValue
	}

	return 0
}

// It will convert a currency to another using the defined datasource
func (b *Bank) GetPaymentData(body, from string) *PaymentData {
	matcher := b.getMatcher(from)

	if matcher == nil {
		return nil
	}

	return matcher.getPayment(body)
}

// Given the set datasource will return the bank handler
func (b *Bank) getMatcher(from string) BankMatcher {
	matched, err := regexp.MatchString("bncontacto@bncr\\.fi\\.cr", from)

	if matched && err == nil {
		return BN{}
	}

	matched, err = regexp.MatchString("popularvisa@bancopopularinforma\\.fi\\.cr", from)

	if matched && err == nil {
		return BP{}
	}

	logger.Error("Bank", "Matcher Not Supported")

	return nil
}

// Given the set datasource will return the bank handler
func (b *Bank) getDataSource() BankHandler {
	if b.CurrentBank == GF_DS {
		return GoogleFinance{}
	}

	logger.Error("Bank", "Data Source Not Supported")

	return nil
}
