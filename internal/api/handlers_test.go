package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"lbc-fizzbuzz/internal/api"
	"lbc-fizzbuzz/internal/api/parameters"
	mock_services "lbc-fizzbuzz/internal/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestAPI_GetFizzBuzz(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)

	log, hook := test.NewNullLogger()
	mockFizzBuzz := mock_services.NewMockIFizzBuzzService(ctrl)

	a := api.NewAPI(log, mockFizzBuzz)

	t.Run("ok", func(t *testing.T) {
		// Create a mock gin context
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		fizzBuzz := parameters.FizzBuzzRequestBody{
			Int1:    3,
			Int2:    5,
			String1: "fizz",
			String2: "buzz",
			Limit:   15,
		}

		// Marshal the request payload to JSON
		reqBodyJSON, _ := json.Marshal(fizzBuzz)

		// Set the request body
		c.Request, _ = http.NewRequest(http.MethodGet, "/fizzbuzz", bytes.NewBuffer(reqBodyJSON))

		mockFizzBuzz.EXPECT().DoFizzBuzz(fizzBuzz.MapToFizzBuzzModel())

		a.GetFizzBuzz(c)

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Len(t, hook.Entries, 1)
	})

	t.Run("nok", func(t *testing.T) {
		t.Run("incorrect-body", func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			reqBodyJSON, _ := json.Marshal(struct {
				Foo int `json:"foo"`
				Bar int `json:"bar"`
			}{
				Foo: 10,
				Bar: 20,
			})

			c.Request, _ = http.NewRequest(http.MethodGet, "/fizzbuzz", bytes.NewBuffer(reqBodyJSON))

			a.GetFizzBuzz(c)

			assert.Equal(t, w.Code, http.StatusBadRequest)
		})

		t.Run("incorrect-request-values", func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Marshal the request payload to JSON
			reqBodyJSON, _ := json.Marshal(parameters.FizzBuzzRequestBody{
				Int1:    3,
				Int2:    5,
				String1: "fizz",
				String2: "buzz",
				Limit:   1000,
			})

			// Set the request body
			c.Request, _ = http.NewRequest(http.MethodGet, "/fizzbuzz", bytes.NewBuffer(reqBodyJSON))

			a.GetFizzBuzz(c)

			assert.Equal(t, w.Code, http.StatusBadRequest)
		})
	})
}
