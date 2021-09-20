package main

import (
	"fmt"

	"tp_mod.com/model"
)

func main() {
	c, err := model.NewChain("TX06ABCDE")

	//_, err := model.analiceChain(&c)

	if err == nil {
		fmt.Println(c)
	} else {
		fmt.Println("Error :", err)
	}

}
