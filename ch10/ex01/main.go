package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	t := flag.String("t", "", "output encoding type: jpeg, png or gif")
	flag.Parse()

	var err error
	switch *t {
	case "jpeg":
		err = toJPEG(os.Stdin, os.Stdout)
		break
	case "png":
		err = toPNG(os.Stdin, os.Stdout)
		break
	case "gif":
		err = toGIF(os.Stdin, os.Stdout)
		break
	default:
		fmt.Fprintf(os.Stderr, "Unsupported encoding type: %s\n", *t)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", *t, err)
		os.Exit(2)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, nil)
}
