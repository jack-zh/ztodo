package rgbterm_test

import (
	"fmt"
	"github.com/jack-zh/ztodo/zterminal/rgbterm"
)

func ExampleBytes() {
	data := []byte("█")

	h, s, l := rgbterm.RGBtoHSL(252, 255, 43)
	for i := 0; i < 80; i++ {
		h += (5.0 / 360.0)
		if h > 1.0 {
			h = 0.0
		}
		r, g, b := rgbterm.HSLtoRGB(h, s, l)
		fmt.Printf("%s", rgbterm.Bytes(data, r, g, b, 0, 0, 0))
	}
	fmt.Println()

}

func ExampleBgBytes() {
	data := []byte(" ")

	h, s, l := rgbterm.RGBtoHSL(252, 255, 43)
	for i := 0; i < 80; i++ {
		h += (5.0 / 360.0)
		if h > 1.0 {
			h = 0.0
		}
		r, g, b := rgbterm.HSLtoRGB(h, s, l)
		fmt.Printf("%s", rgbterm.BgBytes(data, r, g, b))
	}
	fmt.Println()

}

func ExampleString() {
	var r, g, b uint8
	r, g, b = 252, 255, 43
	word := "=)"
	coloredWord := rgbterm.String(word, r, g, b, 0, 0, 0)

	fmt.Println("Oh!", coloredWord, "hello!")
}
