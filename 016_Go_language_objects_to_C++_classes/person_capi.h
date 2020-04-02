/*
 * @Author: gongluck 
 * @Date: 2020-04-02 11:24:49 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 14:27:27
 */

#include <stdint.h>

typedef uintptr_t person_handle_t;

person_handle_t person_new(char* name, int age);
void person_delete(person_handle_t p);

void person_set(person_handle_t p, char* name, int age);
char* person_get_name(person_handle_t p, char* buf, int size);
int person_get_age(person_handle_t p);
