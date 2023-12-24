package golum

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

type ImageBuffer struct {
	w, h int
	data [][]pixelData
}

type pixelData struct {
	r, g, b, cnt int
}

func NewImageBuffer(w, h int) ImageBuffer {
	ret := ImageBuffer{}
	ret.w = w
	ret.h = h
	ret.data = make([][]pixelData, h)
	for i := range ret.data {
		ret.data[i] = make([]pixelData, w)
	}
	return ret
}

func (buf ImageBuffer) AddData(x, y int, color Vec3) {
	buf.data[x][y].r += int(color.X * 255)
	buf.data[x][y].g += int(color.Y * 255)
	buf.data[x][y].b += int(color.Z * 255)
	buf.data[x][y].cnt++
}

func (buf ImageBuffer) toImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, buf.w, buf.h))
	for i := 0; i < buf.h; i++ {
		for j := 0; j < buf.w; j++ {
			cur := buf.data[i][j]
			if cur.cnt == 0 {
				img.Set(i, j, color.RGBA{R: 0, G: 0, B: 0, A: 255})
				continue
			}

			r := uint8(cur.r / cur.cnt)
			g := uint8(cur.g / cur.cnt)
			b := uint8(cur.b / cur.cnt)

			img.Set(i, j, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return img
}

func (buf ImageBuffer) WritePng(filename string) {
	img := buf.toImage()
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func (buf ImageBuffer) WriteJpeg(filename string, quality int) {
	img := buf.toImage()
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = jpeg.Encode(f, img, &jpeg.Options{Quality: quality})
	if err != nil {
		panic(err)
	}
}
