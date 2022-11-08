package hw11

import (
	// "github.com/UofSC-Fall-2022-Math-587-001/homework11/lib"
	"errors"

	"github.com/UofSC-Fall-2022-Math-587-001/homework11/ec"
	// "errors"
	// "math/rand" // no for use in prod :)
)

// A struct to represent the public data for ECDSA
type StandardData struct {
	Prime int
	E ec.EllipticCurve 
	P ec.Point 
	Order int 
}

// Check that 
// - E has nonzero discriminant mod p
// - P is on E 
// - ord is the order of P 
func (D StandardData) Valid() error {
	return nil 
}

// PrivateKey is a type alias for int 
type PrivateKey struct { 
	s int
	D StandardData
}

// PublicKey resturns the public verification key 
// V = sP 
func (K PrivateKey) PublicKey() ec.Point {
	return nil 
}

type Signature struct{
	s, t int 
}

// Sign the document d using the PrivateKey K 
func (K PrivateKey) Sign(d int) Signature {
	// 1. Get a random integer e using rand.Int() 
	// 2. Compute e*K.D.P 
	// 3. Set s1 = eP.X mod Order  
	// 4. Set s2 = (d + K.s*s1) e^{-1} mod Order
	// 5. Return (s1,s2) 
	return Signature{0,0}
}

// Verify that the Signature is valid for the document d 
// using the K.PublicKey()
func (S Signature) Verify(d int, K PrivateKey) error {
	// 1. Compute v1 = d*S.t^{-1} mod K.Order 
	// 2. Compute v2 = S.s*S.t^{-1} mod K.Order 
	// 3. Compute Q := v1*P + v2*K.PublicKey() 
	// 4. Check that Q.X mod Order = S.s
	// return nil 
	return errors.New("Forgery")
}
