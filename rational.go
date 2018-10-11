package hello

import "fmt"

type Rational struct {
	nummerator int
	deminator  int
}

func NewRational(nummerator int, deminator int) Rational {
	if deminator == 0 {
		panic("Teilen durch 0")
	}
	var ggt = gcd(nummerator, deminator)
	r := Rational{}
	r.deminator = deminator / ggt
	r.nummerator = nummerator / ggt
	return r
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func (a Rational) String() string {
	return fmt.Sprintf("(%v/%v", a.nummerator, a.deminator)
}

func (a Rational) Add(b Rational) Rational {
	return NewRational((a.nummerator*b.deminator)+(a.deminator*b.nummerator), a.deminator*b.deminator)
}

func (a Rational) Multilicate(b Rational) Rational {
	return NewRational(a.nummerator*b.nummerator, a.deminator*b.deminator)
}
