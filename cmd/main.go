package main

import (
	"fmt"
	"log"

	"github.com/arav-ind/ffmpego/internal/converter"
)

func main() {
	// Get the FFmpeg path
	ffmpegPath, err := converter.GetFFmpegPath()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Define input and output folders
	inputFolder := "./mp4"
	outputFolder := "./webm"

	fmt.Println("Starting conversion process...")

	// Convert files
	if err := converter.ConvertToWebM(ffmpegPath, inputFolder, outputFolder); err != nil {
		log.Fatalf("Error converting to WebM: %v", err)
	}

	fmt.Println("All files successfully converted to WebM.")
}
