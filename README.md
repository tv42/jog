==========================================
 jog -- Structured logging library for Go
==========================================

Jog is a simple, minimal, logging library for Go. It's meant for
backend services that generate a lot of logs, where ease of processing
is more important than human eyeball friendliness.

- the right way is the easy way; event types are differentiated for
  you
- log rotation, compression, and such belong outside the process
- don't repeat identical values in every log message, e.g. hostname
  (these can be added at indexing time)
- can be (realtime) postprocessed for more readable messages

Usage:

    log := jog.New(nil)

    type cowbell struct {
	    More bool
    }
    log.Event(cowbell{More: true})

    type sound struct {
	    Description string
    }
    log.Event(sound{Description: "dynamite"})
    log.Event(sound{Description: "loud"})
    log.Event(sound{Description: "distracting"})

Define a new type for every type of event you emit. Don't reuse the
same name within the package (Go helps you with this! They're already
identifiers!).
