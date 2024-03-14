package main

import (
	"errors"
	"fmt"
	"log"
)

type (
	IGun interface {
		setName(name string)
		setPower(power int)
		getName() string
		getPower() int
	}
	Gun struct {
		name  string
		power int
	}
	Ak47 struct {
		Gun
	}
	Musket struct {
		Gun
	}
)

// getName implements IGun.
func (g *Gun) getName() string {
	return g.name
}

// getPower implements IGun.
func (g *Gun) getPower() int {
	return g.power
}

// setName implements IGun.
func (g *Gun) setName(name string) {
	g.name = name
}

// setPower implements IGun.
func (g *Gun) setPower(power int) {
	g.power = power
}

func newAk47() IGun {
	return &Ak47{
		Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

func newMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, errors.New("wrong gun type passed")
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ak47, err := getGun("ak47")
	CheckError(err)
	musket, err := getGun("musket")
	CheckError(err)

	printDetails(ak47)
	printDetails(musket)

	sniper, err := getGun("sniper")
	CheckError(err)
	printDetails(sniper)
}
