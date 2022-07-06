package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//Valores pequeños para la imagen 1 -> Tigre
	//Valores grandes para la imagen 2 -> Catedral
	var determinante float64 = 0.1
	imgPath := "Imagenes/catedral.jpg"
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

	img, _, err := image.Decode(f)

	imgPath2 := "Imagenes/tigre.jpg"
	f2, err2 := os.Open(imgPath2)
	check(err2)
	defer f2.Close()

	img2, _, err2 := image.Decode(f2)

	size2 := img2.Bounds().Size()
	rect2 := image.Rect(0, 0, size2.X, size2.Y)
	wImg2 := image.NewRGBA(rect2)

	start := time.Now()
	// Primer buble determinado por el tamaño X de la imagen
	for x := 0; x < size2.X; x++ {
		// Segundo bucle determinado por el tamaño Y de la imagen 2
		for y := 0; y < size2.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			red_1 := float64(originalColor.R)
			green_1 := float64(originalColor.G)
			blue_1 := float64(originalColor.B)

			pixel2 := img2.At(x, y)
			originalColor2 := color.RGBAModel.Convert(pixel2).(color.RGBA)
			red_2 := float64(originalColor2.R)
			green_2 := float64(originalColor2.G)
			blue_2 := float64(originalColor2.B)

			// Aplicando formula
			red_definitivo := uint8((red_1 * determinante) + ((1 - determinante) * red_2))
			green_definitivo := uint8((green_1 * determinante) + ((1 - determinante) * green_2))
			blue_definitivo := uint8((blue_1 * determinante) + ((1 - determinante) * blue_2))

			pixel_definitivo := color.RGBA{
				R: red_definitivo, G: green_definitivo, B: blue_definitivo,
			}
			wImg2.Set(x, y, pixel_definitivo)
		}
	}

	//Tiempo de Ejecucion
	fin := time.Since(start)
	log.Printf("Ejercicio 2: Tiempo de ejecucion %s", fin)
	//Obteniendo extension y nombre para la imagen 1
	ext_fig1 := filepath.Ext(imgPath)
	name_fig1 := strings.TrimSuffix(filepath.Base(imgPath), ext_fig1)
	//Obteniendo extension y nombre para la imagen 2
	ext_fig2 := filepath.Ext(imgPath2)
	name_fig2 := strings.TrimSuffix(filepath.Base(imgPath2), ext_fig2)
	newImagePath2 := fmt.Sprintf("%s/%s_%s_blending%s", filepath.Dir(imgPath2), name_fig2, name_fig1, ext_fig2)
	fg2, err2 := os.Create(newImagePath2)
	defer fg2.Close()
	check(err2)
	err2 = jpeg.Encode(fg2, wImg2, nil)
	check(err2)

}
