/*
 * @Author: gongluck
 * @Date: 2020-03-18 21:50:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-18 22:04:25
 */

package main

//void SayHello(_GoString_ s);
import "C"

import (
	"fmt"
)

func main() {
	C.SayHello("hello, cgo go")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
