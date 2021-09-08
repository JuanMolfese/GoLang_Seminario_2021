package model

import (
	"errors"
	"strconv"
	"strings"
)

const (
	MAXTYPE  int = 2
	MAXVALUE int = 2
)

type Chain struct {
	Type   string
	Length int
	Value  string
}

func AnaliceChain(r Chain) bool {
	return r.Length == (strings.Count(r.Value, " ") - 1)
}

func NewChain(c string) (Chain, error) {
	var r Chain

	r.Type = c[:MAXTYPE]
	d := c[MAXTYPE:(MAXTYPE + MAXVALUE)]
	l, _ := strconv.Atoi(d)
	r.Length = l
	r.Value = c[(MAXTYPE + MAXVALUE):]
	if AnaliceChain(r) {
		return r, nil
	}
	return r, errors.New("La cadena ingresada no es correcta, revise la documentacion")
}
