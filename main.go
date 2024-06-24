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
	vips.Startup(nil)
	defer vips.Shutdown()

	startTime := time.Now()

	inputDir := "../../../Pictures/towebp"
	files, err := os.ReadDir(inputDir)
	checkError(err)

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
		err = os.WriteFile(filepath.Join(inputDir, outputFile), webpBytes, 0644)
		checkError(err)

		fmt.Printf("Image %s compressed and saved to %s in ", file.Name(), outputFile)
	}

	elapsedTime := time.Since(startTime)
	seconds := float64(elapsedTime) / float64(time.Second)
	fmt.Printf("\nAll images compressed and saved in %.1f seconds\n", seconds)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}