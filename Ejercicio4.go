package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	file, err := os.Open("03.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	imgSize := img.Bounds().Size()

	var redSum float64
	var greenSum float64
	var blueSum float64

	start := time.Now()
	for y := 0; y < imgSize.Y; y++ {
		for x := 0; x < imgSize.X; x++ {
			pixel := img.At(x, y)
			col := color.RGBAModel.Convert(pixel).(color.RGBA)
			redSum += float64(col.R)
			greenSum += float64(col.G)
			blueSum += float64(col.B)
		}

	}
	imgArea := float64(imgSize.X * imgSize.Y)
	redAverage := math.Round(redSum / imgArea)
	greenAverage := math.Round(greenSum / imgArea)
	blueAverage := math.Round(blueSum / imgArea)
	elapsed := time.Since(start)
	fmt.Printf("Ejercicio 4 Solución Normal:")
	log.Printf("Tiempo de ejecucion: %s", elapsed)
	fmt.Printf("Promedio de color: r %.0f, g %.0f, b %.0f)", redAverage, greenAverage, blueAverage)
}
