package services_test

import (
	"reflect"
	"testing"

	"lbc-fizzbuzz/internal/models"
	"lbc-fizzbuzz/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestNewFizzBuzzService(t *testing.T) {
	s := services.NewFizzBuzzService()

	assert.NotNil(t, s)
}

func TestFizzBuzzService_DoFizzBuzz(t *testing.T) {
	s := services.NewFizzBuzzService()

	tt := []struct {
		name     string
		fb       models.FizzBuzz
		expected string
	}{
		{
			name: "test-case-1",
			fb: models.FizzBuzz{
				String1: "fizz",
				String2: "buzz",
				Int1:    3,
				Int2:    5,
				Limit:   15,
			},
			expected: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz",
		},
		{
			name: "test-case-2",
			fb: models.FizzBuzz{
				String1: "foo",
				String2: "bar",
				Int1:    2,
				Int2:    7,
				Limit:   10,
			},
			expected: "1,foo,3,foo,5,foo,bar,foo,9,foo",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := s.DoFizzBuzz(&tc.fb)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
