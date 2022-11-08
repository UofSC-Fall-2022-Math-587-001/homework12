package hw12

import (
	"github.com/UofSC-Fall-2022-Math-587-001/homework12/ec"
	"github.com/UofSC-Fall-2022-Math-587-001/homework12/lib"
	"errors"
	"math/rand" // no for use in prod :)
	"fmt"
)

// A struct to represent the public data for ECDSA:
// The prime we are working over, the elliptic curve E, 
// the point P in E(F_p), and the Order of P under 
// addition
type Data struct {
	Prime int
	E ec.EllipticCurve 
	P ec.Point 
	Order int 
}

// Check that 
// - E has nonzero discriminant mod p
// - P is on E 
// - ord is the order of P 
func (D Data) Validate() error {
	return nil 
}

type PrivateKey struct { 
	D Data
	s int // secret key 0 <= s < K.D.Order
}

type PublicKey struct {
	D Data 
	V ec.Point // public key 
}

// PublicKey resturns the public verification key 
// V = sP 
func (K PrivateKey) PublicKey() PublicKey {
	V := ec.Multiple(K.D.E,K.D.Prime,K.s,K.D.P) 
	return PublicKey{K.D, V}
}

type Signature struct{
	s, t int 
}

// Sign the document d using the PrivateKey K 
func (K PrivateKey) Sign(d int) Signature {
	// 1. Get a random integer e using rand.Int() 
	// 2. Compute Q = e*K.D.P 
	// 3. Set s1 = Q.X mod Order  
	// 4. Set s2 = (d + K.s*s1) e^{-1} mod Order
	// 5. Return (s1,s2) 
	return Signature{0,0}
}

// Verify that the Signature is valid for the document d 
// using the K.PublicKey()
func (K PublicKey) Verify(d int, S Signature) error {
	// 1. Compute v1 = d*S.t^{-1} mod K.Order 
	// 2. Compute v2 = S.s*S.t^{-1} mod K.Order 
	// 3. Compute Q := v1*P + v2*K.PublicKey() 
	// 4. Check that Q.X mod Order = S.s
	return errors.New("Forgery")
}
