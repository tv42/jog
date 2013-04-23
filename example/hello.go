package main

import (
	"github.com/tv42/jog"
	"github.com/tv42/jog/example/greetings"
)

type start struct {
	Enthusiasm string
}

type stop struct {
	Style  string
	Points float32
}

func main() {
	log := jog.New(nil)
	log.Event(start{Enthusiasm: "high"})
	greetings.Greet(log)
	log.Event(stop{Style: "commendable", Points: 9.7})
}
