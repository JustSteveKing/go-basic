package factories

import (
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func TestApplication(test *testing.T) {
	config := ConfigFactory()

	logger, err := zap.NewProduction()
	if err != nil {
		test.Error("Failed to create logger for application test")
	}

	router := mux.NewRouter()

	app := &Application{
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Router: router,
		Logger: logger,
		Config: config,
	}

	test.Run("Test Application has been built correctly", func(test *testing.T) {
		want := "Go App"

		if want != app.Config.App.Name {
			test.Error("Failed to set up app correctly")
		}
	})
}
