package shiksha

// Define all Aksharas in the Sanskrit alphabet

var (
    // Vowels
    A       = Akshara{Unicode: "\u0905", Devanagari: "अ", Roman: "a"}
    AA      = Akshara{Unicode: "\u0906", Devanagari: "आ", Roman: "aa"}
    I       = Akshara{Unicode: "\u0907", Devanagari: "इ", Roman: "i"}
    II      = Akshara{Unicode: "\u0908", Devanagari: "ई", Roman: "ii"}
    U       = Akshara{Unicode: "\u0909", Devanagari: "उ", Roman: "u"}
    UU      = Akshara{Unicode: "\u090A", Devanagari: "ऊ", Roman: "uu"}
    
    // Consonants
    KA      = Akshara{Unicode: "\u0915", Devanagari: "क", Roman: "ka"}
    KHA     = Akshara{Unicode: "\u0916", Devanagari: "ख", Roman: "kha"}
    GA      = Akshara{Unicode: "\u0917", Devanagari: "ग", Roman: "ga"}
    GHA     = Akshara{Unicode: "\u0918", Devanagari: "घ", Roman: "gha"}
    NGA     = Akshara{Unicode: "\u0919", Devanagari: "ङ", Roman: "nga"}
    
    // Add remaining vowels and consonants similarly...
)
