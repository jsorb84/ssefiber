# GoFiber + SSE

[![Go Reference](https://pkg.go.dev/badge/github.com/jsorb84/ssefiber.svg)](https://pkg.go.dev/github.com/jsorb84/ssefiber)

My attempt at a basic wrapper to easily allow SSE with a GoFiber application

## Usage

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jsorb84/ssefiber"
)

func main() {
    app := fiber.New(...)
    sse := ssefiber.New(app, "/sse")

    // Create a channel `One`
    channelOne := sse.CreateChannel("Channel One", "/one")

    // Pass some events
    go func() {
        for i:=10; i<10; i++ {
            channelOne.PushEvent("Event Name", "Event Data")
        }
    }()
    app.Listen(":8000")
}


```
