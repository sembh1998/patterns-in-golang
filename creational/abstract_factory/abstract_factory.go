package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type (
	ISportFactory interface {
		makeShoe() IShoe
		makeShirt() IShirt
	}
	Adidas struct{}
	Nike   struct{}
	IShoe  interface {
		setLogo(logo string)
		setSize(size int)
		getLogo() string
		getSize() int
	}
	Shoe struct {
		logo string
		size int
	}
	AdidasShoe struct {
		Shoe
	}
	NikeShoe struct {
		Shoe
	}

	IShirt interface {
		setLogo(logo string)
		setSize(size int)
		getLogo() string
		getSize() int
	}
	Shirt struct {
		logo string
		size int
	}
	AdidasShirt struct {
		Shirt
	}
	NikeShirt struct {
		Shirt
	}
)

// getLogo implements IShirt.
func (s *Shirt) getLogo() string {
	return s.logo
}

// getSize implements IShirt.
func (s *Shirt) getSize() int {
	return s.size
}

// setLogo implements IShirt.
func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

// setSize implements IShirt.
func (s *Shirt) setSize(size int) {
	s.size = size
}

// getLogo implements IShoe.
func (s *Shoe) getLogo() string {
	return s.logo
}

// getSize implements IShoe.
func (s *Shoe) getSize() int {
	return s.size
}

// setLogo implements IShoe.
func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

// setSize implements IShoe.
func (s *Shoe) setSize(size int) {
	s.size = size
}

func GetSportsFactory(brand string) (ISportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, errors.New("wrong brand type passed")
}

//ADIDAS

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

//NIKE

func (a *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (a *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s, Size: %d, Elem: %v, Type: %v\n", s.getLogo(), s.getSize(), reflect.TypeOf(&s).Elem(), reflect.TypeOf(s).Elem())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s, Size: %d, Elem: %v, Type: %v\n", s.getLogo(), s.getSize(), reflect.TypeOf(&s).Elem(), reflect.TypeOf(s).Elem())
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	adidasFactory, err := GetSportsFactory("adidas")
	handleErr(err)
	nikeFactory, err := GetSportsFactory("nike")
	handleErr(err)

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}
