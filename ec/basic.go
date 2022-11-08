package ec

import (
	"fmt"
	"errors"
	"github.com/UofSC-Fall-2022-Math-587-001/homework11/lib"
)

// A type alias for integers when we think of it as a modulus 
type Modulus = int

// A type for elliptic curves in Weiestrass form y^2 = x^3 + a*x + b. 
// Fields are private to focus use of Makrve. 
type EllipticCurve struct {
	a int 
	b int 
}

// A constructor for EllipticCurve. It chs the discriminant over Z 
// and errors if it vanishes. 
func Makrve(a,b int) (EllipticCurve,error) {
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

// Chs if C is smooth over Z/NZ (N assumed prime) by reducing the 
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

// Chs whether p lies on the curve after reducing mod N 
func (C EllipticCurve) ChPoint(N Modulus, p Point) bool { 
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

// For n >=0, computes nP on C mod N 
func positiveMult(C EllipticCurve, N Modulus, n int, p Point) Point {
	if n < 0 {
		n *= -1
	}
	q := Identity()
	for n > 0 {
		if n % 2 == 1 {
			q = Add(C,N,p,q)
		}
		p = Add(C,N,p,p) 
		n = n / 2
	}
	return q 
}

// Computes the minimal n such that nP = Identity() 
func Order(C EllipticCurve, N Modulus, p Point) int {
	if p == Identity() {
		return 1
	}
	i := 1 
	q := p 
	for q != Identity() {
		q = Add(C,N,p,q) 
		i += 1
	}
	return i
}

// For any integer n, computes nP on C mod N  
func Multiple(C EllipticCurve, N Modulus, n int, p Point) Point {
	if n < 0 {
		p = p.Inverse()
		n *= -1 
	}
	return positiveMult(C,N,n,p)
}

// ListPoints computes the points on C mod N. It iterates over all x values and 
// used Tonelli-Shanks to find the roots of x^3 + a*x + b
func ListPoints(C EllipticCurve, N Modulus) []Point {
	points := []Point{Identity()}
	for x := 0; x < N; x++ {
		c := library.ModN(uint(N),library.FastPower(uint(N),x,3)+C.A()*x + C.B())
		if c == 0 {
			pt := FinitePoint{X:x,Y:0}
			points = append(points, &pt)
		}
		exists, root := library.TonelliShanks(N,c)
		if exists {
			pt := FinitePoint{X:x,Y:root[0]}
			points = append(points,&pt)
			pt.Y = root[1]
			points = append(points,&pt)
		}
	}
	return points
}

// Returns the number of points on C mod N 
func NumberPoints(C EllipticCurve, N Modulus) int {
	return len(ListPoints(C,N))
}


