package main

import (
	"testing"
)

// name, a, b, want, isErr
var test = []struct {
	name  string
	a     float32
	b     float32
	want  float32
	isErr bool
}{
	{"valid", 10.0, 2.0, 5.0, false},
	{"invalid", 10.5, 0.0, 0.0, true},
}

func TestDivide(t *testing.T) {
	for _, value := range test {
		want, err := divide(value.a, value.b)

		if (err != nil) != value.isErr {
			t.Errorf("%s: got error %v, want error %v", value.name, err, value.isErr)
		}

		if want != value.want {
			t.Errorf("name: %s, got %v, want %v", value.name, want, value.want)
		}
	}
}
