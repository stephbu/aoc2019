package day8

import (
	"log"
	"strconv"
)

type SIF struct {
	X      int
	Y      int
	Layers []Layer
}

func LoadSIFFrom(pixels string, x int, y int) *SIF {
	result := &SIF{Layers: make([]Layer, 0), X: x, Y: y}

	for index := 0; index < len(pixels); {

		layer := Layer{}

		for currentX := 0; currentX < x; currentX++ {
			for currentY := 0; currentY < y; currentY++ {
				pixel, err := strconv.Atoi(string(pixels[index]))
				if err != nil {
					log.Fatal(err)
				}

				layer.Pixels = append(layer.Pixels, pixel)
				index++
			}
		}

		result.Layers = append(result.Layers, layer)

	}

	return result
}
