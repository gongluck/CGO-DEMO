/*
 * @Author: gongluck
 * @Date: 2020-03-26 11:52:07
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-26 15:51:50
 */

package main

/*
#include <stdlib.h>

void* makeslice(size_t memsize) {
	return malloc(memsize);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const maxsize = 1<<32 - 1

func makeByteSlize(n uint) []byte {
	if n > maxsize {
		panic("can not malloc buf size larger than maxsize")
	}
	p := C.makeslice(C.size_t(n))
	return ((*[maxsize]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func main() {
	fmt.Println("maxsize :", maxsize)
	s := makeByteSlize(maxsize)
	s[len(s)-1] = 255
	print(s[len(s)-1])
	freeByteSlice(s)
}
