package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	const width = 5000
	const height = 5000
	const h_width = width / 2
	const h_height = height / 2
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := x - h_width
			dy := h_height - y
			if (dx*dx)-(abs(dx)*dy)+(dy*dy) <= 5000 {
				m.SetRGBA(x, y, color.RGBA{
					0xff, 0x1f, 0x00, 0xff,
				})
			} else {
				m.SetRGBA(x, y, color.RGBA{
					0xff, 0xff, 0xff, 0xff,
				})
			}
		}
	}
	fmt.Println(m.Bounds())

	outFile, _ := os.Create("output_image.png")
	defer outFile.Close()
	png.Encode(outFile, m)
}
