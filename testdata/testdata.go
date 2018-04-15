package testdata // OMIT

import "fmt" // OMIT

func Print() {
	fmt.Println("test") // MATCH "something is wrong"
}
