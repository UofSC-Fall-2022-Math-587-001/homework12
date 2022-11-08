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
	fmt.Printf("We are signing d = %d with p = %d, E = %s, P = %s, and q = %d\n",d,K.D.Prime,K.D.E,K.D.P,K.D.Order)
	fmt.Printf("The private key is %d\n",K.s) 
	// 1. Get a random integer e using rand.Int() 
	e := 0 
	for e % K.D.Order == 0 {
		e = rand.Int()
	}
	inv := lib.Inverse(uint(K.D.Order),e)
	e = lib.ModN(uint(K.D.Order),e)
	fmt.Printf("We get a random e = %d and its inverse %d mod %d\n",e,inv,K.D.Order) 
	// 2. Compute Q = e*K.D.P 
	Q := ec.Multiple(K.D.E,K.D.Prime,e,K.D.P) 
	fmt.Printf("We compute e*P = %s\n",Q)
	// 3. Set s1 = Q.X mod Order  
	s1 := lib.ModN(uint(K.D.Order),Q.X) 
	fmt.Printf("We compute s1 = Q.X = %d mod %d\n",s1,K.D.Order)
	// 4. Set s2 = (d + K.s*s1) e^{-1} mod Order
	s2 := (d + K.s*s1)*inv
	s2 = lib.ModN(uint(K.D.Order),s2) 
	fmt.Printf("We compute s2 = (d+K.s*s1)e^{-1} = %d mod %d\n",s2,K.D.Order) 
	// 5. Return (s1,s2) 
	return Signature{s1,s2}
}

// Verify that the Signature is valid for the document d 
// using the K.PublicKey()
func (K PublicKey) Verify(d int, S Signature) error {
	// 1. Compute v1 = d*S.t^{-1} mod K.Order 
	fmt.Printf("We are verifying the signature %d on d = %d\n",S,d) 
	fmt.Printf("Our public key is %s\n",K.V) 
	inv := lib.Inverse(uint(K.D.Order),S.t)
	v1 := lib.ModN(uint(K.D.Order),d*inv) 
	fmt.Printf("We first compute d*t^{-1} = %d*%d = %d mod %d\n",d,S.t,v1,K.D.Order)
	// 2. Compute v2 = S.s*S.t^{-1} mod K.Order 
	v2 := lib.ModN(uint(K.D.Order),S.s*inv) 
	fmt.Printf("We next compute s*t^{-1} = %d*%d = %d mod %d\n",S.s,S.t,v2,K.D.Order)
	// 3. Compute Q := v1*P + v2*K.PublicKey() 
	Q1 := ec.Multiple(K.D.E,K.D.Prime,v1,K.D.P)
	Q2 := ec.Multiple(K.D.E,K.D.Prime,v2,K.V)
	Q := ec.Add(K.D.E,K.D.Prime,Q1,Q2)
	fmt.Printf("We then compute Q = v1*P + v2*V = %s\n",Q) 
	// 4. Check that Q.X mod Order = S.s
	b := lib.ModN(uint(K.D.Order),Q.X) == lib.ModN(uint(K.D.Order),S.s)  
	fmt.Printf("We compare %d and %d mod %d. Which is %t\n",Q.X,S.s,K.D.Order,b)
	if b {
		return nil
	}
	return errors.New("Forgery")
}
