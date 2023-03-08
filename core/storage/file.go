package storage

import (
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

// FileExists checks if the specified file exists in the file system.
// It takes in a string argument representing the name of the file to check for existence,
// and returns a boolean value indicating whether the file exists or not.
func FileExists(filename string) bool {
	// Attempt to retrieve information about the file
	info, err := os.Stat(filename)

	// If the file doesn't exist, return false
	if os.IsNotExist(err) {
		return false
	}

	// If the file exists and is not a directory, return true
	return !info.IsDir()
}

// SaveFile saves a file to the specified path.
// It takes in an io.Reader object representing the source file, and a string representing the destination file path.
// If the directory for the destination file does not exist, the function creates it.
// The function creates the destination file and copies the contents of the source file to it.
// If the operation is successful, the function returns nil. Otherwise, it returns an error object with the corresponding error message.
func SaveFile(src io.Reader, dst string) error {
	// Create the directory if it does not exist
	err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return err
	}

	// Create the destination file
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	return nil
}

// CompressImage compresses the specified image for web usage with a target width and quality.
// It takes in the source and destination file paths, target width, and image quality as arguments,
// and returns an error if one occurs during image processing or file I/O.
// The function resizes the image to the target width using the Lanczos3 algorithm, and encodes
// the result in JPEG format with the specified quality level. The resulting image is saved to
// the specified destination file path.
func CompressImage(srcPath, destPath string, targetWidth, quality int) error {
	// Open the source image file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Decode the image from the source file
	srcImg, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}

	// Resize the image to the target width
	destImg := resize.Resize(uint(targetWidth), 0, srcImg, resize.Lanczos3)

	// Create the destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Encode the destination image to JPEG format with the specified quality
	jpegOpts := &jpeg.Options{Quality: quality}
	if err := jpeg.Encode(destFile, destImg, jpegOpts); err != nil {
		return err
	}

	return nil
}
