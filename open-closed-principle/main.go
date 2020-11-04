package main

import "fmt"

// OCP
// open for extension, closed modification
// Specification

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

type Specification interface {
	isSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) isSatisfied(p *Product) bool {
	return a.first.isSatisfied(p) && a.second.isSatisfied(p)
}

func (s SizeSpecification) isSatisfied(p *Product) bool {
	return p.size == s.size
}

func (c ColorSpecification) isSatisfied(p *Product) bool {
	return p.color == c.color
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.isSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green Products \n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green \n", v.name)
	}

	fmt.Printf("Green Products New \n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}

	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green \n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large Green Products \n")

	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s large and green \n", v.name)
	}

}
