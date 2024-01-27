package filehandler

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestLoadImageAndEncodeImage(t *testing.T) {
	// Create a sample image in memory for testing
	sampleImage := createSampleImage()

	// Test encoding as JPEG
	testJPEG := func() {
		tempFile := createTempFile(sampleImage, "jpg")
		defer tempFile.Close()

		// Test LoadImage function
		imgData, err := LoadImage(tempFile.Name())
		if err != nil {
			t.Fatalf("LoadImage failed: %v", err)
		}

		// Test EncodeImage function (for JPEG)
		var encodedImageJPEG bytes.Buffer
		err = EncodeImage(&encodedImageJPEG, sampleImage, "jpeg")
		if err != nil {
			t.Fatalf("EncodeImage (JPEG) failed: %v", err)
		}

		// Debugging: Print the lengths of the loaded image data and encoded data
		t.Logf("Loaded Image Data Length (JPEG): %d", len(imgData))
		t.Logf("Encoded JPEG Data Length: %d", len(encodedImageJPEG.Bytes()))

		// Compare the loaded image data with encoded image data
		if !bytes.Equal(imgData, encodedImageJPEG.Bytes()) {
			t.Error("Loaded image data and encoded JPEG data do not match.")
		}
	}

	// Test encoding as PNG
	testPNG := func() {
		tempFile := createTempFile(sampleImage, "png")
		defer tempFile.Close()

		// Test LoadImage function
		imgData, err := LoadImage(tempFile.Name())
		if err != nil {
			t.Fatalf("LoadImage failed: %v", err)
		}

		// Test EncodeImage function (for PNG)
		var encodedImagePNG bytes.Buffer
		err = EncodeImage(&encodedImagePNG, sampleImage, "png")
		if err != nil {
			t.Fatalf("EncodeImage (PNG) failed: %v", err)
		}

		// Debugging: Print the lengths of the loaded image data and encoded data
		t.Logf("Loaded Image Data Length (PNG): %d", len(imgData))
		t.Logf("Encoded PNG Data Length: %d", len(encodedImagePNG.Bytes()))

		// Compare the loaded image data with encoded image data
		if !bytes.Equal(imgData, encodedImagePNG.Bytes()) {
			t.Error("Loaded image data and encoded PNG data do not match.")
		}
	}

	// Run both tests
	testJPEG()
	testPNG()
}
func createSampleImage() image.Image {
	// Create a simple 100x100 pixel white image for testing
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, image.White)
		}
	}
	return img
}

func createTempFile(img image.Image, format string) *os.File {
	// Create a temporary file and save the image data to it
	tempFile, err := os.CreateTemp("", "test_image*."+format)
	if err != nil {
		panic(err)
	}

	// Encode and write the image to the temporary file
	switch format {
	case "jpg":
		err = jpeg.Encode(tempFile, img, nil)
	case "png":
		err = png.Encode(tempFile, img)
	default:
		panic("Unsupported image format for testing")
	}

	if err != nil {
		panic(err)
	}

	return tempFile
}
