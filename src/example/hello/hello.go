package main

import "fmt"

func main() {
	message1 := "hi"
	var message2 *string = &message1
	a, b, c := 1, false, 3
	fmt.Println(message1, a, b, c, *message2)
}
