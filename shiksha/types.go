package shiksha

// karneeyam: Devanagari transliterations for sanskrit comments is currently unsupported with nvim. Figure this out someday.

// Svara type as an enum
type VedicSvara int

const (
    Anudaatta VedicSvara = iota
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

/***
// Token struct represents a Sanskrit phonetic unit with Svara, Matra, and Aksharas
type Token struct {
    Svara     Svara     `json:"svara"`
    Matra     int       `json:"matra"`
    Aksharas  []Akshara `json:"aksharas"`
}
***/

// || arthavat padaanaam samoohaH vaakyam ||
// Vaakya represents an array of padas that convey meaning.
type Vaakya []Pada

// Pada represents a group of Varnas
// ||varNaanaam samoohaH padam||
type Pada []Token

// || svatantradhvaniH varNaH bhavati ||
//  Independent sounds are Varnas


// || na ksheeyate na ksharati iti vaa aksharam ||
// That which neither diminishes nor decays is called an aksharam


// || svayam raajante iti svaraaH ||
// a svara is that which shines by itself is called a svara




// || vyanjanaani anvag bhavati iti vyanjanam ityaaha patanjali ||
// || svaradharmam anugacchati iti vyanjanam || 
// As per Patanjali - that which is dependent on a vowel is called as "Vyanjana" (Consonant)
// Vyanjanas follow the nature of svaras 


// ayogavaahaaH - letters which are not termed or classified under svaras or vyanjanas but still seen in usage are referred to as ayogavaahaa

// ANUSVARA
// || svaram anuprayujyate iti anusvaaraH || 
// || anusvaaraH svaraatparaH prayujyate ||
// 

// VISARGA 
// || vividhataam srjati iti visargaH ||
// That which manifests itself in many ways is Visarga

 
// JIHVAAMOOLIYA
// || kakaabhyaam praag ardhavisargasadrshaH jihvamooliyaH ityucyate ||


// UPADHMAANIYA 


// yamaaH
// If any of the 4 letters of a varga comes in front of the 5th varna - a similar letter is added as per shiksha called "yamavarna"
// i.e. agni - g and n have come together - so agni is pronounced as "aggni" - this second `g` is a yama varna
// 4 yama varnas in rig veda pratisakhya










 
