package main

import (
	"fmt"

	"tp_mod.com/model"
)

func main() {
	c := model.NewChain("TX03SC$$")

	b, err := model.AnaliceChain(&c)

	if b && err == nil {
		fmt.Println(c)
	} else {
		fmt.Println("Error :", err)
	}

}
