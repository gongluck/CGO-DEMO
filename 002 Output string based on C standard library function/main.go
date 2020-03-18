/*
 * @Author: gongluck
 * @Date: 2020-03-18 21:03:32
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-18 21:05:42
 */

package main

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("hello, cgo"))
}
