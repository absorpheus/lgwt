package greet

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEmptyStrings(t *testing.T) {
	t.Run("defaults to 'Hello, world' when name is empty", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to 'Hello, Jack' when language is empty", func(t *testing.T) {
		got := Hello("Jack", "")
		want := "Hello, Jack"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to 'Hello, world' when both name and language are empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func TestHelloTranslations(t *testing.T) {
	tests := []struct {
		description string
		name     string
		language string
		want     string
	}{
		{"in Afrikaans", "Aaliyah", "Afrikaans", "Hallo, Aaliyah"},
		{"in Albanian", "Fatima", "Albanian", "Përshëndetje, Fatima"},
		{"in Amharic", "Omar", "Amharic", "ሰላም, Omar"},
		{"in Arabic", "Zainab", "Arabic", "مرحبا, Zainab"},
		{"in Armenian", "Hiroshi", "Armenian", "Բարեւ, Hiroshi"},
		{"in Azerbaijani", "Kwame", "Azerbaijani", "Salam, Kwame"},
		{"in Basque", "Leila", "Basque", "Kaixo, Leila"},
		{"in Belarusian", "Anaya", "Belarusian", "Добры дзень, Anaya"},
		{"in Bengali", "Santiago", "Bengali", "নমস্কার, Santiago"},
		{"in Bosnian", "Yara", "Bosnian", "Zdravo, Yara"},
		{"in Bulgarian", "Mei", "Bulgarian", "Здравей, Mei"},
		{"in Catalan", "Tariq", "Catalan", "Hola, Tariq"},
		{"in Cebuano", "Chioma", "Cebuano", "Kumusta, Chioma"},
		{"in Chinese (Mandarin)", "Jin", "Chinese (Mandarin)", "你好, Jin"},
		{"in Croatian", "Amara", "Croatian", "Bok, Amara"},
		{"in Czech", "Mohammed", "Czech", "Ahoj, Mohammed"},
		{"in Danish", "Priya", "Danish", "Hej, Priya"},
		{"in Dutch", "Babajide", "Dutch", "Hallo, Babajide"},
		{"in English", "Linh", "English", "Hello, Linh"},
		{"in Esperanto", "Samira", "Esperanto", "Saluton, Samira"},
		{"in Estonian", "Takashi", "Estonian", "Tere, Takashi"},
		{"in Filipino", "Ayesha", "Filipino", "Kamusta, Ayesha"},
		{"in Finnish", "José", "Finnish", "Hei, José"},
		{"in French", "Ibrahim", "French", "Bonjour, Ibrahim"},
		{"in Galician", "Chinonso", "Galician", "Ola, Chinonso"},
		{"in Georgian", "Nia", "Georgian", "გამარჯობა, Nia"},
		{"in German", "Aliyah", "German", "Hallo, Aliyah"},
		{"in Greek", "Ravi", "Greek", "Γειά σας, Ravi"},
		{"in Gujarati", "Lucía", "Gujarati", "નમસ્તે, Lucía"},
		{"in Haitian C", "Malik", "Haitian Creole", "Bonjou, Malik"},
		{"in Hausa", "Thandiwe", "Hausa", "Sannu, Thandiwe"},
		{"in Hebrew", "Haruto", "Hebrew", "שלום, Haruto"},
		{"in Hindi", "Nour", "Hindi", "नमस्ते, Nour"},
		{"in Hungarian", "Alejandro", "Hungarian", "Szia, Alejandro"},
		{"in Icelandic", "Sana", "Icelandic", "Halló, Sana"},
		{"in Igbo", "Ayush", "Igbo", "Ndewo, Ayush"},
		{"in Indonesian", "Zuleika", "Indonesian", "Halo, Zuleika"},
		{"in Irish", "Idris", "Irish", "Dia dhuit, Idris"},
		{"in Italian", "Bao", "Italian", "Ciao, Bao"},
		{"in Japanese", "Sanele", "Japanese", "こんにちは, Sanele"},
		{"in Javanese", "Yesenia", "Javanese", "Halo, Yesenia"},
		{"in Kannada", "Kofi", "Kannada", "ನಮಸ್ಕಾರ, Kofi"},
		{"in Kazakh", "Soraya", "Kazakh", "Сәлем, Soraya"},
		{"in Khmer", "Tenzin", "Khmer", "សួស្តី, Tenzin"},
		{"in Kinyarwanda", "Imani", "Kinyarwanda", "Muraho, Imani"},
		{"in Korean", "Habib", "Korean", "안녕하세요, Habib"},
		{"in Kurdish", "Kiara", "Kurdish", "Slav, Kiara"},
		{"in Kyrgyz", "Abdul", "Kyrgyz", "Салам, Abdul"},
		{"in Lao", "Noor", "Lao", "ສະບາຍດີ, Noor"},
		{"in Latin", "Jamila", "Latin", "Salve, Jamila"},
		{"in Latvian", "Miguel", "Latvian", "Sveiki, Miguel"},
		{"in Lithuanian", "Zahra", "Lithuanian", "Sveiki, Zahra"},
		{"in Luxembourgish", "Akio", "Luxembourgish", "Moien, Akio"},
		{"in Macedonian", "Tanisha", "Macedonian", "Здраво, Tanisha"},
		{"in Malagasy", "Sibongile", "Malagasy", "Manao ahoana, Sibongile"},
		{"in Malay", "Hassan", "Malay", "Selamat pagi, Hassan"},
		{"in Malayalam", "Indira", "Malayalam", "നമസ്കാരം, Indira"},
		{"in Maltese", "Marisol", "Maltese", "Bongu, Marisol"},
		{"in Maori", "Farah", "Maori", "Kia ora, Farah"},
		{"in Marathi", "Enzo", "Marathi", "नमस्कार, Enzo"},
		{"in Mongolian", "Chike", "Mongolian", "Сайн уу, Chike"},
		{"in Nepali", "Nur", "Nepali", "नमस्ते, Nur"},
		{"in Norwegian", "Sofía", "Norwegian", "Hei, Sofía"},
		{"in Pashto", "Manuel", "Pashto", "سلام, Manuel"},
		{"in Persian", "Taiwo", "Persian", "سلام, Taiwo"},
		{"in Polish", "Aminah", "Polish", "Cześć, Aminah"},
		{"in Portuguese", "Jorge", "Portuguese", "Olá, Jorge"},
		{"in Punjabi", "Aarushi", "Punjabi", "ਸਤ ਸ੍ਰੀ ਅਕਾਲ, Aarushi"},
		{"in Romanian", "Layla", "Romanian", "Salut, Layla"},
		{"in Russian", "Damilola", "Russian", "Привет, Damilola"},
		{"in Scottish Gaelic", "Tariro", "Scottish Gaelic", "Halò, Tariro"},
		{"in Serbian", "Khadija", "Serbian", "Здраво, Khadija"},
		{"in Sesotho", "Shreya", "Sesotho", "Lumela, Shreya"},
		{"in Shona", "Mateo", "Shona", "Mhoro, Mateo"},
		{"in Sindhi", "Fatou", "Sindhi", "سلام, Fatou"},
		{"in Sinhala", "Rashid", "Sinhala", "ආයුබෝවන්, Rashid"},
		{"in Slovak", "Anjali", "Slovak", "Ahoj, Anjali"},
		{"in Slovenian", "Nyasha", "Slovenian", "Živjo, Nyasha"},
		{"in Somali", "Arjun", "Somali", "Iska warran, Arjun"},
		{"in Spanish", "Keisha", "Spanish", "Hola, Keisha"},
		{"in Sundanese", "Nikhil", "Sundanese", "Halo, Nikhil"},
		{"in Swahili", "Mpho", "Swahili", "Jambo, Mpho"},
		{"in Swedish", "Ismael", "Swedish", "Hej, Ismael"},
		{"in Tajik", "Yasmin", "Tajik", "Салом, Yasmin"},
		{"in Tamil", "Nguyen", "Tamil", "வணக்கம், Nguyen"},
		{"in Tatar", "Huda", "Tatar", "Исәнме, Huda"},
		{"in Telugu", "DeShawn", "Telugu", "నమస్కారం, DeShawn"},
		{"in Thai", "Bao", "Thai", "สวัสดี, Bao"},
		{"in Turkish", "Zain", "Turkish", "Merhaba, Zain"},
		{"in Turkmen", "Nandini", "Turkmen", "Salam, Nandini"},
		{"in Ukrainian", "Lamar", "Ukrainian", "Привіт, Lamar"},
		{"in Urdu", "Naima", "Urdu", "سلام, Naima"},
		{"in Uzbek", "Diego", "Uzbek", "Salom, Diego"},
		{"in Vietnamese", "Adama", "Vietnamese", "Xin chào, Adama"},
		{"in Welsh", "Tariqah", "Welsh", "Helo, Tariqah"},
		{"in Xhosa", "Rocio", "Xhosa", "Molo, Rocio"},
		{"in Yiddish", "Sameer", "Yiddish", "שלום, Sameer"},
		{"in Yoruba", "Jalil", "Yoruba", "Bawo, Jalil"},
		{"in Zulu", "Ishita", "Zulu", "Sawubona, Ishita"},
		{"in Twi", "Zahara", "Twi", "Agoo, Zahara"},
	}

		for _, tt := range tests {
			t.Run(tt.description, func(t *testing.T) {
				got := Hello(tt.name, tt.language)
				assertCorrectMessage(t, got, tt.want)
			})
		}
}
