package hello

import "fmt"

const (
	//Languages supported other than english which is default
	spanish  = "Spanish"
	french   = "French"
	japanese = "Japanese"
	hindi    = "Hindi"

	// Hello prefixes for supported languages
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "Konnichiwa, "
	hindiHelloPrefix    = "Namaste, "
)

func hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	name = name + "!"

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return
}

func PrintHelloWorld() {
	fmt.Println(hello("World", ""))
}
