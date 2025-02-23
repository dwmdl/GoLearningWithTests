package helloWorld

import "fmt"

const (
	french  = "French"
	russian = "Russian"

	helloEngPrefix = "Hello, "
	helloRuPrefix  = "Привет, "
	helloFrPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case russian:
		prefix = helloRuPrefix
	case french:
		prefix = helloFrPrefix
	default:
		prefix = helloEngPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Nikita", ""))
}
