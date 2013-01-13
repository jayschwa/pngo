/*
PNGo reads in Images and writes out PNGs.
*/
package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"strings"
)

func convert(name string) {
	fmt.Print(name, " -> ")

	// Open input image file for reading
	fin, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Decode image data
	img, _, err := image.Decode(fin)
	if err != nil {
		log.Fatal(err)
	}

	// Change extension to .png
	ext := path.Ext(name)
	if len(ext) > 0 {
		name = strings.Replace(name, ext, "", -1)
	}
	name += ".png"

	// Create output file for writing
	fout, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := fout.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Write PNG data to output file
	err = png.Encode(fout, img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
}

func main() {
	flag.Parse()
	for _, f := range flag.Args() {
		convert(f)
	}
}
