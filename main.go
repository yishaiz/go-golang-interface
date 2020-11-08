package main

import "fmt"

type englishBot struct{}
type spaishBot struct{}

func main() {

	// printMap(colors)
}


func printGreeting(eb englishBot){
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spaishBot){
	fmt.Println(sb.getGreeting())
}



func (eb englishBot) getGreeting() string{
	reutn "Hi There !"	
}


func (sb sbo) getGreeting() string{
	reutn "Hi There !"	
}

// func printMap(c map[string]string) {
// 	for color, hex := range c {
// 		fmt.Println("Hex code for", color, " is ", hex)
// 	}
// }
 