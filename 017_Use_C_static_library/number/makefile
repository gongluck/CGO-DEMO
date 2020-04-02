#*
#* @Author: gongluck 
#* @Date: 2020-04-02 14:53:52 
#* @Last Modified by:   gongluck 
#* @Last Modified time: 2020-04-02 14:53:52 
#*/

default:
	gcc -c -o number.o number.c
	ar rcs libnumber.a number.o
	-rm *.o

clean:
	-rm *.a *.o
