/*
 * @Author: gongluck
 * @Date: 2020-04-02 14:48:27
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 16:26:19
 */

package main

//If a static library file has been generated
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber

//#cgo CFLAGS: -I./number
//#include "number.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.number_add_mod(10, 5, 12))
}
