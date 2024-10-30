package shiksha

// Svara type as an enum
type Svara int

const (
    Anudaatta Svara = iota
    Udaatta
    Svarita
    DeerghaSvarita
)

// String method to get a readable string representation of Svara
func (s Svara) String() string {
    switch s {
    case Anudaatta:
        return "anudaatta"
    case Udaatta:
        return "udaatta"
    case Svarita:
        return "svarita"
    case DeerghaSvarita:
        return "deergha_svarita"
    default:
        return "unknown"
    }
}

// Akshara struct represents a Sanskrit character
type Akshara struct {
    Unicode    string `json:"unicode"`
    Devanagari string `json:"devanagari"`
    Roman      string `json:"roman"`
}

// Token struct represents a Sanskrit phonetic unit with Svara, Matra, and Aksharas
type Token struct {
    Svara     Svara     `json:"svara"`
    Matra     int       `json:"matra"`
    Aksharas  []Akshara `json:"aksharas"`
}

// Pada represents a sequence of Tokens
type Pada []Token
