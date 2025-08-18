package main

import "flag"

// sayHi calls the Greet function
func sayHi(name string) {
	Greet(name)
}

func main() {
	// Define CLI flag
	name := flag.String("name", "World", "a name to say hello to")
	flag.Parse()

	// Run our sayHi function
	sayHi(*name)
}

// Greet() lives in greet.go.
// sayHi() in main.go calls Greet().
// Both files are in the same package (package main), so they can see each other.