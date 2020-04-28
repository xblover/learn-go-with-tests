package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "hello, "
const spanishHelloPrefix = "hola, "
const frenchHelloPrefix = "Bonjour, "

//现在需要支持第二个参数，指定问候的语言。如果一种不能被识别的语言传进来，就默认为英语。
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("xubo", ""))
}
