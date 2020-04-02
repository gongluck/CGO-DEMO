/*
 * @Author: gongluck 
 * @Date: 2020-04-02 11:23:01 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 14:34:04
 */

#include <stdio.h>
#include "person.h"

extern "C" 
void Main()
{
	auto p = Person::New("gongluck", 10);

	char buf[64];
	char* name = p->GetName(buf, sizeof(buf)-1);
	int age = p->GetAge();

	printf("%s, %d years old.\n", name, age);
	p->Delete();
}
