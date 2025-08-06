package main

import "fmt"

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

func Repeat(character string, loop_time int)  (result string) {
	for i := 0; i < loop_time; i++ {
		result += character
	}

	return result;
}

func main(){
	fmt.Println(Hello("Hung", ""))
}
