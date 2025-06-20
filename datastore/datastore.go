package datastore

import (
	"bufio"
	"os"
)

func SaveChunkToFile(chunk []byte, name string) error {

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(chunk)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// TODO: deduplicate chunk data

// TODO: git-like storage
