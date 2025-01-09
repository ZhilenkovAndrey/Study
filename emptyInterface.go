package main

import (
	. "fmt"
	"reflect"
)

func main() {
	m := map[string]interface{}{}
	m["one"] = 1
	m["two"] = 2.0
	m["three"] = true

	for k, v := range m {
		switch v.(type) {
		case int:
			Printf("%s is integer\n", k)
		case float64:
			Printf("%s is float\n", k)
		// case bool:
		// 	Printf("%s is bool\n", k)
		default:
			Printf("%s is %v\n", k, reflect.TypeOf(v))
		}

	}
}
