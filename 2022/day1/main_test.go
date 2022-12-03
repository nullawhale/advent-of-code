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

	want := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}
	assert.Equal(t, want, got)
}

func TestStar1(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got, err := star1(data)

	if err != nil {
		t.Errorf("got error: %q", err)
	}
	want := 24000
	assert.Equal(t, want, got)
}

func TestStar2(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got, err := star2(data)
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	want := 45000
	assert.Equal(t, want, got)
}

func TestRowSums(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	_, err = rowSums(data)
	if err != nil {
		t.Errorf("got error: %q", err)
	}
}
