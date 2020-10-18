package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		intVal     = 65
		float64Val = 2.2
		stringVal  = "64"
	)
	fmt.Println(intVal, float64Val, stringVal)

	fmt.Printf("%T, %#v\n", float64(intVal), float64(intVal))
	fmt.Printf("%T, %#v\n", int(float64Val), int(float64Val))

	fmt.Println(string(intVal))

	v, err := strconv.Atoi(stringVal)
	fmt.Println(err, v)

	vv, err := strconv.ParseFloat(stringVal, 64)
	fmt.Println(err, vv)

	fmt.Println(strconv.Itoa(intVal))
	fmt.Println(strconv.FormatFloat(float64Val, 'f', 10, 64)) // 2.2000000000 浮点数 10个精度

}

//赋值
