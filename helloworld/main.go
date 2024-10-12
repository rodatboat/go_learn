package main

import (
	"fmt"
)

// Define a type for the callback function
type CallbackFunc func(string)

// Function that takes a callback and applies it to each item in the list
func ProcessItems(items []string, callback CallbackFunc) {
	for _, item := range items {
		callback(item)
	}
}

// Custom callback function
func PrintItem(item string) {
	fmt.Println("Processing item:", item)
}

func main() {
	items := []string{"apple", "banana", "cherry"}

	// Call ProcessItems with PrintItem as the callback
	ProcessItems(items, PrintItem)
}
