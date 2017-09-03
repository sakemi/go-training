package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if _, err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) (bool, error) {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		var z int64
		i := v.Int()
		if i == z {
			return false, nil
		}
		fmt.Fprintf(buf, "%d", i)

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var z uint64
		u := v.Uint()
		if u == z {
			return false, nil
		}
		fmt.Fprintf(buf, "%d", u)

	case reflect.String:
		var z string
		s := v.String()
		if s == z {
			return false, nil
		}
		fmt.Fprintf(buf, "%q", s)

	case reflect.Ptr:
		if v.IsNil() {
			return false, nil
		}
		return encode(buf, v.Elem())

	case reflect.Slice:
		if v.Len() == 0 {
			return false, nil
		}
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if _, err := encode(buf, v.Index(i)); err != nil {
				return true, err
			}
		}
		buf.WriteByte(')')

	case reflect.Array:
		tmp := bytes.NewBuffer([]byte{})
		tmp.WriteByte('(')
		var isEncoded bool
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				tmp.WriteByte(' ')
			}
			enc, err := encode(tmp, v.Index(i))
			if err != nil {
				return true, err
			}
			if enc {
				isEncoded = true
			}
		}
		if isEncoded {
			buf.Write(tmp.Bytes())
			buf.WriteByte(')')
		}

	case reflect.Struct:
		tmp := bytes.NewBuffer([]byte{})
		tmp.WriteByte('(')
		var isEncoded bool
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				tmp.WriteByte(' ')
			}
			val := bytes.NewBuffer([]byte{})
			enc, err := encode(val, v.Field(i))
			if err != nil {
				return true, err
			}
			if enc {
				fmt.Fprintf(tmp, "(%s ", v.Type().Field(i).Name)
				tmp.Write(val.Bytes())
				tmp.WriteByte(')')
				isEncoded = true
			}
		}
		if isEncoded {
			buf.Write(tmp.Bytes())
			buf.WriteByte(')')
		}

	case reflect.Map:
		if v.IsNil() {
			return false, nil
		}
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if _, err := encode(buf, key); err != nil {
				return true, err
			}
			buf.WriteByte(' ')
			if _, err := encode(buf, v.MapIndex(key)); err != nil {
				return true, err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "%s", "t")
		} else {
			return false, nil
		}

	case reflect.Float32, reflect.Float64:
		var z float64
		f := v.Float()
		if f == z {
			return false, nil
		}
		fmt.Fprintf(buf, "%v", f)

	case reflect.Complex64, reflect.Complex128:
		var z complex128
		c := v.Complex()
		if z == c {
			return false, nil
		}
		fmt.Fprintf(buf, "#C(%v %v)", real(c), imag(c))

	case reflect.Chan, reflect.Func:
		if v.IsNil() {
			return false, nil
		}
		fmt.Fprintf(buf, "%v", v.Type())

	case reflect.Interface:
		if v.IsNil() {
			return false, nil
		}
		if _, err := encode(buf, v.Elem()); err != nil {
			return true, err
		}
	default:
		return true, fmt.Errorf("unsupported type: %s", v.Type())
	}
	return true, nil
}
