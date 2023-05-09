package api

import (
	"net/http"

	"lbc-fizzbuzz/internal/api/helpers"
	"lbc-fizzbuzz/internal/api/parameters"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func (a *API) GetFizzBuzz(c *gin.Context) {
	var b parameters.FizzBuzzRequestBody

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := a.val.Struct(b); err != nil {
		// Type assertion on "err" won't fail here because struct method returns an assertable error to
		// type validator.ValidationErrors.
		c.JSON(http.StatusBadRequest, helpers.ParseValidationErrors(err.(validator.ValidationErrors)))
		return
	}

	fb := b.MapToFizzBuzzModel()

	a.log.WithFields(logrus.Fields{
		"int1":    fb.Int1,
		"int2":    fb.Int2,
		"string1": fb.String1,
		"string2": fb.String2,
		"limit":   fb.Limit,
	}).Info("performing fizzbuzz")

	c.JSON(http.StatusOK, a.fizzBuzz.DoFizzBuzz(fb))
}
