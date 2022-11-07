package library

import "math"

// A modification of the built-in modulus % which 
// always returns a positive remainder 
func ModN(N uint, i int) int {
	m := i % int(N) 
	if m < 0 {
		m += int(N) 
	}
	return m
}

// FastPower(N,g,A) computes g^A mod N. Note that it
// assumes that N and A are non-negative. 
func FastPower(N uint, g int, A uint) int {
	var b int 
	a := g 
	b = 1 
	if A < 0 {
		A = -A 
	}
	for A > 0 {
		if A % 2 == 1 {
			b = ModN(N,b*a)
		} 
		a = ModN(N,a*a) 
		A = A / 2 
	}
	return b
}

// A separate type for the output of the Euclidean algorithm 
type EuclidData struct {
	GCD int 
	U int 
	V int
}

// Given two integers a and b, GCD(a,b) returns g,u,v where 
// g is the gcd(a,b) and au+bv = g 
func EuclidAlgo(a, b int) EuclidData {
	// if b = 0, then the gcd is a 
	if b == 0 {
		return EuclidData{a,1,0}
	}

	// Keeps tracks of the sign of a and b and makes sure 
	// a and b are non-negative 
	var na , nb   bool 
	if a < 0 {
		na = true 
		a *= -1 
	}
	if b < 0 {
		nb = true
		b *= -1 
	}

	// Variables we will return 
	var g, u, v int 
	u = 1 
	g = a
	x := 0  // keeps track of the number a's used in the Euclidean algorithm 
	y := b  // keeps track of the denominator in the Euclidean algorithm
	for y != 0 {
		t := ModN(uint(y),g) // find t,q with g = qy + t 
		q := g / y 
		s := u - q*x  
		u = x 
		g = y 
		x = s 
		y = t 
	}
	v = (g-a*u)/b  
	
	if !na && !nb {
		return EuclidData{g, u, v} 
	} else if !na && nb {
		return EuclidData{g, u, -v}
	} else if na && !nb {
		return EuclidData{g, -u, v}
	} else {
		return EuclidData{g , -u, -v}
	}
}

// Computes a such that ax = 1 mod N if gcd(a,N) = 1. Else 
// it returns -1 which serves a "junk value"

func Inverse(N uint, x int) int {
	d := EuclidAlgo(x,int(N))
	if d.GCD == 1 {
		return d.U 
	} else {
		return -1
	}
}

// Checks whether a is a Miller-Rabin witness for N being composite
func MillerRabinTest(N, a int) bool {
	// Input. Integer n to be tested, integer a as potential 
	// witness. 

	if N < 0 {
		N *= -1 
	}

	// 1. If n is even or 1 < gcd(a,n) < n, return Composite
	if N % 2 == 0 {
		return true
	} else if EuclidAlgo(N,a).GCD != 1 {
		return true
	}

	q := N-1 
	k := 0 

	// 2. Write n-1 = 2^k q with q odd
	for q % 2 != 0 {
		q = q / 2 
		k += 1 
	}

	// 3. Set a = a^q mod n. 
	a = ModN(uint(N),FastPower(uint(N),a,uint(q)))

	// 4. If a = 1 mod n, return Test Fails.
	if a == 1 {
		return false
	}

	// 5. Loop i = 0,1,2,...,k-1
	for i := 0; i < k; i++ {
	//	6. If a = -1 mod n, return Test Fails. 
		if (a + 1) % N == 0 {
			return false
		}
	//	7. Set a = a^2 mod n 
		a = FastPower(uint(N),a,2)
	}
	// 8. End i loop.

	// 9. Return Composite
	return true
}

func FactorBase(B int) []int {
	// compute the primes <= B
	// works assuming the extended Riemann hypothesis 
	var primes []int 
	for n := 2; n <= B; n++ {
		upperbound := 2*math.Log(float64(n))*math.Log(float64(n))
		upperbound = math.Min(float64(n),upperbound-1)
		composite := false 
		for a := 2; a <= int(upperbound); a++ {
			b := MillerRabinTest(n,a) 
			if b {
				composite = true 
				break
			}
		}
		if !composite {
			primes = append(primes, n)
		}
	}
	return primes
}

