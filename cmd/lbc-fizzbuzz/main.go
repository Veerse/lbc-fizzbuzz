package main

import (
	"github.com/sirupsen/logrus"
	"lbc-fizzbuzz/internal/api"
	"lbc-fizzbuzz/internal/services"
)

func main() {
	log := logrus.New()
	fizzBuzzService := services.NewFizzBuzzService()

	a := api.NewAPI(log, fizzBuzzService)

	a.Initialize()

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
