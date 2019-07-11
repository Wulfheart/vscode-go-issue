package main

import (
	"fmt"

	geo "github.com/kellydunn/golang-geo"
)

func main() {
	poly := []*geo.Point{
		geo.NewPoint(0, 0), geo.NewPoint(0, 1), geo.NewPoint(1, 1), geo.NewPoint(1, 0),
	}
	gon := geo.NewPolygon(poly)
	fmt.Println("x:0\ty:0\t", "Expecting: true\t", "Got:", gon.Contains(geo.NewPoint(0, 0)))
	fmt.Println("x:0.5\ty:0\t", "Expecting: true\t", "Got:", gon.Contains(geo.NewPoint(0.5, 0)))
	fmt.Println("x:0.5\ty:0.7\t", "Expecting: true\t", "Got:", gon.Contains(geo.NewPoint(0.5, 0.7)))
	fmt.Println("x:2\ty:8\t", "Expecting: false\t", "Got:", gon.Contains(geo.NewPoint(2, 8)))

}
