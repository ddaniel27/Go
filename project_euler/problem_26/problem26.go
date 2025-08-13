/**
* Problem 26: Reciprocal Cycles
* @see {@link https://projecteuler.net/problem=26}
*
* A unit fraction contains 1 in the numerator. The decimal representation
* of the unit fractions with denominators 2 to 10 are given:
* 1/2	= 0.5
* 1/3	= 0.(3)
* 1/4	= 0.25
* 1/5	= 0.2
* 1/6	= 0.1(6)
* 1/7	= 0.(142857)
* 1/8	= 0.125
* 1/9	= 0.(1)
* 1/10	= 0.1
* Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle.
* It can be seen that 1/7 has a 6-digit recurring cycle.
* Find the value of d < 1000 for which 1/d contains the longest recurring cycle
* in its decimal fraction part.
*
* @author ddaniel27
 */
package problem26

import (
	"math/big"

	"github.com/TheAlgorithms/Go/math/gcd"
)

// A051626 comes from the OEIS sequence A051626,
// which gives the length of the recurring cycle in the decimal
// expansion of 1/n, or 0 if n is a power of 2 or 5.
// @see {@link https://oeis.org/A051626}
func A051626(n uint64) uint64 {
	if n <= 1 {
		return 0
	}
	// Remove factors of 2 and 5
	m := n
	for m%2 == 0 {
		m /= 2
	}
	for m%5 == 0 {
		m /= 5
	}
	if m == 1 {
		return 0
	}

	// gcd(10, m) == 1 garantizado tras quitar 2 y 5
	return multiplicativeOrder(10, m)
}

// multiplicativeOrder calcula el menor k>0 tal que a^k ≡ 1 (mod n).
func multiplicativeOrder(a, n uint64) uint64 {
	if gcd.Iterative(int64(a), int64(n)) != 1 {
		return 0
	}

	// φ(n)
	factorsN := factorize(n)
	phi := eulerPhiFromFactors(n, factorsN)

	// Factorizar φ(n)
	factorsPhi := factorize(phi)

	ord := phi
	for p := range factorsPhi {
		for ord%p == 0 {
			if powMod(a, ord/p, n) == 1 {
				ord /= p
			} else {
				break
			}
		}
	}
	return ord
}

// factorize breaks down n into its prime factors and returns a map
func factorize(n uint64) map[uint64]int {
	f := make(map[uint64]int)
	if n == 0 || n == 1 {
		return f
	}
	// Contar 2s
	for n%2 == 0 {
		f[2]++
		n /= 2
	}
	// Impares
	var d uint64 = 3
	for d*d <= n {
		for n%d == 0 {
			f[d]++
			n /= d
		}
		d += 2
	}
	if n > 1 {
		f[n]++
	}
	return f
}

func eulerPhiFromFactors(n uint64, fac map[uint64]int) uint64 {
	result := n
	for p := range fac {
		result = result / p * (p - 1)
	}
	return result
}

// powMod calculates (a^e) mod m using modular exponentiation.
func powMod(a, e, m uint64) uint64 {
	if m == 1 {
		return 0
	}
	A := new(big.Int).SetUint64(a % m)
	E := new(big.Int).SetUint64(e)
	M := new(big.Int).SetUint64(m)
	res := new(big.Int).Exp(A, E, M)
	return res.Uint64()
}

func Problem26(limit uint64) uint64 {
	maxValue := uint64(0)
	d := uint64(1)

	for i := uint64(1); i < limit; i++ {
		res := A051626(i)
		if res > maxValue {
			maxValue = res
			d = i
		}
	}

	return d
}
