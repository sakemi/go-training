package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Decoder struct {
	scan *scanner.Scanner
}

func NewDecoder(r io.Reader) *Decoder {
	s := scanner.Scanner{Mode: scanner.GoTokens}
	return &Decoder{s.Init(r)}
}

func (d *Decoder) Decode(v interface{}) (err error) {
	lex := &lexer{scan: *d.scan}
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(v).Elem())
	return nil
}

func Unmarshal(data []byte, out interface{}) (err error) {
	dec := NewDecoder(bytes.NewReader(data))
	return dec.Decode(out)
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "t" {
			v.SetBool(true)
			lex.next()
			return
		}
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(f)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	case reflect.Interface:
		name, _ := strconv.Unquote(lex.text())
		item := reflect.New(typeMap[name]).Elem()
		lex.next()
		read(lex, item)
		v.Set(item)

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

var typeMap = make(map[string]reflect.Type)

func AddType(name string, typ reflect.Type) {
	typeMap[name] = typ
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
