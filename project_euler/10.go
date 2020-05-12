/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
package main
import "fmt"

func main() {
  // Would use Sieve's, except Golang isn't as great as python is about lists
  // So do a sort of modified version?
  sum := 0
  primes := []int{2}
  for i := 3; i < 2000000; i++ {
    prime := true
    for _, val := range primes {
      if (i % val) == 0 {
        prime = false
        break;
      }
    }
    if prime {
      primes = append(primes, i)
      sum += i
    }
    fmt.Println(i)
  }

  // Now we see all the primes, and print their sum
  fmt.Println(primes)
  fmt.Println(sum)

  // result = 142913828922
  /* Improvements
  This method was actually pretty slow... I wonder if it would have been faster
  with the one I used previously? I think it might have been
  - Could have used a bit manip version -- just make an array of bools and then add
    indices that are marked true for prime -- this would still require a checking
    mechanism though
  */

}
