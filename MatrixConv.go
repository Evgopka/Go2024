package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

const (
	kernelSize = 3 // Ядер на свёртку
)

var kernel = [kernelSize][kernelSize]float64{
	{1 / 16.0, 2 / 50.0, 1 / 16.0}, // 0.0625	0.125	0.0625
	{2 / 16.0, 4 / 50.0, 2 / 16.0}, // 0.125	0.25	0.125
	{1 / 16.0, 2 / 50.0, 1 / 16.0}, // 0.0625	0.125	0.0625
}

func applyFilter(src draw.RGBA64Image, dst draw.RGBA64Image, x, y int, wg *sync.WaitGroup) {
	defer wg.Done()

	bounds := src.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	var r, g, b, a float64

	for ky := 0; ky < kernelSize; ky++ {
		for kx := 0; kx < kernelSize; kx++ {
			nx := x + kx - 1
			ny := y + ky - 1

			if nx >= 0 && nx < width && ny >= 0 && ny < height {
				color := src.RGBA64At(nx, ny) // А - Альфа канал, можно зашакалить
				r += float64(color.R) * kernel[ky][kx]
				g += float64(color.G) * kernel[ky][kx]
				b += float64(color.B) * kernel[ky][kx]
				a += float64(color.A) * kernel[ky][kx]
			}
		}
	}

	dst.SetRGBA64(x, y, color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	})
}

func filterImageConvolution(src draw.RGBA64Image) draw.RGBA64Image {
	bounds := src.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	dst := image.NewRGBA64(bounds)

	var wg sync.WaitGroup

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			wg.Add(1)
			go applyFilter(src, dst, x, y, &wg)
		}
	}

	wg.Wait()

	return dst
}

func main() {
	file, err := os.Open("C:\\Demo.png")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)

	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Неверный формат файла, нужен png:")
		return
	}

	start := time.Now() // Начало замера времени
	resultImg := filterImageConvolution(drawImg)
	elapsed := time.Since(start)
	fmt.Printf("Маска за : %s\n", elapsed) // И конец

	outFile, err := os.Create("МатричнаяДемка.png")
	defer outFile.Close()

	err = png.Encode(outFile, resultImg)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
		return
	}
	fmt.Println("Успешное выполнение")
}
