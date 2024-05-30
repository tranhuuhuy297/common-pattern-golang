package main

// Bridge is a structural design pattern
// that divides business logic or huge class into separate class hierarchies
//that can be developed independently

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
}
