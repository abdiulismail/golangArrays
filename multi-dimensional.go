package main

import "fmt"

func main() {

	countryCity := [3][3]string{
		{"Uk", "France", "Germany"},
		{"London", "Manchester", "Wales"},
		{"Calais", "Munich", "Frankfurt"},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(countryCity[i][j])
		}
	}
}
