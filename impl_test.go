package hw11

import (
	"github.com/UofSC-Fall-2022-Math-587-001/homework11/ec"
	"testing"
	// "fmt"
)

func Test1(t *testing.T) {
	A := 2
	B := 1
	N := 7
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	count := NumberPoints(C,N)
	if count != 5 {
		t.Errorf("The number F_3-points in C is 5 but you said %d.\n",count)
	}
}

func Test2(t *testing.T) {
	A := 2
	B := 0
	N := 11
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:2,Y:10}
	ord := Order(C,N,p)
	want := 6
	if ord != want {
		t.Errorf("The order of %d is %d and not %d\n",p,want,ord)
	}
}

func Test3(t *testing.T) {
	A := 4
	B := 7
	N := 13
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:1,Y:8}
	mult := 4
	q := Multiple(C,N,mult,p)
	r := &ec.FinitePoint{X:5,Y:10}
	if r.X != q.X || r.Y != q.Y {
		t.Errorf("%d*%s = %s but you said = %s",mult,p,r,q)
	}
}

func Test4(t *testing.T) {
	A := 4
	B := 7
	N := 13
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	p := &ec.FinitePoint{X:1,Y:8}
	mult := -4
	q := Multiple(C,N,mult,p)
	r := &ec.FinitePoint{X:5,Y:3}
	if r.X != q.X || r.Y != q.Y {
		t.Errorf("%d*%s = %s but you said = %s",mult,p,r,q)
	}
}

func Test5(t *testing.T) {
	A := 101
	B := 803
	N := 8675309
	C, err := ec.MakeCurve(A,B)
	if err != nil {
		t.Errorf("The discriminant should not be 0.\n")
	}
	count := NumberPoints(C,N)
	want := 8675361
	if count != want {
		t.Errorf("The number of points on %s is %d but you returned %d.\n",C,want,count)
	}
}
