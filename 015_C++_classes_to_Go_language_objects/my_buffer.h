/*
 * @Author: gongluck 
 * @Date: 2020-04-01 10:09:12 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-01 14:39:00
 */

#include <string>

class MyBuffer 
{
public:
	MyBuffer(int size) 
	{
		this->s_ = new std::string(size, char('\0'));
	}
	~MyBuffer() 
	{
		delete this->s_;
	}

	int Size() const 
	{
		return this->s_->size();
	}
	char* Data() 
	{
		return const_cast<char*>(this->s_->data());
	}

private:
	std::string* s_ = nullptr;
};