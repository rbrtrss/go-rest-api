package main

import "fmt"

// The run fucntion is goint to be responsible for the app's instatiation and startup
func Run() error {
	fmt.Print("Starting up app")
	return nil
}

func main() {
	fmt.Println("Ere we go")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
