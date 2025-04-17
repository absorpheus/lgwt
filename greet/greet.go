package greet

const (
	englishHelloPrefix = "Hello, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	if val, exists := greetings[language]; exists {
		return val + ", "
	}

	return englishHelloPrefix
}

// TODO: Language direction (LTR,RTL)
var greetings = map[string]string{
	"Afrikaans":          "Hallo",
	"Albanian":           "Përshëndetje",
	"Amharic":            "ሰላም",
	"Arabic":             "مرحبا",
	"Armenian":           "Բարեւ",
	"Azerbaijani":        "Salam",
	"Basque":             "Kaixo",
	"Belarusian":         "Добры дзень",
	"Bengali":            "নমস্কার",
	"Bosnian":            "Zdravo",
	"Bulgarian":          "Здравей",
	"Catalan":            "Hola",
	"Cebuano":            "Kumusta",
	"Chinese (Mandarin)": "你好",
	"Croatian":           "Bok",
	"Czech":              "Ahoj",
	"Danish":             "Hej",
	"Dutch":              "Hallo",
	"English":            "Hello",
	"Esperanto":          "Saluton",
	"Estonian":           "Tere",
	"Filipino":           "Kamusta",
	"Finnish":            "Hei",
	"French":             "Bonjour",
	"Galician":           "Ola",
	"Georgian":           "გამარჯობა",
	"German":             "Hallo",
	"Greek":              "Γειά σας",
	"Gujarati":           "નમસ્તે",
	"Haitian Creole":     "Bonjou",
	"Hausa":              "Sannu",
	"Hebrew":             "שלום",
	"Hindi":              "नमस्ते",
	"Hungarian":          "Szia",
	"Icelandic":          "Halló",
	"Igbo":               "Ndewo",
	"Indonesian":         "Halo",
	"Irish":              "Dia dhuit",
	"Italian":            "Ciao",
	"Japanese":           "こんにちは",
	"Javanese":           "Halo",
	"Kannada":            "ನಮಸ್ಕಾರ",
	"Kazakh":             "Сәлем",
	"Khmer":              "សួស្តី",
	"Kinyarwanda":        "Muraho",
	"Korean":             "안녕하세요",
	"Kurdish":            "Slav",
	"Kyrgyz":             "Салам",
	"Lao":                "ສະບາຍດີ",
	"Latin":              "Salve",
	"Latvian":            "Sveiki",
	"Lithuanian":         "Sveiki",
	"Luxembourgish":      "Moien",
	"Macedonian":         "Здраво",
	"Malagasy":           "Manao ahoana",
	"Malay":              "Selamat pagi",
	"Malayalam":          "നമസ്കാരം",
	"Maltese":            "Bongu",
	"Maori":              "Kia ora",
	"Marathi":            "नमस्कार",
	"Mongolian":          "Сайн уу",
	"Nepali":             "नमस्ते",
	"Norwegian":          "Hei",
	"Pashto":             "سلام",
	"Persian":            "سلام",
	"Polish":             "Cześć",
	"Portuguese":         "Olá",
	"Punjabi":            "ਸਤ ਸ੍ਰੀ ਅਕਾਲ",
	"Romanian":           "Salut",
	"Russian":            "Привет",
	"Scottish Gaelic":    "Halò",
	"Serbian":            "Здраво",
	"Sesotho":            "Lumela",
	"Shona":              "Mhoro",
	"Sindhi":             "سلام",
	"Sinhala":            "ආයුබෝවන්",
	"Slovak":             "Ahoj",
	"Slovenian":          "Živjo",
	"Somali":             "Iska warran",
	"Spanish":            "Hola",
	"Sundanese":          "Halo",
	"Swahili":            "Jambo",
	"Swedish":            "Hej",
	"Tajik":              "Салом",
	"Tamil":              "வணக்கம்",
	"Tatar":              "Исәнме",
	"Telugu":             "నమస్కారం",
	"Thai":               "สวัสดี",
	"Turkish":            "Merhaba",
	"Turkmen":            "Salam",
	"Ukrainian":          "Привіт",
	"Urdu":               "سلام",
	"Uzbek":              "Salom",
	"Vietnamese":         "Xin chào",
	"Welsh":              "Helo",
	"Xhosa":              "Molo",
	"Yiddish":            "שלום",
	"Yoruba":             "Bawo",
	"Zulu":               "Sawubona",
	"Twi":                "Agoo",
}
