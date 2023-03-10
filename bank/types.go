package bank

type BankHandler interface {
	getCRCValue() (float64, error)
}
