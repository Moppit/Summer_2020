/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that
the 6th prime is 13.

What is the 10 001st prime number?
*/

package main
import "fmt"

// Same prime function as in problem 3
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
  // 2 is the first prime
  val := 2
  idx := 1
  for idx < 10001 {
    if isPrime(val) {
      idx++
    }
    val++
  }

  fmt.Println(val-1)
}

// result = 104743

/* Improvements:
- Could have just checked odd values instead of everything
- Sieve's algorithm -- super useful!
*/
