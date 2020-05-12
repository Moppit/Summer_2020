/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"
import "math"

// Simple method: try heuristic method later on
// 6k +/- 1 omptimization
func isPrime(num int) bool {
  // Check cases that often happen for ease of computation
  // Deals with 1-3 case
  if num <= 3 {
    return num > 1
  } else if num%2 == 0 || num%3 == 0 || num%5 == 0 {
    return false // Handles multiples of 2,3,5
  }

  // Case that deals with more difficult values to evaluate
  i := 5
  for i*i <= num {
    // Check if divisible by i or i+2
    if num%i == 0 || num%(i+2) == 0 {
      return false
    }
    i += 6
  }
  return true
}

func main() {
  // Primality testing typically goes from 2 to sqrt(n) -- from number theory,
  // composite numbers have factors less than or equal to its root
  toFactor := 600851475143
  start := int(math.Sqrt(float64(toFactor)))
  for i := start; i >= 2; i-- {
    if isPrime(i) && toFactor%i == 0 {
      fmt.Println(i)
      break
    }
  }
}

// Result = 6857

// Improvement: again, could use sieve since numbers aren't too large
