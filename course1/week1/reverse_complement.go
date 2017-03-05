package main

import (
  "fmt"
)

func ReverseComplement(text string) string {
  var reverseComplement string
  for i := 0; i < len(text); i++ {
    switch nucleotide := text[i]; nucleotide {
    case 'A':
      reverseComplement += "T"
    case 'C':
      reverseComplement += "G"
    case 'G':
      reverseComplement += "C"
    case 'T':
      reverseComplement += "A"
    }
  }
  // TODO: reverse it before return
  return reverseComplement
}

func main() {
  reverseComplement := ReverseComplement("TTGTGTC")
  fmt.Println(reverseComplement)
}

