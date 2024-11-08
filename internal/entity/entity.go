// Package entity is
package entity

import "github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"

// Mortgage is.
type Mortgage struct {
	Program    mortgage.Program    `json:"program"`
	Params     mortgage.Params     `json:"params"`
	Aggregates mortgage.Aggregates `json:"aggregates"`
}

// CachedMortgage is cache.
type CachedMortgage struct {
	ID         int                 `json:"id"`
	Params     mortgage.Params     `json:"params"`
	Program    mortgage.Program    `json:"program"`
	Aggregates mortgage.Aggregates `json:"aggregates"`
}
