/*
 * @Author: gongluck
 * @Date: 2020-03-27 16:37:48
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-04-02 11:27:19
 */

package main

import "sync"

type ObjectID int

var refs struct {
	sync.Mutex
	objs map[ObjectID]interface{}
	next ObjectID
}

func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[ObjectID]interface{})
	refs.next = 1
}

func NewObjectID(obj interface{}) ObjectID {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++
	refs.objs[id] = obj
	return id
}

func (id ObjectID) IsNil() bool {
	return id == 0
}

func (id ObjectID) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[id]
}

func (id *ObjectID) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[*id]
	delete(refs.objs, *id)
	*id = 0

	return obj
}
