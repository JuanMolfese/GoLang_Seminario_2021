package model

import "testing"

func TestNewChain(t *testing.T) {
	c := NewChain("CF04POLI")
	if c.Type != "CF" || c.Length != 4 || c.Value != "POLI" {
		t.Error("Error de TESTING, uno de los parametros se carga mal")
	}
}
