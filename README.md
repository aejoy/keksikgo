<img src="preview.png"/>

<p align="center">Package for «Keksik» API in VK Mini App</p>

<div align="center">
    <a href="https://pkg.go.dev/github.com/aejoy/keksikgo">
        <img src="https://img.shields.io/static/v1?label=go&message=reference&color=00add8&logo=go" />
    </a>
    <a href="http://www.opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/license-MIT-fuchsia.svg" />
    </a>
    <a href="https://goreportcard.com/report/github.com/aejoy/keksikgo">
        <img src="https://goreportcard.com/badge/github.com/aejoy/keksikgo" />
    </a>
</div>

# Installation

```bash
go get github.com/aejoy/keksikgo
```

# Usage

## Methods

```go
package main

import (
    "fmt"
    "os"

    "github.com/aejoy/keksikgo/api"
)

func main() {
    app := api.New(os.Getenv("TOKEN"), 123)
    fmt.Println(app.Balance())
}
```

## Callback

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/aejoy/keksikgo/api"
    "github.com/aejoy/keksikgo/callback"
    "github.com/aejoy/keksikgo/update"
    "github.com/gofiber/fiber/v3"
    "go.uber.org/zap"
)

func main() {
    logger, err := zap.NewProduction()
    if err != nil {
        panic(err)
    }

    cake := api.New(123, os.Getenv("TOKEN"), logger)

    app := fiber.New()
    session := callback.New(cake)

    session.Donate(func(_ *api.API, donate update.Donate) {
        fmt.Printf("New donate: %+v\n", donate)
    })

    app.Post("/callback/:confirmation", session.Fiber)

    log.Fatalln(app.Listen(":3000"))
}
```