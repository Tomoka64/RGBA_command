package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"strconv"
	"strings"
)

var Rgba = flag.String("RGBA", "256 256 255 255", "put RGBA number between 0 and 256 (uint8)")
var X = flag.Int("X", 255, "put width")
var Y = flag.Int("Y", 255, "put height")

type Image struct {
	width  int
	height int
	RGBA   RGBA
}

type RGBA struct {
	R, G, B, A uint8
}

func main() {
	RG := RGBAHandler()
	m := Image{*X, *Y, RG}
	fmt.Print("\033]1337;")
	fmt.Print("File=inline=1:")
	ShowImage(m)
	fmt.Print("\a\n")
}

func New(r, g, b, a uint8) (RGBA, error) {
	return RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}, nil
}

func RGBAHandler() RGBA {
	flag.Parse()
	Trimmed := strings.TrimSpace(*Rgba)
	slice := strings.Split(Trimmed, " ")
	Adjusted, err := Adjust(slice)
	if err != nil {
		log.Fatalln(err)
	}
	RG, _ := New(Adjusted[0], Adjusted[1], Adjusted[2], Adjusted[3])
	return RG
}

func Adjust(s []string) ([]uint8, error) {
	var ret []uint8
	if len(s) != 4 {
		return ret, errors.New("4 numbers must be included")
	}
	for _, component := range s {
		un, err := strconv.Atoi(component)
		fmt.Println(un)
		if err != nil {
			log.Fatalln("could not adjust type", err)
		}
		ret = append(ret, uint8(un))
	}
	return ret, nil
}
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.RGBA.R, i.RGBA.G, i.RGBA.B, i.RGBA.A}
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Print(enc)
}
