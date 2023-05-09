package parameters_test

import (
	"testing"

	"lbc-fizzbuzz/internal/api/parameters"
	"lbc-fizzbuzz/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzRequestBody_MapToFizzBuzzModel(t *testing.T) {
	assert.Equal(t, parameters.FizzBuzzRequestBody{
		String1: "foo",
		String2: "bar",
		Int1:    1,
		Int2:    2,
		Limit:   3,
	}.MapToFizzBuzzModel(),
		&models.FizzBuzz{
			String1: "foo",
			String2: "bar",
			Int1:    1,
			Int2:    2,
			Limit:   3,
		})
}
