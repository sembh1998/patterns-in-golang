package main

import "fmt"

type (
	Printer interface {
		PrintFile()
	}
	Epson    struct{}
	Hp       struct{}
	Computer interface {
		Print()
		SetPrinter(Printer)
	}
	Mac struct {
		printer Printer
	}
	Windows struct {
		printer Printer
	}
)

// Print implements Computer.
func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

// SetPrinter implements Computer.
func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()

	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
}
