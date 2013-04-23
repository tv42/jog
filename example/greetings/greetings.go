package greetings

import (
	"github.com/tv42/jog"
)

type greet struct {
	Message string
	ID      uint64 `json:",string"`
}

func Greet(log *jog.Logger) {
	log.Event(greet{ID: 43235433, Message: "Hello, world!"})
}
