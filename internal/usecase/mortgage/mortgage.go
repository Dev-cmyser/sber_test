package uc_mortgage

import (
	"context"
	"math"
	"time"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/cache"
)

type MortgageUseCase[K comparable, V entity.CachedMortgage] struct {
	c cache.Cache[K, V]
}

func New[K comparable, V entity.CachedMortgage](c cache.Cache[K, V]) *MortgageUseCase[K, V] {
	return &MortgageUseCase[K, V]{
		c: c,
	}
}

func (uc *MortgageUseCase[K, V]) Execute(ctx context.Context, req mortgage.Request) (entity.Mortgage, error) {
	if req.ObjectCost <= 0 {
		return entity.Mortgage{}, usecase.ErrLowInitPay
	}

	if float64(req.InitialPayment) < 0.2*float64(req.ObjectCost) {
		return entity.Mortgage{}, usecase.ErrLowInitPay
	}

	rate, err := uc.determineProgramRate(req.Program)
	if err != nil {
		return entity.Mortgage{}, err
	}

	loanSum := req.ObjectCost - req.InitialPayment
	monthlyPayment := calcMonthPayment(float64(loanSum), float64(rate), req.Months)
	overpayment := int(monthlyPayment*float64(req.Months) - float64(loanSum))
	lastPaymentDate := calcLastPaymentDate(req.Months)

	result := entity.Mortgage{
		Params: mortgage.Params{
			ObjectCost:     req.ObjectCost,
			InitialPayment: req.InitialPayment,
			Months:         req.Months,
		},
		Program: req.Program,
		Aggregates: mortgage.Aggregates{
			Rate:            int(rate),
			LoanSum:         loanSum,
			MonthlyPayment:  int(monthlyPayment),
			Overpayment:     overpayment,
			LastPaymentDate: lastPaymentDate,
		},
	}

	return result, nil
}

func (uc *MortgageUseCase[K, V]) determineProgramRate(prog mortgage.Program) (int, error) {
	selectedPrograms := 0
	if prog.Salary {
		selectedPrograms++
	}
	if prog.Military {
		selectedPrograms++
	}
	if prog.Base {
		selectedPrograms++
	}

	if selectedPrograms == 0 {
		return 0, usecase.ErrChoosing
	} else if selectedPrograms > 1 {
		return 0, usecase.ErrOnlyOneProgram
	}

	switch {
	case prog.Salary:
		return 8.0, nil
	case prog.Military:
		return 9.0, nil
	case prog.Base:
		return 10.0, nil
	default:
		return 0, usecase.ErrOnlyOneProgram
	}
}

func calcMonthPayment(loanSum float64, rate float64, months int) float64 {
	monthlyRate := rate / 12 / 100
	return loanSum * (monthlyRate / (1 - math.Pow(1+monthlyRate, float64(-months))))
}

func calcLastPaymentDate(months int) time.Time {
	now := time.Now()
	return now.AddDate(0, months, 0)
}

func (uc *MortgageUseCase[K, V]) Cache(ctx context.Context) ([]entity.CachedMortgage, error) {
	var res []entity.CachedMortgage
	keys := uc.c.Keys()

	if len(keys) == 0 {
		return nil, usecase.ErrEmpty
	}

	for _, k := range keys {
		// ignore expiration live
		v, _ := uc.c.Get(k)
		res = append(res, entity.CachedMortgage(v))
	}

	return []entity.CachedMortgage{}, nil
}
