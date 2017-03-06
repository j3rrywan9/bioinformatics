package main

import (
  "errors"
  "fmt"
)

func calculateHammingDistance(text1, text2 string) (int, error) {
  var hemmingDistance int
  var err error
  if len(text1) != len(text2) {
    err = errors.New("Input strings are not of same length!")
  } else {
    for i := 0; i < len(text1); i ++ {
      if text1[i] != text2[i] {
        hemmingDistance += 1
      }
    }
  }
  return hemmingDistance, err
}

func main() {
  text1, text2 := "CAGAAAGGAAGGTCCCCATACACCGACGCACCAGTTTA", "CACGCCGTATGCATAAACGAGCCGCACGAACCAGAGAG"
  hemmingDistance, err := calculateHammingDistance(text1, text2)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(hemmingDistance)
  }
}

