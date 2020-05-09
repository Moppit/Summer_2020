/*
2520 is the smallest number that can be divided by each of the numbers from 1 to
10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
*/
package main

import "fmt"

func isMultiple(num int) bool {
  for i := 1; i < 21; i++ {
    if num % i != 0 {
      return false
    }
  }
  return true
}

func main() {
  factorial := 1
  for i := 2; i < 21; i++ {
    factorial *= i
  }

  for i := 2520; i < factorial; i++ {
    if isMultiple(i) {
      fmt.Println(i)
      break
    }
  }
}

// result: 232792560
