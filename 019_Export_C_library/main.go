/*
 * @Author: gongluck
 * @Date: 2020-04-02 20:22:15
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 20:27:46
 */

package main

import "C"

//go build -buildmode=c-archive -o number.a
//go build -buildmode=c-shared -o number.so
func main() {}

//export number_add_mod
func number_add_mod(a, b, mod C.int) C.int {
	return (a + b) % mod
}
