/*
 * @Author: gongluck
 * @Date: 2020-03-18 21:46:32
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-18 21:47:40
 */

package main

//#include <hello.h>
import "C"

func main() {
	C.SayHelloGo(C.CString("hello, cgo Go\n"))
}
