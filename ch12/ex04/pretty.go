package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	b, _ := MarshalIndent(strangelove)
	fmt.Println(string(b))
}

func MarshalIndent(v interface{}) ([]byte, error) {
	p := printer{width: margin}
	if err := pretty(&p, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return p.Bytes(), nil
}

const margin = 1000

type token struct {
	kind rune
	str  string
	size int
}

type printer struct {
	tokens []*token
	stack  []*token
	rtotal int

	bytes.Buffer
	indents []int
	width   int

	keyLength int
	ind       int
}

func (p *printer) string(str string) {
	tok := &token{kind: 's', str: str, size: len(str)}
	if len(p.stack) == 0 {
		p.print(tok)
	} else {
		p.tokens = append(p.tokens, tok)
		p.rtotal += len(str)
	}
}
func (p *printer) pop() (top *token) {
	last := len(p.stack) - 1
	top, p.stack = p.stack[last], p.stack[:last]
	return
}
func (p *printer) begin() {
	p.ind++
	if len(p.stack) == 0 {
		p.rtotal = 1
	}
	t := &token{kind: '(', size: -p.rtotal}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.string("(")
}
func (p *printer) end(needReturn bool) {
	p.ind--
	if needReturn {
		p.string(")\n")
	} else {
		p.string(")")
	}
	p.tokens = append(p.tokens, &token{kind: ')'})
	x := p.pop()
	x.size += p.rtotal
	if x.kind == ' ' {
		p.pop().size += p.rtotal
	}
	if len(p.stack) == 0 {
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}
func (p *printer) space() {
	last := len(p.stack) - 1
	x := p.stack[last]
	if x.kind == ' ' {
		x.size += p.rtotal
		p.stack = p.stack[:last]
	}
	t := &token{kind: ' ', size: -p.rtotal}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.rtotal++
}
func (p *printer) print(t *token) {
	switch t.kind {
	case 's':
		p.WriteString(t.str)
		p.width -= len(t.str)
	case '(':
		p.indents = append(p.indents, p.width)
	case ')':
		p.indents = p.indents[:len(p.indents)-1]
	case ' ':
		if t.size > p.width {
			p.width = p.indents[len(p.indents)-1] - 1
			fmt.Fprintf(&p.Buffer, "\n%*s", margin-p.width, "")
		} else {
			p.WriteByte(' ')
			p.width--
		}
	}
}
func (p *printer) stringf(format string, args ...interface{}) {
	p.string(fmt.Sprintf(format, args...))
}

func pretty(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		p.stringf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p.stringf("%d", v.Uint())

	case reflect.String:
		p.stringf("%q", v.String())

	case reflect.Array, reflect.Slice:
		p.begin()
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				p.string("\n")
				for i := 0; i < p.keyLength; i++ {
					p.space()
				}
			}
			if err := pretty(p, v.Index(i)); err != nil {
				return err
			}
		}
		p.end(false)

	case reflect.Struct:
		p.begin()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				p.space()
			}
			p.begin()
			p.string(v.Type().Field(i).Name)
			p.keyLength = len(v.Type().Field(i).Name) + 2 + p.ind //"(key "
			p.space()
			if err := pretty(p, v.Field(i)); err != nil {
				return err
			}
			if i != v.NumField()-1 {
				p.end(true)
			} else {
				p.end(false)
			}
		}
		p.end(false)

	case reflect.Map:
		p.begin()
		for i, key := range v.MapKeys() {
			if i > 0 {
				for i := 0; i < p.keyLength; i++ {
					p.space()
				}
			}
			p.begin()
			if err := pretty(p, key); err != nil {
				return err
			}
			p.space()
			if err := pretty(p, v.MapIndex(key)); err != nil {
				return err
			}
			if i != len(v.MapKeys())-1 {
				p.end(true)
			} else {
				p.end(false)
			}
		}
		p.end(false)

	case reflect.Ptr:
		return pretty(p, v.Elem())

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
