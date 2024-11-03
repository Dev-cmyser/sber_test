package mortgage

import "time"

type Params struct {
	ObjectCost     int `json:"object_cost"`
	InitialPayment int `json:"initial_payment"`
	Months         int `json:"months"`
}

type Program struct {
	Salary   bool `json:"salary" binding:"omitempty"`
	Military bool `json:"military" binding:"omitempty"`
	Base     bool `json:"base" binding:"omitempty"`
}

type Aggregates struct {
	Rate            int       `json:"rate"`
	LoanSum         int       `json:"loan_sum"`
	MonthlyPayment  int       `json:"monthly_payment"`
	Overpayment     int       `json:"overpayment"`
	LastPaymentDate time.Time `json:"last_payment_date"`
}

type Request struct {
	ObjectCost     int     `json:"object_cost"`
	InitialPayment int     `json:"initial_payment"`
	Months         int     `json:"months"`
	Program        Program `json:"program"`
}
