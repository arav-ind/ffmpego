package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/arav-ind/ffmpego/internal/converter"
)

func main() {
	mp4Folder := flag.String("input", "", "Path to the folder containing MP4 files")
	webmFolder := flag.String("output", "", "Path to the output folder for WebM files")

	// Parse command-line arguments
	flag.Parse()

	// Ensure both flags are provided
	if *mp4Folder == "" || *webmFolder == "" {
		fmt.Println("Usage: ffmpego --input <input_folder> --output <output_folder>")
		os.Exit(1)
	}
	
	ffmpegPath, err := converter.GetFFmpegPath()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Starting conversion process...")

	// Convert files
	if err := converter.ConvertToWebM(ffmpegPath, *mp4Folder, *webmFolder); err != nil {
		log.Fatalf("Error converting to WebM: %v", err)
	}

	fmt.Println("All files successfully converted to WebM.")
}
