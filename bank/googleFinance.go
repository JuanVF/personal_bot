package bank

import (
	"strconv"

	"github.com/gocolly/colly"
)

type GoogleFinance struct {
}

// It will convert from one currency to another using BN Data
func (gf GoogleFinance) getCRCValue() (value float64, errData error) {
	c := colly.NewCollector()

	BNURL := "https://www.google.com/finance/quote/USD-CRC"

	c.OnError(func(r *colly.Response, err error) {
		logger.Error("Bank Google Finance", err.Error())

		errData = err
	})

	c.OnHTML(".YMlKec.fxKbKc", func(h *colly.HTMLElement) {
		liveValue := h.Text
		value, _ = strconv.ParseFloat(liveValue, 64)
	})

	c.Visit(BNURL)

	return value, errData
}
