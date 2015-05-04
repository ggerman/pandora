package main
// PART0 OMIT
import (
  "fmt"
  "strconv"
)

// PART1 OMIT

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string { 	
  r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// PART2 OMIT

func mod (a, b int) int {
  return a % b // operator % remainder integers 
}

// PART3 OMIT

func convert(a, b int, str *string) {
  if a >= b {
    remainder := mod(a, b)
    if b == 16 {  // Hexadecimal exception for base (16)
      hex := []string{"A", "B", "C", "D", "E", "F"}
      switch remainder {
        case 10,11,12,13,14,15: *str = *str + hex[remainder-10]
        default: *str = *str + strconv.Itoa(remainder)
      }
    } else {
      *str = *str + strconv.Itoa(remainder)
    }
    convert(a/b, b, str) // recursive solution
  } else {
    *str = *str + strconv.Itoa(a)
    *str = Reverse(*str)
  }
}

// PART4 OMIT

func main() {
  var s string
  convert(125, 16, &s)
  fmt.Println(s)
}

// PART5 OMIT

