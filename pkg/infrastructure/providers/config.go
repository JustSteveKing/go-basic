package providers

import (
	"github.com/JustSteveKing/go-basic/pkg/infrastructure/factories"
)

// ConfigProvider will run and create a new Config struct from the Config Factory
func ConfigProvider() *factories.Config {
	return factories.ConfigFactory()
}
