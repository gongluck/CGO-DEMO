/*
 * @Author: gongluck
 * @Date: 2020-04-02 11:23:51
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 14:38:20
 */

package main

// #cgo CXXFLAGS: -std=c++11
// extern void Main();
import "C"

// C++(pointer) -> Go(export function) -> Go(ObjectID)
func main() {
	C.Main()
}
