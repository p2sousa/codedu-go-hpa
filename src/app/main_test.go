package main

import "testing"

func TestGreeting(t *testing.T) {
	html := greeting("Code.education Rocks!")
	if html != "<b>Code.education Rocks!</b>" {
		t.Errorf("Text html was incorrect, got: %s, want: %s.", html, "<b>Code.education Rocks!</b>")
	}
}