package main

import (
  "fmt"
)

func MostFrequentKMers(text string, k int) map[string]int {
  kMers := make(map[string]int)
  for i := 0; i < len(text); i++ {
    if i + k > len(text) {
      break
    }
    if _, ok := kMers[text[i: i + k]]; ok == false {
      kMers[text[i: i + k]] = 0
    }
    kMers[text[i: i + k]] += 1
  }
  return kMers
}

func main() {
  MostFrequent3Mers := MostFrequentKMers("CGCCTAAATAGCCTCGCGGAGCCTTATGTCATACTCGTCCT", 3)
  fmt.Println(MostFrequent3Mers)
}

