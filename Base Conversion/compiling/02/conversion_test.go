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

func TestReverse(t *testing.T) {  text := reverse("En un lugar de la Mancha, de cuyo nombre no quiero acordarme, no ha mucho tiempo que vivía un hidalgo de los de lanza en astillero, adarga antigua, rocín flaco y galgo corredor.")
  if text != ".roderroc oglag y ocalf nícor ,augitna agrada ,orellitsa ne aznal ed sol ed ogladih nu aíviv euq opmeit ohcum ah on ,emradroca oreiuq on erbmon oyuc ed ,ahcnaM al ed ragul nu nE" {
    t.Error("Error in reverse String %v" + text)
  }

  text = reverse("123456789")
  if text != "987654321" {
    t.Error("Error in reverse String %v " + text)
  }

}
