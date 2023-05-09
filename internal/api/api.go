package api

import (
	"lbc-fizzbuzz/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type API struct {
	log      logrus.FieldLogger
	router   *gin.Engine
	fizzBuzz services.IFizzBuzzService
	val      *validator.Validate
}

// NewAPI returns an instance of API.
func NewAPI(log logrus.FieldLogger, fbService services.IFizzBuzzService) *API {
	return &API{
		log:      log,
		router:   gin.Default(),
		fizzBuzz: fbService,
		val:      validator.New(),
	}
}

// Initialize initializes the API endpoints.
func (a *API) Initialize() {
	gin.SetMode(gin.ReleaseMode)

	a.router.GET("/fizzbuzz", a.GetFizzBuzz)
}

// Run runs the API on port 8090.
func (a *API) Run() error {
	return a.router.Run("localhost:8090")
}
