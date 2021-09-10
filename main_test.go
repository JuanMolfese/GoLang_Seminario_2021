package main

import (
	"testing"

	"tp_mod.com/model"
)

func TestAnaliceChain(t *testing.T) {
	c := model.NewChain("CF04POLI")
	_, err := AnaliceChain(&c)
	if (c.Type == "CF" && c.Length == 4 && c.Value == "POLI" && err != nil) || (c.Type != "CF" || c.Length != 4 || c.Value != "POLI" && err == nil) {
		t.Error("La funcion AnaliceChain no esta funcionando bien")
	}
}
