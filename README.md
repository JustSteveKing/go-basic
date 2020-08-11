# Go Basic

![CI pipeline](https://github.com/JustSteveKing/go-basic/workflows/CI%20pipeline/badge.svg?branch=master)

A basic go api framework.

## Usage

To use this basic API framework in your application simply:

```go
package main

import (
    "github.com/JustSteveKing/go-basic/pkg/infrastructure/kernel"
)

func main() {
    // Load env in whichever way works for you

    // Create your application
    app := kernel.Boot()

    // Load in your routing here. Here is an example I use
    manager.SetupHealthCheckService(app)
    manager.SetupDomainSpecificService(app)
    // ... and so on ... //

    // Run our application
    go func () {
        app.Run()
    }()

    // Wait for any termination signals
    app.WaitForShutdown()
}
```

This "framework" uses the following dependencies:

- [Application Logger](https://go.uber.org/zap)
- [Application Router](https://github.com/gorilla/mux)
- [Application CORS Handler](https://github.com/gorilla/handlers)
- [Application API Response for Errors](https://github.com/JustSteveKing/go-http-response)
- [Application General APIError](https://github.com/JustSteveKing/go-api-problem)