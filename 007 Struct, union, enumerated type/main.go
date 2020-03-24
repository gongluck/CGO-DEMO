/*
 * @Author: gongluck
 * @Date: 2020-03-23 09:39:07
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-23 09:57:51
 */

package main

/*
#include <stdint.h>

struct A{
	int i;
	float f;
	int type;
	int   size:10;	// The bit field cannot be accessed
	float arr[];	// Zero-length arrays are also inaccessible
};

union B1 {
	int i;
	float f;
};
union B2 {
	int8_t i8;
	int64_t i64;
};

enum C {
	ONE,
	TWO,
};
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
	fmt.Println(a._type)

	var b1 C.union_B1
	fmt.Printf("%T\n", b1) // [4]uint8
	var b2 C.union_B2
	fmt.Printf("%T\n", b2) // [8]uint8
	fmt.Println("b1.i:", *(*C.int)(unsafe.Pointer(&b1)))
	fmt.Println("b1.f:", *(*C.float)(unsafe.Pointer(&b1)))

	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}
