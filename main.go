package main

import "fmt"

type englishBot struct{}
type spaishBot struct{}

func main() {

	// printMap(colors)
}


func (eb englishBot) getGreeting() string{
	reutn "Hi There !"	
}


func (sb spaishBot) getGreeting() string{
	reutn "Hola !"	
}

// func printMap(c map[string]string) {
// 	for color, hex := range c {
// 		fmt.Println("Hex code for", color, " is ", hex)
// 	}
// }
 