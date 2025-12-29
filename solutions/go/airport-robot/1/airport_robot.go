package airportrobot

import "fmt"

// Greeter defines methods for greeting in a specific language.
type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

// SayHello returns a greeting that includes the language name and the greeting message.
func SayHello(name string, greeter Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// English implements Greeter for English.
type English struct{}

func (English) LanguageName() string {
    return "English"
}

func (English) Greet(name string) string {
    return fmt.Sprintf("Hello %s!", name)
}

// Italian implements Greeter for Italian.
type Italian struct{}

func (Italian) LanguageName() string {
    return "Italian"
}

func (Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

// Portuguese implements Greeter for Portuguese.
type Portuguese struct{}

func (Portuguese) LanguageName() string {
    return "Portuguese"
}

func (Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}
