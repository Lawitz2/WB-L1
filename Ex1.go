package main

import "fmt"

type Human struct {
	AliveStatus bool
}

func (h *Human) checkIfAlive() {
	fmt.Printf("Am i alive? That is %t", h.AliveStatus)
}

type Action struct {
	Human               // встраивание
	SpecificHuman Human // композиция
}

func Ex1() {
	actionExample := Action{}
	actionExample.checkIfAlive()
	actionExample.SpecificHuman.checkIfAlive()
}
