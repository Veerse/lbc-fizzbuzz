package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"lbc-fizzbuzz/internal/api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"lbc-fizzbuzz/internal/api/parameters"
)

func (a *API) GetFizzBuzz(c *gin.Context) {
	var b parameters.FizzBuzzRequestBody

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := a.val.Struct(b); err != nil {
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
