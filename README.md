# GoFiber + SSE

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
    // Channel for Events
    eventsChan := make(chan *ssefiber.FiberSSEEvent)
    defer close(eventsChan)
    // Create a channel `One`
    channelOne := &ssefiber.FiberSSEChannel{
        Name: "Channel One",
        Base: "/one",
        Events: eventsChan, // Events Channel
    }

    // Call New to add the base routes
    ssefiber.New(app, "/sse", channelOne, ...) // Pass in channels

    // SSE at /sse/one

    // Pass some events
    go func() {
        for i:=10; i<10; i++ {
            newEvent := &ssefiber.FiberSSEEvent{
                Event: "name",
                Data: "data",
            }
            eventsChan <- newEvent
        }
    }()
    app.Listen(":8000")
}


```
