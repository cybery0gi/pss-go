package shiksha

// karneeyam: Devanagari transliterations for Sanskrit comments are currently unsupported with nvim. Figure this out someday.

// VedicSvara represents the tonal accent in Vedic Sanskrit (e.g., Anudaatta, Udaatta)
type VedicSvara int

const (
    Anudaatta VedicSvara = iota
    Udaatta
    Svarita
    DeerghaSvarita
)

func (s VedicSvara) String() string {
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

// || na ksheeyate na ksharati iti vaa aksharam ||
// That which neither diminishes nor decays is called an aksharam
// Akshara represents a basic Sanskrit character or sound unit
type Akshara struct {
    Unicode    string `json:"unicode"`    // Unicode representation
    Devanagari string `json:"devanagari"` // Devanagari script representation
    Roman      string `json:"roman"`      // Roman transliteration
}

// || svatantradhvaniH varNaH bhavati ||
// Independent sounds are Varnas
// Varna represents an independent phonetic unit in Sanskrit
type Varna struct {
    Akshara Akshara   `json:"akshara"` // The core sound unit
    Svara   VedicSvara `json:"svara"`   // The tonal accent of the Varna
}

// || vyanjanaani anvag bhavati iti vyanjanam ityaaha patanjali ||
// || svaradharmam anugacchati iti vyanjanam || 
// As per Patanjali - that which is dependent on a vowel is called as "Vyanjana" (Consonant)
// Vyanjanas follow the nature of svaras 
// Vyanjana represents a consonant that depends on a Svara (vowel) for pronunciation
type Vyanjana struct {
    Varna Varna `json:"varna"` // The underlying Varna representing this consonant
}

// ayogavaahaaH - letters which are not termed or classified under svaras or vyanjanas but still seen in usage are referred to as ayogavaahaa
// Ayogavaaha represents auxiliary sounds that are neither Svaras nor Vyanjanas
type Ayogavaaha struct {
    Sound string `json:"sound"` // A sound like a nasal or breath that’s not a vowel or consonant
}

// || svaram anuprayujyate iti anusvaaraH || 
// || anusvaaraH svaraatparaH prayujyate ||
// Anusvara represents a nasal sound used as an extension of Svara
type Anusvara struct {
    Varna Varna `json:"varna"` // The underlying Varna followed by nasalization
}

// || vividhataam srjati iti visargaH ||
// That which manifests itself in many ways is Visarga
// Visarga represents an aspirate sound (ḥ), occurring as a modification of a Svara
type Visarga struct {
    Varna Varna `json:"varna"` // The underlying Varna with an added aspirate sound
}

// || kakaabhyaam praag ardhavisargasadrshaH jihvamooliyaH ityucyate ||
// Jihvamuliya is a special variant of Visarga, pronounced before velar sounds
type Jihvamuliya struct {
    Varna Varna `json:"varna"` // The underlying Varna with a half-visarga sound before velars
}

// Upadhmaaniya is a special variant of Visarga, pronounced before labials
type Upadhmaaniya struct {
    Varna Varna `json:"varna"` // The underlying Varna with a half-visarga sound before labials
}

// yamaaH
// If any of the 4 letters of a varga comes in front of the 5th varna - a similar letter is added as per shiksha called "yamavarna"
// i.e. agni - g and n have come together - so agni is pronounced as "aggni" - this second `g` is a yama varna
// 4 yama varnas in rig veda pratisakhya
// Yama represents a duplicated consonant sound due to sandhi (phonetic combination) rules
type Yama struct {
    Primary   Varna `json:"primary"`   // The original Varna
    Duplicate Varna `json:"duplicate"` // The duplicated Varna following sandhi rules
}
