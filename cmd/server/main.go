package main

import "fmt"

// RUN - is going to be responsible for
// the instantiation and startup of our
// GO Application
func Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {

	fmt.Println("Hello World!!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
