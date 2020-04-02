/*
 * @Author: gongluck 
 * @Date: 2020-04-01 10:11:12 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-01 14:28:34
 */

typedef struct MyBuffer_T MyBuffer_T;

MyBuffer_T* NewMyBuffer(int size);
void DeleteMyBuffer(MyBuffer_T* p);

char* MyBuffer_Data(MyBuffer_T* p);
int MyBuffer_Size(MyBuffer_T* p);