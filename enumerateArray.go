package main

import "fmt"

func main() {

	namesCity := [3]string{"Paris", "Lisbon", "Riga"}

	for index, value := range namesCity {
		fmt.Println(index, value)
	}
}
