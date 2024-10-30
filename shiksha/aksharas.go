package shiksha


// Vowels
var (
    A       = Akshara{Unicode: "\u0905", Devanagari: "अ", Roman: "a"}
    AA      = Akshara{Unicode: "\u0906", Devanagari: "आ", Roman: "aa"}
    I       = Akshara{Unicode: "\u0907", Devanagari: "इ", Roman: "i"}
    II      = Akshara{Unicode: "\u0908", Devanagari: "ई", Roman: "ii"}
    U       = Akshara{Unicode: "\u0909", Devanagari: "उ", Roman: "u"}
    UU      = Akshara{Unicode: "\u090A", Devanagari: "ऊ", Roman: "uu"}
    
    R       = Akshara{Unicode: "\u090B", Devanagari: "ऋ", Roman: "r̥"}
    RR      = Akshara{Unicode: "\u0960", Devanagari: "ॠ", Roman: "r̥̄"}
    L       = Akshara{Unicode: "\u090C", Devanagari: "ऌ", Roman: "l̥"}
    LL      = Akshara{Unicode: "\u0961", Devanagari: "ॡ", Roman: "l̥̄"}

    E       = Akshara{Unicode: "\u090F", Devanagari: "ए", Roman: "e"}
    AI      = Akshara{Unicode: "\u0910", Devanagari: "ऐ", Roman: "ai"}
    O       = Akshara{Unicode: "\u0913", Devanagari: "ओ", Roman: "o"}
    AU      = Akshara{Unicode: "\u0914", Devanagari: "औ", Roman: "au"}
)

// Consonants
var (
    KA      = Akshara{Unicode: "\u0915", Devanagari: "क", Roman: "ka"}
    KHA     = Akshara{Unicode: "\u0916", Devanagari: "ख", Roman: "kha"}
    GA      = Akshara{Unicode: "\u0917", Devanagari: "ग", Roman: "ga"}
    GHA     = Akshara{Unicode: "\u0918", Devanagari: "घ", Roman: "gha"}
    NGA     = Akshara{Unicode: "\u0919", Devanagari: "ङ", Roman: "nga"}

    CA      = Akshara{Unicode: "\u091A", Devanagari: "च", Roman: "ca"}
    CHA     = Akshara{Unicode: "\u091B", Devanagari: "छ", Roman: "cha"}
    JA      = Akshara{Unicode: "\u091C", Devanagari: "ज", Roman: "ja"}
    JHA     = Akshara{Unicode: "\u091D", Devanagari: "झ", Roman: "jha"}
    NYA     = Akshara{Unicode: "\u091E", Devanagari: "ञ", Roman: "nya"}

    TTA     = Akshara{Unicode: "\u091F", Devanagari: "ट", Roman: "tta"}
    TTHA    = Akshara{Unicode: "\u0920", Devanagari: "ठ", Roman: "ttha"}
    DDA     = Akshara{Unicode: "\u0921", Devanagari: "ड", Roman: "dda"}
    DDHA    = Akshara{Unicode: "\u0922", Devanagari: "ढ", Roman: "ddha"}
    NNA     = Akshara{Unicode: "\u0923", Devanagari: "ण", Roman: "nna"}

    TA      = Akshara{Unicode: "\u0924", Devanagari: "त", Roman: "ta"}
    THA     = Akshara{Unicode: "\u0925", Devanagari: "थ", Roman: "tha"}
    DA      = Akshara{Unicode: "\u0926", Devanagari: "द", Roman: "da"}
    DHA     = Akshara{Unicode: "\u0927", Devanagari: "ध", Roman: "dha"}
    NA      = Akshara{Unicode: "\u0928", Devanagari: "न", Roman: "na"}

    PA      = Akshara{Unicode: "\u092A", Devanagari: "प", Roman: "pa"}
    PHA     = Akshara{Unicode: "\u092B", Devanagari: "फ", Roman: "pha"}
    BA      = Akshara{Unicode: "\u092C", Devanagari: "ब", Roman: "ba"}
    BHA     = Akshara{Unicode: "\u092D", Devanagari: "भ", Roman: "bha"}
    MA      = Akshara{Unicode: "\u092E", Devanagari: "म", Roman: "ma"}

    YA      = Akshara{Unicode: "\u092F", Devanagari: "य", Roman: "ya"}
    RA      = Akshara{Unicode: "\u0930", Devanagari: "र", Roman: "ra"}
    LA      = Akshara{Unicode: "\u0932", Devanagari: "ल", Roman: "la"}
    VA      = Akshara{Unicode: "\u0935", Devanagari: "व", Roman: "va"}

    SHA     = Akshara{Unicode: "\u0936", Devanagari: "श", Roman: "sha"}
    SSA     = Akshara{Unicode: "\u0937", Devanagari: "ष", Roman: "ssa"}
    SA      = Akshara{Unicode: "\u0938", Devanagari: "स", Roman: "sa"}
    HA      = Akshara{Unicode: "\u0939", Devanagari: "ह", Roman: "ha"}
)

// Additional Characters
var (
    Anusvara = Akshara{Unicode: "\u0902", Devanagari: "ं", Roman: "m̐"}
    Visarga  = Akshara{Unicode: "\u0903", Devanagari: "ः", Roman: "ḥ"}
)
