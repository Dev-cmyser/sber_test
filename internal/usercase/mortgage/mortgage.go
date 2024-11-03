package uc_mortgage

import (
	"context"
	"errors"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/cache"
)

type MortgageUseCase[K comparable, V any] struct {
	c cache.Cache[K, V]
}

func New[K comparable, V any](c cache.Cache[K, V]) *MortgageUseCase[K, V] {
	return &MortgageUseCase[K, V]{
		c: c,
	}
}

func (uc *MortgageUseCase[K, V]) Execute(ctx context.Context, req mortgage.Request) (entity.Mortgage, error) {
	if req.ObjectCost <= 0 {
		return entity.Mortgage{}, errors.New("object cost must be greater than zero")
	}

	return entity.Mortgage{}, nil
}

func (uc *MortgageUseCase[K, V]) Cache(ctx context.Context) ([]entity.CachedMortgage, error) {
	return []entity.CachedMortgage{}, nil
}
