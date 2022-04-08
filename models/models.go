package models

type CheckWBM struct {
	Id string `json:"id"`
}

type ReplenishB struct {
	Id     string `json:"id"`
	Amount string `json:"amount"`
}

type RespW struct {
	Message string `json:"message"`
}
type RespB struct {
	Id    string `json:"id"`
	Money string `json:"money"`
}

type RespM struct {
	Id               string  `json:"id"`
	NumberOfIncome   float64 `json:"number_of_income"`
	NumberOfExpenses float64 `json:"number_of_expenses"`
	AllActions       int64   `json:"all_actions"`
	AllIncomeSumms   int64   `json:"all_income_summs"`
}
