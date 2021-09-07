package model

import "reflect"

type Chain struct {
	Type   string
	Value  string
	Length int
}

func NewChain(c string) Chain {
	var r Chain

	for i := 0; i < 2; i++ {
		d = reflect.ValueOf(c[i])
		r.Type = reflect.Indirect(d)
	}
	for i := 2; i < 4; i++ {
		e := reflect.ValueOf(c[i])
		r.Length = reflect.Indirect(e).FieldByName("Length")
	}
	for i := 4; i < len(c); i++ {
		f := reflect.ValueOf(c[i])
		r.Value = reflect.Indirect(f).FieldByName("Value")
	}
	return r
}
