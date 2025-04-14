/**
 * Problem 25: 1000-digit Fibonacci number
 * @see {@link https://projecteuler.net/problem=25}
 *
 * The Fibonacci sequence is defined by the recurrence relation:
 * F(n) = F(n−1) + F(n−2), where F(1) = 1 and F(2) = 1.
 * Hence the first 12 terms will be:
 * F(1) = 1
 * F(2) = 1
 * F(3) = 2
 * F(4) = 3
 * F(5) = 5
 * F(6) = 8
 * F(7) = 13
 * F(8) = 21
 * F(9) = 34
 * F(10) = 55
 * F(11) = 89
 * F(12) = 144
 *
 * The 12th term, F(12), is the first term to contain three digits.
 * What is the index of the first term in the Fibonacci sequence
 * to contain 1000 digits?
 *
 * @author ddaniel27
 */
package problem25

import (
	"math"
	"math/big"
)

func Problem25(digits int) int {
	currentFib := big.NewInt(1)          // Current Fibonacci number
	prevFib := big.NewInt(1)             // Previous Fibonacci number
	toCompare := createReference(digits) // Number with enough digits to compare
	currentIdx := 2                      // Current Fibonacci index F(idx)

	// Calculate amount of bits needed to store our target value
	// This could be replaced with toCompare.BitLen()
	// But this looks more mathematician
	triggerFloat := float64(digits-1) / math.Log10(2)
	trigger := int(math.Ceil(triggerFloat))

	res := 0 // response to return

	for {
		// Traditional iterative Fibonacci algorithm
		tmp := big.NewInt(0)
		tmp.Set(currentFib)

		currentFib := currentFib.Add(currentFib, prevFib)
		prevFib.Set(tmp)
		currentIdx++

		// Check if we reached enough bits to start comparations
		bitLen := currentFib.BitLen()
		if bitLen >= trigger {
			// If our current Fibonacci number is greater or equal to toCompare
			if comparationRes := currentFib.Cmp(toCompare); comparationRes >= 0 {
				res = currentIdx // we found our solution
				break
			}
		}
	}

	return res
}

// createReference creates a new big.Int with the given
// digits (i.e. 1e(digits-1))
func createReference(digits int) *big.Int {
	base := big.NewInt(1e0)
	for i := 0; i < digits-1; i++ {
		digit := big.NewInt(1e1)
		base.Mul(base, digit)
	}

	return base
}
