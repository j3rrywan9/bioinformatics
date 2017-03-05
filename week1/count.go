package main

import (
  "fmt"
)

func Count(text, pattern string) (int, []int) {
  lengthOfText, lengthOfPattern, count, positions := len(text), len(pattern), 0, []int{}
  if lengthOfText < lengthOfPattern {
    return 0, positions
  } else {
    for i := 0; i < lengthOfText; i++ {
      if i + lengthOfPattern > lengthOfText {
        break
      } else {
        if text[i: i + lengthOfPattern] == pattern {
          count += 1
          positions = append(positions, i)
        }
      }
    }
  }
  return count, positions
}

func main() {
  text, pattern := "ACTGTACGATGATGTGTGTCAAAG", "TGT"
  count, _ := Count(text, pattern)
  fmt.Printf("Count(%s, %s) = %d\n", text, pattern, count)

  text, pattern = "ATGACTTCGCTGTTACGCGC", "CGC"
  count, positions := Count(text, pattern)
  fmt.Printf("Count(%s, %s) = %d, Positions = %v\n", text, pattern, count, positions)
}

