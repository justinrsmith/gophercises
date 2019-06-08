package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
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

func TestReadFile(t *testing.T) {
    var buffer bytes.Buffer
    buffer.WriteString("fake, csv, data")
    content, err := readFile(&buffer)
    if err != nil {
        t.Error("Failed to read csv data")
    }
    fmt.Print(content)
}