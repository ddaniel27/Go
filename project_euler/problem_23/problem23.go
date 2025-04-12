/**
* Problem 23: Non-abundant sums
* @see {@link https://projecteuler.net/problem=23}
*
* A perfect number is a number for which the sum of
* its proper divisors is exactly equal to the number.
* For example, the sum of the proper divisors of 28
* would be 1 + 2 + 4 + 7 + 14 = 28, which means that
* 28 is a perfect number.
*
* A number n is called deficient if the sum of its proper
* divisors is less than n and it is called abundant if this
* sum exceeds n.
*
* As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
* the smallest number that can be written as the sum of two
* abundant numbers is 24. By mathematical analysis, it can be
* shown that all integers greater than 28123 can be written as
* the sum of two abundant numbers. However, this upper limit
* cannot be reduced any further by analysis even though it is
* known that the greatest number that cannot be expressed as
* the sum of two abundant numbers is less than this limit.
*
* Find the sum of all the positive integers which cannot be
* written as the sum of two abundant numbers.
*
* @author ddaniel27
 */
package problem23

import (
	"sort"

	"github.com/TheAlgorithms/Go/math"
)

func Problem23() uint {
	var upperLimit uint = 28123
	var res uint = 0

	as := abundantSlice(upperLimit)
	currentAbundantIdx := 0

	for i := uint(1); i <= upperLimit; i++ {
		if i < as[currentAbundantIdx] { // Sum all numbers up to current abundant limit
			res += i
			continue
		}

		currentAbundantIdx++ // Increase index when abundant number was reached
	}

	return res
}

func abundantSlice(upperLimit uint) []uint {
	abundantSlice := make(map[uint]struct{})
	tmp := make(map[uint]struct{})

	for i := uint(1); i <= upperLimit; i++ {
		if e := math.SumOfProperDivisors(i); e > i { // Get all proper devisors and check abundance
			abundantSlice[i] = struct{}{}
			for j := range abundantSlice { // Get all sums of two abundant numbers
				if i+j > upperLimit { // If greater than upperLimit, just throw it
					continue
				}

				tmp[i+j] = struct{}{}
			}
		}
	}

	res := []uint{}
	for i := range tmp { // "Simple" change from keys to slice
		res = append(res, i)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	res = append(res, upperLimit+1) // Adding an upper limit

	return res
}
