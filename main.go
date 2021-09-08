package main

import (
	"fmt"
	"tp_mod/model"
)

func main() {
	c, err := model.NewChain("TX04SABCD")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println(c)
}
