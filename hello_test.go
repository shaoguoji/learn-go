package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Jim")
    want := "Hello, Jim"

    if got != want {
        t.Errorf("got '%q' wnat '%q'", got, want)
    }
}
