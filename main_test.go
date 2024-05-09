package main

import "testing"

var tests = []struct {
	nom   string
	a     float32
	b     float32
	want  float32
	isErr bool
}{
	{"valide", 18, 2, 9, false},
	{"invalide", 9, 0, 0, true},
}

func testDivisiion(t *testing.T) {
	for _, tt := range tests {
		want, err := division(tt.a, tt.b)
		if (err != nil) != tt.isErr {
			t.Errorf("%s: got error %v, want error %v", tt.nom, err, tt.isErr)
		}

		if want != tt.want {
			t.Errorf("%s got %v, want %v", tt.nom, want, tt.want)
		}
	}
}