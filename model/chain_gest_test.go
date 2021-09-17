package model

import (
	"strings"
	"testing"
)

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

}
