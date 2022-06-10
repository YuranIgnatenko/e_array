package main

import (
	. "Enum"
	"fmt"
)

var P = fmt.Printf

func info(arg interface{}) {
	fmt.Printf("Array: %v, Type: %v\n", arg, fmt.Sprintf("%T", arg))
}

func main() {
	info(ArrayS("*?@#$%^&*", "a", "abc", "123", "abc123"))
	info(ArrayI(1, 2, 3, 11, 2222, -1, -55, 0))
	info(ArrayF(0.1, 0, -0.1, 0.0000000007878))
	info(ArraySS("1234,arg1,arg2,Hello,#$%^&"))
	info(ArraySI("1,-12,0,12"))
	info(ArraySF("0.1, 0, -0.1, 0.000078"))

}
