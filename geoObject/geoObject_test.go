package geoObject

import "testing"

func TestGeoObject(t *testing.T) {
	s := []Painter{
		GeoObject{Position{1, 2}, 3},
		Circle{GeoObject{Position{2, 3}, 4}, 5},
		Rectangle{GeoObject{Position{3, 4}, 5}, 6, 7},
		Triangle{GeoObject{Position{5, 6}, 6}, Position{1, 2}, Position{3, 4}, Position{5, 6}},
	}
	for _, obj := range s {
		obj.Paint()
	}
}
