package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "hello, "
const spanishHelloPrefix = "hola, "
const frenchHelloPrefix = "Bonjour, "

//如果我们希望稍后添加更多的语言支持，我们可以使用 switch 来重构代码，使代码更易于阅读和扩展。
//你可能会抱怨说也许我们的函数正在变得很臃肿。对此最简单的重构是将一些功能提取到另一个函数中。
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("xubo", ""))
}
