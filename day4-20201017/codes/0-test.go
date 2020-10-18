package main

import "fmt"

func test(stop bool) {
	if stop {
		fmt.Println("stop")
		return
	}

	fmt.Println("running")

}

func main() {
	test()

}
