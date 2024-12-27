package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Initiated")
	ThreeСorners := triangle{
		sideA: 20,
		sideB: 15,
		sideC: 10,
	}
	pointA := point{
		X: 12,
		Y: 10,
	}
	pointB := point{
		X: 18,
		Y: 10,
	}
	shore := shortLine{
		A: pointA,
		B: pointB,
	}
	fmt.Println(shore.A.CenterOffset()) // Расстояние от центра координат до точки
	fmt.Println(shore.B.CenterOffset()) // Расстояние от центра координат до точки
	circ := circle{
		radius: 5,
		center: pointA,
	}
	fmt.Println(circ.Area()) // Площадь круга
	PrintArea(ThreeСorners)
	PrintArea(circ)
}

type point struct {
	X int
	Y int
}
type shortLine struct {
	A point
	B point
}

func (p point) CenterOffset() float64 {
	lat := math.Sqrt(float64((p.X * p.X) + (p.Y * p.Y)))
	return (lat)
}

type triangle struct {
	sideA int
	sideB int
	sideC int
}

func (s triangle) Area() float64 {
	halfP := (s.sideA + s.sideB + s.sideC) / 2
	area := math.Sqrt(float64(halfP * (halfP - s.sideA) * (halfP - s.sideB) * (halfP - s.sideC)))
	fmt.Println(area)
	return area
}
func (s triangle) isReal() bool {
	if s.sideA+s.sideB < s.sideC || s.sideA+s.sideC < s.sideB || s.sideB+s.sideC < s.sideA {
		return false
	} else {
		return true
	}
}

type circle struct {
	radius int
	center point
}

func (s circle) Area() float64 {
	pi := math.Pi
	area := pi * float64(s.radius) * float64(s.radius)
	return area
}

type shape interface {
	Area() float64
}

func PrintArea(s shape) {
	A := s.Area()
	fmt.Printf("Площадь = %.2f\n", A)
}
