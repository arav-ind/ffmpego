package converter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/arav-ind/ffmpego/internal/utils"
)

// ConvertResolution processes all MP4 and WebM files in the input folder
// and converts each file to multiple resolutions specified in the resolutions slice.
func ConvertResolution(ffmpegPath, inputFolder, outputFolder string, resolutions []string) error {
	// Create the output folder, if it doesn't exist.
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		utils.LogError(fmt.Sprintf("Failed to create output folder: %v", err))
		return err
	}

	// Read all files in the input folder.
	files, err := os.ReadDir(inputFolder)
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to read input folder: %v", err))
		return err
	}

	var wg sync.WaitGroup

	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == ".mp4" || ext == ".webm" {
			inputFile := filepath.Join(inputFolder, file.Name())
			// Remove the extension from the file name.
			baseName := file.Name()[:len(file.Name())-len(ext)]

			// For each resolution, create an output file and process conversion concurrently.
			for _, res := range resolutions {
				// For example, if res = "240" the output file name becomes "video_240p.mp4" or "video_240p.webm" based on the input.
				outputFile := filepath.Join(outputFolder, fmt.Sprintf("%s_%sp%s", baseName, res, ext))

				wg.Add(1)
				go func(input, output, resolution string) {
					defer wg.Done()

					utils.LogInfo(fmt.Sprintf("Resizing: %s → %s", input, output))

					// Use scale filter "-vf" with "-2" to enforce even width.
					cmd := exec.Command(ffmpegPath, "-i", input,
						"-vf", "scale=-2:"+resolution,
						"-c:v", "libx264", "-preset", "fast", "-crf", "23",
						"-c:a", "aac", "-b:a", "128k",
						output)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					if err := cmd.Run(); err != nil {
						utils.LogError(fmt.Sprintf("Error resizing %s: %v", input, err))
					} else {
						utils.LogSuccess(fmt.Sprintf("Successfully resized: %s", output))
					}
				}(inputFile, outputFile, res)
			}
		} else {
			utils.LogInfo(fmt.Sprintf("Skipping unsupported file: %s", file.Name()))
		}
	}

	wg.Wait()
	utils.LogInfo("Resolution conversion completed.")
	return nil
}
