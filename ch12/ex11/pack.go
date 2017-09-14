package main

import (
	"net/url"
	"reflect"
	"strconv"
)

func Pack(ptr interface{}) url.URL {

	values := url.Values{}
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			continue
		}

		var addParam func(string, reflect.Value)
		addParam = func(n string, val reflect.Value) {
			switch val.Kind() {
			case reflect.String:
				values.Add(n, val.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				values.Add(n, strconv.FormatInt(val.Int(), 10))
			case reflect.Bool:
				values.Add(n, strconv.FormatBool(val.Bool()))
			case reflect.Slice, reflect.Array:
				for i := 0; i < val.Len(); i++ {
					addParam(n, val.Index(i))
				}
			}
		}

		addParam(name, v.Field(i))
	}

	u := url.URL{}
	u.RawQuery = values.Encode()
	return u
}
