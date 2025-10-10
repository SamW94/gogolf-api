package main

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Sam"
	want := regexp.MustCompile(`\b`+name+`\b`)
	greeting := greet(name)
	if !want.MatchString(greeting) {
		t.Errorf(`greet("Sam") = %q, want match for %#q`, greeting, want)
	}
}