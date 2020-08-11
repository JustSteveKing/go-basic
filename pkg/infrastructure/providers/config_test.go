package providers

import (
	"testing"
)

func TestConfigProvider(test *testing.T) {
	config := ConfigProvider()

	test.Run("Test App Name is set", func(test *testing.T) {
		want := "Go App"

		if want != config.App.Name {
			test.Error("App Name is incorrect, config provider failed to create the config struct")
		}
	})
}
