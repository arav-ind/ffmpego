package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetFFmpegPath returns the absolute path to FFmpeg based on the OS.
func GetFFmpegPath() (string, error) {
	// Check if user has set FFMPEG_PATH environment variable
	if envPath := os.Getenv("FFMPEG_PATH"); envPath != "" {
		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}
	}

	// Fallback to local ffmpeg binary
	var ffmpegPath string

	if runtime.GOOS == "windows" {
		ffmpegPath, _ = filepath.Abs("./ffmpeg/bin/ffmpeg.exe")
	} else {
		ffmpegPath, _ = filepath.Abs("./ffmpeg/bin/ffmpeg")
	}

	// Check if FFmpeg binary exists
	if _, err := os.Stat(ffmpegPath); os.IsNotExist(err) {
		return "", fmt.Errorf("FFmpeg binary not found at: %s", ffmpegPath)
	}

	return ffmpegPath, nil
}
