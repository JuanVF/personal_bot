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
