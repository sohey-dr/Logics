package main

import (
    "fmt"

    "github.com/deltam/go-lsd-parametrized"
)

func main() {
    a, b := "kitten", "face"

    // standard
    fmt.Println(lsdp.Lsd(a, b))
    // Output:
    // 4

    // weighted
    wd := lsdp.Weights{Insert: 0.1, Delete: 1, Replace: 0.01}
    fmt.Println(wd.Distance(a, b))
    // Output:
    // 0.22

    // weighted and normalized
    nd := lsdp.Normalized(wd)
    fmt.Println(nd.Distance(a, b))
    // Output:
    // 0.0275

    // weighted by rune
    wr := lsdp.ByRune(&lsdp.Weights{1, 1, 1}).
        Insert("g", 0.1).
        Insert("h", 0.01).
        Replace("k", "s", 0.001).
        Replace("e", "i", 0.0001)
    fmt.Println(wr.Distance(a, b))
    // Output:
    // 0.1111
}