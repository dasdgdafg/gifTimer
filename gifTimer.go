package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"os"
)

var repititions int
var totalTime int
var inFilename string
var outFilename string
var force bool

func main() {
	flag.StringVar(&inFilename, "i", "", "input file")
	flag.IntVar(&repititions, "r", 1, "repititions of the gif")
	flag.IntVar(&totalTime, "t", 100, "total time in 1/100 seconds")
	flag.StringVar(&outFilename, "o", "", "output file")
	flag.BoolVar(&force, "f", false, "overwrite output file if it already exists")
	flag.Parse()

	if inFilename == "" {
		fmt.Println("input file (-i) is required")
		return
	}
	if outFilename == "" {
		fmt.Println("output file (-o) is required")
		return
	}
	if totalTime < 2 {
		fmt.Println("total time (-t) must be at least 2 (0.02s)")
		return
	}
	if repititions < 1 {
		fmt.Println("repititions (-r) must be at least 1")
		return
	}

	inputFile, err := os.Open(inFilename)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	inputImage, err := gif.DecodeAll(inputFile)
	if err != nil {
		panic(err)
	}

	prevTime := 0
	numFrames := len(inputImage.Delay) * repititions
	newDelays := []int{}
	newFrames := []*image.Paletted{}
	newDisposals := []byte{}
	for i := 0; i < numFrames; i++ {
		thisTime := (i + 1) * totalTime / numFrames
		delay := thisTime - prevTime
		// delay 1 (0.01s, or 100 fps) gifs don't display correctly on my computer, so limit to 2 (0.02s, or 50 fps)
		if delay >= 2 {
			newDelays = append(newDelays, delay)
			newFrames = append(newFrames, inputImage.Image[i%len(inputImage.Image)])
			newDisposals = append(newDisposals, inputImage.Disposal[i%len(inputImage.Disposal)])
			prevTime = thisTime
		}
	}
	// if the last frame was dropped, extend the frame before it to make the total time correct
	if prevTime != totalTime {
		newDelays[len(newDelays)-1] += totalTime - prevTime
	}
	outputImage := inputImage
	outputImage.Image = newFrames
	outputImage.Delay = newDelays
	outputImage.Disposal = newDisposals

	openFlags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	if !force {
		openFlags |= os.O_EXCL
	}
	outputFile, err := os.OpenFile(outFilename, openFlags, 0666)
	if os.IsExist(err) {
		fmt.Printf("output file %v exists, use -f to overwrite\n", outFilename)
		return
	} else if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = gif.EncodeAll(outputFile, outputImage)
	if err != nil {
		panic(err)
	}
}
