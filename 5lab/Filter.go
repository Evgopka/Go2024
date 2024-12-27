package main

import (
	"fmt"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"
)

// Функция filter, которая применяет преобразование к пикселям изображения
func filter(img draw.RGBA64Image) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Маска - отсерение (зато видно)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			oldColor := img.RGBA64At(x, y)
			grayValue := (oldColor.R + oldColor.G + oldColor.B) / 3
			newColor := color.RGBA64{grayValue, grayValue, grayValue, oldColor.A}
			img.SetRGBA64(x, y, newColor)
		}
	}
}

func main() {
	file, err := os.Open("C:\\Demo.png")
	if err != nil {
		fmt.Println("Файл отсутствует", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Неверный формат файла, нужен png:", err)
		return
	}

	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Преобразование в draw.RGBA64Image не удалось")
		return
	}

	start := time.Now() // Начало замера времени
	filter(drawImg)
	elapsed := time.Since(start)
	fmt.Printf("Маска за : %s\n", elapsed) // И конец

	outFile, err := os.Create("СерыйШакал.png")
	defer outFile.Close()

	err = png.Encode(outFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
		return
	}

	fmt.Println("На файл наложили маску файл теперь СерыйШакал.png")
}
