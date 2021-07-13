package model

type Account struct {
	AccountNum       string   `json:"accountNum"`
	CustomerNum      string   `json:"customerNum"`
	Message          string   `json:"message"`
	Customer         Customer `json:"customer"`
	Amount           int      `json:"amount"`
}
