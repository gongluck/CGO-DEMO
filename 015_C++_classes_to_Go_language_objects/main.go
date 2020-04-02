/*
 * @Author: gongluck
 * @Date: 2020-04-01 14:33:37
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-01 14:36:19
 */

package main

//#include <stdio.h>
import "C"

import (
	"unsafe"
)

func main() {
	buf := NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello, cgo"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
