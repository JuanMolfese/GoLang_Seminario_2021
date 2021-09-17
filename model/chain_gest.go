package model

import (
	"errors"
	"strconv"
	"strings"
)

const (
	MAXTYPE   int = 2
	MAXLENGTH int = 2
)

type Chain struct {
	Type   string
	Length int
	Value  string
}

func NewChain(c string) Chain {
	var r Chain
	r.Type = c[:MAXTYPE]
	d := c[MAXTYPE:(MAXTYPE + MAXLENGTH)]
	l, _ := strconv.Atoi(d)
	r.Length = l
	r.Value = c[(MAXTYPE + MAXLENGTH):]
	return r
}

func AnaliceChain(r *Chain) (bool, error) {
	if r.Length == (strings.Count(r.Value, "") - 1) {
		return true, nil
	} else {
		return false, errors.New("la cadena ingresada no es correcta, revise la documentacion")
	}
}
