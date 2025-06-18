package chunker

import (
	"os"
	"testing"

	gopath "path"
)

func TestSplitFile(t *testing.T) {
	testPath := "/Users/jean/code_practice/go_projects/FileUploadSystem/test"
	repoPath := "/Users/jean/code_practice/go_projects/FileUploadSystem/repo"
	filepath := gopath.Join(testPath, "1M.bin")

	// Open test file
	file, err := os.Open(filepath)
	if err != nil {
		t.Errorf("Failed to open test file. error: %s", err.Error())
	}
	defer file.Close()

	// Define chunk size
	chunkSize := 256 * 1024 // 256k

	// Split file into chunks
	err = SplitFile(file, chunkSize, repoPath)
	if err != nil {
		t.Errorf("Failed to split the file. error: %s", err.Error())
	}
}
