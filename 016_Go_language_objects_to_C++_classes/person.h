/*
 * @Author: gongluck 
 * @Date: 2020-04-02 11:25:05 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 14:27:47
 */

extern "C" 
{
#include "./person_capi.h"
}

struct Person 
{
public:
	static Person* New(const char* name, int age) 
	{
		return reinterpret_cast<Person*>(person_new(const_cast<char*>(name), age));
	}
	void Delete() 
	{
		person_delete(reinterpret_cast<person_handle_t>(this));
	}

	void Set(char* name, int age) 
	{
		person_set(reinterpret_cast<person_handle_t>(this), name, age);
	}
	char* GetName(char* buf, int size) 
	{
		return person_get_name(reinterpret_cast<person_handle_t>(this), buf, size);
	}
	int GetAge() 
	{
		return person_get_age(reinterpret_cast<person_handle_t>(this));
	}
};
