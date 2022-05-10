package main

import "fmt"

func main() {

	countryCity := [3][3]string{
		{"Uk", "France", "Germany"},
		{"London", "Manchester", "Wales"},
		{"Calais", "Munich", "Frankfurt"},
	}

	for i := 0; i < len(countryCity); i++ {
		for j := 0; j < len(countryCity); j++ {
			fmt.Println(countryCity[i][j])
		}
	}
}
