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

// Example output:
//
//     {"Time":"2013-04-23T20:51:17.632909Z","Type":"main#start","Data":{"Enthusiasm":"high"}}
//     {"Time":"2013-04-23T20:51:17.633016Z","Type":"github.com/tv42/jog/example/greetings#greet","Data":{"Message":"Hello, world!","ID":"43235433"}}
//     {"Time":"2013-04-23T20:51:17.633047Z","Type":"main#stop","Data":{"Style":"commendable","Points":9.7}}
func main() {
	log := jog.New(nil)
	log.Event(start{Enthusiasm: "high"})
	greetings.Greet(log)
	log.Event(stop{Style: "commendable", Points: 9.7})
}
