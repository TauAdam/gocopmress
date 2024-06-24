package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	config := &vips.Config{
		ConcurrencyLevel: 6,
	}
	vips.Startup(config)
	defer vips.Shutdown()

	startTime := time.Now()

	inputDir := "../../../Pictures/towebp"
	files, err := os.ReadDir(inputDir)
	checkError(err)

	outputDir := filepath.Join(inputDir, "webp_images")
	err = os.MkdirAll(outputDir, 0755)
	checkError(err)

	fmt.Println("Compressing images to webp format...")
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".jpg" {
			continue
		}

		imgPath := filepath.Join(inputDir, file.Name())
		img, err := vips.NewImageFromFile(imgPath)
		checkError(err)

		params := vips.NewWebpExportParams()
		webpBytes, _, err := img.ExportWebp(params)
		checkError(err)

		outputFile := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".webp"
		err = os.WriteFile(filepath.Join(outputDir, outputFile), webpBytes, 0644)
		checkError(err)

	}

	elapsedTime := time.Since(startTime)
	seconds := float64(elapsedTime) / float64(time.Second)
	fmt.Printf("Total time taken: %.1f seconds", seconds)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
