package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"sync"
	"time"
)

var infile = flag.String("infile", "img.png", "path to image (gif, jpeg, png)")

func main() {

	flag.Parse()
	reader, err := os.Open("01.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	var histogram [16][4]int
	wg := new(sync.WaitGroup)
	start := time.Now()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		y := y
		go func() {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				histogram[r>>12][0]++
				histogram[g>>12][1]++
				histogram[b>>12][2]++
				histogram[a>>12][3]++
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	// Print the results.
	fin := time.Since(start)
	log.Printf("Ejercicio 3  Solución con concurrencia:")
	log.Printf("Tiempo de ejecucion %s", fin)
	fmt.Printf("%6s %6s %6s\n", "red", "green", "blue")
	for i, x := range histogram {
		fmt.Printf("%6d %6d %6d\n", x[0], x[1], x[2])
		i = i + 1
	}
}
