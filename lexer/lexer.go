package lexer

import (
    "pss/shiksha"
    "fmt"
)

func Tokenize(input string) ([]shiksha.Akshara, error) {
    var tokens []shiksha.Akshara

    for _, char := range input {
        // Convert the rune to a string for mapping
        chStr := string(char)
        if chStr != " "{
          fmt.Println(chStr)
        }
    }

    return tokens, nil
}
// Mapping Baraha encodings to Aksharas
var barahaToAkshara = map[string]shiksha.Akshara{
    // Vowels
    "a": shiksha.A, "aa": shiksha.AA, "i": shiksha.I, "ii": shiksha.II,
    "u": shiksha.U, "uu": shiksha.UU, "e": shiksha.E, "E": shiksha.E,"ai": shiksha.AI,
    "o": shiksha.O, "au": shiksha.AU, "R": shiksha.R, "RR": shiksha.RR,
    "L": shiksha.L, "LL": shiksha.LL,

    // Consonants with inherent vowels
    "ka": shiksha.KA, "kha": shiksha.KHA, "ga": shiksha.GA, 
    "gha": shiksha.GHA, "nga": shiksha.NGA, "ya": shiksha.YA, "ma": shiksha.MA,

    // Common consonant-vowel pairs
    "yO": shiksha.Akshara{Unicode: "यो", Devanagari: "यो", Roman: "yO"},
    "vE": shiksha.Akshara{Unicode: "वे", Devanagari: "वे", Roman: "vE"},
    "prO": shiksha.Akshara{Unicode: "प्रो", Devanagari: "प्रो", Roman: "prO"},
    "vEq": shiksha.Akshara{Unicode: "वेक", Devanagari: "वेक", Roman: "vEq"},

    // Special symbols
    "#": shiksha.Akshara{Unicode: "#", Devanagari: "#", Roman: "#"}, // Syllable boundary
    "|": shiksha.Akshara{Unicode: "|", Devanagari: "|", Roman: "|"}, // Line boundary
}
