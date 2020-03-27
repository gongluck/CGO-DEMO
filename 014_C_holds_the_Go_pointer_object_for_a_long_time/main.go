/*
 * @Author: gongluck
 * @Date: 2020-03-27 16:45:17
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-27 17:21:11
 */

package main

/*
extern int NewGoStringObj(char*);
extern void FreeGoStringObj(int);
extern void PrintGoStringObj(int);

static void printString(char* s) {
	int gs = NewGoStringObj(s);
	PrintGoStringObj(gs);
	FreeGoStringObj(gs);
}

static void Main(){
	printString("hello, cgo");
}
*/
import "C"
import (
	"014_C_holds_the_Go_pointer_object_for_a_long_time/object"
)

//export NewGoStringObj
func NewGoStringObj(s *C.char) C.int {
	gs := C.GoString(s)
	id := object.NewObjectID(gs)
	return (C.int)(id)
}

//export FreeGoStringObj
func FreeGoStringObj(obj C.int) {
	id := object.ObjectID(obj)
	id.Free()
}

//export PrintGoStringObj
func PrintGoStringObj(obj C.int) {
	id := object.ObjectID(obj)
	gs := id.Get().(string)
	print(gs)
}

func main() {
	C.Main()
}
