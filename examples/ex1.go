package main 

import (
    "fmt"
    "pss/shiksha"
)

func main() {
    // Create an example Token with a Svara, Matra, and Aksharas
    token := shiksha.Token{
        Svara:    shiksha.Udaatta,
        Matra:    1,
        Aksharas: []shiksha.Akshara{shiksha.A, shiksha.I}, // Example usage of multiple Aksharas
    }

    // Display the token details
    fmt.Printf("Token Details:\n")
    fmt.Printf("  Svara: %s\n", token.Svara.String())
    fmt.Printf("  Matra: %d\n", token.Matra)
    fmt.Printf("  Aksharas:\n")
    for _, akshara := range token.Aksharas {
        fmt.Printf("    Unicode: %s, Devanagari: %s, Roman: %s\n", akshara.Unicode, akshara.Devanagari, akshara.Roman)
    }

    // Create a Pada with multiple tokens
    pada := shiksha.Pada{token, token}

    // Display Pada details
    fmt.Printf("\nPada Details:\n")
    for i, tok := range pada {
        fmt.Printf("  Token %d - Svara: %s, Matra: %d\n", i+1, tok.Svara.String(), tok.Matra)
    }
}
