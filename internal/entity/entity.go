package entity

import "github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"

type Mortgage struct {
	Params     mortgage.Params     `json:"params"`
	Program    mortgage.Program    `json:"program"`
	Aggregates mortgage.Aggregates `json:"aggregates"`
}

type CachedMortgage struct {
	ID         int                 `json:"id"`
	Params     mortgage.Params     `json:"params"`
	Program    mortgage.Program    `json:"program"`
	Aggregates mortgage.Aggregates `json:"aggregates"`
}
