/*
 * @Author: gongluck
 * @Date: 2020-03-18 21:16:49
 * @Last Modified by:   gongluck
 * @Last Modified time: 2020-03-18 21:16:49
 */

package main

/*
#include <stdio.h>

static void SayHello(const char* s)
{
	puts(s);
}
*/
import "C"

func main() {
	C.SayHello(C.CString("hello, cgo\n"))
}
