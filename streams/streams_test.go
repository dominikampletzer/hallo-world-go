package streams

import (
	"fmt"
	"strings"
	"testing"
)

func TestMapFilterReduce(t *testing.T) {
	// array of generic interfaces.
	stringSlice := []Any{"a", "b", "c", "1", "D"}

	toUpperCase := func(a Any) Any {
		return strings.ToUpper(a.(string))
	}

	notDigit := func(a Any) bool {
		s := a.(string)
		result := true
		for _, v := range s {
			if v >= '0' && v <= '9' {
				result = false
				break
			}
		}
		return result
		//temp := []rune(a.(string))
		//return unicode.IsDigit(temp[0])
	}

	concat := func(a, b Any) Any {
		return a.(string) + "," + b.(string)
	}

	// Map/Reduce
	result := ToStream(stringSlice).
		Map(toUpperCase).
		Filter(notDigit).
		Reduce(concat).(string)

	if result != "A,B,C,D" {
		t.Error(fmt.Sprintf("Result should be 'A,B,C,D' but is: %v", result))
	}
}

func TestMapFilterReduceWord(t *testing.T) {
	// array of generic interfaces.
	stringSlice := []Any{"a", "a", "b", "b", "D", "a"}

	mapper := func(a Any) Any {
		result := []Pair{{a.(string), 1}}
		return result
	}

	concat := func(a, b Any) Any {
		x := a.([]Pair)
		y := b.([]Pair)

		for i, el := range x {
			for _, pair := range y {
				if el.name == pair.name {
					x[i].count = el.count + pair.count
					return x
				}
			}
		}
		result := append(x, y...)
		return result
	}
	// Map/Reduce
	result := ToStream(stringSlice).
		Map(mapper).
		Reduce(concat).([]Pair)

	for _, e := range result {
		fmt.Printf("%v:%v, ", e.name, e.count)
	}
}
