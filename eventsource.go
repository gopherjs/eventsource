// Package eventsource provides low-level bindings for the browser's EventSource API,
// which is used to receive server-sent events.
package eventsource

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/util"
)

// EventSource is used to receive server-sent events.
//
// It connects to a server over HTTP and receives events in text/event-stream format without closing the connection.
type EventSource struct {
	o *js.Object

	// Available event types are "open", "error", "message".
	util.EventTarget

	ReadyState ReadyState `js:"readyState"`
	URL        string     `js:"url"`
}

// New creates a new EventSource object. It returns immediately, without waiting to connect.
func New(url string) *EventSource {
	es := js.Global.Get("EventSource").New(url)
	return &EventSource{
		o:           es,
		EventTarget: util.EventTarget{Object: es},
	}
}

// Close the connection, if any, and set the readyState attribute to Closed.
//
// If the connection is already closed, the method does nothing.
func (es *EventSource) Close() {
	es.o.Call("close")
}

// ReadyState is the state of the connection.
type ReadyState int

const (
	Connecting ReadyState = 0 // The connection is being established.
	Open       ReadyState = 1 // The connection is open and dispatching events.
	Closed     ReadyState = 2 // The connection is not being established, has been closed or there was a fatal error.
)

// String returns a human-readable representation of the ReadyState.
func (rs ReadyState) String() string {
	switch rs {
	case Connecting:
		return "Connecting"
	case Open:
		return "Open"
	case Closed:
		return "Closed"
	default:
		return "Unknown"
	}
}
