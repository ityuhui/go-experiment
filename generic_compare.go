package main

import "fmt"

func simpleEqual(a, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

func sliceEqual(a, b []interface{}) bool {
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
		value1:   1,
		value2:   2,
	}
	if true == intcv.compFunc(intcv.value1, intcv.value2) {
		fmt.Println("The values of <" + intcv.name + "> are the same.")
	} else {
		fmt.Println("The values of <" + intcv.name + "> are the same.")
	}

	strslicecv := &ConfigValue{
		name:     "string slice compare",
		compFunc: sliceEqual,
		value1:   []string{"a", "b", "c"},
		value2:   []string{"d", "d"},
	}
	if true == strslicecv.compFunc(strslicecv.value1, strslicecv.value2) {
		fmt.Println("The values of <" + strslicecv.name + "> are the same.")
	} else {
		fmt.Println("The values of <" + strslicecv.name + "> are the same.")
	}
}
