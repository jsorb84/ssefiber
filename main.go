// A very basic implementation of SSE for GoFiber

package ssefiber

import (
	"bufio"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)
type FiberSSEEvent struct {
	Timestamp time.Time `json:"timestamp"`
	ID string `json:"id"`
	Event string `json:"event"`
	Data string `json:"data"`
	Retry string `json:"retry"`
}
type FiberSSEChannel struct {
	Name   string
	Base   string
	Events chan *FiberSSEEvent
	
}


// Write the event to the writer `w` - formats according to SSE standard
func (e *FiberSSEEvent) WriteEvent(w *bufio.Writer) {
	fmt.Fprintf(w, "event: %s\ndata: %s\n\n", e.Event, e.Data)
	err := w.Flush()
	if err!= nil {
		panic(err)
	}
}
// Prints the channel information to the console
func (c *FiberSSEChannel) Print() {
	fmt.Printf("==CHANNEL CREATED==\nName: %s\nRoute: %s\n===================", c.Name, c.Base)
}

type FiberSSEHandler func(c *fiber.Ctx, w *bufio.Writer) error

// New initializes a base SSE route group at `base`.
//
// The base route is the base path for all channels.
//
// The channels parameter is a list of channels that will be created.
// Each channel has a name, a base route, and a channel for sending events.
// 
// ### Example:
//
//  ```go 
//  app := fiber.New()
//  eventChan := make(chan *ssefiber.FiberSSEEvent)
//  
//  chanOne := &ssefiber.FiberSSEChannel{
//  	Name: "Channel One",
//  	Base: "/one",
//  	Events: eventChan,
//  } // Create the channel at /sse/one
//  defer close(eventChan)
//  ssefiber.New(app, "/sse", chanOne)
//  ```
func New(app *fiber.App, base string, channels ...*FiberSSEChannel) {
	app.Add("GET", base, func(c *fiber.Ctx) error {
		return nil
	})
	for _, channel := range channels {
		channel.Print()
		app.Add("GET", base+channel.Base, channel.CreateRoute())
	}
}
// CreateRoute returns a fiber.Handler for the channel.
func (fChan *FiberSSEChannel) CreateRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache")
		c.Set("Content-Type", "text/event-stream")
		c.Set("Connection", "keep-alive")
		c.Set("Access-Control-Allow-Origin", "*")
		c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
			for {
				event := <-fChan.Events
				// fmt.Fprintf(w, "event: %s\ndata: %s\n\n", string(event.Event), string(event.Data))
				// w.Flush()
				event.WriteEvent(w)
			}
		})

		return nil
	}
}