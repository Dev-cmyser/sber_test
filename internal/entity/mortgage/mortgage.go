// Package mortgage s
package mortgage

import "time"

// Params s.
type Params struct {
	ObjectCost     int `json:"object_cost"`
	InitialPayment int `json:"initial_payment"`
	Months         int `json:"months"`
}

// Program s.
type Program struct {
	Salary   *bool `json:"salary,omitempty"`
	Military *bool `json:"military,omitempty"`
	Base     *bool `json:"base,omitempty"`
}

// Aggregates s.
type Aggregates struct {
	Rate            int       `json:"rate"`
	LoanSum         int       `json:"loan_sum"`
	MonthlyPayment  int       `json:"monthly_payment"`
	Overpayment     int       `json:"overpayment"`
	LastPaymentDate time.Time `json:"last_payment_date"`
}

// Request s.
type Request struct {
	ObjectCost     int     `json:"object_cost"`
	InitialPayment int     `json:"initial_payment"`
	Months         int     `json:"months"`
	Program        Program `json:"program"`
}
