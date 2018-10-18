package geoObject

import "fmt"

type Position struct {
	x int
	y int
}

type Painter interface {
	Paint()
}

type GeoObject struct {
	Position
	color int
}

func (geoObject GeoObject) Paint() {
	fmt.Printf("GeoObject: x: %v, y: %v, color: %v \n", geoObject.x, geoObject.y, geoObject.color)
}

type Circle struct {
	GeoObject
	radius int
}

func (circle Circle) Paint() {
	fmt.Printf("Circle: x: %v, y: %v, color: %v, radius: %v \n", circle.x, circle.y, circle.color, circle.radius)

}

type Rectangle struct {
	GeoObject
	width  int
	height int
}

func (rec Rectangle) Paint() {
	fmt.Printf("Rectangle: x:%v, y:%v, color:%v, width:%v, height:%v \n", rec.x, rec.y, rec.color, rec.width, rec.height)
}

type Triangle struct {
	GeoObject
	posA Position
	posB Position
	posC Position
}

func (tri Triangle) Paint() {
	fmt.Printf("Triangle: x: %v, y: %v, color: %v, posA: (x:%v,y:%v),  posB: (x:%v,y:%v),  posC: (x:%v,y:%v) \n",
		tri.x, tri.y, tri.color,
		tri.posA.x, tri.posA.y,
		tri.posB.x, tri.posB.y,
		tri.posC.x, tri.posC.y)
}
