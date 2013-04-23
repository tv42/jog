package jog

import (
	"bytes"
	"testing"
	"time"
)

func FixedClock() time.Time {
	return time.Date(2012, 11, 24, 16, 32, 56, 123456789, time.UTC)
}

var FIXED_TIME = FixedClock().Format(time.RFC3339Nano)

func TestExplicit(t *testing.T) {
	// one test avoiding the formatting helpers, just to show they're
	// not buggy
	buf := new(bytes.Buffer)
	conf := Config{
		Out: buf,
	}
	log := New(&conf)
	log.clock = FixedClock
	type xyzzy struct {
		Quux string
		Thud int
	}
	log.Event(xyzzy{Quux: "foo", Thud: 42})
	got := buf.String()
	want := `{"Time":"` + FIXED_TIME + `","Type":"github.com/tv42/jog#xyzzy","Data":{"Quux":"foo","Thud":42}}` + "\n"
	if got != want {
		t.Errorf("wrong output: %q != %s", got, want)
	}
}

func testEvent(t *testing.T, data interface{}, type_ string, want string) {
	buf := new(bytes.Buffer)
	conf := Config{
		Out: buf,
	}
	log := New(&conf)
	log.clock = FixedClock
	log.Event(data)
	got := buf.String()
	want = `{"Time":"` + FIXED_TIME + `","Type":"` + type_ + `","Data":` + want + `}` + "\n"
	if got != want {
		t.Errorf("wrong output: %q != %s", got, want)
	}
}

func TestSimple(t *testing.T) {
	type frob struct {
		Quux string
		Thud int
	}
	testEvent(t, frob{Quux: "foo", Thud: 42},
		"github.com/tv42/jog#frob", `{"Quux":"foo","Thud":42}`)
}

func TestPointer(t *testing.T) {
	type frob struct {
		Quux string
		Thud int
	}
	testEvent(t, &frob{Quux: "foo", Thud: 42},
		"github.com/tv42/jog#frob", `{"Quux":"foo","Thud":42}`)
}

func TestEmpty(t *testing.T) {
	type justMyPresence struct {
	}
	testEvent(t, justMyPresence{},
		"github.com/tv42/jog#justMyPresence", `{}`)
}

func TestNilPointer(t *testing.T) {
	// still a typed nil, not an interface{} nil
	type justMyPresence struct {
	}
	testEvent(t, &justMyPresence{},
		"github.com/tv42/jog#justMyPresence", `{}`)
}
