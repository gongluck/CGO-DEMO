/*
 * @Author: gongluck
 * @Date: 2020-03-18 21:43:19
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-18 21:48:09
 */

package main

import "C"

import "fmt"

//export SayHelloGo
func SayHelloGo(s *C.char) {
	fmt.Print(C.GoString(s))
}
