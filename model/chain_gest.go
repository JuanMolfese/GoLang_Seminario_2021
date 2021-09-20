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
	Value  string
	Length int
}

func analiceChain(r Chain) (Chain, error) {
	if ((!IsNumber(r.Value) && (strings.ToUpper(r.Type) == "TX")) ||
		(IsNumber(r.Value) && (strings.ToUpper(r.Type) == "NN"))) &&
		(r.Length == (strings.Count(r.Value, "") - 1)) {
		return r, nil
	} else {
		return Chain{"", "", 0}, errors.New("la cadena ingresada no es correcta, revise la documentacion")
	}
}

func NewChain(c string) (Chain, error) {
	var r Chain
	r.Type = c[:MAXTYPE]
	d := c[MAXTYPE:(MAXTYPE + MAXLENGTH)]
	l, _ := strconv.Atoi(d)
	r.Length = l
	r.Value = c[(MAXTYPE + MAXLENGTH):]
	x, err := analiceChain(r)
	return x, err
}

func IsNumber(c string) bool {
	_, err := strconv.Atoi(c)
	return err == nil
}
