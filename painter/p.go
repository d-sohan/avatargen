package painter

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func hexToRGB(hexColor string) (color.Color, error) {
	rhex, bhex, ghex := hexColor[0:2], hexColor[2:4], hexColor[4:6]
	red, errRed := strconv.ParseUint(rhex, 16, 8)
	green, errGreen := strconv.ParseUint(ghex, 16, 8)
	blue, errBlue := strconv.ParseUint(bhex, 16, 8)
	if errRed != nil || errGreen != nil || errBlue != nil {
		return nil, errors.New("failed to parse hex color")
	}
	c := color.RGBA{uint8(red), uint8(green), uint8(blue), 0xff}
	return c, nil
}

func paintBlock(img *image.RGBA, startX, startY, block int, c color.Color) {
	if img.Bounds().Min.X > startX || img.Bounds().Min.Y > startY {
		return
	}
	for x := startX; x < startX+block && x < img.Bounds().Max.X; x++ {
		for y := startY; y < startY+block && y < img.Bounds().Max.Y; y++ {
			img.Set(x, y, c)
		}
	}
}

func Paint(count, block int, hexColorFG, hexColorBG, output string) error {

	// error if invalid hex color

	colourFG, err := hexToRGB(hexColorFG)
	if err != nil {
		return err
	}

	colourBG, err := hexToRGB(hexColorBG)
	if err != nil {
		return err
	}

	// error if output path is invalid
	f, perr := os.Create(output)
	if perr != nil {
		return perr
	}

	rand.Seed(time.Now().UnixNano())

	border := block / 2
	width := count*block + 2*border
	height := width

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// paint border
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if l, r := border, border+count*block; x < l || x >= r || y < l || y >= r {
				img.Set(x, y, colourBG)
			}
		}
	}

	// paint count x count grid
	for i := 0; i < count; i++ {
		x := border + i*block
		for j := 0; j < count; j++ {
			y := border + j*block
			if i <= (count-1)/2 {
				if rand.Int63()%2 == 0 {
					paintBlock(img, x, y, block, colourFG)
				} else {
					paintBlock(img, x, y, block, colourBG)
				}
			} else {
				c := img.At(border+(count-i-1)*block, y)
				paintBlock(img, x, y, block, c)
			}
		}
	}

	// Encode as PNG.
	return png.Encode(f, img)
}
