package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchInput(t *testing.T) {
	got, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}

	want := []string{"A Y", "B X", "C Z"}
	assert.Equal(t, want, got)
}

func TestStar1(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got := star1(data)
	want := 15
	assert.Equal(t, want, got)
}

func TestStar2(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got := star2(data)
	want := 12
	assert.Equal(t, want, got)
}

func TestParsePairs(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got := parsePairs(data)
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	want := []Pair{{elf: 0, me: 1}, {elf: 1, me: 0}, {elf: 2, me: 2}}
	for n, g := range got {
		assert.Equal(t, g, want[n])
	}
}
