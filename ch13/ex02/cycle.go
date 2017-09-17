package main

import (
	"reflect"
	"unsafe"
)

type instance struct {
	p unsafe.Pointer
	t reflect.Type
}

func IsCycle(ptr interface{}) bool {
	seen := map[instance]bool{}
	return isCycle(reflect.ValueOf(ptr), seen)
}

func isCycle(v reflect.Value, seen map[instance]bool) bool {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	ptr := unsafe.Pointer(v.UnsafeAddr())
	c := instance{ptr, v.Type()}
	if seen[c] {
		return true
	}
	seen[c] = true
	for i := 0; i < v.NumField(); i++ {
		if isCycle(v.Field(i), seen) {
			return true
		}
	}
	return false
}
