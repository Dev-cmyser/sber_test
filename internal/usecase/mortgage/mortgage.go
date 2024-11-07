package uc_mortgage

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/cache"
)

type MortgageUseCase[K comparable, V entity.CachedMortgage] struct {
	c cache.Cache[K, V]
	// Closure function but will be better use uuid func
	nextID func() int
}

func New[K comparable, V entity.CachedMortgage](c cache.Cache[K, V]) *MortgageUseCase[K, V] {
	id := 0

	return &MortgageUseCase[K, V]{
		c: c,
		nextID: func() int {
			id++
			return id
		},
	}
}

func (uc *MortgageUseCase[K, V]) Execute(ctx context.Context, req mortgage.Request) (entity.Mortgage, error) {
	if req.ObjectCost <= 0 {
		return entity.Mortgage{}, usecase.ErrLowInitPay
	}

	if float64(req.InitialPayment) < 0.2*float64(req.ObjectCost) {
		return entity.Mortgage{}, usecase.ErrLowInitPay
	}

	rate, err := uc.chooseProgramRate(req.Program)

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
			Rate:            rate,
			LoanSum:         loanSum,
			MonthlyPayment:  int(monthlyPayment),
			Overpayment:     overpayment,
			LastPaymentDate: lastPaymentDate,
		},
	}

	uc.saveToCache(result)

	return result, nil
}

func (uc *MortgageUseCase[K, V]) saveToCache(prog entity.Mortgage) error {
	id := uc.nextID()

	cachedMortgage := entity.CachedMortgage{
		ID:         id,
		Params:     prog.Params,
		Program:    prog.Program,
		Aggregates: prog.Aggregates,
	}

	var key K
	if k, ok := any(id).(K); ok {
		key = k
	} else {
		return errors.New("invalid key type")
	}
	var value V
	if v, ok := any(cachedMortgage).(V); ok {
		value = v
	} else {
		return errors.New("invalid value type")
	}

	uc.c.Add(key, value)

	return nil
}

func (uc *MortgageUseCase[K, V]) chooseProgramRate(prog mortgage.Program) (int, error) {
	selectedPrograms := 0

	if prog.Salary != nil && *prog.Salary {
		selectedPrograms++
	}
	if prog.Military != nil && *prog.Military {
		selectedPrograms++
	}
	if prog.Base != nil && *prog.Base {
		selectedPrograms++
	}

	if selectedPrograms == 0 {
		return 0, usecase.ErrChoosing
	} else if selectedPrograms > 1 {
		return 0, usecase.ErrOnlyOneProgram
	}

	switch {
	case prog.Salary != nil && *prog.Salary:
		return 8.0, nil
	case prog.Military != nil && *prog.Military:
		return 9.0, nil
	case prog.Base != nil && *prog.Base:
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

	return res, nil
}
