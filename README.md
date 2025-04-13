# ffmpego

A go wrapper around ffmpeg for fast, batch video format conversion, resolution adjustments, and custom transcoding processes.

## Prerequisites

You must have **ffmpeg** installed and accessible in one of the following ways:

- **Option 1:** ffmpeg is available in your system’s `PATH` (i.e., you can run `ffmpeg` from your terminal or command prompt).  
- **Option 2:** The `ffmpeg` file (folder) that you download must be located in the same directory as the `ffmpego` executable.

Download ffmpeg from the official site:  
[https://www.ffmpeg.org/download.html](https://www.ffmpeg.org/download.html)

## Features

- Convert `.mp4` files to `.webm` format  
- Convert videos (both `.mp4` and `.webm`) to multiple resolutions (`240p`, `360p`, `480p`, etc.)

---

## Usage

```bash
ffmpego --input <input_folder> --output <output_folder> --action <convertToWebm|convertResolution> --resolution 480,360,240
```

---

### Examples

#### 1. Convert all `.mp4` files in `./mp4` folder to `.webm` format in `./webm` folder:
```bash
ffmpego --input ./mp4 --output ./webm --action convertToWebm
```

#### 2. Convert all `.mp4` or `.webm` files in `./mp4` folder to multiple resolutions (`480p`, `360p`, `240p`) and output to `./webm`:
```bash
ffmpego --input ./mp4 --output ./webm --action convertResolution --resolution 480,360,240
```

> ℹ️ The resolution output will match the input format (i.e., input `.mp4` → output `.mp4`, input `.webm` → output `.webm`).
