// Jog is a structured logging library outputting JSON messages one
// per line.
//
// Each logged event will have the following fields:
//
//     - Time: ISO 8601 format
//     - Type: type of the Go variable representing the event
//     - Data: details of the event
//
// The format for the type is "IMPORTPATH#TYPE", for example
// "github.com/tv42/jog#unencodable".
package jog

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"sync"
	"time"
)

type Config struct {
	// Destination to which log events will be written.
	Out io.Writer

	// This is mostly useful for unit tests.
	Clock func() time.Time
}

func nowUTC() time.Time {
	return time.Now().UTC()
}

// New creates a new Logger. Configuration is optional.
func New(conf *Config) *Logger {
	l := &Logger{}
	if conf != nil {
		// Config is copied so callers can't mess things up.
		l.conf = *conf
	}
	if l.conf.Out == nil {
		l.conf.Out = os.Stdout
	}
	if l.conf.Clock == nil {
		l.conf.Clock = nowUTC
	}
	return l
}

// A Logger represents a logger that takes Go variables as events and
// writes them to an io.Writer as JSON. Each logging operation makes a
// single call to the Writer's Write method and contains one complete
// log event. For high throughput logging, you may want to use a
// bufio.Writer. A Logger can be used simultaneously from multiple
// goroutines, Write calls happen one at a time.
type Logger struct {
	// immutable for the lifetime of the logger
	conf Config

	// prevents interleaved writes
	mu sync.Mutex
}

type header struct {
	Time time.Time
	Type string
	Data interface{} `json:",omitempty"`
}

func (h *header) Set(data interface{}) {
	h.Data = data
	// indirect because we don't want "*foo" for pointer types
	t := reflect.Indirect(reflect.ValueOf(data)).Type()
	h.Type = t.PkgPath() + "#" + t.Name()
}

// Record an event on this Logger. Data may be any value that can be
// JSON encoded.
//
// Do not rely synchronicity of JSON encoding or output. These may
// change.
func (l *Logger) Event(data interface{}) {
	if data == nil {
		panic("nil events are pointless")
	}
	event := header{
		Time: l.conf.Clock(),
	}
	event.Set(data)
	buf, err := json.Marshal(event)
	if err != nil {
		type unencodable struct {
			Error string
		}
		event.Set(unencodable{Error: err.Error()})
		buf, err = json.Marshal(event)
		if err != nil {
			// this really is not expected to happen
			panic("jog: cannot log internal error: " + err.Error())
		}
	}
	buf = append(buf, '\n')
	l.mu.Lock()
	// TODO config option to keep retrying until e.g. space is freed?
	// see daemontools/svlogd for all the possible ugliness
	_, _ = l.conf.Out.Write(buf)
	defer l.mu.Unlock()
}
