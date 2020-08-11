package factories

import "testing"

func TestConfig(test *testing.T) {
	config := ConfigFactory()

	test.Run("Test App Config", func(test *testing.T) {
		want := "Go App"
		if want != config.App.Name {
			test.Error("App Name does not match expected value")
		}

		want = "v1.0"
		if want != config.App.Version {
			test.Error("App Version does not match expected value")
		}

		want = ":8080"
		if want != config.App.Port {
			test.Error("App Port does not match expected value")
		}
	})

	test.Run("Test HTTP Config", func(test *testing.T) {
		want := "application/vnd.api+json"
		if want != config.HTTP.Content {
			test.Error("HTTP Content does not match expected value")
		}

		want = "application/problem+json"
		if want != config.HTTP.Problem {
			test.Error("HTTP Problem does not match expected value")
		}
	})
}
