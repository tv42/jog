package jog_test

import (
	"io/ioutil"
	"testing"

	"github.com/tv42/jog"
)

func BenchmarkSimple(b *testing.B) {
	type xyzzy struct {
		Quux string
		Thud int
	}
	conf := jog.Config{
		Out: ioutil.Discard,
	}
	log := jog.New(&conf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Event(xyzzy{Quux: "foo", Thud: 42})
	}
}
