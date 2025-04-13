package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arav-ind/ffmpego/internal/converter"
)

func main() {
	mp4Folder := flag.String("input", "", "Path to the folder containing MP4 files")
	webmFolder := flag.String("output", "", "Path to the output folder for WebM files")
	action := flag.String("action", "", "Operation to perform: convertToWebm or convertResolution")
	resolutions := flag.String("resolution", "1280", "Comma-separated resolutions for convertResolution (e.g., 480,360,240)")

	// Parse command-line arguments
	flag.Parse()

	// Ensure required flags are provided
	if *mp4Folder == "" || *webmFolder == "" || *action == "" {
		fmt.Println("Usage: ffmpego --input <input_folder> --output <output_folder> --action <convertToWebm|convertResolution> [--resolution 480,360,240]")
		os.Exit(1)
	}

	ffmpegPath, err := converter.GetFFmpegPath()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Starting process...")

	switch *action {
		case "convertToWebm":
			if err := converter.ConvertToWebM(ffmpegPath, *mp4Folder, *webmFolder); err != nil {
				log.Fatalf("Error converting to WebM: %v", err)
			}
		case "convertResolution":
			resolutionList := strings.Split(*resolutions, ",")
			if err := converter.ConvertResolution(ffmpegPath, *mp4Folder, *webmFolder, resolutionList); err != nil {
				log.Fatalf("Error converting resolution: %v", err)
			}
		default:
			fmt.Println("Invalid action. Use convertToWebm or convertResolution.")
			os.Exit(1)
	}

	fmt.Println("Process completed successfully.")
}
