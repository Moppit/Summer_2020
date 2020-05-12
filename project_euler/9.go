/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
package main

import "fmt"
import "math"

func main() {
  // This can be solved by brute force: I'm sure there's a more clever way, but here goes
  /*
  ------- Brute Force -------
  We iterate through all possibilities of a and b -- possibilities have to be less than 1000
  - We also know that a and b have to be less than c, and that c must be positive.
    therefore, we can just iterate to 500 for both of them
    * Bc if we either a or b is > 500, then c must be some value greater than that, and sum > 1000
  And then we check if it's a Pythagorean triplet
  - We do this by comparing the int(float) to float, and seeing if equal
    * Of c = sqrt(a^2+b^2)
  And then we see if the sum is 1000
  - If so, we return the product
  */
  for a := 1; a < 500; a++ {
    for b := 1; b < 500; b++ {
      // Check if it's a triplet
      c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
      if c == float64(int(c)) {
        // Check if sum is 1000
        if (float64(a)+float64(b)+c) == 1000.0 {
          // Print values for sanity check
          fmt.Println("Values:",a,b,c)
          fmt.Println(int(float64(a)*float64(b)*c))
          break
        }
      }
    }
  }
  // result: 31875000

  // Looking at the tread, people pretty much did the same thing
  // They did have smarter ranges though: instead of 500, they used 400
  // Would also be beneficial by calculating a c as 1000-a-b instead of all this float stuff
  // People also solved it as a system of equations
}
