package main

import (
	"reflect"
	"testing"
)

func TestFetchInput(t *testing.T) {
	got, err := fetchInput("i_test")
	if err != nil {
		t.Errorf("got error: %q", err)
	}

	want := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
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
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
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
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
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
