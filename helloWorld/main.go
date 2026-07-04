package main

const English = "English"
const EnglishGreeting = "Hello "

const Hindi = "Hindi"
const HindiGreeting = "Namaste "

const Japanees = "Japanees"
const JapaneesGreeting = "Konnichiwa "

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return  greetPrefix(language ) +name
}

func greetPrefix(language string ) (greet string){
	switch language {
	case Hindi:
		greet = HindiGreeting
	case Japanees:
		greet = JapaneesGreeting
	default:
		greet = EnglishGreeting
	}
	return
}

func main() {
	// fmt.Println(hello())
}
