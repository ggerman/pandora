package main
// PART0 OMIT
import (
  "fmt"
  "strconv"
)

// PART1 OMIT

// Reverse returns its argument string reversed rune-wise left to right.
func reverse(s string) string { 	
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

func hexSpecialChars(n int) string {
  hex := []string{"A", "B", "C", "D", "E", "F"}
  if(n > 10) {
    return hex[n-10]
  } 
  return strconv.Itoa(n)
}

// PART3 OMIT

func convert(a, b int) string {
  var remainder, aux int
  var str string
  aux = a
  for aux > b {
    remainder = aux % b 
    if b == 16 {
      str = str + hexSpecialChars(remainder)
    } else {
      str  = str + strconv.Itoa(remainder)
    }
    aux = aux / b
  }
  str = str + strconv.Itoa(aux)
  return reverse(str)
}

// PART4 OMIT

func main() {
  s := convert(125, 16)
  fmt.Println(s)
}

// PART5 OMIT

