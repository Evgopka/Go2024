package main

import (
	"fmt"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

func filterParallel(img draw.RGBA64Image, y int, wg *sync.WaitGroup) {
	defer wg.Done() // Указываем, что горутина завершена
	bounds := img.Bounds()
	width := bounds.Max.X

	for x := 0; x < width; x++ {
		oldColor := img.RGBA64At(x, y)
		// Применяем преобразование к цветам (оттенки серого)
		grayValue := (oldColor.R + oldColor.G + oldColor.B) / 3
		newColor := color.RGBA64{grayValue, grayValue, grayValue, oldColor.A}
		img.SetRGBA64(x, y, newColor)
	}
}

// Функция для обработки изображения с использованием параллельной обработки
func filterImageParallel(drawImg draw.RGBA64Image) {
	bounds := drawImg.Bounds()
	height := bounds.Max.Y

	var wg sync.WaitGroup

	start := time.Now() // Начало замера времени

	for y := 0; y < height; y++ {
		wg.Add(1)                          // Увеличиваем счетчик горутин
		go filterParallel(drawImg, y, &wg) // Запускаем горутину для обработки строки
	}

	wg.Wait() // Ждем завершения всех горутин

	elapsed := time.Since(start) // Вычисляем затраченное время
	fmt.Printf("Параллельная обработка завершена за: %s\n", elapsed)
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

	filterImageParallel(drawImg)
	outFile, err := os.Create("ОработкаСерым.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
		return
	}

	fmt.Println("Изображение сохранено в месте где вы запустили программу, под названием: ОработкаСерым.png")
}
