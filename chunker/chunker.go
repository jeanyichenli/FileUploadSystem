package chunker

import (
	"fmt"
	"io"
	"os"

	datastore "github.com/jeanyichenli/FileUploadSystem/datastore"
)

func SplitFile(file io.Reader, chunksize int) error {
	// Define variables
	chunkIndex := 0 // Start chunk index with 0
	repoPath := os.Getenv("REPO")

	for {
		chunk := make([]byte, chunksize)

		// Read chunk
		byteRead, err := file.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		// write chunk

		name := fmt.Sprintf("%s/%s_%d", repoPath, "outputChunk", chunkIndex)
		err = datastore.SaveChunkToFile(chunk[:byteRead], name)
		if err != nil {
			return err
		}

		chunkIndex++
	}

	return nil
}
