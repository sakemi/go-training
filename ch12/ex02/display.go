package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type st struct {
	i int
	s string
}

type S []S

func main() {
	var s = make(S, 1)
	s[0] = s
	Display("s", s)
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	safeDisplay(name, reflect.ValueOf(x))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Array:
		elem := make([]string, v.Len())
		for i := 0; i < v.Len(); i++ {
			elem[i] = formatAtom(v.Index(i))
		}
		return "[" + strings.Join(elem, ",") + "]"
	case reflect.Struct:
		fields := make([]string, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			name := v.Type().Field(i).Name
			value := formatAtom(v.Field(i))
			fields[i] = fmt.Sprintf("%s:%s", name, value)
		}
		return "[" + strings.Join(fields, ",") + "]"
	default:
		return v.Type().String() + " value"
	}
}

const maxDepth = 5

func safeDisplay(path string, v reflect.Value) {
	depth := 0
	var display func(string, reflect.Value)
	display = func(path string, v reflect.Value) {
		depth++
		if depth > maxDepth {
			depth--
			return
		}
		switch v.Kind() {
		case reflect.Invalid:
			fmt.Printf("%s = invalid\n", path)
		case reflect.Slice, reflect.Array:
			for i := 0; i < v.Len(); i++ {
				display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
			}
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
				display(fieldPath, v.Field(i))
			}
		case reflect.Map:
			for _, key := range v.MapKeys() {
				display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
			}
		case reflect.Ptr:
			if v.IsNil() {
				fmt.Printf("%s = nil\n", path)
			} else {
				display(fmt.Sprintf("(*%s)", path), v.Elem())
			}
		case reflect.Interface:
			if v.IsNil() {
				fmt.Printf("%s = nil\n", path)
			} else {
				fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
				display(path+".value", v.Elem())
			}
		default:
			fmt.Printf("%s = %s\n", path, formatAtom(v))
		}
		depth--
	}

	display(path, v)
}
