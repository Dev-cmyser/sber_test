package uc_mortgage

import (
	"context"
	"errors"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/cache"
)

type MortgageUseCase struct {
	cache cache.Cache[int, entity.CachedMortgage]
}

func New() *MortgageUseCase {
	return &MortgageUseCase{}
}

func (m *MortgageUseCase) Execute(ctx context.Context, req mortgage.Request) (entity.Mortgage, error) {
	if req.ObjectCost <= 0 {
		return entity.Mortgage{}, errors.New("object cost must be greater than zero")
	}

	return entity.Mortgage{}, nil
}

func (m *MortgageUseCase) Cache(ctx context.Context) ([]entity.CachedMortgage, error) {
	return []entity.CachedMortgage{}, nil
}