func MaxPowerDiv(p, n int) int {
	e := 0
	for (n % p) == 0 {
		n = n / p
		e += 1
	}
	return e 
}

func IsBSmooth(n int,l []int) bool {
	// Return true if n is B-smooth and false if not 
	primes := l
	m := n 
	for _, p := range(primes) {
		if p > m {
			break
		}
		e := MaxPowerDiv(p,m)
		m = m / int(math.Pow(float64(p),float64(e)))
	}
	if m == 1 {
		return true
	}
	return false
}

func EulerCrit(p, a int) bool {
	m := FastPower(uint(p),a,uint((p-1)/2))
	if m == 1 {
		return true 
	}
	return false 
}

/// GetQuadNonRes(p) returns a quadratic non-residue 
func GetQuadNonRes(p int) int {
	for i := 2; i < p; i++ {
		if !EulerCrit(p,i) {
			return i
		}
	}
	return 0 
}

// TonelliShanks(p,a) returns the solutions of x^2 = a mod p for a prime p
func TonelliShanks(p, a int) (bool,[]int) {
	// If p = 2, then a^2 = a mod 2 so return true, [a]. Next 
	// check using Euler's criteria that a is a quadratic residue mod p 
	// if so return false and the empty slice
	// If p = 3 mod 4, then check that a^{(p+1)/4} is a square root of 
	// a. If not, return false and the empty slice. Otherwise 
	// return true, [r,(-r) mod p] 
	if p == 2{ 
		return true, []int{a}
	} else if !EulerCrit(p,a) {
		return false, []int{}
	} else if p % 4 == 3 {
		m := (p+1) / 4
		r := FastPower(uint(p),a,uint(m))
		if FastPower(uint(p),r,2) == a {
			return true, []int{r,p-r}
		}
		return false, []int{}
	}
	// Find a quadratic non-residue modulo p 
	z := GetQuadNonRes(p)
	// Factor p-1 into 2^s*q for q odd 
	s := 0 
	q := p-1 
	for q % 2 == 0 {
		q = q / 2 
		s += 1
	}
	// Initialize d = (q+1)/2, x = a^q mod p, c = z^q mod p and 
	// r = a^d mod p 
	d := (q + 1)/2 
	x := FastPower(uint(p),a,uint(q))
	c := FastPower(uint(p),z,uint(q))
	r := FastPower(uint(p),a,uint(d))
	// Loop while x != 1
	//  - if x = 0, then return true, [0]
	//  - else compute the minimal i such that x^{2^i} = 1 
	//    let b = c^{2*s - i -1} mod p, s = i, c = b^2 mod p, 
	//    x = x*c mod p, and r = r*b mod p 
	for x != 1 {
		if x == 0 {
			return true, []int{0}
		}
		i := 0
		y := x 
		for y != 1 {
			y = FastPower(uint(p),y,2)
			i += 1
		}
		b := FastPower(uint(p),c,uint(2*s-i-1))
		s = i 
		c = FastPower(uint(p),b,2)
		x = ModN(uint(p),x*c) 
		r = ModN(uint(p),r*b) 
	}
	return true, []int{r,p-r}
}

// GenTonelliShanks(p,e,a) returns the solutions of x^2 = a mod p^e for a 
// prime p and a with gcd(a,p) = 1
func GenTonelliShanks(p, e, a int) (bool,[]int) {
	if e == 0 {
		return TonelliShanks(p,a)
	}
	b, x := GenTonelliShanks(p,e-1,a)
	if b {
		roots := []int{}
		q := int(math.Pow(float64(p),float64(e-1)))
		for _, x := range(x) {
			c := (a - x*x)/q 
			y := ModN(uint(p),Inverse(uint(p),2*x)*c)
			r := ModN(uint(p*q),x + q*y) 
			roots = append(roots, r)
		}
		return true, roots 
	}
	return false, []int{}
}

