package main

import (
	"testing"
)

//nolint
func TestHelloWorld(t *testing.T) {
	var expected Response
	var result Response
	expected.Text = "HelloWorld"
	result = HelloWordTxt()
	if expected.Text != result.Text {
		t.Errorf("Phrase was incorrect. Got : %s, want: %s.", result.Text, expected.Text)
	}
}
