/*
 * @Author: gongluck
 * @Date: 2020-03-27 16:18:08
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-27 16:28:09
 */

package main

/*
#include <stdio.h>

void printString(const char* s, int n){
	for(int i = 0; i< n; ++i){
		putchar(s[i]);
	}
	putchar('\n');
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func printString(s string) {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	C.printString((*C.char)(unsafe.Pointer(p.Data)), C.int(len(s)))
}

func main() {
	s := "gongluck"
	printString(s)
}
