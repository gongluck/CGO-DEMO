/*
 * @Author: gongluck
 * @Date: 2020-04-02 11:24:03
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 11:27:13
 */

package main

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) Set(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) Get() (name string, age int) {
	return p.name, p.age
}
