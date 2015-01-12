package rainbow

import (
	"bytes"
	"github.com/jack-zh/ztodo/zterminal/rgbterm"
	"io"
)

type Rainbow struct {
	wrap    io.Writer
	h, s, l float64
}

func New(w io.Writer, r, g, b uint8) *Rainbow {
	h, s, l := rgbterm.RGBtoHSL(r, g, b)
	return &Rainbow{wrap: w, h: h, s: s, l: l}
}

func (r *Rainbow) Write(p []byte) (int, error) {

	buf := bytes.NewBuffer(nil)
	for i := range p {
		r.h += (0.5 / 360)
		if r.h > 1.0 {
			r.h = 0
		}
		r, g, b := rgbterm.HSLtoRGB(r.h, r.s, r.l)
		_, _ = buf.Write(rgbterm.Byte(p[i], r, g, b))
	}

	_, err := buf.WriteTo(r.wrap)
	return len(p), err
}
