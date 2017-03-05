package main

import (
  "fmt"
)

func Count(text, pattern string) int {
  count, lengthOfText, lengthOfPattern := 0, len(text), len(pattern)
  if lengthOfText < lengthOfPattern {
    return 0
  } else {
    for i := 0; i < lengthOfText; i++ {
      if i + lengthOfPattern > lengthOfText {
        break
      } else {
        if text[i: i + lengthOfPattern] == pattern {
          count += 1
        }
      }
    }
  }
  return count
}

func main() {
  count := Count("ACTGTACGATGATGTGTGTCAAAG", "TGT")
  fmt.Printf("Count(ACTGTACGATGATGTGTGTCAAAG, TGT) = %d\n", count)
}

