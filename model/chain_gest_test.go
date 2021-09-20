package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func TestNewChain(t *testing.T) {
	c := NewChain("CF04POLI")
	if c.Type != "CF" || c.Length != 4 || c.Value != "POLI" {
		t.Error("Error de TESTING, uno de los parametros se carga mal")
	}
}

func TestAnaliceChain(t *testing.T) {
	c := NewChain("CF04POLI")
	_, err := AnaliceChain(&c)
	if (c.Type == "CF" && c.Length == 4 && c.Value == "POLI" && c.Length != (strings.Count(c.Value, "")-1) && err != nil) || (c.Type != "CF" || c.Length != 4 || c.Value != "POLI" && c.Length == (strings.Count(c.Value, "")-1) && err == nil) {
		t.Error("La funcion AnaliceChain no esta funcionando bien")

	}

}
func TestAnaliceChainFalse(t *testing.T) {
	c := NewChain("CF03POLI")
	_, err := AnaliceChain(&c)
	if (c.Type == "CF" && c.Length == 3 && c.Value == "POLI" && c.Length != (strings.Count(c.Value, "")-1) && err == nil) || (c.Type != "CF" || c.Length != 3 || c.Value != "POLI" && c.Length == (strings.Count(c.Value, "")-1) && err != nil) {
		t.Error("La funcion AnaliceChain no esta funcionando bien")
	}

}*/

func TestNewChain(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {

		r, err := NewChain(testData.Input)

		assert.Equal(t, err == nil, testData.Success)
		if err != nil {
			assert.Equal(t, testData.Type, r.Type)
		}

		//
	}
}
