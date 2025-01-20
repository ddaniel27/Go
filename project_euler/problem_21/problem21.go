/**
* Problem 21 - Amicable numbers
* @see {@link https://projecteuler.net/problem=21}
*
* Let d(n) be defined as the sum of proper divisors of n
* (numbers less than n which divide evenly into n).
* If d(a) = b and d(b) = a, where a ≠ b, then a and b are an
* amicable pair and each of a and b are called amicable numbers.
*
* For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11,
* 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors
* of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
*
* Evaluate the sum of all the amicable numbers under 10000.
*
* @author ddaniel27
 */
package problem21

func Problem21(input uint) uint {
	divisorMap := make(map[uint]uint)
	sum := uint(0)

	for i := uint(1); i < input; i++ {
		divisorMap[i] = divisorsSum(i)
	}

	amicablePairs := getAmicablePairs(divisorMap)

	for _, pair := range amicablePairs {
		sum += pair
	}

	return sum
}

// getAmicablePairs returns a slice of amicable pairs from the divisorMap
func getAmicablePairs(divisorMap map[uint]uint) []uint {
	pairs := []uint{}

	for k, v := range divisorMap {
		if v != k && divisorMap[v] == k {
			pairs = append(pairs, k)
		}
	}

	return pairs
}

// divisorsSum returns the sum of the proper divisors of n
func divisorsSum(n uint) uint {
	sum := uint(1)
	limit := n / 2 // n/2 is the largest possible divisor

	for i := uint(2); i < limit; i++ {
		if n%i == 0 {
			sum += i      // add the divisor
			sum += n / i  // add the divisor's pair
			limit = n / i // set the new limit
		}
	}

	return sum
}
