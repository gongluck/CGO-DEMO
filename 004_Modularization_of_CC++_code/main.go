/*
 * @Author: gongluck
 * @Date: 2020-03-18 21:31:27
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-18 21:33:25
 */

package main

/*
#include "hello.h"
*/
import "C"

func main() {
	C.SayHelloC(C.CString("hello cgo c\n"))
	C.SayHelloCPP(C.CString("hello cgo cpp\n"))
}
