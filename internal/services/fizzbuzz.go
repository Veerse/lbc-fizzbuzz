package services

import (
	"strconv"

	"lbc-fizzbuzz/internal/models"
)

//go:generate mockgen --destination ./mocks/fizzbuzz.go . IFizzBuzzService
type IFizzBuzzService interface {
	// DoFizzBuzz performs the fizzbuzz operation
	DoFizzBuzz(fizzBuzz *models.FizzBuzz) string
}

type FizzBuzzService struct{}

func NewFizzBuzzService() IFizzBuzzService {
	return &FizzBuzzService{}
}

func (s *FizzBuzzService) DoFizzBuzz(fb *models.FizzBuzz) string {
	var result string

	for i := 1; i <= fb.Limit; i++ {
		if i%fb.Int1 == 0 && i%fb.Int2 == 0 {
			result = result + fb.String1 + fb.String2 + ","
		} else if i%fb.Int1 == 0 {
			result = result + fb.String1 + ","
		} else if i%fb.Int2 == 0 {
			result = result + fb.String2 + ","
		} else {
			result = result + strconv.Itoa(i) + ","
		}
	}

	return result[:len(result)-1]
}
