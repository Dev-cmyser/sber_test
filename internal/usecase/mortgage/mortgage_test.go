package uc_mortgage_test

import (
	"context"
	"testing"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase"
	uc_mortgage "github.com/Dev-cmyser/calc_ipoteka/internal/usecase/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase/mortgage/mocks"
	gin "github.com/onsi/ginkgo"
	gom "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = gin.Describe("MortgageUseCase", func() {
	var (
		mockCache *mocks.Cache[int, entity.CachedMortgage]
		uc        *uc_mortgage.MortgageUseCase[int, entity.CachedMortgage]
	)

	pointerFalse := new(bool)
	pointerTrue := new(bool)
	*pointerTrue = true

	gin.BeforeEach(func() {
		mockCache = mocks.NewCache[int, entity.CachedMortgage](gin.GinkgoT())
		uc = uc_mortgage.New[int, entity.CachedMortgage](mockCache)

	})

	gin.Describe("Execute", func() {
		gin.It("should return an error if the initial payment is too low", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 15000,
				Months:         120,
				Program:        mortgage.Program{Salary: new(bool)},
			}

			_, err := uc.Execute(context.Background(), req)

			gom.Expect(err).To(gom.Equal(usecase.ErrLowInitPay))
		})

		gin.It("should return an error if the amount programs too many", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Salary: pointerTrue, Military: pointerTrue},
			}

			_, err := uc.Execute(context.Background(), req)

			gom.Expect(err).To(gom.Equal(usecase.ErrOnlyOneProgram))
		})

		gin.It("should return an error if the amount programs too many", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Salary: pointerTrue, Base: pointerTrue, Military: pointerTrue},
			}

			_, err := uc.Execute(context.Background(), req)

			gom.Expect(err).To(gom.Equal(usecase.ErrOnlyOneProgram))
		})

		gin.It("should return an error if the amount programs too low", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Salary: pointerFalse},
			}

			_, err := uc.Execute(context.Background(), req)

			gom.Expect(err).To(gom.Equal(usecase.ErrChoosing))
		})

		gin.It("should calculate mortgage successfully for salary program", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Salary: pointerTrue},
			}

			mockCache.On("Add", mock.Anything, mock.Anything).Return(true)

			result, err := uc.Execute(context.Background(), req)

			gom.Expect(err).NotTo(gom.HaveOccurred())
			gom.Expect(result.Aggregates.Rate).To(gom.Equal(8))
			gom.Expect(result.Aggregates.LoanSum).To(gom.Equal(80000))
			gom.Expect(result.Aggregates.MonthlyPayment).To(gom.Equal(970))
		})

		gin.It("should calculate mortgage successfully for military program", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Military: pointerTrue},
			}

			mockCache.On("Add", mock.Anything, mock.Anything).Return(true)

			result, err := uc.Execute(context.Background(), req)

			gom.Expect(err).NotTo(gom.HaveOccurred())
			gom.Expect(result.Aggregates.Rate).To(gom.Equal(9))
			gom.Expect(result.Aggregates.LoanSum).To(gom.Equal(80000))
			gom.Expect(result.Aggregates.MonthlyPayment).To(gom.Equal(1013))
		})

		gin.It("should calculate mortgage successfully for base program", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Base: pointerTrue},
			}

			mockCache.On("Add", mock.Anything, mock.Anything).Return(true)

			result, err := uc.Execute(context.Background(), req)

			gom.Expect(err).NotTo(gom.HaveOccurred())
			gom.Expect(result.Aggregates.Rate).To(gom.Equal(10))
			gom.Expect(result.Aggregates.LoanSum).To(gom.Equal(80000))
			gom.Expect(result.Aggregates.MonthlyPayment).To(gom.Equal(1057))
		})

	})
})

func TestMortgageUseCase(t *testing.T) {
	gom.RegisterFailHandler(gin.Fail)
	gin.RunSpecs(t, "MortgageUseCase Suite")
}
