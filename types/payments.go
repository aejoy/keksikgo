package types

type CreatePayments struct {
	General
	System string `json:"system,omitempty"`
	Purse  string `json:"purse,omitempty"`
	Name   string `json:"name,omitempty"`
	Amount int    `json:"amount,omitempty"`
}
