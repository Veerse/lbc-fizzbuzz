package main

import (
	"lbc-fizzbuzz/internal/api"
	"lbc-fizzbuzz/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	log := logrus.New()
	fizzBuzzService := services.NewFizzBuzzService()

	a := api.NewAPI(log, fizzBuzzService)

	a.Initialize()

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
