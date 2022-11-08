package hw12

import (
	"github.com/UofSC-Fall-2022-Math-587-001/homework12/ec"
	"testing"
	"fmt"
)

func Test1(t *testing.T) {
	A := 100
	B := 101
	N := 223
	X := 170 
	Y := 169 
	s := 50 
	d := 40 
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:X,Y:Y}
	fmt.Printf("P = %s\n",p) 
	ord := ec.Order(C,N,p)
	fmt.Printf("The order of %s is %d\n",p,ord) 
	D := Data{N,C,p,ord}
	K := PrivateKey{D,s}
	V := K.PublicKey()
	signed := K.Sign(d) 
	checked := V.Verify(d,signed) 
	if checked != nil {
		t.Errorf("The signature cannot be verified\n")
	}
}

func Test2(t *testing.T) {
	A := 1
	B := 1	
	N := 1627
	X := 1502 
	Y := 219 
	s := 45 
	d := 31 
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:X,Y:Y}
	fmt.Printf("P = %s\n",p) 
	ord := ec.Order(C,N,p)
	fmt.Printf("The order of %s is %d\n",p,ord) 
	D := Data{N,C,p,ord}
	K := PrivateKey{D,s}
	V := K.PublicKey()
	signed := K.Sign(d) 
	checked := V.Verify(d,signed) 
	if checked != nil {
		t.Errorf("The signature cannot be verified\n")
	}
}

func Test3(t *testing.T) {
	A := 200
	B := 300	
	N := 7919
	X := 7893 
	Y := 6660  
	s := 1541 
	d := 1776
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:X,Y:Y}
	fmt.Printf("P = %s\n",p) 
	ord := ec.Order(C,N,p)
	fmt.Printf("The order of %s is %d\n",p,ord) 
	D := Data{N,C,p,ord}
	K := PrivateKey{D,s}
	V := K.PublicKey()
	signed := K.Sign(d) 
	checked := V.Verify(d,signed) 
	if checked != nil {
		t.Errorf("The signature cannot be verified\n")
	}
}

func Test4(t *testing.T) {
	A := 200
	B := 300	
	N := 7919
	X := 7893 
	Y := 6660  
	s := 1541 
	d := 1776
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:X,Y:Y}
	fmt.Printf("P = %s\n",p) 
	ord := ec.Order(C,N,p)
	fmt.Printf("The order of %s is %d\n",p,ord) 
	D := Data{N,C,p,ord}
	K := PrivateKey{D,s}
	V := K.PublicKey()
	signed := Signature{440,440} 
	checked := V.Verify(d,signed) 
	if checked == nil {
		t.Errorf("The signature should not be verified\n")
	}
}

func Test5(t *testing.T) {
	A := 1
	B := 1	
	N := 1627
	X := 1502 
	Y := 219 
	s := 45 
	d := 31 
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:X,Y:Y}
	fmt.Printf("P = %s\n",p) 
	ord := ec.Order(C,N,p)
	fmt.Printf("The order of %s is %d\n",p,ord) 
	D := Data{N,C,p,ord}
	K := PrivateKey{D,s}
	V := K.PublicKey()
	signed := Signature{234,444}
	checked := V.Verify(d,signed) 
	if checked == nil {
		t.Errorf("The signature should not be verified\n")
	}
}
