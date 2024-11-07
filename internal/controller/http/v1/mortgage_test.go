package v1_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/Dev-cmyser/calc_ipoteka/internal/controller/http/v1"
	"github.com/Dev-cmyser/calc_ipoteka/internal/controller/http/v1/mocks"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	gink "github.com/gin-gonic/gin"
	gin "github.com/onsi/ginkgo"
	gom "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = gin.Describe("Mortgage Routes", func() {
	const (
		exPath = "/v1/mortgage/execute"
	)

	var (
		mockUC  *mocks.UseCase
		mockLog *mocks.Interface
		router  *gink.Engine
	)

	type executeResp struct {
		Result entity.Mortgage `json:"result"`
	}
	// pointerFalse := new(bool)
	pointerTrue := new(bool)
	*pointerTrue = true

	gin.BeforeEach(func() {
		mockUC = &mocks.UseCase{}
		mockLog = &mocks.Interface{}
		gink.SetMode(gink.ReleaseMode)
		router = gink.Default()
		v1.NewRouter(router, mockLog, mockUC)
	})

	gin.Describe("GET /mortgage/cache", func() {
		gin.It("should get err", func() {
			cachedData := []entity.CachedMortgage{}
			mockUC.On("Cache", mock.Anything).Return(cachedData, v1.ErrEmpty)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/mortgage/cache", nil)
			router.ServeHTTP(w, req)

			gom.Expect(w.Code).To(gom.Equal(http.StatusNotFound))
		})
	})

	gin.Describe("POST /mortgage/execute", func() {

		gin.It("should return calculated mortgage details", func() {
			request := mortgage.Request{
				ObjectCost:     100000,
				InitialPayment: 20000,
				Months:         120,
				Program:        mortgage.Program{Salary: pointerTrue},
			}

			expectedResult := entity.Mortgage{
				Aggregates: mortgage.Aggregates{LoanSum: 80000, Rate: 8, MonthlyPayment: 970},
			}

			mockUC.On("Execute", mock.Anything, request).Return(expectedResult, nil)
			body, _ := json.Marshal(request)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, exPath, bytes.NewReader(body))

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			gom.Expect(w.Code).To(gom.Equal(http.StatusOK))

			var res executeResp
			_ = json.Unmarshal(w.Body.Bytes(), &res)

			gom.Expect(res.Result).To(gom.Equal(expectedResult))

		})

		gin.It("should return an error when request data is invalid", func() {
			invalidBody := `{"ObjectCost": "invalid"}`

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, exPath, bytes.NewBufferString(invalidBody))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			gom.Expect(w.Code).To(gom.Equal(http.StatusBadRequest))
			var res map[string]string
			_ = json.Unmarshal(w.Body.Bytes(), &res)
			gom.Expect(res["error"]).To(gom.ContainSubstring("invalid"))
		})
	})
})

func TestMortgageRoutes(t *testing.T) {
	gom.RegisterFailHandler(gin.Fail)
	gin.RunSpecs(t, "MortgageRoutes Suite")
}
