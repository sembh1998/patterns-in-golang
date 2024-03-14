package main

import "fmt"

type (
	House struct {
		windowType string
		doorType   string
		floor      int
	}
	Builder interface {
		setWindowType()
		setDoorType()
		setNumFloor()
		getHouse() House
	}
	Director struct {
		builder Builder
	}
	NormalBuilder struct {
		windowType string
		doorType   string
		floor      int
	}
	IglooBuilder struct {
		windowType string
		doorType   string
		floor      int
	}
)

// getHouse implements Builder.
func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// setDoorType implements Builder.
func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

// setNumFloor implements Builder.
func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

// setWindowType implements Builder.
func (b *IglooBuilder) setWindowType() {
	b.windowType = "no windows"
}

// getHouse implements Builder.
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// setDoorType implements Builder.
func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

// setNumFloor implements Builder.
func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// setWindowType implements Builder.
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func newNormalBuilder() Builder {
	return &NormalBuilder{}
}

func newIglooBuilder() Builder {
	return &IglooBuilder{}
}

func getBuilder(builderType string) Builder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
