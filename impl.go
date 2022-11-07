package hw11

import (
	// "github.com/UofSC-Fall-2022-Math-587-001/homework11/lib"
	"github.com/UofSC-Fall-2022-Math-587-001/homework11/ec"
)

// Check out ec/basic.go for the types and methods you have available for use 

// For n >=0, computes nP on C mod N 
func positiveMult(C ec.EllipticCurve, N ec.Modulus, n int, p ec.Point) ec.Point {
	// Use ec.Add(C,N,-,-) to add points on C mod N 
	return ec.Identity() 
}

// Computes the minimal n such that nP = Identity() 
func Order(C ec.EllipticCurve, N ec.Modulus, p ec.Point) int {
	// The order of the identity is 1. For an non-identity element, 
	// add it to itself iteratively until it = Identity() 
	return 0 
}

// For any integer n, computes nP on C mod N  
func Multiple(C ec.EllipticCurve, N ec.Modulus, n int, p ec.Point) ec.Point {
	// If n >= 0, we can use positiveMult. For n < 0, we need to take (-n)(-P). 
	return ec.Identity() 
}

// ListPoints computes the points on C mod N. 
func ListPoints(C ec.EllipticCurve, N ec.Modulus) []ec.Point {
	// I will get you started with the identity point 
	points := []ec.Point{ec.Identity()}
	// Iterates over all x values and use Tonelli-Shanks to find 
	// the roots of x^3 + a*x + b mod N 
	return points
}

// Returns the number of points on C mod N 
func NumberPoints(C ec.EllipticCurve, N ec.Modulus) int {
	return len(ListPoints(C,N))
}
