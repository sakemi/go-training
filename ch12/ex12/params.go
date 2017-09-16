package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type paramInfo struct {
	field     reflect.Value
	paramType string
}

// param types
const (
	ht      = "http"
	email   = "email"
	card    = "card"
	zipcode = "zipcode"
)

var paramTypes [4]string = [4]string{ht, email, card, zipcode}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := map[string]paramInfo{}
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag

		for _, t := range paramTypes {
			name, ok := tag.Lookup(t)
			if !ok {
				continue
			}
			if name == "" {
				name = strings.ToLower(fieldInfo.Name)
			}
			fields[name] = paramInfo{v.Field(i), t}
			break
		}
	}

	for name, values := range req.Form {
		f := fields[name].field
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := validate(value, fields[name].paramType); err != nil {
					return err
				}
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func validate(value, paramType string) error {
	switch paramType {
	case email: // RFC完全準拠ではないが実用的っぽいバリデーション
		if l := len(value); l > 255 {
			return fmt.Errorf("email address is too long: %d", l)
		}

		parts := strings.Split(value, "@")
		if len(parts) != 2 {
			return fmt.Errorf("malformed email address: %s", value)
		}
		if local := len(parts[0]); local == 0 || local > 64 {
			return fmt.Errorf("malformed email address: %s", value)
		}
		if domain := len(parts[1]); domain == 0 || domain > 254 {
			return fmt.Errorf("malformed email address: %s", value)
		}
	case card:
		_, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("malformed credit card number: %s", value)
		}
		if l := len(value); l < 14 || 16 < l {
			return fmt.Errorf("malformed credit card number: %s", value)
		}
	case zipcode:
		_, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("malformed credit card number: %s", value)
		}
		if l := len(value); l != 5 {
			return fmt.Errorf("malformed credit card number: %s", value)
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