// Computes n^e as an integer
func IntPow(n, e int) int {
	prod := 1 
	for i := 0; i < e; i++ {
		prod *= n 
	}
	return prod
}

// Given a slice of integers ints and another slice of ints exps, Prod 
// returns the product of all the primes raised to the corresponding 
// exponents. 
func Prod(ints []int, exps []int) int {
	prod := 1 
	for key, num := range(ints) {
		prod *= IntPow(num, exps[key])
	}
	return prod
}

// Reduce the entries of a matrix modulo N 
func MatModN(N int, matrix [][]int) {
	for _ , vect := range(matrix) {
		for _ , entry := range(vect) {
			entry = ModN(uint(N),entry)
		}
	}	
}

func Dimensions(matrix [][]int) (int,int) { 
	rows := len(matrix) 
	cols := 0 
	for _, vect := range(matrix) {
		cols = len(vect) 
		break
	}
	return rows, cols 
}

// Scalar multiplication of vectors  
func ScalarMul(c int, matrix []int) {
	for _ , component := range(matrix) {
		component = c*component 
	}	
}

// Adding vectors 
func AddVectsModN(N int, v, w []int) []int {
	sum := []int{}
	for key, val := range(v) {
		term := ModN(uint(N),val + w[key])
		sum = append(sum,term)
	}
	return sum 
}

// Get the index of the first row whose index is > frozenrows 
// and with a nonzero entry in column i. If there is none, 
// then return -1
func GetPivot(matrix [][]int, i, frozenrows int) int {
	for key, row := range(matrix) {
		if key <= frozenrows {
			continue 
		}
		for ind, entry := range(row) {
			if ind == i && entry != 0 {
				return key 
			}
		} 
	}
	return -1 
}

// Reduces a matrix over F_2 into row echelon form 
func REForm (matrix [][]int) {
	MatModN(2,matrix) 
	rows, cols := Dimensions(matrix)
	frozenrows := 0 
	for i := 0; i < cols; i++ {
		pivot := GetPivot(matrix,i,frozenrows) 
		switch pivot {
		case -1: 
			continue  
		case frozenrows + 1: 
			pivotrow := matrix[pivot] 
			for j := pivot + 1; j < rows; j++ {
				if matrix[j][i] != 0 {
					matrix[j] = AddVectsModN(2,matrix[j],pivotrow)
				}
			}
		default: 
			pivotrow := matrix[pivot] 
			matrix[pivot] = matrix[frozenrows+1] 
			matrix[frozenrows+1] = pivotrow
			for j := pivot + 1; j < rows; j++ {
				if matrix[j][i] != 0 {
					matrix[j] = AddVectsModN(2,matrix[j],pivotrow)
				}
			}
		}
		frozenrows += 1
	}
}

// Computes the kernel of matrix over F_2. 
func Kernel (matrix [][]int) [][]int {
	REForm(matrix)
	rows , cols := Dimensions(matrix)
	boundvars, freevars := []int{}, []int{} 
	for i := 0; i < cols; i++ {
		for _, row := range(matrix) {
			for j:=0; j < rows; j++ {
				if row[i] != 0 {
					freevars = append(freevars, i)
					break 
				}
			}
		}
	} 
	for i := cols-1 ; i >= 0; i-- {
		for _, val := range(boundvars) {
			if i == val {
				continue 
			}
		}
	}
	for _, var := range(boundvars) {
	
	}
	return [][]int{}
}

// Multiplies the matrix with the vector to produce a new slice
func MatrixMul (matrix [][]int, vector []int) []int {
	output := []int{}
	for _ , vect := range(matrix) {
		component := 0 
		for ind, val := range(vector) {
			component += vect[ind]*val 
		}
		output = append(output, component)
	}
	return output
}
