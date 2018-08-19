
package main

import (
   "fmt"
   "testing"
)

func TestCI(t *testing.T) {
   a := 2
   b := 3
   c := a + b
   if c != 5 {
      panic("wrong c result")
   }
}

func TestCISecond(t *testing.T) {
   a := 2
   b := 3
   c := a * b
   if c != 6 {
      panic("wrong c result")
   }
}

func TestCalculations(t *testing.T) {
   var result int64 = 23*3219 + 4342 + 32*32 + (1+4)*5 + 2*3 + ((3 * 3) + 4*4)
   if result != 79459 {
      panic(fmt.Sprintf("we receive wrong result: %d", result))
   }
}