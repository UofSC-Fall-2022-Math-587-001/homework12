package ec

import (
	"fmt"
	"errors"
	"github.com/UofSC-Fall-2022-Math-587-001/homework11/lib"
)

// A type alias for integers when we think of it as a modulus 
type Modulus = int

// A type for elliptic curves in Weiestrass form y^2 = x^3 + a*x + b. 
// Fields are private to focus use of MakeCurve. 
type EllipticCurve struct {
	a int 
	b int 
}

// A constructor for EllipticCurve. It checks the discriminant over Z 
// and errors if it vanishes. 
func MakeCurve(a,b int) (EllipticCurve,error) {
	C := new(EllipticCurve)
	C.a = a 
	C.b = b 
	if 4*a*a*a+27*b*b == 0 {
		return EllipticCurve{a,b},errors.New("The curve is singular")
	} 
	return EllipticCurve{a,b},nil
}

// The discriminant of the cubic equation 
func (C EllipticCurve) Discriminant() int {
	return 4*C.A()*C.A()*C.A()+27*C.B()*C.B()
}

// A method to extract the a coefficient in y^2 = x^3 + a*x + b
func (C EllipticCurve) A() int {
	return C.a 
}

// A method to extract the b coefficient in y^2 = x^3 + a*x + b
func (C EllipticCurve) B() int {
	return C.b 
}

// Checks if C is smooth over Z/NZ (N assumed prime) by reducing the 
// discriminant mod N 
func (C EllipticCurve) IsSmooth(N Modulus) bool {
	if C.Discriminant() % N == 0 {
		return false 
	} 
	return true 
}

// Satisfies the Stringer interface 
func (C EllipticCurve) String() string {
	return fmt.Sprintf("Elliptic curve in Weierstrass form given by the equation y^2 = x^3 + %v*x + %v", C.a, C.b)
}

// A type for the points of an elliptic curve excluding the identity. 
// Note that it does not guarantee that the point is actually on the curve 
type FinitePoint struct {
	X int 
	Y int 
}

// A type for all points on an elliptic curve. The nil pointer is used to 
// model the point at infinity. 
type Point = *FinitePoint

// Returns the identity point of the elliptic curve
func Identity() Point {
	var pt Point
	return pt 
}

// Returns the inverse of a point of an elliptic curve
func (p Point) Inverse() Point {
	if p == Identity() {
		return Identity()
	} 
	p.Y = -p.Y
	return p 
}

// Checks whether p lies on the curve after reducing mod N 
func (C EllipticCurve) CheckPoint(N Modulus, p Point) bool { 
	if p == Identity() {
		return true
	}
	c := library.FastPower(uint(N),p.X,3)+C.A()*p.X + C.B() 
	if library.FastPower(uint(N),p.Y,2) == library.ModN(uint(N),c){
		return true 
	}
	return false 
}

// Satisfies the Stringer interface
func (p Point) String() string {
	if p != nil {
		return fmt.Sprintf("(%d,%d)",p.X,p.Y)
	}
	return fmt.Sprintf("𝒪")
}

// Addition of points on EllipticCurve. Assumes that p and q are on C 
// mod N
func Add(C EllipticCurve, N Modulus, p, q Point) Point {
	if p == Identity() {
		return q 
	} else if q == Identity() {
		return p 
	}
	switch library.ModN(uint(N),p.X) == library.ModN(uint(N),q.X) {
	case true: 
		if library.ModN(uint(N),p.Y) == library.ModN(uint(N),-q.Y) {
			return Identity()
		} 
		num := library.ModN(uint(N),3*library.FastPower(uint(N),p.X,2)+C.A())
		denom := library.ModN(uint(N),2*p.Y) 
		lam := library.ModN(uint(N),library.Inverse(uint(N),denom)*num)
		x := library.ModN(uint(N),library.FastPower(uint(N),lam,2)-2*p.X)
		y := library.ModN(uint(N), lam*(p.X-x) - p.Y) 
		pt := FinitePoint{X:x,Y:y}
		return &pt 
	default: 
		num := library.ModN(uint(N),q.Y - p.Y) 
		denom := library.ModN(uint(N),q.X - p.X)
		lam := library.ModN(uint(N),library.Inverse(uint(N),denom)*num)
		x := library.ModN(uint(N),library.FastPower(uint(N),lam,2)-p.X-q.X) 
		y := library.ModN(uint(N),lam*(p.X-x)-p.Y) 
		pt := FinitePoint{X:x,Y:y}
		return &pt 
	}
}


