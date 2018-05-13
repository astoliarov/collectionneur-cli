package entities

import "time"

type SpendInfo struct {
	Currency string
	Sum      float32

	Card        string
	Date        time.Time
	Description string

	Category string
}

func NewSpendInfo(currency string, sum float32, card string, date time.Time, description string) *SpendInfo {
	return &SpendInfo{
		Currency:    currency,
		Sum:         sum,
		Card:        card,
		Date:        date,
		Description: description,
	}
}
