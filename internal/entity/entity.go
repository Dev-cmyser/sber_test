// Package entity is
package entity

import "github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"

// Mortgage is.
type Mortgage struct {
	Program    mortgage.Program    `json:"program"`
	Aggregates mortgage.Aggregates `json:"aggregates"`
	Params     mortgage.Params     `json:"params"`
}

// CachedMortgage is cache.
type CachedMortgage struct {
	Program    mortgage.Program    `json:"program"`
	Aggregates mortgage.Aggregates `json:"aggregates"`
	ID         int                 `json:"id"`
	Params     mortgage.Params     `json:"params"`
}
