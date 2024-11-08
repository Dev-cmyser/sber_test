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
	LastPaymentDate time.Time `json:"last_payment_date"`
	Rate            int       `json:"rate"`
	LoanSum         int       `json:"loan_sum"`
	MonthlyPayment  int       `json:"monthly_payment"`
	Overpayment     int       `json:"overpayment"`
}

// Request s.
type Request struct {
	Program        Program `json:"program"`
	ObjectCost     int     `json:"object_cost"`
	InitialPayment int     `json:"initial_payment"`
	Months         int     `json:"months"`
}
