package painter

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

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

func Paint(count, block int) {
	rand.Seed(time.Now().UnixNano())

	border := block / 2
	width := count*block + 2*border
	height := width

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// paint border
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if l, r := border, border+count*block; x < l || x >= r || y < l || y >= r {
				img.Set(x, y, color.White)
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
					paintBlock(img, x, y, block, cyan)
				} else {
					paintBlock(img, x, y, block, color.White)
				}
			} else {
				c := img.At(border+(count-i-1)*block, y)
				paintBlock(img, x, y, block, c)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("avatar.png")
	png.Encode(f, img)
}
