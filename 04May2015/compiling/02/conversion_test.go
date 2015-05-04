package main

// PART1 OMIT

import (
  "testing"
)

func TestBinary(t *testing.T) {
  var str string
  divisor, dividendo, convertido := 25, 2, "11001"
  convert(divisor, dividendo, &str)
  if str != convertido {
    t.Error("Error in conversion 25 in binary is 11001 is not " + str + " %v")
  }

}

func TestHexa(t *testing.T) {
  str := ""
  divisor, dividendo, convertido := 15, 16, "F"
  convert(divisor, dividendo, &str)
  if str != convertido {
    t.Error("Error in conversion 15 in hex is F is not " + str + " %v")
  }

// PART2 OMIT
  str = ""
  divisor, dividendo, convertido = 1024, 16, "400"
  convert(divisor, dividendo, &str)
  if str != convertido {
    t.Error("Error in conversion 1024 in hex is 400 is not " + str + " %v") 
  }

  str = ""
  divisor, dividendo, convertido = 125, 16, "7D"
  convert(divisor, dividendo, &str)
  if str != convertido {
    t.Error("Error in conversion 125 in hex is 7D is not " + str + " %v")
  }
// PART3 OMIT
}


