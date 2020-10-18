// 12 os 包写的add程序
package main


import (
	"fmt"
	"os"

)

func main() {
	args := os.Args[1:]

	usage := func() {

	}

	if len(args) < 2{
		usage()
	}

	total := 0

	for _, value := range args {
		if intvalue, err := strconv.Ato
	}

	if 
	fmt.Println()
}