package kernel

import (
	"net/http"
	"time"

	"github.com/JustSteveKing/go-basic/pkg/infrastructure/factories"
	"github.com/JustSteveKing/go-basic/pkg/infrastructure/providers"
	gohandlers "github.com/gorilla/handlers"
)

// Boot our application and it's services
func Boot() *factories.Application {
	// Build our configuration
	config := providers.ConfigProvider()

	// Build our router
	router := providers.RouteProvider()

	// Build our logger
	logger, err := providers.LoggerProvider()
	if err != nil {
		panic(err)
	}

	// Set up CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// Create and return our Application to be ran
	return &factories.Application{
		Server: &http.Server{
			Addr:         ":" + config.App.Port,
			Handler:      corsHandler(router),
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Router: router,
		Logger: logger,
		Config: config,
	}
}
