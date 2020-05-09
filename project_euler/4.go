/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main
import "fmt"
import "strconv"

func isPalindrome(num int) bool {
  strVer := strconv.Itoa(num)
  isPalind := true
  for i := 0; i < len(strVer)/2; i++ {
    if strVer[i] != strVer[len(strVer)-1-i] {
      isPalind = false
    }
  }
  return isPalind
}

func main() {
  // Doesn't get the correct response...
  max := 1
  for i := 100; i < 1000; i++ {
    for j := 100; j < 1000; j++ {
      product := i*j
      if isPalindrome(product) && max < product {
        max = i*j
      }
    }
  }
  fmt.Println(max)

  // result = 906609
}
