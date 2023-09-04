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

	want := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
	assert.Equal(t, want, got)
}

func TestStar1(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got := star1(data)
	want := 2
	assert.Equal(t, want, got)
}

func TestStar2(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got := star2(data)
	want := 4
	assert.Equal(t, want, got)
}

func TestCheckFullyOverlap(t *testing.T) {
	right := checkFullyOverlap([]string{"1", "4"}, []string{"2", "3"})
	wrong := checkFullyOverlap([]string{"1", "4"}, []string{"2", "5"})
	assert.True(t, right)
	assert.False(t, wrong)
}

func TestCheckOverlap(t *testing.T) {
	right := checkOverlap([]string{"1", "4"}, []string{"3", "5"})
	wrong := checkOverlap([]string{"1", "4"}, []string{"5", "7"})
	assert.True(t, right)
	assert.False(t, wrong)
}

func TestParsePairs(t *testing.T) {
	data, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	got, err := convertData(data)
	if err != nil {
		t.Errorf("got error: %q", err)
	}
	want := []Pair{
		{left: "2-4", right: "6-8"},
		{left: "2-3", right: "4-5"},
		{left: "5-7", right: "7-9"},
		{left: "2-8", right: "3-7"},
		{left: "6-6", right: "4-6"},
		{left: "2-6", right: "4-8"},
	}

	assert.Equal(t, len(want), len(got))

	if len(want) == len(got) {
		for n, g := range got {
			assert.Equal(t, want[n], g)
		}
	}
}
