/*
 * @Author: gongluck
 * @Date: 2020-03-24 21:57:43
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-24 22:16:09
 */

package main

//extern int go_qsort_compare(void* a, void* b);
import "C"

import (
	"fmt"
	"unsafe"

	"009_Export_the_qsort_function_from_the_Go_package/qsort"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.Sort(unsafe.Pointer(&values[0]),
		len(values), int(unsafe.Sizeof(values[0])),
		qsort.CompareFunc(C.go_qsort_compare),
	)
	fmt.Println(values)
}

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}
