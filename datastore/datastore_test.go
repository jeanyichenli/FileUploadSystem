package datastore

import "testing"

func TestSaveChunkToFile(t *testing.T) {
	testChunk := []byte("This is a test message.")
	name := "test.txt"

	err := SaveChunkToFile(testChunk, name)
	if err != nil {
		t.Errorf("Error occurs in SaveChunkToFile. error: %s", err.Error())
	}
}
