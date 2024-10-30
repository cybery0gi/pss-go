package main

import (
    "fmt"
    "pss/lexer"
)

func main() {
    input := "yO vEdAdau sva#raH prOqktOq vEqdAntE# ca praqtiShThi#taH" // | tasya# praqkRuti#lInaqsyaq yaqH para#H sa maqhESva#raH |"
    
    // Tokenize the input string
    tokens, err := lexer.Tokenize(input)
    if err != nil {
        fmt.Printf("Error tokenizing input: %v\n", err)
        return
    }
    
    // Display the tokens
    fmt.Println("Tokens:")
    for _, token := range tokens {
        fmt.Printf("Unicode: %s, Devanagari: %s, Roman: %s\n", token.Unicode, token.Devanagari, token.Roman)
    }
}
