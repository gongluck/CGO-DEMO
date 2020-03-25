/*
 * @Author: gongluck
 * @Date: 2020-03-25 10:24:34
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-25 10:47:35
 */

package main

import (
	"010_The_closure_function_ACTS_as_a_comparison_function/qsort"
	"fmt"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.Sort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)

	fmt.Println(values)
}
