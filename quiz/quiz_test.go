package main

import (
	"encoding/csv"
	"testing"
	"strings"
	"reflect"
)

func TestParseLines(t *testing.T) {
	in := `5+5,10
7+3,10
1+1,2
`
	expected := []problem{
		{question: "5+5", answer: "10"},
		{question: "7+3", answer: "10"},
		{question: "1+1", answer: "2"},
	}

	r := csv.NewReader(strings.NewReader(in))
	lines, _ := r.ReadAll()
	got := parseLines(lines)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("parselines did not return expected results")
	}
}