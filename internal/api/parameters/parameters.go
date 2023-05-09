package parameters

import (
	"lbc-fizzbuzz/internal/models"
)

type FizzBuzzRequestBody struct {
	String1 string `json:"string1" validate:"required"`
	String2 string `json:"string2" validate:"required"`
	Int1    int    `json:"int1" validate:"required,gte=0"`
	Int2    int    `json:"int2" validate:"required,gte=0"`
	Limit   int    `json:"limit" validate:"required,lte=100"`
}

func (fb FizzBuzzRequestBody) MapToFizzBuzzModel() *models.FizzBuzz {
	return &models.FizzBuzz{
		String1: fb.String1,
		String2: fb.String2,
		Int1:    fb.Int1,
		Int2:    fb.Int2,
		Limit:   fb.Limit,
	}
}
