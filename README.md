# ffmpego!

A command-line tool that utilizes FFmpeg to process video files. 

## 📦 Features

- Convert `.mp4` files to `.webm` format  
- Convert videos (both `.mp4` and `.webm`) to multiple resolutions (`240p`, `360p`, `480p`, etc.)

---

## 🧪 Usage

```bash
ffmpego --input <input_folder> --output <output_folder> --action <convertToWebm|convertResolution> --resolution 480,360,240
```

---

### ✅ Examples

#### 1. Convert all `.mp4` files in `./mp4` folder to `.webm` format in `./webm` folder:
```bash
ffmpego --input ./mp4 --output ./webm --action convertToWebm
```

#### 2. Convert all `.mp4` or `.webm` files in `./mp4` folder to multiple resolutions (`480p`, `360p`, `240p`) and output to `./webm`:
```bash
ffmpego --input ./mp4 --output ./webm --action convertResolution --resolution 480,360,240
```

> ℹ️ The resolution output will match the input format (i.e., input `.mp4` → output `.mp4`, input `.webm` → output `.webm`).
