package main

import (
	"fmt"
	"reflect"
)

func simpleEqual(a, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

func sliceEqual(a, b interface{}) bool {
	slice_a, ok_a := generateInterfaceSlice(a)
	if !ok_a {
		return false
	}
	slice_b, ok_b := generateInterfaceSlice(b)
	if !ok_b {
		return false
	}

	return sliceEqual_(slice_a, slice_b)
}

func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return val, ok
}

func generateInterfaceSlice(arg interface{}) ([]interface{}, bool) {
	val, ok := isSlice(arg)
	if !ok {
		return nil, false
	}
	sliceLen := val.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out, true
}

func sliceEqual_(a, b []interface{}) bool {
	if (a == nil) && (b == nil) {
		return true
	}

	if (a == nil) && (len(b) == 0) {
		return true
	}

	if (b == nil) && (len(a) == 0) {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

type ConfigValue struct {
	name string
	//kind     reflect.Kind
	compFunc func(interface{}, interface{}) bool
	value1   interface{}
	value2   interface{}
}

func testgcomp() {
	intcv := &ConfigValue{
		name:     "integer compare",
		compFunc: simpleEqual,
		value1:   91,
		value2:   9,
	}
	if true == intcv.compFunc(intcv.value1, intcv.value2) {
		fmt.Println("The values of <" + intcv.name + "> are the same.")
	} else {
		fmt.Println("The values of <" + intcv.name + "> are not the same.")
	}

	strslicecv := &ConfigValue{
		name:     "string slice compare",
		compFunc: sliceEqual,
		value1:   []string{"a", "b", "c"},
		value2:   []string{"d", "d", "c"},
	}
	if true == strslicecv.compFunc(strslicecv.value1, strslicecv.value2) {
		fmt.Println("The values of <" + strslicecv.name + "> are the same.")
	} else {
		fmt.Println("The values of <" + strslicecv.name + "> are not the same.")
	}
}
