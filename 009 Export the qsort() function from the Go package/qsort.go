/*
 * @Author: gongluck
 * @Date: 2020-03-24 21:50:09
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-24 22:12:51
 */

package main

/*
 #include <stdlib.h>
 typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
*/
import "C"

import "unsafe"

type CompareFunc C.qsort_cmp_func_t

func Sort(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}
