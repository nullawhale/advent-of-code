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

	want := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	assert.Equal(t, want, got)
}

func TestStar1(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got := star1(data)
	want := 157
	assert.Equal(t, want, got)
}

func TestStar2(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got := star2(data)
	want := 70
	assert.Equal(t, want, got)
}

func TestStringContainsRune(t *testing.T) {
	got := stringContainsRune("advent", 'd')
	assert.True(t, got)
}

func TestPriority(t *testing.T) {
	got := priority('f')
	assert.Equal(t, 6, got)
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
	want := []Pair{
		{left: "vJrwpWtwJgWr", right: "hcsFMMfFFhFp"},
		{left: "jqHRNqRjqzjGDLGL", right: "rsFMfFZSrLrFZsSL"},
		{left: "PmmdzqPrV", right: "vPwwTWBwg"},
		{left: "wMqvLMZHhHMvwLH", right: "jbvcjnnSBnvTQFn"},
		{left: "ttgJtRGJ", right: "QctTZtZT"},
		{left: "CrZsJsPPZsGz", right: "wwsLwLmpwMDw"},
	}

	assert.Equal(t, len(want), len(got))

	if len(want) == len(got) {
		for n, g := range got {
			assert.Equal(t, want[n], g)
		}
	}
}
