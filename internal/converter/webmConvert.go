package converter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/arav-ind/ffmpego/internal/utils"
)

// convertFile converts a single MP4 file to WebM.
func convertFile(ffmpegPath, inputFile, outputFile string, wg *sync.WaitGroup) {
	defer wg.Done()

	utils.LogInfo(fmt.Sprintf("Converting: %s → %s", inputFile, outputFile))

	cmd := exec.Command(ffmpegPath, "-i", inputFile, "-c:v", "libvpx-vp9", "-b:v", "1M", "-c:a", "libopus", outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		utils.LogError(fmt.Sprintf("Error converting %s: %v", inputFile, err))
	} else {
		utils.LogSuccess(fmt.Sprintf("Successfully converted: %s", outputFile))
	}
}

// ConvertToWebM processes all MP4 files concurrently from an input folder to an output folder.
func ConvertToWebM(ffmpegPath, inputFolder, outputFolder string) error {
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output folder: %v", err)
	}

	files, err := os.ReadDir(inputFolder)
	if err != nil {
		return fmt.Errorf("failed to read input folder: %v", err)
	}

	var wg sync.WaitGroup

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".mp4" {
			inputFile := filepath.Join(inputFolder, file.Name())
			outputFile := filepath.Join(outputFolder, file.Name()[:len(file.Name())-4]+".webm")

			wg.Add(1)
			go convertFile(ffmpegPath, inputFile, outputFile, &wg)
		}
	}

	wg.Wait()
	utils.LogSuccess("All conversions completed.")
	return nil
}
