// Package filehandler provides functionality for handling files, specifically images.
package filehandler

import (
	"bytes"
	"files-management/pkg/logger"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

// FileHandler is an interface that defines the methods a file handler should have.
type FileHandler interface {
	// LoadImage loads an image from the provided file path and returns it as a byte slice.
	LoadImage(filePath string) ([]byte, error)
}

// DefaultFileHandler is a struct that implements the FileHandler interface.
type DefaultFileHandler struct{}

var ErrUnsupportedFormat = fmt.Errorf("unsupported image format")

// LoadImage is a method of DefaultFileHandler that opens a file at the provided path,
// decodes the image and returns it as a byte slice.
func (d DefaultFileHandler) LoadImage(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		if err == ErrUnsupportedFormat {
			return nil, fmt.Errorf("unsupported image format: %s", format)
		}
		return nil, err
	}

	// Encode the image to bytes
	return EncodeImage(img, format)
}

// encodeImage is a helper function that encodes an image to a byte slice based on its format.
func EncodeImage(img image.Image, format string) ([]byte, error) {
	buf := new(bytes.Buffer)
	switch format {
	case "jpeg":
		err := jpeg.Encode(buf, img, nil)
		return buf.Bytes(), err
	case "png":
		err := png.Encode(buf, img)
		return buf.Bytes(), err
	default:
		logger.Error("Unsupported image format: ", format)
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
}
