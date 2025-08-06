package main

import (
	"fmt"
	"strings"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish"{
		return "Hola, " + name
	}

	if language == "French" {
		return "Bonjour, " + name
	}


	return "Hello, " + name
}

func Repeat(character string, loop_time int)  string {
	var repeatedString strings.Builder
	for i := 0; i < loop_time; i++ {
		repeatedString.WriteString(character)
	}

	return repeatedString.String();
}

func main(){
	fmt.Println(Hello("Hung", ""))
}
