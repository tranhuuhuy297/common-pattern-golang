package main

import "fmt"

// Factory method is a creational design pattern
// which solves the problem of creating product objects
// without specifying their concrete classes

func main() {
	ak47, _ := GetGun("ak47")
	printDetails(ak47)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
