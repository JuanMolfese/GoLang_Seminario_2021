package main

import (
	"errors"
	"fmt"
	"strings"

	"tp_mod.com/model"
)

func AnaliceChain(r *model.Chain) (bool, error) {
	if r.Length == (strings.Count(r.Value, "") - 1) {
		return true, nil
	} else {
		return false, errors.New("la cadena ingresada no es correcta, revise la documentacion")
	}
}

func main() {
	c := model.NewChain("TX04SC$$")

	b, err := AnaliceChain(&c)

	if b && err == nil {
		fmt.Println(c)
	} else {
		fmt.Println("Error :", err)
	}

}
