/*
 * @Author: gongluck
 * @Date: 2020-04-02 11:24:36
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 14:29:07
 */

package main

//#include "./person_capi.h"
import "C"
import "unsafe"

//export person_new
func person_new(name *C.char, age C.int) C.person_handle_t {
	id := NewObjectID(NewPerson(C.GoString(name), int(age)))
	return C.person_handle_t(id)
}

//export person_delete
func person_delete(h C.person_handle_t) {
	((*ObjectID)((unsafe.Pointer)(&h))).Free()
}

//export person_set
func person_set(h C.person_handle_t, name *C.char, age C.int) {
	p := ObjectID(h).Get().(*Person)
	p.Set(C.GoString(name), int(age))
}

//export person_get_name
func person_get_name(h C.person_handle_t, buf *C.char, size C.int) *C.char {
	p := ObjectID(h).Get().(*Person)
	name, _ := p.Get()

	n := int(size) - 1
	bufSlice := ((*[1 << 31]byte)(unsafe.Pointer(buf)))[0:n:n]
	n = copy(bufSlice, []byte(name))
	bufSlice[n] = 0

	return buf
}

//export person_get_age
func person_get_age(h C.person_handle_t) C.int {
	p := ObjectID(h).Get().(*Person)
	_, age := p.Get()
	return C.int(age)
}
