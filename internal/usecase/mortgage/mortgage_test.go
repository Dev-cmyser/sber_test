package uc_mortgage_test

import (
	"context"
	"testing"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase"
	uc_mortgage "github.com/Dev-cmyser/calc_ipoteka/internal/usecase/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase/mortgage/mocks"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("MortgageUseCase", func() {
	var (
		mockCache *mocks.Cache[int, entity.CachedMortgage]
		uc        *uc_mortgage.MortgageUseCase[int, entity.CachedMortgage]
	)

	ginkgo.BeforeEach(func() {
		mockCache = mocks.NewCache[int, entity.CachedMortgage](ginkgo.GinkgoT())
		uc = uc_mortgage.New[int, entity.CachedMortgage](mockCache)

	})

	ginkgo.Describe("Execute", func() {
		ginkgo.It("should return an error if the initial payment is too low", func() {
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 15000,
				Months:         120,
				Program:        mortgage.Program{Salary: new(bool)}, // Example program
			}

			_, err := uc.Execute(context.Background(), req)

			gomega.Expect(err).To(gomega.Equal(usecase.ErrLowInitPay))
		})

		ginkgo.It("should calculate mortgage successfully for salary program", func() {
			pointer := new(bool)
			*pointer = true
			req := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Salary: pointer},
			}

			result, err := uc.Execute(context.Background(), req)

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Aggregates.Rate).To(gomega.Equal(8))
			gomega.Expect(result.Aggregates.LoanSum).To(gomega.Equal(80000))
		})
	})
})

func TestMortgageUseCase(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "MortgageUseCase Suite")
}
