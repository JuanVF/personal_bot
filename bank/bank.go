package bank

import (
	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/repositories"
)

var logger *common.Logger = common.GetLogger()

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

// Given the set datasource will return the bank handler
func (b *Bank) getDataSource() BankHandler {
	if b.CurrentBank == GF_DS {
		return GoogleFinance{}
	}

	logger.Error("Bank", "Data Source Not Supported")

	return nil
}
