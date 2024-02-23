package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_findMax(t *testing.T) {
	slice := []string{"0.99", "2.2", "1.3"}
	want := 2.2
	got, _ := strconv.ParseFloat(findMax(slice), 64)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func Test_findMin(t *testing.T) {
	slice := []string{"0.99", "2.2", "1.3"}
	want := 0.99
	got, _ := strconv.ParseFloat(findMin(slice), 64)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func Test_extractAndParse(t *testing.T) {
	want := []float64{0.99, 2.2, 1.3}
	strfloats := []string{"0.99", "2.2", "1.3"}
	got := extractAndParse(strfloats)
	if !reflect.DeepEqual(got, want) {
		// slices.Equal[float64]()
		t.Errorf("got %f, wanted %f", got, want)
	}
}
