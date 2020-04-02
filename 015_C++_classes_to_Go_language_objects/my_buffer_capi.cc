/*
 * @Author: gongluck 
 * @Date: 2020-04-01 10:13:05 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-01 14:38:27
 */

#include "./my_buffer.h"

extern "C" 
{
	#include "./my_buffer_capi.h"
}

struct MyBuffer_T : public MyBuffer 
{
public:
	MyBuffer_T(int size): MyBuffer(size)
	{

	}
	~MyBuffer_T() 
	{

	}
};

MyBuffer_T* NewMyBuffer(int size) 
{
	auto p = new MyBuffer_T(size);
	return p;
}
void DeleteMyBuffer(MyBuffer_T* p) 
{
	delete p;
}

char* MyBuffer_Data(MyBuffer_T* p) 
{
	return p->Data();
}
int MyBuffer_Size(MyBuffer_T* p) 
{
	return p->Size();
}
