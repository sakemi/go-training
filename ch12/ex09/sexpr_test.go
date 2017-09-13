package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	data := `((foo ("bar" 1)))`
	dec := NewDecoder(strings.NewReader(data))
	tok := dec.Token()
	if _, ok := tok.(StartList); !ok {
		t.Errorf("%T", tok)
	}

	tok = dec.Token()
	if _, ok := tok.(StartList); !ok {
		t.Errorf("%T", tok)
	}

	tok = dec.Token()
	want := "foo"
	if tok, ok := tok.(Symbol); ok {
		if tok.Name != want {
			t.Errorf("want:%s but got:%s", want, tok.Name)
		}
	} else {
		t.Errorf("want:Symbol but got:%T", tok)
	}

	tok = dec.Token()
	if _, ok := tok.(StartList); !ok {
		t.Errorf("%T", tok)
	}

	tok = dec.Token()
	want = "bar"
	if tok, ok := tok.(String); ok {
		if tok.Value != want {
			t.Errorf("want:%s but got:%s", want, tok.Value)
		}
	} else {
		t.Errorf("want:Symbol but got:%T", tok)
	}

	tok = dec.Token()
	numWant := 1
	if tok, ok := tok.(Int); ok {
		if tok.Value != numWant {
			t.Errorf("want:%d but got:%d", numWant, tok.Value)
		}
	} else {
		t.Errorf("want:Symbol but got:%T", tok)
	}

	tok = dec.Token()
	if _, ok := tok.(EndList); !ok {
		t.Errorf("%T", tok)
	}

	tok = dec.Token()
	if _, ok := tok.(EndList); !ok {
		t.Errorf("%T", tok)
	}
}

func Test(t *testing.T) {
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

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	dec := NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Decode() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Unmarshal
	var movie2 Movie
	if err := Unmarshal(data, &movie2); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie2)

	// Check equality.
	if !reflect.DeepEqual(movie2, strangelove) {
		t.Fatal("not equal")
	}
}
