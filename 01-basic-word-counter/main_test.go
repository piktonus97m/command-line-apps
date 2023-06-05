package main

import (
	"bytes"
	"testing"
)

// Testcountwords test the count function set to count words

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 world2 word3 word4\n")

	exp := 4

	res := count(b)

	if res != exp {
		t.Errorf("Expectd %d, got %d instead.\n", exp, res)
	}
}
