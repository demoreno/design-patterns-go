package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {
	panic("Implement me")
}

func (m MultiFunctionPrinter) Fax(d Document) {
	panic("Implement me")
}

func (m MultiFunctionPrinter) Scan(d Document) {
	panic("Implement me")
}

type OldFashionedPrinter struct {
}

// ISP

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Print(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {
	panic("implement me")
}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	panic("implement me")
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func main() {

}
