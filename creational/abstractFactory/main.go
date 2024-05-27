package main

import "fmt"

// Abstract Factory is a creational design pattern
// which solves the problem of creating entire product families
// without specifying their concrete classes.

func main() {
	adidasFactory, _ := GetSportFactory("adidas")
	adidasShoe := adidasFactory.makeShoe()
	AdidasShirt := adidasFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShirtDetails(AdidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
