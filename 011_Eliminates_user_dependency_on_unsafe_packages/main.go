/*
 * @Author: gongluck
 * @Date: 2020-03-25 21:05:26
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-25 21:07:25
 */

package main

import (
	"fmt"

	"011_Eliminates_user_dependency_on_unsafe_packages/qsort"
)

func main() {
	values := []int64{42, 9, 101, 95, 27, 25}

	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
}
