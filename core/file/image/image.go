package image

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Quality constants for image compression
const (
	SuperHighestQuality = 100
	HighQuality         = 80
	MediumQuality       = 60
	LowQuality          = 40
	LowestQuality       = 20
)

// CompressImage compresses an input image using ImageMagick.
// The compressed image is saved to the "public/storage" directory.
//
// Arguments:
// - inputPath: the path to the input image to compress.
// - outputFilename: the name of the compressed output image file.
// - quality: the desired quality of the compressed image.
//
// Returns:
// - string: the path to the compressed output image file.
// - error: an error if the compression operation failed.
func CompressImage(inputPath string, outputFilename string, quality int) (string, error) {
	// Construct the output path for the compressed image
	outputPath := filepath.Join("public", "storage", outputFilename)

	// Construct the ImageMagick command to compress the image
	var command string
	if runtime.GOOS == "windows" {
		command = fmt.Sprintf("magick -quality %d %s %s", quality, inputPath, outputPath)
	} else {
		command = fmt.Sprintf("convert -quality %d %s %s", quality, inputPath, outputPath)
	}

	// Run the command to compress the image
	cmd := exec.Command(command)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, nil
}

// ResizeImage resizes an input image using ImageMagick.
// The resized image is saved to the "public/storage" directory.
//
// Arguments:
// - inputPath: the path to the input image to resize.
// - outputFilename: the name of the resized output image file.
// - width: the desired width of the resized image.
// - height: the desired height of the resized image.
//
// Returns:
// - string: the path to the resized output image file.
// - error: an error if the resize operation failed.
func ResizeImage(inputPath string, outputFilename string, width int, height int) (string, error) {
	// Construct the output path for the resized image
	outputPath := filepath.Join("public", "storage", outputFilename)

	// Construct the ImageMagick command to resize the image
	var command string
	if runtime.GOOS == "windows" {
		command = fmt.Sprintf("magick %s -resize %dx%d %s", inputPath, width, height, outputPath)
	} else {
		command = fmt.Sprintf("convert %s -resize %dx%d %s", inputPath, width, height, outputPath)
	}

	// Run the command to resize the image
	cmd := exec.Command(command)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
