/*
 * @Author: gongluck 
 * @Date: 2020-03-18 21:29:07 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-18 21:31:20
 */

#include <iostream>
extern "C" 
{
    #include "hello.h"
}

void SayHelloCPP(const char* s)
{
    std::cout << s;
}