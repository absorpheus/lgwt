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

// TODO: Language direction (LTR,RTL)
func greetingPrefix(language string) string {
	if val, exists := greetings[language]; exists {
		return val + ", "
	}

	return englishHelloPrefix
}
