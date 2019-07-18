package main

import "testing"

func TestLoop(t *testing.T) {
    number := loop()
    if number != 10 {
       t.Errorf("Number was incorrect, got: %d, want: %d.", number, 10)
    }
}