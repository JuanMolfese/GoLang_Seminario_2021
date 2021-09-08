package model

import (
	"strconv"
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
