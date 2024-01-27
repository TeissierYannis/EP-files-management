package filehandler

import (
	"testing"
)

func TestLoadImage(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{"Valid JPEG", "tests_files/valid.JPEG", false},
		{"Valid PNG", "tests_files/valid.png", false},
		{"Valid PNG", "tests_files/valid.JPG", false},
		{"Invalid File", "tests_files/nonexistent.jpg", true},
		//{"Unsupported Format", "tests_files/invalid.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfh := DefaultFileHandler{}
			_, err := dfh.LoadImage(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadImage() for %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
