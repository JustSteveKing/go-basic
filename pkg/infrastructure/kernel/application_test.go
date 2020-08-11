package kernel

import "testing"

func TestApplicationCanBeBooted(test *testing.T) {
	app := Boot()

	want := "Go App"

	if want != app.Config.App.Name {
		test.Error("Application failed to be booted up")
	}
}
