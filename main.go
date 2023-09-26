package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, world!")
    fmt.Println("\nHeres a quote from Go:")
    fmt.Println("\n"+quote.Go())
}