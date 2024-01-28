package filehandler

import (
	"os"
	"testing"
)

func TestLoadImage(t *testing.T) {
	// Create a new DefaultFileHandler instance
	dfh := DefaultFileHandler{}

	// Test loading a valid JPEG image
	_, err := dfh.LoadImage("testdata/valid.JPG")
	if err != nil {
		t.Fatalf("Failed to load valid JPEG image: %v", err)
	}

	_, err = dfh.LoadImage("testdata/valid.png")
	if err != nil {
		t.Fatalf("Failed to load valid PNG image: %v", err)
	}

	// Test loading an unknown image format
	_, err = dfh.LoadImage("testdata/invalid.bmp")
	if err == nil {
		t.Fatal("Expected error for invalid image format")
	}
	if err.Error() != "image: unknown format" {
		t.Fatalf("%v", err)
	}

	// Test loading an unsupported image format
	/*_, err = dfh.LoadImage("testdata/invalid.gif")
	if err == nil {
		t.Fatal("Expected error for unsupported image format")
	}
	if err != ErrUnsupportedFormat {
		t.Fatalf("Expected %v, got %v", ErrUnsupportedFormat, err)
	}*/

	// Test error handling for invalid file path
	_, err = dfh.LoadImage("/nonexistent/file.jpg")
	if err == nil {
		t.Fatal("Expected error for nonexistent file path")
	}
	if !os.IsNotExist(err) {
		t.Fatalf("Expected 'IsNotExist' error, got %v", err)
	}
}
