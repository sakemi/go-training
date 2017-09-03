package main

import "testing"

type Movie struct {
	Title, Subtitle string
	Year            int
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	F               float64
	C               complex128
	B               bool
	Ch              chan string
	Func            func(int)
	I               interface{}
}

func TestZeroValue(t *testing.T) {
	var z Movie
	data, err := Marshal(z)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	if "" != string(data) {
		t.Errorf("Zero value are encoded: %s", string(data))
	}
}

func Test(t *testing.T) {
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
		F:    3.1415,
		C:    complex(1, 2),
		B:    true,
		Ch:   make(chan string),
		Func: func(x int) {},
		I:    struct{ x int }{1},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	// var movie Movie
	// if err := Unmarshal(data, &movie); err != nil {
	// 	t.Fatalf("Unmarshal failed: %v", err)
	// }
	// t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	// if !reflect.DeepEqual(movie, strangelove) {
	// 	t.Fatal("not equal")
	// }

	// Pretty-print it:
	// data, err = MarshalIndent(strangelove)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Logf("MarshalIdent() = %s\n", data)
}
